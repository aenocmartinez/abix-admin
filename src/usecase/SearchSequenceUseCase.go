package usecase

import (
	"abix360/src/dao/mysql"
	"abix360/src/domain"
	"abix360/src/view/dto"
)

type SearchSequenceUseCase struct{}

func (useCase *SearchSequenceUseCase) Execute(name string) (sequences []dto.SequenceDto, err error) {
	var repository domain.RepositorySequence = mysql.NewSequenceDao()
	allSequences, err := domain.SearchSequenceByName(name, repository)
	if err != nil {
		return sequences, err
	}
	for _, sequence := range allSequences {
		sequences = append(sequences, dto.SequenceDto{
			Id:           sequence.Id(),
			Name:         sequence.Name(),
			Prefix:       sequence.Prefix(),
			Value:        sequence.Value(),
			CurrentValue: sequence.CurrentValue(),
			CreateAt:     sequence.CreatedAt(),
			UpdatedAt:    sequence.UpdatedAt(),
		})
	}
	return sequences, nil
}
