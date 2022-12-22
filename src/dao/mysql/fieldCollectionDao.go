package mysql

import (
	"abix360/database"
	"abix360/src/domain"
	"abix360/src/view/dto"
	"bytes"
	"database/sql"
	"log"
)

type FieldCollectionDao struct {
	db *database.ConnectDB
}

func NewFieldCollectionDao() *FieldCollectionDao {
	return &FieldCollectionDao{
		db: database.Instance(),
	}
}

func (f *FieldCollectionDao) Add(fc domain.FieldCollection) error {
	var strSQL bytes.Buffer
	strSQL.WriteString("INSERT INTO fields_collection(collection_id, field_id, isUnique, required, editable, sequence_id, list_id) VALUES (?, ?, ?, ?, ?, ?, ?)")
	stmt, err := f.db.Source().Conn().Prepare(strSQL.String())
	if err != nil {
		log.Println("abix-admin / dao / FieldCollectionDao / Add / Conn().Prepare: ", err)
	}

	var idSequence sql.NullInt64
	if fc.Sequence().Exists() {
		idSequence = sql.NullInt64{Int64: fc.Sequence().Id(), Valid: true}
	}

	var idList sql.NullInt64
	if fc.List().Exists() {
		idList = sql.NullInt64{Int64: fc.List().Id(), Valid: true}
	}

	_, err = stmt.Exec(fc.IdCollection(), fc.IdField(), fc.Unique(), fc.Required(), fc.Editable(), idSequence, idList)
	if err != nil {
		log.Println("abix-admin / dao / FieldCollectionDao / Add / stmt.Exec: ", err)
	}

	return err
}

func (f *FieldCollectionDao) Remove(fc domain.FieldCollection) error {
	var strSQL bytes.Buffer
	strSQL.WriteString("DELETE FROM fields_collection WHERE collection_id = ? AND field_id = ?")
	stmt, err := f.db.Source().Conn().Prepare(strSQL.String())
	if err != nil {
		log.Println("abix-admin / dao / FieldCollectionDao / Remove / Conn().Prepare: ", err)
	}

	_, err = stmt.Exec(fc.IdCollection(), fc.IdField())
	if err != nil {
		log.Println("abix-admin / dao / FieldCollectionDao / Remove / stmt.Exec: ", err)
	}

	return err
}

func (f *FieldCollectionDao) AllFields(idCollection int64) (fields []domain.FieldCollection) {
	var field domain.IField
	var collection domain.Collection

	var strSQL bytes.Buffer
	strSQL.WriteString("SELECT ")
	strSQL.WriteString("c.id, c.name, f.id, f.name, f.description, f.abbreviation, ")
	strSQL.WriteString("fc.isUnique, fc.required, fc.editable, IF (count(s.field_id)>0, 1, 0) as hasSubfields, fc.id as idFieldCollection, ")
	strSQL.WriteString("fc.sequence_id, fc.list_id ")
	strSQL.WriteString("FROM fields f ")
	strSQL.WriteString("INNER JOIN fields_collection fc ON f.id = fc.field_id ")
	strSQL.WriteString("INNER JOIN collections c ON c.id = fc.collection_id ")
	strSQL.WriteString("LEFT JOIN subfields s ON s.field_id = f.id ")
	strSQL.WriteString("WHERE c.id = ? ")
	strSQL.WriteString("GROUP BY c.id, c.name, f.id, f.name, f.description, f.abbreviation, fc.isUnique, fc.required, fc.editable, fc.id ")
	strSQL.WriteString("ORDER BY f.name ")

	stmt, err := f.db.Source().Conn().Prepare(strSQL.String())
	if err != nil {
		log.Println("abix-admin / dao / FieldCollectionDao / AllFields / Conn().Prepare: ", err)
	}

	rows, err := stmt.Query(idCollection)
	if err != nil {
		log.Println("abix-admin / dao / FieldCollectionDao / AllFields / stmt.Query: ", err)
	}
	defer rows.Close()

	idCollection = 0

	for rows.Next() {
		var idField, idFieldCollection int64
		var nameField, nameCollection, description, abbreviation string
		var unique, required, editable, hasSubfields bool
		var idSequence sql.NullInt64
		var idList sql.NullInt64

		rows.Scan(&idCollection, &nameCollection, &idField, &nameField, &description, &abbreviation, &unique, &required, &editable, &hasSubfields, &idFieldCollection, &idSequence, &idList)

		var fieldCollection domain.FieldCollection = domain.FieldCollection{}

		fieldCollection.WithId(idFieldCollection)
		fieldCollection.WithUnique(unique).WithEditable(editable).WithRequired(required)

		if idSequence.Valid {
			sequence := domain.Sequence{}
			sequence.WithId(idSequence.Int64)
			fieldCollection.WithSequence(sequence)
		}

		if idList.Valid {
			list := domain.List{}
			list.WithId(idList.Int64)
			fieldCollection.WithList(list)
		}

		field = domain.NewSingleField(nameField)
		if hasSubfields {
			field = domain.NewCompositeField(nameField, f.getSubfields(idField))
		}
		field.WithId(idField).WithAbbreviation(abbreviation).WithDescription(description)

		collection = *domain.NewCollection(nameCollection).WithId(idCollection)

		fieldCollection.WithField(field)
		fieldCollection.WithCollection(collection)

		fields = append(fields, fieldCollection)
	}

	return fields
}

