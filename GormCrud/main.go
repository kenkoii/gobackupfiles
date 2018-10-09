package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Greeting struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func greet(c echo.Context) error {
	return c.JSON(http.StatusOK, &Greeting{Message: "Hello World", Code: 200})
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", greet)

	u := e.Group("/users")
	u.GET("", getAllUser)

	e.Logger.Fatal(e.Start(":8080"))
}
