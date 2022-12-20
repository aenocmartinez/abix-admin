package domain

func FindFieldById(id int64, repository FieldRepository) IField {
	return repository.FindById(id)
}

func FindFieldByName(name string, search bool, repository FieldRepository) IField {
	return repository.FindByName(name, search)
}

func AllFields(repository FieldRepository) []IField {
	return repository.AllFields()
}
