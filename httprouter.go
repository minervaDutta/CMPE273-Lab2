/*
*
* CMPE273 Lab2 - Golang, Hello World using REST API
* Minerva Dutta; 009678123
*
*/

package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"encoding/json"
	)

type Username struct {
	Name string `json:"name"`
}

type Message struct {
	Greet string `json:"greeting"`
}

func hello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
}

//To handle POST requests
func postHandler(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	var usr Username
	json.NewDecoder(req.Body).Decode(&usr)
	var msg Message
	msg.Greet = "Hello, " + usr.Name +"!"
	j, _ := json.Marshal(msg)

	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(rw, "%s", j)
}

//main`
func main() {
	mux := httprouter.New()
	mux.GET("/hello/:name", hello)
	mux.POST("/hello", postHandler)
	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
