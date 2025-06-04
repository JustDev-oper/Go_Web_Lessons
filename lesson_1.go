package main

import (
	"fmt"
	"net/http"
)

// Отслеживание URL адресов - https://youtu.be/c6lBYNqrxRk?si=HsdULLKmR6m04Mz2

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About Page")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/about/", aboutPage)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()
}
