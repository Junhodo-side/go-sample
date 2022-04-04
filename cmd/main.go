package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/", handlerFunc)
	log.Fatal(http.ListenAndServe("0.0.0.0:9090", nil))
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	log.Println("hi my server invoked!")
	fmt.Fprintln(w, "Hello, WebServer is listening on  9090")
}