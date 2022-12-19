package usecase

import (
	"abix360/src/dao/mysql"
	"abix360/src/domain"
	"errors"
)

type DeleteCollectionUseCase struct {
}

func (useCase *DeleteCollectionUseCase) Execute(id int64) (code int, err error) {
	var repository domain.CollectionRepository = mysql.NewCollectionDao()
	collection := domain.FindCollectionById(id, repository)
	if !collection.Exists() {
		return 202, errors.New("la colección no existe")
	}

	collection.WithId(id).WithRepository(repository)
	err = collection.Delete()
	if err != nil {
		return 500, err
	}

	return 200, nil
}
