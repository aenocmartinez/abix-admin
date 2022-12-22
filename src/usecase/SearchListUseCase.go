package usecase

import (
	"abix360/src/dao/mysql"
	"abix360/src/domain"
	"abix360/src/view/dto"
)

type SearchListUseCase struct{}

func (useCase *SearchListUseCase) Execute(name string) (dtolists []dto.ListDto, err error) {
	var repository domain.RepositoryList = mysql.NewListDao()
	lists, err := domain.SearchList(name, repository)
	if err != nil {
		return dtolists, err
	}

	for _, list := range lists {
		dtolists = append(dtolists, dto.ListDto{
			Id:     list.Id(),
			Name:   list.Name(),
			Values: list.Values(),
		})
	}

	return dtolists, err
}
