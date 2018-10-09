package routers

import (
	"github.com/gorilla/mux"
	"github.com/kenkoii/TOEICWebSystem/api/handlers"
)

func SetPackagesRoutes(router *mux.Router) *mux.Router {
	packagesRouter := mux.NewRouter()
	packagesRouter.HandleFunc("/api/v1/packages", handlers.PostPackageEndpoint).Methods("POST")
	packagesRouter.HandleFunc("/api/v1/packages/{id}", handlers.GetPackageEndpoint).Methods("GET")
	packagesRouter.HandleFunc("/api/v1/packages", handlers.GetPackagesEndpoint).Methods("GET")
	router.PathPrefix("/api/v1/packages").Handler(packagesRouter)
	return router
}
