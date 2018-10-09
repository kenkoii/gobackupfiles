package routers

import(
	"github.com/gorilla/mux"
	"github.com/kenkoii/GAE-practice/api/handlers"
)

func SetCategoriesRoutes(router *mux.Router) *mux.Router{
	categoriesRouter := mux.NewRouter()
	categoriesRouter.HandleFunc("/categories", handlers.GetCategoriesEndpoint).Methods("GET")
	categoriesRouter.HandleFunc("/categories/{id}", handlers.GetCategoryEndpoint).Methods("GET")
	categoriesRouter.HandleFunc("/categories/{id}", handlers.DeleteCategoryEndpoint).Methods("DELETE")
	categoriesRouter.HandleFunc("/categories/{id}", handlers.UpdateCategoryEndpoint).Methods("PUT")
	categoriesRouter.HandleFunc("/categories", handlers.PostCategoryEndpoint).Methods("POST")
	router.PathPrefix("/categories").Handler(categoriesRouter)
	//router.
	return router
}