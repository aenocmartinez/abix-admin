package database

import (
	"database/sql"
	"log"
)

type PostgreSQL struct {
	conn *sql.DB
}

func NewPostgreSQL() *MySQL {
	strConnect := "host=" + host + " port=" + port + " user=" + user + "password=" + pass + " dbname=" + database + " sslmode=disable"
	db, err := sql.Open("postgres", strConnect)
	if err != nil {
		log.Fatal("connectMysql / New / sql.Open: ", err)
	}

	return &MySQL{
		conn: db,
	}
}

func (m *PostgreSQL) Close() {
	if m.conn != nil {
		m.conn.Close()
	}
}

func (m *PostgreSQL) Conn() *sql.DB {
	return m.conn
}
