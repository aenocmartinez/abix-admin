package usecase

import (
	"abix360/src/dao/mysql"
	"abix360/src/domain"
	"abix360/src/view/dto"
	"errors"
)

type ViewSequenceUseCase struct{}

func (useCase *ViewSequenceUseCase) Execute(id int64) (dtoSequence dto.SequenceDto, err error) {
	var repository domain.RepositorySequence = mysql.NewSequenceDao()
	sequence := domain.FindSequenceById(id, repository)
	if !sequence.Exists() {
		return dtoSequence, errors.New("la secuencia no existe")
	}

	dtoSequence.Id = sequence.Id()
	dtoSequence.Name = sequence.Name()
	dtoSequence.CurrentValue = sequence.CurrentValue()
	dtoSequence.Value = sequence.Value()
	dtoSequence.Prefix = sequence.Prefix()
	dtoSequence.CreateAt = sequence.CreatedAt()
	dtoSequence.UpdatedAt = sequence.UpdatedAt()

	return dtoSequence, nil
}
