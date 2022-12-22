package usecase

import (
	"abix360/src/dao/mysql"
	"abix360/src/domain"
	"abix360/src/view/dto"
	"errors"
)

type ViewListUseCase struct{}

func (useCase *ViewListUseCase) Execute(id int64) (dtoList dto.ListDto, err error) {
	var repository domain.RepositoryList = mysql.NewListDao()
	list, err := domain.FindListById(id, repository)
	if err != nil {
		return dtoList, err
	}

	if !list.Exists() {
		return dtoList, errors.New("la lista no existe")
	}

	dtoList.Id = list.Id()
	dtoList.Name = list.Name()
	dtoList.Values = list.Values()

	return dtoList, err
}
