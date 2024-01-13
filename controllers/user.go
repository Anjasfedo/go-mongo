package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Anjasfedo/go-mongo/models" // Importing the models package which contains user model
	"github.com/julienschmidt/httprouter"  // Importing the httprouter package for routing
	"gopkg.in/mgo.v2"                      // Importing the mgo package for MongoDB interaction
	"gopkg.in/mgo.v2/bson"                 // Importing bson package for working with BSON data
)

// UserController represents the controller for user-related operations
type UserController struct {
	session *mgo.Session // MongoDB session
}

// NewUserController creates and returns a new UserController with the provided MongoDB session
func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

// GetUserById handles GET requests to retrieve a user by ID
func (uc UserController) GetUserById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ID := p.ByName("id")

	// Checking if the provided ID is a valid ObjectIdHex
	if !bson.IsObjectIdHex(ID) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	oid := bson.ObjectIdHex(ID)

	user := models.User{}

	// Querying the database to find the user by ObjectId
	if err := uc.session.DB("go-mongo").C("users").FindId(oid).One(&user); err != nil {
		w.WriteHeader(404)
		return
	}

	// Marshaling user object to JSON
	userJson, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}

	// Setting response headers and writing the JSON response
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", userJson)
}

// CreateUser handles POST requests to create a new user
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	user := models.User{}

	// Decoding the JSON request body into the user object
	json.NewDecoder(r.Body).Decode(&user)

	// Generating a new ObjectId for the user
	user.ID = bson.NewObjectId()

	// Inserting the user into the "users" collection in the database
	uc.session.DB("go-mongo").C("users").Insert(user)

	// Marshaling user object to JSON
	userJson, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}

	// Setting response headers and writing the JSON response
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", userJson)
}

// DeleteUserById handles DELETE requests to delete a user by ID
func (uc UserController) DeleteUserById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ID := p.ByName("id")

	// Checking if the provided ID is a valid ObjectIdHex
	if !bson.IsObjectIdHex(ID) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	oid := bson.ObjectIdHex(ID)

	// Removing the user from the "users" collection in the database
	if err := uc.session.DB("go-mongo").C("users").RemoveId(oid); err != nil {
		w.WriteHeader(404)
		return
	}

	// Setting response status and writing a success message
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Deleted User with ID: %s\n", ID)
}

// GetUsers handles GET requests to retrieve all users
func (uc UserController) GetUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var users []models.User

	// Querying the database to retrieve all users from the "users" collection
	if err := uc.session.DB("go-mongo").C("users").Find(bson.M{}).All(&users); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Setting response headers and status
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusOK)

	// Iterating over users and writing JSON responses
	for _, user := range users {
		userJSON, err := json.Marshal(user)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "%s,\n", userJSON)
	}
}

// UpdateUserById handles PUT requests to update a user by ID
func (uc UserController) UpdateUserById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ID := p.ByName("id")

	// Checking if the provided ID is a valid ObjectIdHex
	if !bson.IsObjectIdHex(ID) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	oid := bson.ObjectIdHex(ID)

	user := models.User{}

	// Decoding the JSON request body into the user object
	json.NewDecoder(r.Body).Decode(&user)

	user.ID = oid

	// Creating an update document for the user
	updatedUser := bson.M{"$set": bson.M{
		"name":   user.Name,
		"gender": user.Gender,
		"age":    user.Age,
	}}

	// Updating the user in the "users" collection in the database
	if err := uc.session.DB("go-mongo").C("users").UpdateId(oid, updatedUser); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Marshaling user object to JSON
	userJson, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}

	// Setting response headers and writing the JSON response
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", userJson)
}
