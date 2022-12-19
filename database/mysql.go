package database

import (
	"database/sql"
	"log"
)

type MySQL struct {
	conn *sql.DB
}

func NewMySQL() *MySQL {

	strConnect := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + database
	db, err := sql.Open("mysql", strConnect)
	if err != nil {
		log.Fatal("connectMysql / New / sql.Open: ", err)
	}

	return &MySQL{
		conn: db,
	}
}

func (m *MySQL) Close() {
	if m.conn != nil {
		m.conn.Close()
	}
}

func (m *MySQL) Conn() *sql.DB {
	return m.conn
}
