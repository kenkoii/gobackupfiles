package handlers

import (
	"encoding/json"
	"github.com/kenkoii/GAE-practice/api/models"
	"net/http"
	"google.golang.org/appengine"
	//"github.com/kenkoii/GAE-practice/api/common"
	"github.com/kenkoii/GAE-practice/api/common"
	"log"
)

// Handler for HTTP Post - "/users/register"
// Add a new User document
func Register(w http.ResponseWriter, r *http.Request) {

	context := appengine.NewContext(r)
	user, err := models.NewUser(context, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if j, err := json.Marshal(UserResource{Data: *user}); err != nil {
		common. DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	} else {
		user.HashPassword = nil
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		//json.NewEncoder(w).Encode(user)
		w.Write(j)
	}
}

func Login(w http.ResponseWriter, r *http.Request){
	context := appengine.NewContext(r)
	user, err := models.LoginUser(context, r.Body)
	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		log.Println("AAAAAA")

		token, err := common.GenerateJWT(user.Email, "member")

		user.HashPassword = nil
		authUser := AuthUserModel{
			User: *user,
			Token: token,
		}
		j, err := json.Marshal(AuthUserResource{Data: authUser})
		log.Println(token)
		if err != nil {

			common.DisplayAppError(
				w,
				err,
				"An unexpected error has occurred",
				500,
			)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}
/*
func Login(w http.ResponseWriter, r *http.Request) {
	var dataResource LoginResource
	var token string
	// Decode the incoming Login json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid Login data",
			500,
		)
		return
	}
	loginModel := dataResource.Data
	loginUser := models.User{
		Email: loginModel.Email,
		Password: loginModel.Password,
	}
	context := appengine.NewContext()

	c := context.DbCollection("users")
	repo := &data.UserRepository{c}
	// Authenticate the login user
	if user, err := repo.Login(loginUser); err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid login credentials",
			401,
		)
		return
	} else { //if login is successful
		// Generate JWT token
		token, err = common.GenerateJWT(user.Email, "member")
		if err != nil {
			common.DisplayAppError(
				w,
				err,
				"Eror while generating the access token",
				500,
			)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		user.HashPassword = nil
		authUser := AuthUserModel{
			User: user,
			Token: token,
		}
		j, err := json.Marshal(AuthUserResource{Data: authUser})
		if err != nil {
			common.DisplayAppError(
				w,
				err,
				"An unexpected error has occurred",
				500,
			)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}*/