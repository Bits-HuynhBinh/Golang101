package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"fmt"
	"log"
)

func main()  {

	router := httprouter.New();

	router.GET("/", Index);
	router.POST("/ADD/:uid", AddUser)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func AddUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// uid := r.FormValue("uid")
	uid := ps.ByName("uid")

	r.ParseForm();

	username := r.Form["username"];
	pass := r.Form["password"];

	fmt.Fprintf(w, "you are add user %s, %s, %s", uid, username, pass);
}
