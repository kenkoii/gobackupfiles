package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kenkoii/TOEICWebSystem/api/controllers"
	"github.com/kenkoii/TOEICWebSystem/api/handlers"
	"github.com/kenkoii/TOEICWebSystem/api/middlewares"
	"github.com/kenkoii/TOEICWebSystem/api/routers"
)

func init() {
	http.Handle("/", GetMainEngine())
}

func GetMainEngine() *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.CORSMiddleware())
	router.GET("/", handlers.Greetings)
	router = routers.InitGinRoutes(router)
	router.LoadHTMLGlob("templates/*")
	router.GET("/index", controllers.HomePageController)
	return router
}
