package mysql

import (
	"abix360/database"
	"abix360/src/domain"
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
	return nil
}

func (l *ListDao) Update(list domain.List) error {
	return nil
}

func (l *ListDao) Delete(list domain.List) error {
	return nil
}

func (l *ListDao) FindById(id int64) (list domain.List, err error) {
	return list, nil
}

func (l *ListDao) FindByName(name string) (list domain.List, err error) {
	return list, err
}

func (l *ListDao) Search(name string) (lists []domain.List, err error) {
	return lists, err
}

func (l *ListDao) All() (lists []domain.List, err error) {
	return lists, err
}

func (l *ListDao) AddValue(idList int64, value string) error {
	return nil
}

func (l *ListDao) RemoveValue(idList int64, value string) error {
	return nil
}

func (l *ListDao) Values(idList int64) (values []string, err error) {
	return values, err
}
