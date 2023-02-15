package main

import (
	"fmt"
	"net/http"
	"html/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, ", APA LOE!!!!!")
}

func main() {
	port := "4321"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data = map[string]interface{}{
			"Name": "Rifan",
			"Age": 24,
			"Hobbies": []string{"CODING", "JOGGING", "BADMINTON"},
			"Detail": map[string]interface{}{
				"Address": "BEKASI",
				"Phone": "08137437434",
				"Bio": "Explore new technology is my passion.",
			},
		}
		var t, err = template.ParseFiles("client.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		t.Execute(w, data)
		// fmt.Fprintln(w, "HALO!!!!!!!")
	})

	http.HandleFunc("/index", index)

	fmt.Println("Start Web Server at http://localhost:" + port)
	http.ListenAndServe(":" + port, nil)
}