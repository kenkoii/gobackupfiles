package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kenkoii/Analytics/api/models/EnglishStory"
	"google.golang.org/appengine"
)

// PostStageEndpoint handles POST request on Stage endpoint
func PostStageEndpoint(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	stage, err := models.NewStage(ctx, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(stage)
}

// GetStagesEndpoint handles GET all request on Stages endpoint
func GetStagesEndpoint(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	stages, err := models.GetAllStages(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(stages)
}
