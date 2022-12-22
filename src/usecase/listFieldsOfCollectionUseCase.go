package usecase

import (
	"abix360/src/dao/mysql"
	"abix360/src/domain"
	"abix360/src/view/dto"
	"errors"
)

type ListFieldsOfCollectionUseCase struct{}

func (useCase *ListFieldsOfCollectionUseCase) Execute(idCollection int64) (fields []dto.FieldCollectionDto, err error) {
	var repositoryFieldCollection domain.FieldCollectionRepository = mysql.NewFieldCollectionDao()
	var repositoryCollection domain.CollectionRepository = mysql.NewCollectionDao()

	collection := domain.FindCollectionById(idCollection, repositoryCollection)
	if !collection.Exists() {
		return fields, errors.New("la colección no existe")
	}

	collection.WithRepositoryFieldCollection(repositoryFieldCollection)
	for _, fieldCollection := range collection.AllFields() {
		fields = append(fields, dto.FieldCollectionDto{
			Id:           fieldCollection.Id(),
			IdCollection: fieldCollection.IdCollection(),
			IdField:      fieldCollection.IdField(),
			Unique:       fieldCollection.Unique(),
			Editable:     fieldCollection.Editable(),
			Required:     fieldCollection.Required(),
			Field:        fieldCollection.Field().Name(),
			Collection:   fieldCollection.Collection().Name(),
			Sequence:     fieldCollection.Sequence().Id(),
			NameSequence: fieldCollection.Sequence().Name(),
			List:         fieldCollection.List().Id(),
			NameList:     fieldCollection.List().Name(),
		})
	}

	return fields, nil
}
