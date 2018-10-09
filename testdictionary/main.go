package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/kenkoii/testdictionary/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

var topics models.Topics
var words models.Words
var session, err = mgo.Dial("mongodb://user:pass@ds149437.mlab.com:49437/testdict")


func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"Hello World")
}

func GetTopicsEndpoint(w http.ResponseWriter, r *http.Request){
	c := session.DB("testdict").C("Topics")
	result := models.Topics{}
	err = c.Find(nil).All(&result)
	json.NewEncoder(w).Encode(result)

}

func GetTopicEndpoint(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	id := vars["topic"]
	w.Header().Set("Content-Type", "application/json")
	c := session.DB("testdict").C("Topics")
	result := models.Topic{}
	err = c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&result)
	if err != nil{
		panic(err)
	}
	json.NewEncoder(w).Encode(result)

}

func DeleteTopicEndpoint(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["topic"]
	c := session.DB("testdict").C("Topics")
	err := c.RemoveId(bson.ObjectIdHex(id))
	if err != nil{
		panic(err)
	}
}

func PostTopicEndpoint(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var t models.Topic
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil{
		http.Error(w, err.Error(),400)
		return
	}

	c := session.DB("testdict").C("Topics")
	topic :=  models.Topic{
		Name:		t.Name,
		AddedAt:	time.Now(),
	}
	err = c.Insert(&topic)
	json.NewEncoder(w).Encode(topic)
}


func GetWordsEndpoint(w http.ResponseWriter, r *http.Request){
	c := session.DB("testdict").C("Words")
	result := models.Words{}
	err = c.Find(nil).All(&result)
	json.NewEncoder(w).Encode(result)
}

func PostWordEndpoint(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var t models.Word
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil{
		http.Error(w, err.Error(),400)
		return
	}
	c := session.DB("testdict").C("Words")
	word :=  models.Word{
		Topic: t.Topic,
		Name:	t.Name,
		AddedAt:	time.Now(),
	}
	go func(){
		c := session.DB("testdict").C("Topics")
		err := c.Update(bson.M{"_id": bson.ObjectIdHex(t.Topic)},bson.M{"$push" :bson.M{"words": word}})
		if err != nil{
			http.Error(w, err.Error(),400)
			return
		}
	}()
	err = c.Insert(&word)
	json.NewEncoder(w).Encode(word)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/",handler)
	//Categories

	//topics
	router.HandleFunc("/topics",GetTopicsEndpoint).Methods("GET")
	router.HandleFunc("/topics",PostTopicEndpoint).Methods("POST")
	router.HandleFunc("/topics/{topic}",GetTopicEndpoint).Methods("GET")
	router.HandleFunc("/topics/{topic}",DeleteTopicEndpoint).Methods("DELETE")
	router.HandleFunc("/topics/{topic}/words",GetWordsEndpoint).Methods("GET")
	//Words
	router.HandleFunc("/words",GetWordsEndpoint).Methods("GET")
	router.HandleFunc("/words",PostWordEndpoint).Methods("POST")

	log.Fatal(http.ListenAndServe(":12345", router))
}
