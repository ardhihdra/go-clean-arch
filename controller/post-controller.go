package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/ardhihdra/go-clean-arch/cache"
	"github.com/ardhihdra/go-clean-arch/entity"
	"github.com/ardhihdra/go-clean-arch/errors"
	"github.com/ardhihdra/go-clean-arch/models"
)

type PostController interface {
	GetPosts(res http.ResponseWriter, req *http.Request)
	GetPostByID(res http.ResponseWriter, req *http.Request)
	AddPost(res http.ResponseWriter, req *http.Request)
}

type postController struct{}

var (
	postCache cache.PostCache
	postModel models.PostModels
)

func NewPostController(model models.PostModels, cache cache.PostCache) PostController {
	/** model is passed as parameter for testability */
	postModel = model
	postCache = cache
	return &postController{}
}

func (*postController) GetPosts(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	posts, err := postModel.FindAll()
	// result, err := json.Marshal(posts)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		payload := errors.ServiceError{Message: "Error getting posts"}
		json.NewEncoder(res).Encode(payload)
		return
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(posts)
	// res.Write(result)
}

func (*postController) GetPostByID(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	postID := strings.Split(req.URL.Path, "/")[2]
	var post *entity.Post = postCache.Get(postID)
	if post == nil {
		id, err := strconv.Atoi(postID)
		post, err := postModel.FindByID(int64(id))
		if err != nil {
			res.WriteHeader(http.StatusNotFound)
			payload := errors.ServiceError{Message: "No posts found"}
			json.NewEncoder(res).Encode(payload)
			return
		}
		postCache.Set(postID, post)
		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(post)
	} else {
		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(post)
	}
}

func (*postController) AddPost(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	var post *entity.Post
	if err := json.NewDecoder(req.Body).Decode(&post); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		payload := errors.ServiceError{Message: "Error marshalling the posts array"}
		json.NewEncoder(res).Encode(payload)
		return
	}
	if err := postModel.Validate(post); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		payload := errors.ServiceError{Message: err.Error()}
		json.NewEncoder(res).Encode(payload)
		return
	}

	result, err2 := postModel.Create(post)
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		payload := errors.ServiceError{Message: "Error saving post"}
		json.NewEncoder(res).Encode(payload)
		return
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(result)
}
