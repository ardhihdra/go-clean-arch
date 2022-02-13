package controller

import (
	"encoding/json"
	"net/http"

	"github.com/ardhihdra/go-clean-arch/entity"
	"github.com/ardhihdra/go-clean-arch/errors"
	"github.com/ardhihdra/go-clean-arch/models"
)

type PostController interface {
	GetPosts(res http.ResponseWriter, req *http.Request)
	AddPost(res http.ResponseWriter, req *http.Request)
}

type postController struct{}

var postModel models.PostModels

func NewPostController(model models.PostModels) PostController {
	/** model is passed as parameter for testability */
	postModel = model
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
