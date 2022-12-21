package mysql

import (
	"abix360/database"
	"abix360/src/domain"
	"bytes"
	"log"
)

type SequenceDao struct {
	db *database.ConnectDB
}

func NewSequenceDao() *SequenceDao {
	return &SequenceDao{
		db: database.Instance(),
	}
}

func (s *SequenceDao) Create(sequence domain.Sequence) error {
	var strSQL bytes.Buffer
	strSQL.WriteString("INSERT INTO sequences(name, prefix, value) VALUES (?, ?, ?)")

	stmt, err := s.db.Source().Conn().Prepare(strSQL.String())
	if err != nil {
		log.Println("abix-admin / dao / sequenceDao / Create / Conn().Prepared(): ", err)
	}

	_, err = stmt.Exec(sequence.Name(), sequence.Prefix(), sequence.Value())
	if err != nil {
		log.Println("abix-admin / dao / sequenceDao / Create / stmt.Exec: ", err)
	}

	return err
}

func (s *SequenceDao) Update(sequence domain.Sequence) error {
	var strSQL bytes.Buffer
	strSQL.WriteString("UPDATE sequences SET name=?, prefix=?, value=?, updatedAt=NOW() WHERE id=?")

	stmt, err := s.db.Source().Conn().Prepare(strSQL.String())
	if err != nil {
		log.Println("abix-admin / dao / sequenceDao / Update / Conn().Prepared(): ", err)
	}

	_, err = stmt.Exec(sequence.Name(), sequence.Prefix(), sequence.Value(), sequence.Id())
	if err != nil {
		log.Println("abix-admin / dao / sequenceDao / Update / stmt.Exec: ", err)
	}

	return err
}

func (s *SequenceDao) Delete(sequence domain.Sequence) error {
	var strSQL bytes.Buffer
	strSQL.WriteString("DELETE FROM sequences WHERE id=?")

	stmt, err := s.db.Source().Conn().Prepare(strSQL.String())
	if err != nil {
		log.Println("abix-admin / dao / sequenceDao / Delete / Conn().Prepared(): ", err)
	}

	_, err = stmt.Exec(sequence.Id())
	if err != nil {
		log.Println("abix-admin / dao / sequenceDao / Delete / stmt.Exec: ", err)
	}

	return err
}

func (s *SequenceDao) FindById(id int64) (sequence domain.Sequence, err error) {
	var name, prefix, createdAt, updatedAt string
	var strSQL bytes.Buffer
	var value int64

	strSQL.WriteString("SELECT id, name, prefix, value, createdAt, updatedAt FROM sequences WHERE id = ?")

	stmt, err := s.db.Source().Conn().Prepare(strSQL.String())
	if err != nil {
		log.Println("abix-admin / dao / sequenceDao / FindById / Conn().Prepared(): ", err)
	}

	row := stmt.QueryRow(id)
	id = 0

	row.Scan(&id, &name, &prefix, &value, &createdAt, &updatedAt)

	sequence = *domain.NewSequence(name)
	sequence.WithId(id).WithPrefix(prefix).WithValue(value).WithCreatedAt(createdAt).WithUpdatedAt(updatedAt)

	return sequence, nil
}

func (s *SequenceDao) FindByName(name string) (sequence domain.Sequence, err error) {
	var prefix, createdAt, updatedAt string
	var strSQL bytes.Buffer
	var id, value int64

	strSQL.WriteString("SELECT id, name, prefix, value, createdAt, updatedAt FROM sequences WHERE name = ?")

	stmt, err := s.db.Source().Conn().Prepare(strSQL.String())
	if err != nil {
		log.Println("abix-admin / dao / sequenceDao / FindByName / Conn().Prepared(): ", err)
	}

	row := stmt.QueryRow(name)
	if err != nil {
		log.Println("abix-admin / dao / sequenceDao / FindByName / stmt.Query: ", err)
	}
	row.Scan(&id, &name, &prefix, &value, &createdAt, &updatedAt)

	sequence = *domain.NewSequence(name)
	sequence.WithId(id).WithPrefix(prefix).WithValue(value).WithCreatedAt(createdAt).WithUpdatedAt(updatedAt)

	return sequence, nil
}

func (s *SequenceDao) Search(name string) (sequences []domain.Sequence, err error) {
	var prefix, createdAt, updatedAt string
	var strSQL bytes.Buffer
	var id, value int64

	strSQL.WriteString("SELECT id, name, prefix, value, createdAt, updatedAt FROM sequences WHERE name like ? order by name")

	stmt, err := s.db.Source().Conn().Prepare(strSQL.String())
	if err != nil {
		log.Println("abix-admin / dao / sequenceDao / Search / Conn().Prepared(): ", err)
	}

	rows, err := stmt.Query("%" + name + "%")
	if err != nil {
		log.Println("abix-admin / dao / sequenceDao / Search / stmt.Query: ", err)
	}

	for rows.Next() {
		rows.Scan(&id, &name, &prefix, &value, &createdAt, &updatedAt)
		sequence := *domain.NewSequence(name)
		sequence.WithId(id).WithPrefix(prefix).WithValue(value).WithCreatedAt(createdAt).WithUpdatedAt(updatedAt)
		sequences = append(sequences, sequence)
	}

	return sequences, nil
}

func (s *SequenceDao) AllSequences() (sequences []domain.Sequence, err error) {
	var name, prefix, createdAt, updatedAt string
	var strSQL bytes.Buffer
	var id, value int64

	strSQL.WriteString("SELECT id, name, prefix, value, createdAt, updatedAt FROM sequences order by name")

	stmt, err := s.db.Source().Conn().Prepare(strSQL.String())
	if err != nil {
		log.Println("abix-admin / dao / sequenceDao / AllSequences / Conn().Prepared(): ", err)
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Println("abix-admin / dao / sequenceDao / AllSequences / stmt.Query: ", err)
	}

	for rows.Next() {
		rows.Scan(&id, &name, &prefix, &value, &createdAt, &updatedAt)
		sequence := *domain.NewSequence(name)
		sequence.WithId(id).WithPrefix(prefix).WithValue(value).WithCreatedAt(createdAt).WithUpdatedAt(updatedAt)
		sequences = append(sequences, sequence)
	}

	return sequences, nil
}

func (s *SequenceDao) NextValue(idSequence int64) error {
	var strSQL bytes.Buffer
	strSQL.WriteString("UPDATE sequences SET value value+1 WHERE id=?")

	stmt, err := s.db.Source().Conn().Prepare(strSQL.String())
	if err != nil {
		log.Println("abix-admin / dao / sequenceDao / NextValue / Conn().Prepared(): ", err)
	}

	_, err = stmt.Exec(idSequence)
	if err != nil {
		log.Println("abix-admin / dao / sequenceDao / NextValue / stmt.Exec: ", err)
	}

	return err
}
