package main

import (
	"fmt"
	"html/template" // Новый импорт !!!
	"net/http"
)

// Работа с HTML шаблонами - https://youtu.be/p9rBjSXopsw?si=iYOUsR3LU6Leh5wT

type User struct {
	Name      string
	Age       int
	Happiness float64
	Hobbies   []string
}

func homePage(w http.ResponseWriter, r *http.Request) {
	bob := User{Name: "Bob", Age: 25, Happiness: 1.99, Hobbies: []string{"games", "tennis", "skate"}}
	temp, err := template.ParseFiles("lesson_3/templates/home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = temp.Execute(w, bob)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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
