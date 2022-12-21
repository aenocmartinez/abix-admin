package domain

import (
	"abix360/shared"
	"log"
)

type Sequence struct {
	id         int64
	name       string
	prefix     string
	value      int64
	createAt   string
	updatedAt  string
	repository RepositorySequence
}

func NewSequence(name string) *Sequence {
	return &Sequence{
		name:  name,
		value: 1,
	}
}

func (s *Sequence) WithRepository(repository RepositorySequence) *Sequence {
	s.repository = repository
	return s
}

func (s *Sequence) WithCreatedAt(createdAt string) *Sequence {
	s.createAt = createdAt
	return s
}

func (s *Sequence) WithId(id int64) *Sequence {
	s.id = id
	return s
}

func (s *Sequence) WithName(name string) *Sequence {
	s.name = name
	return s
}

func (s *Sequence) WithPrefix(prefix string) *Sequence {
	s.prefix = prefix
	return s
}

func (s *Sequence) WithValue(value int64) *Sequence {
	if value > 1 {
		s.value = value
	}
	return s
}

func (s *Sequence) Id() int64 {
	return s.id
}

func (s *Sequence) Name() string {
	return s.name
}

func (s *Sequence) Prefix() string {
	return s.prefix
}

func (s *Sequence) Value() int64 {
	return s.value
}

func (s *Sequence) CurrentValue() string {
	currentValue := shared.CompleteWithZero(s.value)
	if s.prefix != "" {
		currentValue = s.prefix + currentValue
	}
	return currentValue
}

func (s *Sequence) Create() error {
	return s.repository.Create(*s)
}

func (s *Sequence) Delete() error {
	return s.repository.Delete(*s)
}

func (s *Sequence) Update() error {
	return s.repository.Update(*s)
}

func (s *Sequence) NextValue() error {
	return s.repository.NextValue(s.id)
}

func (s *Sequence) CreatedAt() string {
	return s.createAt
}

func (s *Sequence) UpdatedAt() string {
	return s.updatedAt
}

func (s *Sequence) WithUpdatedAt(updatedAt string) *Sequence {
	s.updatedAt = updatedAt
	return s
}

func (s *Sequence) Exists() bool {
	return s.id > 0
}

func FindSequenceById(id int64, repository RepositorySequence) Sequence {
	sequence, err := repository.FindById(id)
	if err != nil {
		log.Println("abix-admin / domain / sequence / FindSequenceById: ", err)
	}
	return sequence
}

func FindSequenceByName(name string, repository RepositorySequence) Sequence {
	sequence, err := repository.FindByName(name)
	if err != nil {
		log.Println("abix-admin / domain / sequence / FindSequenceByName: ", err)
	}
	return sequence
}

func SearchSequenceByName(name string, repository RepositorySequence) ([]Sequence, error) {
	sequences, err := repository.Search(name)
	if err != nil {
		log.Println("abix-admin / domain / sequence / FindSequenceByName: ", err)
	}
	return sequences, err
}

func AllSequences(repository RepositorySequence) ([]Sequence, error) {
	sequences, err := repository.AllSequences()
	if err != nil {
		log.Println("abix-admin / domain / sequence / AllSequences: ", err)
	}
	return sequences, err
}
