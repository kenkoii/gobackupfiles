package app

import (
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"time"

	mgo "gopkg.in/mgo.v2"

	"github.com/codegangsta/negroni"
	"github.com/kenkoii/AnalyticsMongoDB/api/handlers"
	"github.com/kenkoii/AnalyticsMongoDB/api/routers"
	"github.com/rs/cors"
)

const (
	MongoDBHosts = "ds127731.mlab.com:27731"
	AuthDatabase = "frecre-analytics"
	AuthUserName = "kenkoii"
	AuthPassword = "frecre"
)

func init() {
	c := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT"},
	})

	router := routers.InitRoutes()
	router.HandleFunc("/", handlers.Handler)
	n := negroni.Classic()
	handler := c.Handler(router)
	n.UseHandler(handler)
	http.Handle("/", n)

	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{MongoDBHosts},
		Timeout:  60 * time.Second,
		Database: AuthDatabase,
		Username: AuthUserName,
		Password: AuthPassword,
	}

	mongoDBDialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), nil)
		return conn, err
	}

	// _, err := mgo.Dial("mongodb://kenkoii:frecre@ds127731.mlab.com:27731/frecre-analytics")
	session, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Println("cannot dial mongo: ", err)
	} else {
		log.Println("Connected!")
	}
	defer session.Close()
}
