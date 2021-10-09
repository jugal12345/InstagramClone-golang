package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"

	"github.com/jugal12345/go-lang/controllers"
)

func main() {
	r := httprouter.New()
	uc := controllers.ExtraUserController(getSession())
	pc := controllers.ReplicaPostController(getSession())
	r.GET("/user/:id", uc.Getuser)
	r.POST("/user", uc.Make_user)
	r.GET("/post/:id",pc.Get_post)
	r.POST("/post",pc.Make_post)
	r.GET("/posts/users/:id", pc.Get_user_post)
	http.ListenAndServe("localhost:9000", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	return s
}
