package usecase

import (
	"abix360/src/dao/mysql"
	"abix360/src/domain"
	"abix360/src/view/dto"
	"errors"
)

type UpdateSequenceUseCase struct{}

func (useCase *UpdateSequenceUseCase) Execute(dtoSequence dto.SequenceDto) (code int, err error) {
	var repository domain.RepositorySequence = mysql.NewSequenceDao()

	sequence := domain.FindSequenceById(dtoSequence.Id, repository)
	if !sequence.Exists() {
		return 202, errors.New("la secuencia no existe")
	}

	sequence.WithName(dtoSequence.Name).WithPrefix(dtoSequence.Prefix).WithValue(dtoSequence.Value).WithRepository(repository)

	err = sequence.Update()
	if err != nil {
		return 500, err
	}
	return 200, nil
}
