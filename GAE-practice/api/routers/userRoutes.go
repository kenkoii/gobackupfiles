package routers

import(
	"github.com/gorilla/mux"
	"github.com/kenkoii/GAE-practice/api/handlers"
)

func SetUserRoutes(router *mux.Router) *mux.Router{
	r := mux.NewRouter()
	r.HandleFunc("/users/register", handlers.Register).Methods("POST")
	r.HandleFunc("/users/login", handlers.Login).Methods("POST")

	router.PathPrefix("/users").Handler(r)
	//router.
	return router
}