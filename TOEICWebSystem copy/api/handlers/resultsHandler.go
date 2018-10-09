package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kenkoii/TOEICWebSystem/api/models"
	"google.golang.org/appengine"
)

func PostResultEndpoint(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	result, err := models.NewTestResult(ctx, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func PostFeedbackEndpoint(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	ctx := appengine.NewContext(r)
	result, err := models.NewFeedbackResult(ctx, id, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}
