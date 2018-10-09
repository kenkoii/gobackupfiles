package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/kenkoii/GAE-practice/api/models"
	"google.golang.org/appengine"
	"net/http"
	"strconv"
)

func PostTopicEndpoint(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	category, err := models.NewTopic(ctx, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(category)
}

func GetTopicsEndpoint(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	categories, err := models.GetTopics(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(categories)
}

func GetTopicEndpoint(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	ctx := appengine.NewContext(r)
	category, err := models.GetTopic(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(category)
}

func DeleteTopicEndpoint(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	ctx := appengine.NewContext(r)
	category, err := models.RemoveTopic(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(category)

}

func UpdateTopicEndpoint(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	ctx := appengine.NewContext(r)
	category, err := models.UpdateTopic(ctx, id, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(category)
}
