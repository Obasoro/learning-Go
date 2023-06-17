package main

import (
	"fmt"
	"log"
	"net/http"
)

// HelloServer responds to request with the given URL patg

func HelloServer (w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Hello, your request: %s", r.URL.Path)
	log.Printf("Recieved request for path: %s", r.URL.Path)
}

func main() {
	var addr string = ":8180"
	handler := http.HandlerFunc(HelloServer)
	if err := http.ListenAndServe(addr, handler); 
	err != nil {
		log.Fatalf("could not listen to port %s %v", addr, err)
	}
}