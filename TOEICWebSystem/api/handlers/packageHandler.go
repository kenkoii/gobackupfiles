package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/kenkoii/TOEICWebSystem/api/models"
	"google.golang.org/appengine"
)

// func PostPackageEndpoint(w http.ResponseWriter, r *http.Request) {
// 	ctx := appengine.NewContext(r)
// 	pack, err := models.NewPackage(ctx, r.Body)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	json.NewEncoder(w).Encode(pack)
// }
func PostPackageEndpoint(c *gin.Context) {
	r := c.Request
	ctx := appengine.NewContext(r)
	pack, err := models.NewPackage(ctx, r.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
	}
	c.JSON(http.StatusOK, pack)
}

// GetPackageEndpoint handles the /api/v1/packages/{id} {GET} method
// func GetPackageEndpoint(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, _ := strconv.ParseInt(vars["id"], 10, 64)
// 	ctx := appengine.NewContext(r)
// 	user, err := models.GetPackage(ctx, id)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	json.NewEncoder(w).Encode(user)
// }
func GetPackageEndpoint(c *gin.Context) {
	r := c.Request
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	ctx := appengine.NewContext(r)
	pack, err := models.GetPackage(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
	}
	c.JSON(http.StatusOK, pack)
}

func DeletePackageEndpoint(c *gin.Context) {
	r := c.Request
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	ctx := appengine.NewContext(r)
	pack, err := models.DeletePackage(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
	}
	c.JSON(http.StatusOK, pack)
}

// func GetPackagesEndpoint(w http.ResponseWriter, r *http.Request) {
// 	ctx := appengine.NewContext(r)
// 	user, err := models.GetPackages(ctx)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	json.NewEncoder(w).Encode(user)
// }

func GetPackagesEndpoint(c *gin.Context) {
	r := c.Request
	ctx := appengine.NewContext(r)
	packs, err := models.GetPackages(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
	}
	c.JSON(http.StatusOK, packs)
}
