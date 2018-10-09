package app

import (
	"net/http"

	"github.com/kenkoii/WordCosmosStaticServer/api/models"
	"github.com/labstack/echo"
	"google.golang.org/appengine"
)

func init() {
	e := echo.New()
	e.GET("/", indexHandler)
	e.GET("/review", reviewHandler)

	http.Handle("/", e)
}

func indexHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func reviewHandler(c echo.Context) error {
	ctx := appengine.NewContext(c.Request())
	appReviewInfo, err := models.NewAppReviewInfo(ctx, c.Request().Body)
	if err != nil {
		//Handle Error
		c.JSON(http.StatusInternalServerError, err.Error)
		return err
	}

	// Return status 200 for success
	c.JSON(http.StatusOK, appReviewInfo)
	return nil
}
