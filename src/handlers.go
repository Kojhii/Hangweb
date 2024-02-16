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
func FirstHistory(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "firsthistory")
}
func Character2(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "character2")
}
func Character1(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "character1")
}
func Character15(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "character15")
}
func IntroHangman(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "introhangman")
}
func FirstHangman(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "firsthangman")
}
func EndFirstHangman(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "endfirsthangman")
}
func EndOfChapter1(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "endofchapter1",)
}
func Image1to2(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "image1to2",)
}
func Image0to1(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "image0to1",)
}




func RenderTemplate(w http.ResponseWriter, html string) {
	t, err := template.ParseFiles("./templates_HTML/" + html + ".page.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}