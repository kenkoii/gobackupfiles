package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kenkoii/Analytics/api/models/EnglishStory"
	"google.golang.org/appengine"
)

// PostEventEndpoint handles POST requests on Event endpoint
func PostEventEndpoint(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	event, err := models.NewEvent(ctx, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(event)
}

// GetEventsEndpoint handles GET all requests on Events endpoint
func GetEventsEndpoint(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	events, err := models.GetAllEvents(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(events)
}
