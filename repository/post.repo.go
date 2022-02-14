package repository

import (
	"github.com/ardhihdra/go-clean-arch/entity"
)

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
	FindByID(id int64) (entity.Post, error)
}
