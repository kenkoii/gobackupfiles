package routers

import (
	"github.com/gorilla/mux"
)

// InitRoutes is for initializing all routes/endpoints
func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = SetPackagesRoutes(router)
	// router = SetWordsRoutes(router)
	router = SetUsersRoutes(router)
	router = SetResultsRoutes(router)
	return router.StrictSlash(false)
}
