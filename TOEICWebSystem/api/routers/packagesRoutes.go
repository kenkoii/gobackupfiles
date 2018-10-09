package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/kenkoii/TOEICWebSystem/api/handlers"
)

// func SetPackagesRoutes(router *mux.Router) *mux.Router {
// 	packagesRouter := mux.NewRouter()
// 	packagesRouter.HandleFunc("/api/v1/packages", handlers.PostPackageEndpoint).Methods("POST")
// 	packagesRouter.HandleFunc("/api/v1/packages/{id}", handlers.GetPackageEndpoint).Methods("GET")
// 	packagesRouter.HandleFunc("/api/v1/packages", handlers.GetPackagesEndpoint).Methods("GET")
// 	router.PathPrefix("/api/v1/packages").Handler(packagesRouter)
// 	return router
// }

func SetPackagesRoutes(router *gin.Engine) *gin.Engine {
	packages := router.Group("/api/v1/packages")
	packages.POST("", handlers.PostPackageEndpoint)
	packages.GET("/:id", handlers.GetPackageEndpoint)
	packages.DELETE("/:id", handlers.DeletePackageEndpoint)
	packages.GET("", handlers.GetPackagesEndpoint)
	return router
}