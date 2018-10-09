package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kenkoii/Analytics/api/models/EnglishStory"
	"google.golang.org/appengine"
)

// PostTutorialEndpoint handles POST request on Tutorial endpoint
func PostTutorialEndpoint(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	tutorial, err := models.NewTutorial(ctx, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tutorial)
}

// GetTutorialsEndpoint handles GET all request on Tutorials endpoint
func GetTutorialsEndpoint(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	tutorials, err := models.GetAllTutorials(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tutorials)
}
