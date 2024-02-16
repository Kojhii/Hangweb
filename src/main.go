package main

import (
	"fmt"
	"net/http"
)

const port = ":8080"

func main() {
	http.Handle("/templates_HTML/", http.StripPrefix("/templates_HTML/", http.FileServer(http.Dir("templates_HTML"))))

	http.HandleFunc("/", Home)
	http.HandleFunc("/startgame",Startgame)
	http.HandleFunc("/character",Character)
	http.HandleFunc("/firsthistory",FirstHistory)
	http.HandleFunc("/character2",Character2)
	http.HandleFunc("/character1",Character1)
	http.HandleFunc("/character15",Character15)
	http.HandleFunc("/introhangman",IntroHangman)
	http.HandleFunc("/firsthangman",FirstHangman)
	http.HandleFunc("/endfirsthangman",EndFirstHangman)
	http.HandleFunc("/endofchapter1",EndOfChapter1)
	http.HandleFunc("/image1to2",Image1to2)
	http.HandleFunc("/image0to1",Image0to1)

	fmt.Println("(http://localhost:8080) - Server started on port ", port)
	http.ListenAndServe(port, nil)
}




