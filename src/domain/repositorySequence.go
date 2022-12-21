package domain

type RepositorySequence interface {
	Create(sequence Sequence) error
	Update(sequence Sequence) error
	Delete(sequence Sequence) error
	FindById(id int64) (Sequence, error)
	FindByName(name string) (Sequence, error)
	AllSequences() ([]Sequence, error)
	NextValue(idSequence int64) error
	Search(name string) ([]Sequence, error)
}
