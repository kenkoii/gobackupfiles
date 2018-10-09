package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kenkoii/TOEICWebSystem/api/models"
	"google.golang.org/appengine"
)

func GetUserEndpoint(c *gin.Context) {
	id := c.Param("id")
	ctx := appengine.NewContext(c.Request)
	user, err := models.GetUser(ctx, id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, user)
}

func UpdateUserEndpoint(c *gin.Context) {
	id := c.Param("id")
	ctx := appengine.NewContext(c.Request)
	user, err := models.UpdateUser(ctx, id, c.Request.Body)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, user)
}

func LoginUserEndpoint(c *gin.Context) {
	id := c.Param("id")
	ctx := appengine.NewContext(c.Request)
	user, err := models.GetUser(ctx, id)
	if err != nil {
		user, err = models.NewUser(ctx, c.Request.Body)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
	}
	c.JSON(http.StatusOK, user)
}

// GetUserEndpoint handles the /api/v1/users/{id} {GET} method
// func GetUserEndpoint(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id := vars["id"]
// 	ctx := appengine.NewContext(r)
// 	user, err := models.GetUser(ctx, id)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	json.NewEncoder(w).Encode(user)
// }

// UpdateUserEndpoint handles the /api/v1/users/{id} {GET} method
// func UpdateUserEndpoint(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id := vars["id"]

// 	ctx := appengine.NewContext(r)
// 	user, err := models.UpdateUser(ctx, id, r.Body)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	json.NewEncoder(w).Encode(user)
// }

// func LoginUserEndpoint(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id := vars["id"]

// 	ctx := appengine.NewContext(r)
// 	user, err := models.GetUser(ctx, id)
// 	if err != nil {
// 		user, err = models.NewUser(ctx, r.Body)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 	}
// 	json.NewEncoder(w).Encode(user)
// }
