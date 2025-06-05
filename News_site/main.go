package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("News_site/templates/index.html", "News_site/templates/header.html", "News_site/templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	temp.ExecuteTemplate(w, "index", nil)
}

func handleFunc() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleFunc()
}
