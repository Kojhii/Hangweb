package main

import (
	"fmt"
	"net/http"
)

const port = ":8080"

func main() {
	
	http.HandleFunc("/", Home)
	http.HandleFunc("/startgame",Startgame)
	http.HandleFunc("/character",Character)
	http.HandleFunc("/firsthistory",FirstHistory)
	http.HandleFunc("/character2",Character2)
	fs := http.FileServer(http.Dir("templates_HTML"))
	http.Handle("/templates_HTML", http.StripPrefix("/templates_HTML/", fs))

	fmt.Println("(htpp://localhost:8080) - Server started on port ", port)
	http.ListenAndServe(port, nil)
}




