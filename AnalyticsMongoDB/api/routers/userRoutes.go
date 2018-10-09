package routers

import (
	"github.com/gorilla/mux"
	"github.com/kenkoii/AnalyticsMongoDB/api/handlers/EnglishStory"
)

// SetUserRoutes initializes routes pertaining to users
func SetUserRoutes(router *mux.Router) *mux.Router {
	r := mux.NewRouter()

	// GET requests
	r.HandleFunc("/analytics/englishstory/userproperties", handlers.GetUserPropertiesEndpoint).Methods("GET")
	r.HandleFunc("/analytics/englishstory/userpurchases", handlers.GetUserPurchasesEndpoint).Methods("GET")
	r.HandleFunc("/analytics/englishstory/userdailyproperties", handlers.GetUserDailyPropertiesEndpoint).Methods("GET")

	// POST requests
	r.HandleFunc("/analytics/englishstory/userproperties", handlers.PostUserPropertyEndpoint).Methods("POST")
	r.HandleFunc("/analytics/englishstory/userpurchases", handlers.PostUserPurchaseEndpoint).Methods("POST")
	r.HandleFunc("/analytics/englishstory/userdailyproperties", handlers.PostUserDailyPropertyEndpoint).Methods("POST")

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
