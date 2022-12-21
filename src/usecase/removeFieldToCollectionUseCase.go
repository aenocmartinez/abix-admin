package usecase

import (
	"abix360/src/dao/mysql"
	"abix360/src/domain"
	"errors"
)

type RemoveFieldToCollectionUseCase struct{}

func (useCase *RemoveFieldToCollectionUseCase) Execute(idCollection, idField int64) (code int, err error) {

	var repositoryFieldCollection domain.FieldCollectionRepository = mysql.NewFieldCollectionDao()
	var repositoryCollection domain.CollectionRepository = mysql.NewCollectionDao()
	var repositoryField domain.FieldRepository = mysql.NewFieldDao()

	collection := domain.FindCollectionById(idCollection, repositoryCollection)
	if !collection.Exists() {
		return 202, errors.New("la colección no existe")
	}

	field := domain.FindFieldById(idField, repositoryField)
	if !field.Exists() {
		return 202, errors.New("el campo no existe")
	}

	fieldCollection := domain.NewFieldCollection(collection, field)
	collection.WithRepositoryFieldCollection(repositoryFieldCollection)

	err = collection.RemoveField(*fieldCollection)
	if err != nil {
		return 500, err
	}

	return 200, nil
}
