package repository

import (
	"github.com/ardhihdra/go-clean-arch/entity"
)

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}
