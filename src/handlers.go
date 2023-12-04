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


func RenderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("./templates_HTML/" + tmpl + ".page.tmpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}