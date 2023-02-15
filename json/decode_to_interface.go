package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonString := `[{"Name": "Rifan", "Age": 24}, {"Name": "Novita", "Age": 22}]`
	jsonData := []byte(jsonString) // CONVERT/CASTING TO BYTE BEFORE DECODE

	var data []map[string]interface{} // VARIABEL PENAMPUNG RESULT DATA JSON

	err := json.Unmarshal(jsonData, &data) //PARSE TO STRUCT USER (DECODE)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, user := range data {
		fmt.Println("USER: ", user["Name"])
	}

	// fmt.Println("USER: ", data["Name"])
	// fmt.Println("AGE: ", data["Age"])
}