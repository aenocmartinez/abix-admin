package usecase

import (
	"abix360/src/dao/mysql"
	"abix360/src/domain"
	"abix360/src/view/dto"
)

type ListFieldsUseCase struct{}

func (useCase *ListFieldsUseCase) Execute() []dto.FieldDto {

	var dtoFields []dto.FieldDto

	var repository domain.FieldRepository = mysql.NewFieldDao()
	fields := domain.AllFields(repository)

	for _, field := range fields {
		dtoField := dto.FieldDto{
			Id:          field.Id(),
			Name:        field.Name(),
			Description: field.Description(),
			CreatedAt:   field.CreatedAt(),
			UpdatedAt:   field.UpdatedAt(),
		}
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
		dtoFields = append(dtoFields, dtoField)
	}

	return dtoFields
}
