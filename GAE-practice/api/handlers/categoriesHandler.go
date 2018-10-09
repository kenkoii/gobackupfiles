package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kenkoii/GAE-practice/api/models"
	"google.golang.org/appengine"
	"net/http"
	"strconv"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func PostCategoryEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	ctx := appengine.NewContext(r)
	category, err := models.NewCategory(ctx, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(category)
}

func GetCategoriesEndpoint(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	categories, err := models.GetCategories(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(categories)
}

func GetCategoryEndpoint(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	ctx := appengine.NewContext(r)
	category, err := models.GetCategory(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(category)
}

func DeleteCategoryEndpoint(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	ctx := appengine.NewContext(r)
	category, err := models.RemoveCategory(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(category)

}

func UpdateCategoryEndpoint(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	ctx := appengine.NewContext(r)
	category, err := models.UpdateCategory(ctx, id, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(category)
}

/*
func PostCategoriesEndpoint(w http.ResponseWriter, r *http.Request){
	ctx := appengine.NewContext(r)
	usr := user.Current(ctx)
	var t models.Category
	err := json.NewDecoder(r.Body).Decode(&t)
	if usr != nil{
		t.Posted_By = usr.String()
	} else {
		t.Posted_By = "Anonymous"
	}
	t.Timestamp = time.Now()
	if err != nil{
		http.Error(w, err.Error(),400)
		return
	}

	key := datastore.NewIncompleteKey(ctx, "category", nil)
	_, err = datastore.Put(ctx, key, &t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(t)
	//fmt.Printf("%+v\n", employees)
}
*/
/*
func GetCategoriesEndpoint(w http.ResponseWriter, r *http.Request){
	ctx := appengine.NewContext(r)
	var categories []models.Category
	q := datastore.NewQuery("category")

	_, err := q.GetAll(ctx, &categories)

	if err != nil {
		http.Error(w, err.Error(),http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(categories)
}*/
