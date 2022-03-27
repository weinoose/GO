package advanced

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type API struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func Get() {
	response, err := http.Get("http://jsonplaceholder.typicode.com/todos/17") // Get the API.

	if err != nil {
		fmt.Println("API does not found!", err) // Check the Error situation.
	}

	defer response.Body.Close() // Close the response whatever happens

	bodyBytes, _ := ioutil.ReadAll(response.Body) // Read the API

	bodyString := string(bodyBytes) // Trasform Bytes format into the String type

	// Import the API into the Struct.
	var api API
	json.Unmarshal(bodyBytes, &api)
	fmt.Println(api)
	fmt.Println(bodyString)
}

func Put(a int, b int, c string, d bool) {
	backendapi := API{a, b, c, d}
	jsonBackendapi, _ := json.Marshal(backendapi)

	response, err := http.Post("https://jsonplaceholder.typicode.com/todos/", "application/json;charset=utf-8", bytes.NewBuffer(jsonBackendapi))

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body) // Read the API

	bodyString := string(bodyBytes) // Trasform Bytes format into the String type

	// Import the API into the Struct.
	var backendapiResponse API
	json.Unmarshal(bodyBytes, &backendapiResponse)
	fmt.Println(backendapiResponse)
	fmt.Println(bodyString)
}
