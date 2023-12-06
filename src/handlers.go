package main

import (
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "home")
}
func Startgame(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "startgame")
}
func Character(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "character")
}
func Style(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "character")
}


func RenderTemplate(w http.ResponseWriter, html string) {
	t, err := template.ParseFiles("./templates_HTML/" + html + ".page.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}