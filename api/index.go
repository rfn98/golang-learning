package main

import (
	"encoding/json"
	"net/http"
	"fmt"
)

type employee struct {
	Nip  	string
	Name 	string
	Age 	int
}

func main() {
	http.HandleFunc("/users", users)
	http.HandleFunc("/user", user)

	fmt.Println("Starting At PORT 4321")
	http.ListenAndServe(":4321", nil)
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := []employee{
		{"A01", "Pamela", 22},
		{"A02", "Selvy", 20},
		{"A03", "Novita", 21},
	}

	if r.Method == "GET" {
		result, err := json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := []employee{
		{"A01", "Pamela", 22},
		{"A02", "Selvy", 20},
		{"A03", "Novita", 21},
	}

	if r.Method == "GET" {
		id := r.FormValue("id")
		for _, row := range data {
			if row.Nip == id {
				result, err := json.Marshal(row)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				w.Write(result)
				return
			} /*else {
				w.Write([]byte("Employee Not Found"))
				return
			}*/
		}
	}

	message, _ := json.Marshal(map[string]string{"msg": "Employee not found."})

	w.Write(message)

	// http.Error(w, "", http.StatusBadRequest)
}