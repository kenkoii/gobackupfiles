package routers

import (
	"github.com/gorilla/mux"
	"github.com/kenkoii/TOEICWebSystem/api/handlers"
)

// SetWordsRoutes sets routing for Words Endpoint
func SetUsersRoutes(router *mux.Router) *mux.Router {
	usersRouter := mux.NewRouter()
	// usersRouter.HandleFunc("/api/v1/users", handlers.PostWordEndpoint).Methods("POST")
	usersRouter.HandleFunc("/api/v1/users/{id}", handlers.LoginUserEndpoint).Methods("POST")
	usersRouter.HandleFunc("/api/v1/users/{id}", handlers.UpdateUserEndpoint).Methods("PUT")
	router.PathPrefix("/api/v1/users").Handler(usersRouter)
	//router.
	return router
}
