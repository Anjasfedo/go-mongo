package main

import (
	"net/http"

	"github.com/Anjasfedo/go-mongo/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()

	uc := controllers.NewUserController(getSession())

	r.GET("/user/:id", uc.GetUserById)

	r.POST("/user", uc.CreateUser)

	r.DELETE("/user/:id", uc.DeleteUserById)

	http.ListenAndServe(":8000", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27107")
	if err != nil {
		panic(err)
	}

	return s
}
