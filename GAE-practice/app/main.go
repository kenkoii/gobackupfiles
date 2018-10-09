package main

import (
	"github.com/kenkoii/GAE-practice/api/handlers"
	"github.com/kenkoii/GAE-practice/api/routers"
	"github.com/codegangsta/negroni"
	"net/http"
	"github.com/rs/cors"
	//"github.com/kenkoii/GAE-practice/api/common"
	"github.com/kenkoii/GAE-practice/api/common"
)

func init() {
	//router := mux.NewRouter()
	/*http.Handle("/", router)*/
	//router.HandleFunc("/", handlers.Handler)*/
	//Categories
	// Get the mux router object
	common.StartUp()
	router := routers.InitRoutes()
	router.HandleFunc("/",handlers.Handler);
	n := negroni.Classic()
	handler := cors.Default().Handler(router)
	n.UseHandler(handler)
	http.Handle("/", n)
}
