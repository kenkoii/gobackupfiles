package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

// InitRoutes is for initializing all routes/endpoints
func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	// router = SetPackagesRoutes(router)
	// // router = SetWordsRoutes(router)
	// // router = SetUsersRoutes(router)
	// router = SetResultsRoutes(router)
	return router.StrictSlash(false)
}

func InitGinRoutes(router *gin.Engine) *gin.Engine {
	router = SetUsersRoutes(router)
	router = SetResultsRoutes(router)
	router = SetPackagesRoutes(router)
	return router
}
