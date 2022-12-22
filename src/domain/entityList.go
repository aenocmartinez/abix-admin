package domain

import "log"

type List struct {
	id         int64
	name       string
	values     []string
	repository RepositoryList
}

func NewList(name string) *List {
	return &List{
		name:   name,
		values: []string{},
	}
}

func (l *List) WithId(id int64) *List {
	l.id = id
	return l
}

func (l *List) WithName(name string) *List {
	l.name = name
	return l
}

func (l *List) WithValues(values []string) *List {
	l.values = values
	return l
}

func (l *List) WithRepository(repository RepositoryList) *List {
	l.repository = repository
	return l
}

func (l *List) Id() int64 {
	return l.id
}

func (l *List) Name() string {
	return l.name
}

func (l *List) Create() error {
	return l.repository.Create(*l)
}

func (l *List) Update() error {
	return l.repository.Update(*l)
}

func (l *List) Delete() error {
	return l.repository.Delete(*l)
}

func (l *List) Exists() bool {
	return l.id > 0
}

func (l *List) AddValue(value string) error {
	return l.repository.AddValue(l.id, value)
}

func (l *List) RemoveValue(value string) error {
	return l.repository.RemoveValue(l.id, value)
}

func (l *List) Values() []string {
	var err error
	if len(l.values) > 0 {
		return l.values
	}

	l.values, err = l.repository.Values(l.id)
	if err != nil {
		log.Println("abix-admin / domain / List / values(): ", err)
	}
	return l.values
}

func FindListById(id int64, repository RepositoryList) (List, error) {
	return repository.FindById(id)
}

func FindListByName(name string, repository RepositoryList) (List, error) {
	return repository.FindByName(name)
}

func SearchList(name string, repository RepositoryList) ([]List, error) {
	return repository.Search(name)
}

func AllList(repository RepositoryList) ([]List, error) {
	return repository.All()
}
