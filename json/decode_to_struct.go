package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FullName string `json:"Name"` //Property Json Name: Name
	Age int // Default Property Json Name: Age
}

func main() {
	jsonString := `[{"Name": "Rifan", "Age": 24}, {"Name": "Novita", "Age": 22}]`
	// jsonString := `{"Name": "Rifan", "Age": 24}`
	jsonData := []byte(jsonString) // CONVERT/CASTING TO BYTE BEFORE DECODE

	var data []User // VARIABEL PENAMPUNG RESULT DATA JSON

	err := json.Unmarshal(jsonData, &data) //PARSE TO STRUCT USER (DECODE)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for idx, user := range data {
		fmt.Printf("USER %d: " + user.FullName + "\n", idx + 1)
	}

	// fmt.Println("USER 1: ", data[0].FullName)
	// fmt.Println("USER 2: ", data[1].FullName)
	// fmt.Println("AGE: ", data.Age)
}