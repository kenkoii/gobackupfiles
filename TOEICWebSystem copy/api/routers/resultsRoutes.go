package routers

import (
	"github.com/gorilla/mux"
	"github.com/kenkoii/TOEICWebSystem/api/handlers"
)

// SetResultsRoutes sets routing for Words Endpoint
func SetResultsRoutes(router *mux.Router) *mux.Router {
	resultsRouter := mux.NewRouter()
	resultsRouter.HandleFunc("/api/v1/results", handlers.PostResultEndpoint).Methods("POST")
	resultsRouter.HandleFunc("/api/v1/results/{id}/feedback", handlers.PostFeedbackEndpoint).Methods("POST")
	router.PathPrefix("/api/v1/results").Handler(resultsRouter)
	return router
}
