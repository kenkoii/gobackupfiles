package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/kenkoii/GAE-practice/api/models"
	"google.golang.org/appengine"
	"net/http"
	"strconv"
)

func PostWordEndpoint(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	category, err := models.NewWord(ctx, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(category)
}

func GetWordsEndpoint(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	categories, err := models.GetWords(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(categories)
}

func GetWordEndpoint(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	ctx := appengine.NewContext(r)
	category, err := models.GetWord(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(category)
}

func DeleteWordEndpoint(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	ctx := appengine.NewContext(r)
	category, err := models.RemoveWord(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(category)

}

func UpdateWordEndpoint(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	ctx := appengine.NewContext(r)
	category, err := models.UpdateWord(ctx, id, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(category)
}
