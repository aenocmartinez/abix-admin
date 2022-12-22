package usecase

import (
	"abix360/src/dao/mysql"
	"abix360/src/domain"
	"errors"
)

type DeleteListUseCase struct{}

func (useCase *DeleteListUseCase) Execute(id int64) (code int, err error) {
	var repository domain.RepositoryList = mysql.NewListDao()
	list, err := domain.FindListById(id, repository)
	if err != nil {
		return 500, err
	}

	if !list.Exists() {
		return 202, errors.New("la lista no existe")
	}

	if err = list.Delete(); err != nil {
		return 500, err
	}

	return 200, nil
}
