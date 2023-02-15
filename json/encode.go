package main

import (
	"fmt"
	"encoding/json"
)

type Detail struct {
	Address string
	Phone string
}

type User struct {
	Name string
	Age int
	Detail Detail
}

func main() {
	/*object := []map[string]interface{}{
		map[string]interface{}{
			"Name": "Pamela",
			"Age": 22,
		},
		map[string]interface{}{
			"Name": "Dewi",
			"Age": 20,
		},
	}*/
	object := []User{{"Pamela", 20, Detail{"Bekasi", "0813273382782"}}, {"Anisa", 18, Detail{"Jakarta", "0858438473847"}}}

	jsonData, err := json.Marshal(object)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(jsonData))
}