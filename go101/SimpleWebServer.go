package main

import (
	"net/http"
	"log"
	"fmt"
	"strings"
	"html/template"
	"net"
)

func main() {

	http.HandleFunc("/", sayHello);
	http.HandleFunc("/login", login);


	var err = http.ListenAndServe(":8080", nil);

	if err != nil {
		log.Fatal(err);
	}
}

func login(w http.ResponseWriter, r *http.Request)  {


	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	}

	if r.Method == "POST" {

		r.ParseForm();

		fmt.Println(r.Form["username"]);
		fmt.Println(r.Form["password"]);

	}

}

func sayHello(res http.ResponseWriter, req *http.Request)  {

	req.ParseForm();
	fmt.Println(req.Form);
	fmt.Println(req.URL.Path);
	fmt.Println(req.URL.Scheme);

	fmt.Println(req.Form["url_long"])

	for k, v := range req.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	fmt.Fprintf(res, "Hello astaxie!") // send data to client side
}
