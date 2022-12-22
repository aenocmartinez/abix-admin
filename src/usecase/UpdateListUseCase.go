package usecase

import (
	"abix360/src/dao/mysql"
	"abix360/src/domain"
	"abix360/src/view/dto"
	"errors"
)

type UpdateListUseCase struct{}

func (useCase *UpdateListUseCase) Execute(dtoList dto.ListDto) (code int, err error) {
	var repository domain.RepositoryList = mysql.NewListDao()
	list, err := domain.FindListById(dtoList.Id, repository)
	if err != nil {
		return 500, err
	}

	if !list.Exists() {
		return 202, errors.New("la lista no existe")
	}

	list.WithRepository(repository).WithName(dtoList.Name).WithValues(dtoList.Values)

	if err = list.Update(); err != nil {
		return 500, err
	}

	return 200, nil
}
