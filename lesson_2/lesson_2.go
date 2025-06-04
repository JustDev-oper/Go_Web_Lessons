package main

import (
	"fmt"
	"net/http"
)

// Создание структур (модели данных) - https://youtu.be/n18K8kCaXyg?si=Ie7_g1HXqcrjUD18

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
