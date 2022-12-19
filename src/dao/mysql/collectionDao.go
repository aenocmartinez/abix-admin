package mysql

import (
	"abix360/database"
	"abix360/src/domain"
	"bytes"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type CollectionDao struct {
	db *database.ConnectDB
}

func NewCollectionDao() *CollectionDao {
	return &CollectionDao{
		db: database.Instance(),
	}
}

func (c *CollectionDao) Create(collection domain.Collection) error {
	var strQuery bytes.Buffer
	strQuery.WriteString("INSERT INTO collections(name) VALUES (?)")

	stmt, err := c.db.Source().Conn().Prepare(strQuery.String())
	if err != nil {
		log.Println("abix-admin / CollectionDao / Create / conn.Prepare: ", err.Error())
	}

	_, err = stmt.Exec(collection.Name())
	if err != nil {
		log.Println("abix-admin / CollectionDao / Create / stmt.Exec: ", err.Error())
	}

	return err
}

func (c *CollectionDao) Update(collection domain.Collection) error {
	var strQuery bytes.Buffer
	strQuery.WriteString("UPDATE collections SET name=?, updated_at=NOW() WHERE id=?")

	stmt, err := c.db.Source().Conn().Prepare(strQuery.String())
	if err != nil {
		log.Println("abix-admin / CollectionDao / Update / conn.Prepare: ", err.Error())
	}

	_, err = stmt.Exec(collection.Name(), collection.Id())
	if err != nil {
		log.Println("abix-admin / CollectionDao / Update / stmt.Exec: ", err.Error())
	}

	return err
}

func (c *CollectionDao) Delete(id int64) error {
	var strQuery bytes.Buffer
	strQuery.WriteString("DELETE FROM collections WHERE id=?")

	stmt, err := c.db.Source().Conn().Prepare(strQuery.String())
	if err != nil {
		log.Println("abix-admin / CollectionDao / Delete / conn.Prepare: ", err.Error())
	}

	_, err = stmt.Exec(id)
	if err != nil {
		log.Println("abix-admin / CollectionDao / Delete / stmt.Exec: ", err.Error())
	}

	return err
}

func (c *CollectionDao) FindById(id int64) domain.Collection {
	var collection domain.Collection
	var cad bytes.Buffer

	cad.WriteString("SELECT id, name, created_at, updated_at FROM collections WHERE id = ?")
	row := c.db.Source().Conn().QueryRow(cad.String(), id)

	var name, createdAt, updatedAt string

	row.Scan(&id, &name, &createdAt, &updatedAt)
	collection = *domain.NewCollection(name)
	collection.WithId(id).WithCreatedAt(createdAt).WithUpdatedAt(updatedAt)

	return collection
}

func (c *CollectionDao) FindByName(name string) domain.Collection {
	var collection domain.Collection
	var createdAt, updatedAt string
	var id int64
	var cad bytes.Buffer

	cad.WriteString("SELECT id, name, created_at, updated_at FROM collections WHERE name = ?")
	row := c.db.Source().Conn().QueryRow(cad.String(), name)

	row.Scan(&id, &name, &createdAt, &updatedAt)
	collection = *domain.NewCollection(name)
	collection.WithId(id).WithCreatedAt(createdAt).WithUpdatedAt(updatedAt)

	return collection
}

func (c *CollectionDao) AllCollections() []domain.Collection {
	var collections []domain.Collection
	var strQuery bytes.Buffer

	strQuery.WriteString("SELECT id, name, created_at, updated_at FROM collections order by name")
	rows, err := c.db.Source().Conn().Query(strQuery.String())
	if err != nil {
		log.Println("abix-admin / CollectionDao / AllCollections / c.db.Source().Conn().Query: ", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var name, updatedAt, createdAt string
		var id int64
		rows.Scan(&id, &name, &createdAt, &updatedAt)
		collection := domain.NewCollection(name)
		collection.WithId(id).WithUpdatedAt(updatedAt).WithCreatedAt(createdAt)
		collections = append(collections, *collection)
	}

	return collections
}
