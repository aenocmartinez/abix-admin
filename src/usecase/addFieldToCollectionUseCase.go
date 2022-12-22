package usecase

import (
	"abix360/src/dao/mysql"
	"abix360/src/domain"
	"abix360/src/view/dto"
	"errors"
)

type AddFieldToCollectionUseCase struct {
}

func (useCase *AddFieldToCollectionUseCase) Execute(dtoFieldCollection dto.FieldCollectionDto) (code int, err error) {

	var repositoryFieldCollection domain.FieldCollectionRepository = mysql.NewFieldCollectionDao()
	var repositoryCollection domain.CollectionRepository = mysql.NewCollectionDao()
	var repositoryField domain.FieldRepository = mysql.NewFieldDao()
	var repositorySequence domain.RepositorySequence = mysql.NewSequenceDao()

	collection := domain.FindCollectionById(dtoFieldCollection.IdCollection, repositoryCollection)
	if !collection.Exists() {
		return 202, errors.New("la colección no existe")
	}

	field := domain.FindFieldById(dtoFieldCollection.IdField, repositoryField)
	if !field.Exists() {
		return 202, errors.New("el campo no existe")
	}

	fieldCollection := domain.NewFieldCollection(collection, field)
	fieldCollection.WithEditable(dtoFieldCollection.Editable)
	fieldCollection.WithRequired(dtoFieldCollection.Required)
	fieldCollection.WithUnique(dtoFieldCollection.Unique)

	sequence := domain.FindSequenceById(dtoFieldCollection.Sequence, repositorySequence)
	if sequence.Exists() {
		fieldCollection.WithSequence(sequence)
	}

	collection.WithRepositoryFieldCollection(repositoryFieldCollection)
	err = collection.AddField(*fieldCollection)
	if err != nil {
		return 500, err
	}

	return 200, nil

}
