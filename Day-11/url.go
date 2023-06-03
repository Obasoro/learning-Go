package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://cnn.com"

func main() {
	fmt.Println("Learning url on Golang")
	fmt.Println(myurl)

	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Host)
	fmt.Println(result.Port())

	qparams := result.Query()

	fmt.Printf("The type of query params are: %T\n", qparams)

	fmt.Println(qparams["coursemate"])

}