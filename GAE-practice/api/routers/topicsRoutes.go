package routers

import(
	"github.com/gorilla/mux"
	"github.com/kenkoii/GAE-practice/api/handlers"
)

func SetTopicsRoutes(router *mux.Router) *mux.Router{
	topicRouter := mux.NewRouter()
	//Topics
	topicRouter.HandleFunc("/topics", handlers.GetTopicsEndpoint).Methods("GET")
	topicRouter.HandleFunc("/topics/{id}", handlers.GetTopicEndpoint).Methods("GET")
	topicRouter.HandleFunc("/topics/{id}", handlers.DeleteTopicEndpoint).Methods("DELETE")
	topicRouter.HandleFunc("/topics/{id}", handlers.UpdateTopicEndpoint).Methods("PUT")
	topicRouter.HandleFunc("/topics", handlers.PostTopicEndpoint).Methods("POST")
	router.PathPrefix("/topics").Handler(topicRouter)
	//router.
	return router
}