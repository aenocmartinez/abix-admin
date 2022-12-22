package usecase

import (
	"abix360/src/dao/mysql"
	"abix360/src/domain"
	"abix360/src/view/dto"
	"errors"
)

type CreateListUseCase struct {
}

func (useCase *CreateListUseCase) Execute(dtoList dto.ListDto) (code int, err error) {
	var repository domain.RepositoryList = mysql.NewListDao()

	list, err := domain.FindListByName(dtoList.Name, repository)
	if err != nil {
		return 500, err
	}

	if list.Exists() {
		return 202, errors.New("la lista ya existe")
	}

	list = *domain.NewList(dtoList.Name).WithValues(dtoList.Values)
	list.WithRepository(repository)

	if err = list.Create(); err != nil {
		return 500, err
	}

	return 200, nil
}
