package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kenkoii/Analytics/api/models/EnglishStory"
	"google.golang.org/appengine"
)

// PostQuestEndpoint handles POST requests on Quest endpoint
func PostQuestEndpoint(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	quest, err := models.NewQuest(ctx, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(quest)
}

// GetQuestsEndpoint handles GET all requests on Quests endpoint
func GetQuestsEndpoint(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	quests, err := models.GetAllQuests(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(quests)
}
