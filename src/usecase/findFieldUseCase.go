package usecase

import (
	"abix360/src/dao/mysql"
	"abix360/src/domain"
	"abix360/src/view/dto"
	"errors"
)

type FindFieldUseCase struct{}

func (useCase *FindFieldUseCase) Execute(id int64) (dtoField dto.FieldDto, err error) {
	var field domain.IField
	var repository domain.FieldRepository = mysql.NewFieldDao()

	field = domain.FindFieldById(id, repository)
	if !field.Exists() {
		return dtoField, errors.New("el campo no existe")
	}

	dtoField.Name = field.Name()
	dtoField.Abbreviation = field.Abbreviation()
	dtoField.Description = field.Description()
	dtoField.CreatedAt = field.CreatedAt()
	dtoField.UpdatedAt = field.UpdatedAt()
	dtoField.Id = field.Id()

	if field.IsComposite() {
		compositeField, ok := field.(*domain.CompositeField)
		if ok {

			for _, subfield := range compositeField.Subfield() {
				dtoField.Subfields = append(dtoField.Subfields, dto.SubfieldDto{
					Name:         subfield.Field.Name(),
					Description:  subfield.Field.Description(),
					Abbreviation: subfield.Field.Abbreviation(),
					Id:           subfield.Field.Id(),
					Order:        subfield.Order,
				})
			}
		}
	}

	return dtoField, nil
}
