package usecase

import (
	"abix360/src/dao/mysql"
	"abix360/src/domain"
	"abix360/src/view/dto"
)

type ListCollectionsUseCase struct {
}

func (useCase *ListCollectionsUseCase) Execute() []dto.CollectionDto {
	dtoCollections := []dto.CollectionDto{}
	var repository domain.CollectionRepository = mysql.NewCollectionDao()

	collections := domain.AllCollections(repository)

	for _, collection := range collections {
		dtoCollections = append(dtoCollections, dto.CollectionDto{
			Id:        collection.Id(),
			Name:      collection.Name(),
			CreatedAt: collection.CreatedAt(),
			UpdatedAt: collection.UpdatedAt(),
		})
	}

	return dtoCollections
}
