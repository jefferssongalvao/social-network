package router

import (
	"social-network/src/router/routes"

	"github.com/gorilla/mux"
)

func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.Configure(r)
}