func (f *FieldCollectionDao) FindByIdCollectionAndIdField(idCollection, idField int64) domain.FieldCollection {
	var fieldCollection domain.FieldCollection
	var collection domain.Collection
	var field domain.IField

	var strSQL bytes.Buffer
	strSQL.WriteString("SELECT ")
	strSQL.WriteString("c.id, c.name, f.id, f.name, f.description, f.abbreviation, fc.isUnique, fc.required, fc.editable, count(s.id) as hasSubfields, fc.id, fc.sequence_id, fc.list_id ")
	strSQL.WriteString("FROM fields f ")
	strSQL.WriteString("INNER JOIN fields_collection fc ON f.id = fc.field_id ")
	strSQL.WriteString("INNER JOIN collections c ON c.id = fc.collection_id ")
	strSQL.WriteString("LEFT JOIN subfields s ON s.field_id = f.id ")
	strSQL.WriteString("WHERE f.id = ? AND c.id = ? ")
	strSQL.WriteString("GROUP BY c.id, c.name, f.id, f.name, f.description, f.abbreviation, fc.isUnique, fc.required, fc.editable")

	stmt, err := f.db.Source().Conn().Prepare(strSQL.String())
	if err != nil {
		log.Println("abix-admin / dao / FieldCollectionDao / FindById / Conn().Prepare: ", err)
	}

	row := stmt.QueryRow(idField, idCollection)

	idCollection = 0
	idField = 0
	var nameField, nameCollection, description, abbreviation string
	var unique, required, editable, hasSubfields bool
	var idSequence, idList sql.NullInt64
	var id int64

	row.Scan(&idCollection, &nameCollection, &idField, &nameField, &description, &abbreviation, &unique, &required, &editable, &hasSubfields, &id, &idSequence, &idList)

	collection = *domain.NewCollection(nameCollection).WithId(idCollection)
	field = domain.NewSingleField(nameField)
	if hasSubfields {
		field = domain.NewCompositeField(nameCollection, f.getSubfields(idField))
	}

	field.WithId(idField).WithAbbreviation(abbreviation).WithDescription(description)

	fieldCollection.WithUnique(unique).WithEditable(editable).WithRequired(required).WithCollection(collection).WithField(field).WithId(id)
	if idSequence.Valid {
		sequence := domain.Sequence{}
		sequence.WithId(idSequence.Int64)
		fieldCollection.WithSequence(sequence)
	}

	if idList.Valid {
		list := domain.List{}
		list.WithId(idList.Int64)
		fieldCollection.WithList(list)
	}

	return fieldCollection
}

func (f *FieldCollectionDao) Update(fc domain.FieldCollection) error {
	var strSQL bytes.Buffer
	strSQL.WriteString("UPDATE fields_collection SET isUnique=?, required=?, editable=?, sequence_id=?, list_id=?, updatedAt=NOW() WHERE collection_id = ? AND field_id = ?")
	stmt, err := f.db.Source().Conn().Prepare(strSQL.String())
	if err != nil {
		log.Println("abix-admin / dao / FieldCollectionDao / Update / Conn().Prepare: ", err)
	}

	var idSequence sql.NullInt64
	if fc.Sequence().Exists() {
		idSequence = sql.NullInt64{Int64: fc.Sequence().Id(), Valid: true}
	}

	var idList sql.NullInt64
	if fc.List().Exists() {
		idList = sql.NullInt64{Int64: fc.List().Id(), Valid: true}
	}

	_, err = stmt.Exec(fc.Unique(), fc.Required(), fc.Editable(), idSequence, idList, fc.IdCollection(), fc.IdField())
	if err != nil {
		log.Println("abix-admin / dao / FieldCollectionDao / Update / stmt.Exec: ", err)
	}

	return err
}

