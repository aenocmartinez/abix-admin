package mysql

import (
	"abix360/database"
	"abix360/src/domain"
	"bytes"
	"database/sql"
	"log"
	"strings"
)

type ListDao struct {
	db *database.ConnectDB
}

func NewListDao() *ListDao {
	return &ListDao{
		db: database.Instance(),
	}
}

func (l *ListDao) Create(list domain.List) error {
	var strSQL bytes.Buffer
	strSQL.WriteString("INSERT INTO lists(name, strValues) VALUES (?, ?)")

	stmt, err := l.db.Source().Conn().Prepare(strSQL.String())
	if err != nil {
		log.Println("abix-admin / dao / mysql / ListDao / Create / Conn().Prepare: ", err.Error())
	}

	_, err = stmt.Exec(list.Name(), l.getValuesFormatNullString(list))
	if err != nil {
		log.Println("abix-admin / dao / mysql / ListDao / Create / stmt.Exec: ", err.Error())
	}

	return err
}

func (l *ListDao) Update(list domain.List) error {
	var strSQL bytes.Buffer
	strSQL.WriteString("UPDATE lists SET name=?, strValues=?, updatedAt=NOW() WHERE id=?")

	stmt, err := l.db.Source().Conn().Prepare(strSQL.String())
	if err != nil {
		log.Println("abix-admin / dao / mysql / ListDao / Update / Conn().Prepare: ", err.Error())
	}

	_, err = stmt.Exec(list.Name(), l.getValuesFormatNullString(list), list.Id())
	if err != nil {
		log.Println("abix-admin / dao / mysql / ListDao / Update / stmt.Exec: ", err.Error())
	}

	return err
}

func (l *ListDao) Delete(list domain.List) error {
	var strSQL bytes.Buffer
	strSQL.WriteString("DELETE FROM lists WHERE id=?")

	stmt, err := l.db.Source().Conn().Prepare(strSQL.String())
	if err != nil {
		log.Println("abix-admin / dao / mysql / ListDao / Delete / Conn().Prepare: ", err.Error())
	}

	_, err = stmt.Exec(list.Id())
	if err != nil {
		log.Println("abix-admin / dao / mysql / ListDao / Delete / stmt.Exec: ", err.Error())
	}

	return err
}

func (l *ListDao) FindById(id int64) (list domain.List, err error) {
	var strSQL bytes.Buffer
	strSQL.WriteString("SELECT id, name, strValues, createdAt, updatedAt FROM lists WHERE id = ?")

	stmt, err := l.db.Source().Conn().Prepare(strSQL.String())
	if err != nil {
		log.Println("abix-admin / dao / mysql / ListDao / FindById / Conn().Prepare: ", err.Error())
	}

	row := stmt.QueryRow(id)
	var name, createdAt, updatedAt string
	var strValues []string
	var sqlValues sql.NullString
	id = 0

	row.Scan(&id, &name, &sqlValues, &createdAt, &updatedAt)
	if sqlValues.Valid {
		strValues = strings.Split(sqlValues.String, ",")
	}

	list = *domain.NewList(name).WithCreatedAt(createdAt).WithUpdatedAt(updatedAt).WithValues(strValues).WithId(id)

	return list, nil
}

func (l *ListDao) FindByName(name string) (list domain.List, err error) {
	var strSQL bytes.Buffer
	strSQL.WriteString("SELECT id, name, strValues, createdAt, updatedAt FROM lists WHERE name = ?")

	stmt, err := l.db.Source().Conn().Prepare(strSQL.String())
	if err != nil {
		log.Println("abix-admin / dao / mysql / ListDao / FindByName / Conn().Prepare: ", err.Error())
	}

	row := stmt.QueryRow(name)
	var createdAt, updatedAt string
	var strValues []string
	var sqlValues sql.NullString
	var id int64

	row.Scan(&id, &name, &sqlValues, &createdAt, &updatedAt)
	if sqlValues.Valid {
		strValues = strings.Split(sqlValues.String, ",")
	}

	list = *domain.NewList(name).WithCreatedAt(createdAt).WithUpdatedAt(updatedAt).WithValues(strValues).WithId(id)

	return list, nil
}

func (l *ListDao) Search(name string) (lists []domain.List, err error) {
	var strSQL bytes.Buffer
	strSQL.WriteString("SELECT id, name, strValues, createdAt, updatedAt FROM lists WHERE name like ?")

	stmt, err := l.db.Source().Conn().Prepare(strSQL.String())
	if err != nil {
		log.Println("abix-admin / dao / mysql / ListDao / Search / Conn().Prepare: ", err.Error())
	}

	rows, err := stmt.Query("%" + name + "%")
	if err != nil {
		log.Println("abix-admin / dao / mysql / ListDao / Search / stmt.Query: ", err.Error())
	}

	for rows.Next() {
		var createdAt, updatedAt string
		var strValues []string
		var sqlValues sql.NullString
		var id int64

		rows.Scan(&id, &name, &sqlValues, &createdAt, &updatedAt)
		if sqlValues.Valid {
			strValues = strings.Split(sqlValues.String, ",")
		}

		list := *domain.NewList(name).WithCreatedAt(createdAt).WithUpdatedAt(updatedAt).WithValues(strValues).WithId(id)
		lists = append(lists, list)

	}

	return lists, nil
}

func (l *ListDao) All() (lists []domain.List, err error) {
	var strSQL bytes.Buffer
	strSQL.WriteString("SELECT id, name, strValues, createdAt, updatedAt FROM lists order by name")

	stmt, err := l.db.Source().Conn().Prepare(strSQL.String())
	if err != nil {
		log.Println("abix-admin / dao / mysql / ListDao / All / Conn().Prepare: ", err.Error())
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Println("abix-admin / dao / mysql / ListDao / All / stmt.Query: ", err.Error())
	}

	for rows.Next() {
		var name, createdAt, updatedAt string
		var strValues []string
		var sqlValues sql.NullString
		var id int64
		rows.Scan(&id, &name, &sqlValues, &createdAt, &updatedAt)
		if sqlValues.Valid {
			strValues = strings.Split(sqlValues.String, ",")
		}
		list := *domain.NewList(name).WithCreatedAt(createdAt).WithUpdatedAt(updatedAt).WithValues(strValues).WithId(id)
		lists = append(lists, list)
	}

	return lists, nil
}

func (l *ListDao) getValuesFormatNullString(list domain.List) sql.NullString {
	var sqlValues sql.NullString

	if len(list.Values()) > 0 {
		var strValues string
		for _, value := range list.Values() {
			strValues = strValues + "," + value
		}
		sqlValues = sql.NullString{String: strValues, Valid: true}
	}
	return sqlValues
}
