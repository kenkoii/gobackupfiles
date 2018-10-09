package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/kenkoii/TOEICWebSystem/api/models"
	"google.golang.org/appengine"
)

// func PostResultEndpoint(w http.ResponseWriter, r *http.Request) {
// 	ctx := appengine.NewContext(r)
// 	result, err := models.NewTestResult(ctx, r.Body)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	json.NewEncoder(w).Encode(result)
// }
func PostResultEndpoint(c *gin.Context) {
	r := c.Request
	ctx := appengine.NewContext(r)
	result, err := models.NewTestResult(ctx, r.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
	}
	c.JSON(http.StatusOK, result)
}

// func PostFeedbackEndpoint(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, err := strconv.ParseInt(vars["id"], 10, 64)
// 	ctx := appengine.NewContext(r)
// 	result, err := models.NewFeedbackResult(ctx, id, r.Body)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	json.NewEncoder(w).Encode(result)
// }
func PostFeedbackEndpoint(c *gin.Context) {
	r := c.Request
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	ctx := appengine.NewContext(r)
	result, err := models.NewFeedbackResult(ctx, id, r.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
	}
	c.JSON(http.StatusOK, result)
}

// func GetResultEndpoint(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, err := strconv.ParseInt(vars["id"], 10, 64)
// 	ctx := appengine.NewContext(r)
// 	results, err := models.GetResult(ctx, id)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	json.NewEncoder(w).Encode(results)
// }
func GetResultEndpoint(c *gin.Context) {
	r := c.Request
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	ctx := appengine.NewContext(r)
	result, err := models.GetResult(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
	}
	c.JSON(http.StatusOK, result)
}

// func GetResultsByPackageEndpoint(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, err := strconv.ParseInt(vars["id"], 10, 64)
// 	ctx := appengine.NewContext(r)
// 	results, err := models.GetResultsByPackage(ctx, id)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	json.NewEncoder(w).Encode(results)
// }
func GetResultsByPackageEndpoint(c *gin.Context) {
	r := c.Request
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	ctx := appengine.NewContext(r)
	var results interface{}
	if c.Query("filter") == "first" {
		results, err = models.GetFirstResultsByPackage(ctx, id)
	} else {
		results, err = models.GetResultsByPackage(ctx, id)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
	}
	c.JSON(http.StatusOK, results)
}
