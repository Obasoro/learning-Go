package main

import (
	"fmt"
	"net/url"
)

const url = "https://insightsforliving.com"

func main() {
	fmt.Println("web request on Golang")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Println("response is of type: %T\n", response)

	defer response.Body.Close()

	databyte, err := ioutil.ReadAll(response.body)

	if err != nil {
		panic(err)
	}

	content := string(databyte)

	fmt.Println(content)
}