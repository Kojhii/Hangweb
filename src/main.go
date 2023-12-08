package main

import (
	"fmt"
	"net/http"
)

const port = ":8877"

func main() {
	
	http.HandleFunc("/", Home)
	http.HandleFunc("/startgame",Startgame)
	http.HandleFunc("/character",Character)
	
	fs := http.FileServer(http.Dir("templates_HTML"))
	http.Handle("/templates_HTML/", http.StripPrefix("/templates_HTML/", fs))

	fmt.Println("(htpp://localhost:8877) - Server started on port ", port)
	http.ListenAndServe(port, nil)
}




