package repository

const ()

func NewMongoRepository() PostRepository {
	return &repo{}
}
