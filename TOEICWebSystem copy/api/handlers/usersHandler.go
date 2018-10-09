package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kenkoii/TOEICWebSystem/api/models"
	"google.golang.org/appengine"
)

// GetUserEndpoint handles the /api/v1/users/{id} {GET} method
func GetUserEndpoint(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	ctx := appengine.NewContext(r)
	user, err := models.GetUser(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

// UpdateUserEndpoint handles the /api/v1/users/{id} {GET} method
func UpdateUserEndpoint(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	ctx := appengine.NewContext(r)
	user, err := models.UpdateUser(ctx, id, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func LoginUserEndpoint(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	ctx := appengine.NewContext(r)
	user, err := models.GetUser(ctx, id)
	if err != nil {
		user, err = models.NewUser(ctx, r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	json.NewEncoder(w).Encode(user)
}
