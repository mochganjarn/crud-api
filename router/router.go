package router

import (
	"net/http"
	"os"

	"github.com/rysmaadit/go-template/handler"
	"github.com/rysmaadit/go-template/service"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func NewRouter(dependencies service.Dependencies) http.Handler {
	r := mux.NewRouter()

	setAuthRouter(r, dependencies.AuthService)
	showMovie(r)
	createMovie(r)
	updateMovie(r)
	deleteMovie(r)
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	return loggedRouter
}

func setAuthRouter(router *mux.Router, dependencies service.AuthServiceInterface) {
	router.Methods(http.MethodGet).Path("/auth/token").Handler(handler.GetToken(dependencies))
	router.Methods(http.MethodPost).Path("/auth/token/validate").Handler(handler.ValidateToken(dependencies))
}

func showMovie(router *mux.Router) {
	router.Methods(http.MethodGet).Path("/movie/{slug}").Handler(handler.ShowMovie())
}

func createMovie(router *mux.Router) {
	router.Methods(http.MethodPost).Path("/movie").Handler(handler.CreateMovie())
}

func updateMovie(router *mux.Router) {
	router.Methods(http.MethodPut).Path("/movie/{slug}").Handler(handler.UpdateMovie())
}

func deleteMovie(router *mux.Router) {
	router.Methods(http.MethodDelete).Path("/movie/{slug}").Handler(handler.DeleteMovie())
}