func (f *FieldCollectionDao) getSubfields(idField int64) []dto.SubfieldDto {
	var subfields []dto.SubfieldDto
	var strQuery bytes.Buffer

	strQuery.WriteString("SELECT ")
	strQuery.WriteString("f2.id, f2.name, f2.description, f2.abbreviation, f2.createdAt, f2.updatedAt, s.orderBy ")
	strQuery.WriteString("FROM subfields s ")
	strQuery.WriteString("INNER JOIN fields f ON f.id = s.field_id ")
	strQuery.WriteString("INNER JOIN fields f2 ON f2.id = s.subfield_id ")
	strQuery.WriteString("WHERE s.field_id = ? ")

	rows, err := f.db.Source().Conn().Query(strQuery.String(), idField)
	if err != nil {
		log.Println("abix-admin / FieldDao / getSubfields / c.db.Source().Conn().Query: ", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var name, description, abbreviation, createdAt, updatedAt string
		var orderBy int
		var id int64
		rows.Scan(&id, &name, &description, &abbreviation, &createdAt, &updatedAt, &orderBy)
		field := dto.SubfieldDto{
			Id:           id,
			Name:         name,
			Description:  description,
			Abbreviation: abbreviation,
			CreatedAt:    createdAt,
			UpdatedAt:    updatedAt,
			Order:        orderBy,
		}

		subfields = append(subfields, field)
	}

	return subfields
}

func (f *FieldCollectionDao) FindById(id int64) domain.FieldCollection {
	var nameField, nameCollection, description, abbreviation string
	var unique, required, editable, hasSubfields bool
	var fieldCollection domain.FieldCollection
	var collection domain.Collection
	var idCollection, idField int64
	var field domain.IField
	var idSequence, idList sql.NullInt64

	var strSQL bytes.Buffer
	strSQL.WriteString("SELECT ")
	strSQL.WriteString("c.id, c.name, f.id, f.name, f.description, f.abbreviation, fc.isUnique, fc.required, fc.editable, count(s.id) as hasSubfields, fc.id, fc.sequence_id, fc.list_id ")
	strSQL.WriteString("FROM fields f ")
	strSQL.WriteString("INNER JOIN fields_collection fc ON f.id = fc.field_id ")
	strSQL.WriteString("INNER JOIN collections c ON c.id = fc.collection_id ")
	strSQL.WriteString("LEFT JOIN subfields s ON s.field_id = f.id ")
	strSQL.WriteString("WHERE fc.id = ? ")
	strSQL.WriteString("GROUP BY c.id, c.name, f.id, f.name, f.description, f.abbreviation, fc.isUnique, fc.required, fc.editable")

	stmt, err := f.db.Source().Conn().Prepare(strSQL.String())
	if err != nil {
		log.Println("abix-admin / dao / FieldCollectionDao / FindById / Conn().Prepare: ", err)
	}

	row := stmt.QueryRow(id)
	id = 0

	row.Scan(&idCollection, &nameCollection, &idField, &nameField, &description, &abbreviation, &unique, &required, &editable, &hasSubfields, &id, &idSequence, &idList)

	collection = *domain.NewCollection(nameCollection).WithId(idCollection)
	field = domain.NewSingleField(nameField)
	if hasSubfields {
		field = domain.NewCompositeField(nameCollection, f.getSubfields(idField))
	}

	field.WithId(idField).WithAbbreviation(abbreviation).WithDescription(description)

	fieldCollection.WithUnique(unique).WithEditable(editable).WithRequired(required).WithCollection(collection).WithField(field).WithId(id)
	if idSequence.Valid {
		sequence := domain.Sequence{}
		sequence.WithId(idSequence.Int64)
		fieldCollection.WithSequence(sequence)
	}

	if idList.Valid {
		list := domain.List{}
		list.WithId(idList.Int64)
		fieldCollection.WithList(list)
	}

	return fieldCollection
}
