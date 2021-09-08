package api

import (
	"net/http"

	"github.com/Devcon4/grackle/packages/framework"
)

func NewMainRouter() *framework.Router {
	r := framework.NewRouter("/main")

	r.Handle("/info", infoHandler())
	return r
}

func infoHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
