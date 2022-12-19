package usecase

import (
	"abix360/src/dao/mysql"
	"abix360/src/domain"
	"errors"
)

type CreateCollectionUseCase struct{}

func (usecase *CreateCollectionUseCase) Execute(name string) (code int, err error) {
	var repository domain.CollectionRepository = mysql.NewCollectionDao()

	collection := domain.FindCollectionByName(name, repository)
	if collection.Exists() {
		return 202, errors.New("la colección ya existe")
	}

	collection.WithName(name).WithRepository(repository)
	err = collection.Create()
	if err != nil {
		return 500, err
	}

	return 200, nil
}
