package usecase

import (
	"abix360/src/dao/mysql"
	"abix360/src/domain"
	"abix360/src/view/dto"
	"errors"
)

type UpdateCollectionUseCase struct {
}

func (useCase *UpdateCollectionUseCase) Execute(dtoCollection dto.CollectionDto) (code int, err error) {
	var repository domain.CollectionRepository = mysql.NewCollectionDao()
	collection := domain.FindCollectionById(dtoCollection.Id, repository)
	if !collection.Exists() {
		return 202, errors.New("la colección no existe")
	}

	collection.WithName(dtoCollection.Name).WithId(dtoCollection.Id).WithRepository(repository)

	err = collection.Update()
	if err != nil {
		return 500, err
	}

	return 200, nil
}
