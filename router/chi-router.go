package router

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type chaiRouter struct{}

var (
	chiDispatcher = chi.NewRouter()
)

func NewChiRouter() Router {
	return &chaiRouter{}
}

func (*chaiRouter) GET(uri string, f func(res http.ResponseWriter, req *http.Request)) {
	chiDispatcher.Get(uri, f)
}

func (*chaiRouter) POST(uri string, f func(res http.ResponseWriter, req *http.Request)) {
	chiDispatcher.Post(uri, f)
}

func (*chaiRouter) SERVE(port string) {
	log.Println("Server is up and running on port", port)
	log.Fatalln(http.ListenAndServe(port, chiDispatcher))
}
