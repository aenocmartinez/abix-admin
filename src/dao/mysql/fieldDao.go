package mysql

import (
	"abix360/database"
	"abix360/src/domain"
	"abix360/src/view/dto"
	"bytes"
	"errors"
	"log"
	"strings"
)

type FieldDao struct {
	db *database.ConnectDB
}

func NewFieldDao() *FieldDao {
	return &FieldDao{
		db: database.Instance(),
	}
}

func (f *FieldDao) Create(field domain.IField) error {
	var strQuery bytes.Buffer
	strQuery.WriteString("INSERT INTO fields(name, description, abbreviation) VALUES (?, ?, ?)")

	stmt, err := f.db.Source().Conn().Prepare(strQuery.String())
	if err != nil {
		log.Println("abix-admin / FieldDao / Create / conn.Prepare: ", err.Error())
		return err
	}

	res, err := stmt.Exec(field.Name(), field.Description(), field.Abbreviation())
	if err != nil {
		log.Println("abix-admin / FieldDao / Create / stmt.Exec: ", err.Error())
		return err
	}

	idField, err := res.LastInsertId()
	if err != nil {
		log.Println("abix-admin / FieldDao / Create / res.LastInsertId(): ", err.Error())
		return err
	}

	if field.IsComposite() {
		compositeField, ok := field.(*domain.CompositeField)
		if !ok {
			return errors.New("conversion a campo compuesto")
		}
		for _, subfieldvo := range compositeField.Subfield() {
			f.addSubfield(idField, subfieldvo.Field.Id(), subfieldvo.Order)
		}
	}

	return err
}

func (f *FieldDao) Delete(field domain.IField) error {
	var strQuery bytes.Buffer
	strQuery.WriteString("DELETE FROM fields WHERE id = ?")

	stmt, err := f.db.Source().Conn().Prepare(strQuery.String())
	if err != nil {
		log.Println("abix-admin / FieldDao / Delete / conn.Prepare: ", err.Error())
		return err
	}

	_, err = stmt.Exec(field.Id())
	if err != nil {
		log.Println("abix-admin / FieldDao / Delete / stmt.Exec: ", err.Error())
		return err
	}
	return err
}

func (f *FieldDao) Update(field domain.IField) error {
	var strQuery bytes.Buffer
	strQuery.WriteString("UPDATE fields SET name=?, description=?, abbreviation=?, updatedAt=NOW() WHERE id=?")

	stmt, err := f.db.Source().Conn().Prepare(strQuery.String())
	if err != nil {
		log.Println("abix-admin / FieldDao / Update / conn.Prepare: ", err.Error())
		return err
	}

	_, err = stmt.Exec(field.Name(), field.Description(), field.Abbreviation(), field.Id())
	if err != nil {
		log.Println("abix-admin / FieldDao / Update / stmt.Exec: ", err.Error())
		return err
	}

	if field.IsComposite() {
		compositeField, ok := field.(*domain.CompositeField)
		if !ok {
			return errors.New("conversion a campo compuesto")
		}
		f.deleteSubfields(field.Id())
		for _, subfieldvo := range compositeField.Subfield() {
			f.addSubfield(field.Id(), subfieldvo.Field.Id(), subfieldvo.Order)
			// f.updateOrderSubfield(field.Id(), subfieldvo.Field.Id(), subfieldvo.Order)
		}
	}

	return err
}

func (f *FieldDao) FindById(id int64) domain.IField {
	var field domain.IField
	var strSQL bytes.Buffer

	strSQL.WriteString("SELECT ")
	strSQL.WriteString("f.id, f.name, f.description, f.abbreviation, f.createdAt, f.updatedAt, IF(count(s.id) > 0, 1, 0) as hasSubfields ")
	strSQL.WriteString("FROM fields f ")
	strSQL.WriteString("LEFT JOIN subfields s ON f.id = s.field_id ")
	strSQL.WriteString("WHERE f.id = ? GROUP BY f.id")

	stmt, err := f.db.Source().Conn().Prepare(strSQL.String())
	if err != nil {
		log.Println("abix-admin / FieldDao / FindById / conn.Prepare: ", err.Error())
	}

	row := stmt.QueryRow(id)
	id = 0
	var name, description, abbreviation, createdAt, updatedAt string
	var hasSubfields bool

	row.Scan(&id, &name, &description, &abbreviation, &createdAt, &updatedAt, &hasSubfields)

	field = domain.NewSingleField(name)
	if hasSubfields {
		field = domain.NewCompositeField(name, f.getSubfields(id))
	}

	field.WithId(id).WithAbbreviation(abbreviation).WithDescription(description).WithCreatedAt(createdAt).WithUpdatedAt(updatedAt)

	return field
}

