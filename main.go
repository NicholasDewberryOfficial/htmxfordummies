package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	var err error
	if err != nil {
		log.Fatalf("Error parsing templates: %v", err)
	}

	http.HandleFunc("/simplegetexample", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("templates/simplegetexample.html")
		if err != nil {
			panic(err)
		}
		t.Execute(w, nil)

	})

	http.HandleFunc("/showexample", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("templates/getchangewithme.html")
		if err != nil {
			panic(err)
		}
		t.Execute(w, nil)
	})

	http.HandleFunc("/simplepostexample", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("templates/simplepostexample.html")
		if err != nil {
			panic(err)
		}
		t.Execute(w, nil)
	})

	fmt.Println("Server running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
