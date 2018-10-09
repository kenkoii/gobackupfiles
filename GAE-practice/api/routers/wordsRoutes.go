package routers

import(
	"github.com/gorilla/mux"
	"github.com/kenkoii/GAE-practice/api/handlers"
)

func SetWordsRoutes(router *mux.Router) *mux.Router{
	wordsRouter := mux.NewRouter()
	wordsRouter.HandleFunc("/words", handlers.GetWordsEndpoint).Methods("GET")
	wordsRouter.HandleFunc("/words/{id}", handlers.GetWordEndpoint).Methods("GET")
	wordsRouter.HandleFunc("/words/{id}", handlers.DeleteWordEndpoint).Methods("DELETE")
	wordsRouter.HandleFunc("/words/{id}", handlers.UpdateWordEndpoint).Methods("PUT")
	wordsRouter.HandleFunc("/words", handlers.PostWordEndpoint).Methods("POST")
	router.PathPrefix("/words").Handler(wordsRouter)
	//router.
	return router
}