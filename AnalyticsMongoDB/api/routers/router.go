package routers

import (
	"log"

	"github.com/gorilla/mux"
)

// InitRoutes is for initializing all routes/endpoints
func InitRoutes() *mux.Router {
	log.Println("Initialized!")
	router := mux.NewRouter()
	//Initiate User Routes
	router = SetUserRoutes(router)
	//Initiate other routes
	router = SetOtherRoutes(router)
	return router
}
