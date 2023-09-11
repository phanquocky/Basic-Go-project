package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	// if r.Method != "GET" {
	// 	http.Error(w, "404 not found", http.StatusNotFound)
	// 	return
	// }

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parse Form (): %v", err)
	}

	fmt.Fprintf(w, "Post Request Success!")

	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "name: %s, address: %s", name, address)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
