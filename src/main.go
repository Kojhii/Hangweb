package main

import (
	"fmt"
	"net/http"
)

const port = ":8887"

func main() {
	
	http.HandleFunc("/", Home)
	http.HandleFunc("/startgame",Startgame)

	fmt.Println("(htpp://localhost:8877) - Server started on port ", port)
	http.ListenAndServe(port, nil)
}




