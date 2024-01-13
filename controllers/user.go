package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Anjasfedo/go-mongo/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func (uc UserController) GetUserById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ID := p.ByName("id")

	if !bson.IsObjectIdHex(ID) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	userID := bson.ObjectIdHex(ID)

	user := models.User{}

	if err := uc.session.DB("go-mongo").C("users").FindId(userID).One(&user); err != nil {
		w.WriteHeader(404)
		return
	}

	userJson, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", userJson)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	user := models.User{}

	json.NewDecoder(r.Body).Decode(&user)

	user.ID = bson.NewObjectId()

	uc.session.DB("go-mongo").C("users").Insert(user)

	userJson, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", userJson)
}

func (uc UserController) DeleteUserById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ID := p.ByName("id")

	if !bson.IsObjectIdHex(ID) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	userID := bson.ObjectIdHex(ID)

	if err := uc.session.DB("go-mongo").C("users").RemoveId(userID); err != nil {
		w.WriteHeader(404)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Deleted user %s\n", userID)
}
