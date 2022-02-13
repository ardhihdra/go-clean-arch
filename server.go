package main

import (
	"net/http"
	"os"

	"github.com/ardhihdra/go-clean-arch/controller"
	"github.com/ardhihdra/go-clean-arch/models"
	"github.com/ardhihdra/go-clean-arch/repository"
	"github.com/ardhihdra/go-clean-arch/router"
	"github.com/joho/godotenv"
)

var (
	PORT           string
	postRepo       repository.PostRepository = repository.NewFirestoreRepository()
	postModel      models.PostModels         = models.NewPostModels(postRepo)
	postController controller.PostController = controller.NewPostController(postModel)
	httpRouter                               = router.NewMuxRouter()
)

func main() {
	loadEnv()
	httpRouter.GET("/", func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		res.Write([]byte("Hello from server!"))
	})
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)
	httpRouter.SERVE(PORT)
}

func loadEnv() {
	godotenv.Load()
	PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = ":8000"
	}
}
