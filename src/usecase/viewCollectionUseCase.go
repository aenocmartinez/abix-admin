package usecase

import (
	"abix360/src/dao/mysql"
	"abix360/src/domain"
	"abix360/src/view/dto"
	"errors"
)

type ViewCollectionUseCase struct {
}

func (useCase *ViewCollectionUseCase) Execute(id int64) (dtoCollection dto.CollectionDto, err error) {
	var repository domain.CollectionRepository = mysql.NewCollectionDao()
	collection := domain.FindCollectionById(id, repository)
	if !collection.Exists() {
		return dtoCollection, errors.New("la colección no existe")
	}

	dtoCollection.Id = collection.Id()
	dtoCollection.Name = collection.Name()
	dtoCollection.CreatedAt = collection.CreatedAt()
	dtoCollection.UpdatedAt = collection.UpdatedAt()

	return dtoCollection, nil
}
