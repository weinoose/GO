package fakebackend

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Execute the command below to deploy json server.
// json-server --watch db.json

type DriversStruct struct {
	Id    int    `json:id`
	Title string `json:title`
	Team  string `json:team`
}

type ManufacturersStruct struct {
	Id    int    `json:id`
	Title string `json:title`
}

func GetDrivers() {
	response, err := http.Get("http://localhost:3000/drivers")

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)

	var ferrari []DriversStruct
	json.Unmarshal(bodyBytes, &ferrari)
	fmt.Println(ferrari)
}

func GetManufacturer() {
	response, err := http.Get("http://localhost:3000/manufacturer")

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)

	var ferrari []ManufacturersStruct
	json.Unmarshal(bodyBytes, &ferrari)
	fmt.Println(ferrari)
}

func GetAll() {
	GetDrivers()
	GetManufacturer()
}

func PutDriver(a int, b string, c string) {
	backendapi := DriversStruct{a, b, c}
	jsonBackendapi, _ := json.Marshal(backendapi)

	response, err := http.Post("http://localhost:3000/drivers", "application/json;charset=utf-8", bytes.NewBuffer(jsonBackendapi))

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body) // Read the API

	bodyString := string(bodyBytes) // Trasform Bytes format into the String type

	// Import the API into the Struct.
	var backendapiResponse DriversStruct
	json.Unmarshal(bodyBytes, &backendapiResponse)
	fmt.Println(backendapiResponse)
	fmt.Println(bodyString)
}

func PutManufacturer(a int, b string) {
	backendapi := ManufacturersStruct{a, b}
	jsonBackendapi, _ := json.Marshal(backendapi)

	response, err := http.Post("http://localhost:3000/manufacturer", "application/json;charset=utf-8", bytes.NewBuffer(jsonBackendapi))

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body) // Read the API

	bodyString := string(bodyBytes) // Trasform Bytes format into the String type

	// Import the API into the Struct.
	var backendapiResponse ManufacturersStruct
	json.Unmarshal(bodyBytes, &backendapiResponse)
	fmt.Println(backendapiResponse)
	fmt.Println(bodyString)
}
