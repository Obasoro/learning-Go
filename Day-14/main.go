package main

import (
	"encoding/json"
	"fmt"

)

type course struct {
	Name      string `json:"kunletraining"`
	Price     int
	platrform string `json:"website"`
	password  string `json:"-"`
	Tags      []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to Json in Golang")
	EncodeJson()
	decodeJson()

}

func EncodeJson() {
	kunleTraining := []course{
		{"ReactJs Bootcamps", 300, "LearningCodeOnline.in", "tunde123", []string{"web-dev", "React-dev"}},
		{"Golang Bootcamps", 600, "LearningCodeOnline.in", "tunde123", []string{"web-dev", "Go-dev"}},
		{"Python Bootcamps", 250, "LearningCodeOnline.in", "Gbenga123", []string{"web-dev", "Python-dev"}},

	}
		// package this data in json format

		finalJson, _ := json.Marshal(kunleTraining, " ", "\t")
		if err != nil {
			panic(err)
		}
		fmt.Println("%s\n", finalJson)

	
}

func decodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "Golang",
		"price":300
		"website": "kunlelearn.com"
		"tag": ["web-dev,"js"],
	}

	`)

	var kunleTraining course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFromWeb, &kunleTraining)
		fmt.Println("%#v\n", kunleTraining)
	}else {
		fmt.Println("Json is not valid")
	}

	// some cases where you just want to add data
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", kunleTraining)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is %T")
	}

}