func (f *FieldDao) AllFields() []domain.IField {
	var fields []domain.IField
	var field domain.IField
	var strSQL bytes.Buffer

	strSQL.WriteString("SELECT ")
	strSQL.WriteString("f.id, f.name, f.description, f.abbreviation, f.createdAt, f.updatedAt, IF(count(s.id) > 0, 1, 0) as hasSubfields ")
	strSQL.WriteString("FROM fields f ")
	strSQL.WriteString("LEFT JOIN subfields s ON f.id = s.field_id ")
	strSQL.WriteString("GROUP BY f.id order by f.name")

	stmt, err := f.db.Source().Conn().Prepare(strSQL.String())
	if err != nil {
		log.Println("abix-admin / FieldDao / FindByName / conn.Prepare: ", err.Error())
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Println("abix-admin / FieldDao / AllFields / stmt.Query(): ", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var name, description, abbreviation, createdAt, updatedAt string
		var id int64
		var hasSubfields bool
		rows.Scan(&id, &name, &description, &abbreviation, &createdAt, &updatedAt, &hasSubfields)
		field = domain.NewSingleField(name)
		if hasSubfields {
			field = domain.NewCompositeField(name, f.getSubfields(id))
		}
		field.WithId(id).WithAbbreviation(abbreviation).WithDescription(description).WithCreatedAt(createdAt).WithUpdatedAt(updatedAt)
		fields = append(fields, field)
	}

	return fields
}

func (f *FieldDao) FindByName(name string, search bool) domain.IField {
	var field domain.IField
	var strSQL bytes.Buffer

	name = strings.TrimSpace(name)

	strSQL.WriteString("SELECT ")
	strSQL.WriteString("f.id, f.name, f.description, f.abbreviation, f.createdAt, f.updatedAt, IF(count(s.id) > 0, 1, 0) as hasSubfields ")
	strSQL.WriteString("FROM fields f ")
	strSQL.WriteString("LEFT JOIN subfields s ON f.id = s.field_id ")
	if !search {
		strSQL.WriteString("WHERE f.name = ? GROUP BY f.id")
	} else {
		strSQL.WriteString("WHERE f.name like %?% GROUP BY f.id")
	}

	stmt, err := f.db.Source().Conn().Prepare(strSQL.String())
	if err != nil {
		log.Println("abix-admin / FieldDao / FindByName / conn.Prepare: ", err.Error())
	}

	row := stmt.QueryRow(name)
	var description, abbreviation, createdAt, updatedAt string
	var id int64
	var hasSubfields bool

	row.Scan(&id, &name, &description, &abbreviation, &createdAt, &updatedAt, &hasSubfields)

	field = domain.NewSingleField(name)
	if hasSubfields {
		field = domain.NewCompositeField(name, f.getSubfields(id))
	}

	field.WithId(id).WithAbbreviation(abbreviation).WithDescription(description).WithCreatedAt(createdAt).WithUpdatedAt(updatedAt)

	return field
}

func (f *FieldDao) addSubfield(idField, idSubfield int64, order int) error {
	var strQuery bytes.Buffer
	strQuery.WriteString("INSERT INTO subfields(field_id, subfield_id, orderBy) VALUES (?, ?, ?)")

	stmt, err := f.db.Source().Conn().Prepare(strQuery.String())
	if err != nil {
		log.Println("abix-admin / FieldDao / addSubfield / conn.Prepare: ", err.Error())
	}

	_, err = stmt.Exec(idField, idSubfield, order)
	if err != nil {
		log.Println("abix-admin / FieldDao / addSubfield / stmt.Exec: ", err.Error())
	}

	return err
}

// func (f *FieldDao) updateOrderSubfield(idField, idSubfield int64, order int) error {
// 	var strQuery bytes.Buffer
// 	strQuery.WriteString("UPDATE subfields SET orderBy=?, updatedAt=NOW() WHERE field_id=? AND subfield_id=?")

// 	stmt, err := f.db.Source().Conn().Prepare(strQuery.String())
// 	if err != nil {
// 		log.Println("abix-admin / FieldDao / updateOrderSubfield / conn.Prepare: ", err.Error())
// 	}

// 	_, err = stmt.Exec(order, idField, idSubfield)
// 	if err != nil {
// 		log.Println("abix-admin / FieldDao / updateOrderSubfield / stmt.Exec: ", err.Error())
// 	}

// 	return err
// }

func (f *FieldDao) deleteSubfields(idField int64) error {
	var strQuery bytes.Buffer
	strQuery.WriteString("DELETE FROM subfields WHERE field_id=?")

	stmt, err := f.db.Source().Conn().Prepare(strQuery.String())
	if err != nil {
		log.Println("abix-admin / FieldDao / deleteSubfields / conn.Prepare: ", err.Error())
	}

	_, err = stmt.Exec(idField)
	if err != nil {
		log.Println("abix-admin / FieldDao / deleteSubfields / stmt.Exec: ", err.Error())
	}

	return err
}

func (f *FieldDao) getSubfields(idField int64) []dto.SubfieldDto {
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
