package cache

import "github.com/ardhihdra/go-clean-arch/entity"

type PostCache interface {
	Set(key string, value entity.Post)
	Get(key string) *entity.Post
}
