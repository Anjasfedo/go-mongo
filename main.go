package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Anjasfedo/go-mongo/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()

	uc := controllers.NewUserController(getSession())

	r.GET("/user", uc.GetUsers)

	r.GET("/user/:id", uc.GetUserById)

	r.POST("/user", uc.CreateUser)

	r.DELETE("/user/:id", uc.DeleteUserById)

	fmt.Printf("Start Server on Port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	return s
}
