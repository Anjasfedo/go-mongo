package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Anjasfedo/go-mongo/controllers" // Importing the controllers package which contains user controllers
	"github.com/julienschmidt/httprouter"       // Importing the httprouter package for routing
	"gopkg.in/mgo.v2"                           // Importing the mgo package for MongoDB interaction
)

func main() {
	r := httprouter.New() // Creating a new httprouter instance

	uc := controllers.NewUserController(getSession()) // Creating a new UserController with a MongoDB session

	r.GET("/user", uc.GetUsers)              // Route to handle GET requests for fetching all users
	r.GET("/user/:id", uc.GetUserById)       // Route to handle GET requests for fetching a user by ID
	r.POST("/user", uc.CreateUser)           // Route to handle POST requests for creating a new user
	r.PUT("/user/:id", uc.UpdateUserById)    // Route to handle PUT requests for updating a user by ID
	r.DELETE("/user/:id", uc.DeleteUserById) // Route to handle DELETE requests for deleting a user by ID

	fmt.Printf("Start Server on Port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r)) // Starting the server on port 8000
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017") // Establishing a MongoDB session
	if err != nil {
		panic(err) // Panic if there is an error establishing the session
	}

	return s // Returning the MongoDB session
}
