package routers

import (
	"github.com/gorilla/mux"
	"github.com/kenkoii/AnalyticsMongoDB/api/handlers/EnglishStory"
)

// SetOtherRoutes initializes routes pertaining to users
func SetOtherRoutes(router *mux.Router) *mux.Router {
	r := mux.NewRouter()

	// GET requests
	r.HandleFunc("/analytics/englishstory/events", handlers.GetEventsEndpoint).Methods("GET")
	r.HandleFunc("/analytics/englishstory/quests", handlers.GetQuestsEndpoint).Methods("GET")
	r.HandleFunc("/analytics/englishstory/stages", handlers.GetStagesEndpoint).Methods("GET")
	r.HandleFunc("/analytics/englishstory/tutorials", handlers.GetTutorialsEndpoint).Methods("GET")

	// POST requests
	r.HandleFunc("/analytics/englishstory/events", handlers.PostEventEndpoint).Methods("POST")
	r.HandleFunc("/analytics/englishstory/quests", handlers.PostQuestEndpoint).Methods("POST")
	r.HandleFunc("/analytics/englishstory/stages", handlers.PostStageEndpoint).Methods("POST")
	r.HandleFunc("/analytics/englishstory/tutorials", handlers.PostTutorialEndpoint).Methods("POST")

	router.PathPrefix("/analytics/englishstory").Handler(r)
	return router
}
