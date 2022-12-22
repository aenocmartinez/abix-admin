package usecase

import (
	"abix360/src/dao/mysql"
	"abix360/src/domain"
	"abix360/src/view/dto"
	"errors"
)

type ViewFieldCollectionUseCase struct{}

func (useCase *ViewFieldCollectionUseCase) Execute(idFieldCollection int64) (dtoField dto.FieldCollectionDto, err error) {
	var repository domain.FieldCollectionRepository = mysql.NewFieldCollectionDao()
	fieldCollection := domain.FindByIdFieldCollection(idFieldCollection, repository)

	if !fieldCollection.Exists() {
		return dtoField, errors.New("el campo no se encuentra asignado a esta colección")
	}

	dtoField.Id = fieldCollection.Id()
	dtoField.Collection = fieldCollection.Collection().Name()
	dtoField.IdCollection = fieldCollection.Collection().Id()
	dtoField.Field = fieldCollection.Field().Name()
	dtoField.IdField = fieldCollection.Field().Id()
	dtoField.Unique = fieldCollection.Unique()
	dtoField.Editable = fieldCollection.Editable()
	dtoField.Required = fieldCollection.Required()
	dtoField.Sequence = fieldCollection.Sequence().Id()

	return dtoField, nil
}
