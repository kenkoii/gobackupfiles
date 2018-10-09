package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kenkoii/AnalyticsMongoDB/api/models/EnglishStory"
	"google.golang.org/appengine"
)

// PostUserPropertyEndpoint handles POST requests on UserProperty endpoint
func PostUserPropertyEndpoint(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	userProperty, err := models.NewUserProperty(ctx, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(userProperty)
}

// GetUserPropertiesEndpoint handles GET all requests on UserProperty endpoint
func GetUserPropertiesEndpoint(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	userProperties, err := models.GetAllUserProperties(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(userProperties)
}

// PostUserPurchaseEndpoint handles POST requests on UserPurchase endpoint
func PostUserPurchaseEndpoint(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	userPurchase, err := models.NewUserPurchase(ctx, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(userPurchase)
}

// GetUserPurchasesEndpoint handles GET all requests on UserProperty endpoint
func GetUserPurchasesEndpoint(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	userPurchases, err := models.GetAllUserPurchases(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(userPurchases)
}

// PostUserDailyPropertyEndpoint handles POST requests on UserDailyProperty endpoint
func PostUserDailyPropertyEndpoint(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	userPurchase, err := models.NewUserDailyProperty(ctx, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(userPurchase)
}

// GetUserDailyPropertiesEndpoint handles GET all requests on UserDailyProperty endpoint
func GetUserDailyPropertiesEndpoint(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	userDailyProperties, err := models.GetAllUserDailyProperties(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(userDailyProperties)
}
