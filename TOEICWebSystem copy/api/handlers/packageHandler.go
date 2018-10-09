package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kenkoii/TOEICWebSystem/api/models"
	"google.golang.org/appengine"
)

func PostPackageEndpoint(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	pack, err := models.NewPackage(ctx, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(pack)
}

// GetPackageEndpoint handles the /api/v1/packages/{id} {GET} method
func GetPackageEndpoint(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	ctx := appengine.NewContext(r)
	user, err := models.GetPackage(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func GetPackagesEndpoint(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	user, err := models.GetPackages(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}
