package models

import (
	//"github.com/asaskevich/govalidator"
	"golang.org/x/net/context"
	//"google.golang.org/appengine/datastore"
	"time"
	//"io"
	//"encoding/json"
	"github.com/asaskevich/govalidator"
	"google.golang.org/appengine/datastore"
	"io"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	//"errors"
	//"log"
	"log"
	"errors"
)

type User struct {
	//Id           int64     `json:"id" datastore:"-"`
	FirstName    string     `json:"firstname"`
	LastName     string    `json:"lastname"`
	Email        string    `json:"email"`
	Timestamp    time.Time `json:"timestamp"`
	Password     string    `json:"password,omitempty" datastore:"-"`
	HashPassword []byte    `json:"hashpassword,omitempty"`
}

func (user *User) save(c context.Context) error {
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		return err
	}

	//k, err := datastore.Put(c, user.key(c), user)
	_, err = datastore.Put(c, user.key(c), user)
	if err != nil {
		return err
	}
	//user.Id = k.IntID()
	return nil
}

func (user *User) key(c context.Context) *datastore.Key {
	//if user.Id == 0 {
	//	log.Println("tidert")
	//	return datastore.NewIncompleteKey(c, "User", nil)
	//}
	return datastore.NewKey(c, "User", user.Email, 0, nil)
}

func (user *User) ifExist(c context.Context) (*User,bool){
	k := user.key(c)
	var u User
	_ = datastore.Get(c, k, &u)
	if(u.Email != ""){
		return &u, true
	}
	return nil, false
}

func NewUser(c context.Context, r io.ReadCloser) (*User, error) {
	type Resource struct {
		Data User `json:"data"`
	}
	var dataResource Resource
	err := json.NewDecoder(r).Decode(&dataResource)
	if err != nil {
		return nil,err
	}
	user := &dataResource.Data //get data from json object to user struct
	_, ok := user.ifExist(c)
	if(ok){ //if email is already used
		return nil, errors.New("Email already exists")
	}
	user.Timestamp = time.Now()

	hpass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	user.HashPassword = hpass
	//clear the incoming text password
	user.Password = ""
	err = user.save(c)
	if err != nil {
		return nil, err
	}
	return user, nil

}

func LoginUser(c context.Context, r io.ReadCloser) (*User, error) {
	type Resource struct {
		Data User `json:"data"`
	}
	var dataResource Resource
	err := json.NewDecoder(r).Decode(&dataResource)
	if err != nil {
		return nil,err
	}
	user := &dataResource.Data
	u, ok := user.ifExist(c)
	if(ok){ //if user exists in database
		err = bcrypt.CompareHashAndPassword(u.HashPassword, []byte(user.Password))
		if err == nil { //if users password is correct
			return u, nil
		}
		log.Println("wrong password")
		return nil, errors.New("Incorrect Email/Password")
	}else{ //else if user does not exist
		return nil, errors.New("Incorrect Email/Password")
	}
	return user, nil
}

