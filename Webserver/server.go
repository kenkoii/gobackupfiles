package main

import (
	"log"
	"net/http"
	"github.com/kenkoii/Webserver/routers"
)

func main() {
	router := routers.NewRouter()
	//router.HandleFunc("/", Index)
	//router.HandleFunc("/todos", TodoIndex)
	//router.HandleFunc("todos/{todoId}", TodoShow)
	//http.Handle("/",router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
