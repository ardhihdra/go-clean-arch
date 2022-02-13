package repository

const ()

func NewPostgresRepository() PostRepository {
	return &repo{}
}
