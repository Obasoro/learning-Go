package main

import (
	"fmt"
	"net/url"
)

func PerformPostJsaonRequest() {

}

func PerformPostFormRequest() {
	const myurl = "hhtp://localhost:8000/postform"

	//formdata

	data := url.Values{}
	data.Add("firstname", "olakunle")
	data.Add("Position", "DevOps/DevSecOps")
	data.Add("Email", "klawobasoro@gmail.com")
	

	// http

	response, err := http.postform(myurl, data)

	if err := nil {
		panic
	}
}