package usecase

import (
	"abix360/src/dao/mysql"
	"abix360/src/domain"
	"abix360/src/view/dto"
	"errors"
	"fmt"
)

type UpdateFieldUseCase struct {
}

func (useCase *UpdateFieldUseCase) Execute(dtoField dto.FieldDto) (code int, err error) {
	var field domain.IField
	var repository domain.FieldRepository = mysql.NewFieldDao()

	field = domain.FindFieldById(dtoField.Id, repository)
	if !field.Exists() {
		return 202, errors.New("el campo no existe")
	}

	fmt.Println("UseCase / dtoField.Subfields: ", dtoField.Subfields)

	if len(dtoField.Subfields) > 0 {
		field = domain.NewCompositeField(dtoField.Name, dtoField.Subfields)
	}

	field.WithId(dtoField.Id)
	field.WithName(dtoField.Name)
	field.WithAbbreviation(dtoField.Abbreviation)
	field.WithDescription(dtoField.Description)
	field.WithRepository(repository)

	err = field.Update()
	if err != nil {
		return 500, err
	}

	return 200, nil
}
