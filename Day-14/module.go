package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello world in golang")
}

func greeter() {
	fmt.Println("Hey ther mod users")
}

func serveHhome(w http.ResponseWriter, r *http.Request) {
	w.Write[]byte("<h1>Welcome to golang<")
}