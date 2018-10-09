package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/kenkoii/TOEICWebSystem/api/handlers"
)

// SetWordsRoutes sets routing for Words Endpoint
// func SetUsersRoutes(router *mux.Router) *mux.Router {
// 	usersRouter := mux.NewRouter()
// 	// usersRouter.HandleFunc("/api/v1/users", handlers.PostWordEndpoint).Methods("POST")
// 	usersRouter.HandleFunc("/api/v1/users/{id}", handlers.LoginUserEndpoint).Methods("POST")
// 	usersRouter.HandleFunc("/api/v1/users/{id}", handlers.UpdateUserEndpoint).Methods("PUT")
// 	router.PathPrefix("/api/v1/users").Handler(usersRouter)
// 	//router.
// 	return router
// }

func SetUsersRoutes(router *gin.Engine) *gin.Engine {
	user := router.Group("/api/v1/users")
	user.POST("/:id", handlers.LoginUserEndpoint)
	user.PUT("/:id", handlers.UpdateUserEndpoint)
	return router
}
