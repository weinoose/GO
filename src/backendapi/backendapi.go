package backendapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Backendapi struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func Get() {
	response, err := http.Get("http://jsonplaceholder.typicode.com/todos/2") // Get the API.

	if err != nil {
		fmt.Println("API does not found!", err) // Check the Error situation.
	}

	defer response.Body.Close() // Close the response whatever happens

	bodyBytes, _ := ioutil.ReadAll(response.Body) // Read the API

	bodyString := string(bodyBytes) // Trasform Bytes format into the String type

	// Import the API into the Struct.
	var backendapi Backendapi
	json.Unmarshal(bodyBytes, &backendapi)
	fmt.Println(backendapi)
	fmt.Println(bodyString)
}

func Put(a int, b int, c string, d bool) {
	backendapi := Backendapi{a, b, c, d}
	jsonBackendapi, _ := json.Marshal(backendapi)

	response, err := http.Post("https://jsonplaceholder.typicode.com/todos/", "application/json;charset=utf-8", bytes.NewBuffer(jsonBackendapi))

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body) // Read the API

	bodyString := string(bodyBytes) // Trasform Bytes format into the String type

	// Import the API into the Struct.
	var backendapiResponse Backendapi
	json.Unmarshal(bodyBytes, &backendapiResponse)
	fmt.Println(backendapiResponse)
	fmt.Println(bodyString)
}
