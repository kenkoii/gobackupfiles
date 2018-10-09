package routers

import (
	"github.com/gorilla/mux"
)
func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	// Routes for the Topic entity
	router = SetTopicsRoutes(router)
	// Routes for the Category entity
	router = SetCategoriesRoutes(router)
	// Routes for the Word entity
	router = SetWordsRoutes(router)
	// Routes for the Word entity
	router = SetUserRoutes(router)
	return router
}