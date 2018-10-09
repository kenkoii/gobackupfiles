package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"

	"github.com/kenkoii/WordCosmosStaticServer/api/models"
)

func init() {
	router := gin.Default()
	// Routes
	router.GET("/", indexHandler)
	router.POST("/SaveReview", saveReview)

	// Register echo context to root HTTP
	http.Handle("/", router)
}

// Handler
func indexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello World!")
}

func saveReview(c *gin.Context) {
	//Instantiate new AppReviewInfo
	ctx := appengine.NewContext(c.Request)

	// New App Review Info
	appReviewInfo, err := models.NewAppReviewInfo(ctx, c.Request.Body)
	if err != nil {
		//Handle Error
		c.JSON(http.StatusInternalServerError, err.Error)
	}

	// Return status 200 for success
	c.JSON(http.StatusOK, appReviewInfo)
}
