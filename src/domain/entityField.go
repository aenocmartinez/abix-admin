package domain

func FindFieldById(id int64, repository FieldRepository) IField {
	return repository.FindById(id)
}

func FindFieldByName(name string, repository FieldRepository) IField {
	return repository.FindByName(name)
}

func SearchFieldByName(name string, repository FieldRepository) []IField {
	return repository.SearchByName(name)
}

func AllFields(repository FieldRepository) []IField {
	return repository.AllFields()
}
