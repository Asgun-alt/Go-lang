package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
)

type MyMux struct{}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		helloWorld(w, r)
		return
	}

	http.NotFound(w, r)

}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // parse arguments, you have to call this by yourselft
	fmt.Println(r.Form) // print form information on the server side
	fmt.Println("path: ", r.URL.Path)
	fmt.Println("scheme: ", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for key, val := range r.Form {
		fmt.Println("key: ", key)
		fmt.Println("value: ", strings.Join(val, ""))
	}
	fmt.Fprintf(w, "Hello world!") // send data to client side
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method) // get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()

		// logic part of login?
		fmt.Println("username: ", r.Form["username"])
		fmt.Println("password: ", r.Form["password"])
	}
}

func main() {
	http.HandleFunc("/", helloWorld) // setting router rule
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Listen and Serve: ", err)
	}

	// mux := &MyMux{}
	// http.ListenAndServe(":8080", mux)

	// http.HandleFunc("/", helloWorld)         // set router
	// err := http.ListenAndServe(":8080", nil) // set listen port
	// if err != nil {
	// 	log.Fatal("Listen and Serve: ", err)
	// }
}
