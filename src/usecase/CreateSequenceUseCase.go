package usecase

import (
	"abix360/src/dao/mysql"
	"abix360/src/domain"
	"abix360/src/view/dto"
	"errors"
)

type CreateSequenceUseCase struct {
}

func (useCase *CreateSequenceUseCase) Execute(dtoSequence dto.SequenceDto) (code int, err error) {
	var repository domain.RepositorySequence = mysql.NewSequenceDao()

	sequence := domain.FindSequenceByName(dtoSequence.Name, repository)
	if sequence.Exists() {
		return 202, errors.New("la secuencia ya existe")
	}

	sequence = *domain.NewSequence(dtoSequence.Name).WithPrefix(dtoSequence.Prefix).WithValue(dtoSequence.Value).WithRepository(repository)

	err = sequence.Create()
	if err != nil {
		return 500, err
	}

	return 200, nil
}
