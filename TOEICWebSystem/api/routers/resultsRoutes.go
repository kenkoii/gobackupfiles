package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/kenkoii/TOEICWebSystem/api/handlers"
)

// SetResultsRoutes sets routing for Words Endpoint
// func SetResultsRoutes(router *mux.Router) *mux.Router {
// 	resultsRouter := mux.NewRouter()
// 	resultsRouter.HandleFunc("/api/v1/results", handlers.PostResultEndpoint).Methods("POST")
// 	resultsRouter.HandleFunc("/api/v1/results/{id}/feedback", handlers.PostFeedbackEndpoint).Methods("POST")
// 	resultsRouter.HandleFunc("/api/v1/results/{id}/package", handlers.GetResultsByPackageEndpoint).Methods("GET")
// 	resultsRouter.HandleFunc("/api/v1/results/{id}", handlers.GetResultEndpoint).Methods("GET")
// 	router.PathPrefix("/api/v1/results").Handler(resultsRouter)
// 	return router
// }

func SetResultsRoutes(router *gin.Engine) *gin.Engine {
	results := router.Group("/api/v1/results")
	results.POST("", handlers.PostResultEndpoint)
	results.POST("/:id/feedback", handlers.PostFeedbackEndpoint)
	results.GET("/:id/package", handlers.GetResultsByPackageEndpoint)
	results.GET("/:id", handlers.GetResultEndpoint)
	return router
}
