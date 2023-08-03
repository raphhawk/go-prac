package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNoContent)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNoContent)
		return
	}
	fmt.Fprintf(w, "Hello!\n")
}
func formHandler(w http.ResponseWriter, r *http.Request) {
	//	if r.URL.Path != "/form" {
	//		http.Error(w, "404 not found", http.StatusNotFound)
	//		return
	//	}
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Error parsing form: %v\n", err)
		return
	}
	fmt.Fprintf(w, "POST request successfull\n")
	name, address := r.FormValue("name"), r.FormValue("address")
	fmt.Fprintf(w, "Name: %v\tAddress: %v\n", name, address)
}

func main() {
	server := http.FileServer(http.Dir("./static"))
	http.Handle("/", server)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Println("Starting server at port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error Listening to the server %v\n", err)
	}

}
