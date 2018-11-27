package main

import (
	"fmt"
	"log"
	"net/http"
)

func printhelloWorld() {
	fmt.Println("hello World, I am starting the web server")
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {

	printhelloWorld()
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
