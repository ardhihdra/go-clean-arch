package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type muxRouter struct{}

var (
	muxDispatcher = mux.NewRouter()
)

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) GET(uri string, f func(res http.ResponseWriter, req *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods(http.MethodGet)
}

func (*muxRouter) POST(uri string, f func(res http.ResponseWriter, req *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods(http.MethodPost)
}

func (*muxRouter) SERVE(port string) {
	log.Println("Server is up and running on port", port)
	log.Fatalln(http.ListenAndServe(port, muxDispatcher))
}
