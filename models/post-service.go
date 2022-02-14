package models

import (
	"errors"
	"math/rand"

	"github.com/ardhihdra/go-clean-arch/entity"
	"github.com/ardhihdra/go-clean-arch/repository"
)

type PostModels interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
	FindByID(id int64) (entity.Post, error)
}

type models struct{}

var (
	repo repository.PostRepository
)

func NewPostModels(repos repository.PostRepository) PostModels {
	repo = repos
	return &models{}
}

func (*models) Validate(post *entity.Post) error {
	if post == nil {
		err := errors.New("post cannot be empty")
		return err
	}
	if post.Title == "" {
		return errors.New("post title cannot be empty")
	}
	return nil
}

func (*models) Create(post *entity.Post) (*entity.Post, error) {
	post.ID = int64(rand.Intn(10000))
	return repo.Save(post)
}

func (*models) FindAll() ([]entity.Post, error) {
	return repo.FindAll()
}

func (*models) FindByID(id int64) (entity.Post, error) {
	return repo.FindByID(id)
}
