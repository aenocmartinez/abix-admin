package usecase

import (
	"abix360/src/dao/mysql"
	"abix360/src/domain"
	"abix360/src/view/dto"
	"errors"
)

type CreateFieldUseCase struct {
}

func (useCase *CreateFieldUseCase) Execute(dtoField dto.FieldDto) (code int, err error) {
	var field domain.IField
	var repository domain.FieldRepository = mysql.NewFieldDao()

	field = domain.FindFieldByName(dtoField.Name, false, repository)
	if field.Exists() {
		return 202, errors.New("el campo ya existe")
	}

	field = domain.NewSingleField(dtoField.Name)
	if len(dtoField.Subfields) > 0 {
		field = domain.NewCompositeField(dtoField.Name, dtoField.Subfields)
	}

	field.WithDescription(dtoField.Description)
	field.WithAbbreviation(dtoField.Abbreviation)
	field.WithRepository(repository)

	err = field.Create()
	if err != nil {
		return 500, err
	}

	return 200, nil
}
