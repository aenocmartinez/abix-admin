package usecase

import (
	"abix360/src/dao/mysql"
	"abix360/src/domain"
	"errors"
)

type DeleteFieldUseCase struct{}

func (useCase *DeleteFieldUseCase) Execute(id int64) (code int, err error) {
	var field domain.IField
	var repository domain.FieldRepository = mysql.NewFieldDao()

	field = domain.FindFieldById(id, repository)
	if !field.Exists() {
		return 202, errors.New("el campo no existe")
	}

	field.WithRepository(repository)

	err = field.Delete()
	if err != nil {
		return 500, err
	}

	return 200, nil
}
