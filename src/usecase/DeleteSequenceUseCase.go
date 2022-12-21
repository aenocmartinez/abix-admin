package usecase

import (
	"abix360/src/dao/mysql"
	"abix360/src/domain"
	"errors"
)

type DeleteSequenceUseCase struct {
}

func (useCase *DeleteSequenceUseCase) Execute(id int64) (code int, err error) {
	var repository domain.RepositorySequence = mysql.NewSequenceDao()

	sequence := domain.FindSequenceById(id, repository)
	if !sequence.Exists() {
		return 202, errors.New("la secuencia no existe")
	}

	sequence.WithRepository(repository)

	err = sequence.Delete()
	if err != nil {
		return 500, err
	}
	return 200, nil
}
