package main

import (
	"fmt"
	"net/http"
)

const port = ":8877"

func main() {
	
	http.HandleFunc("/", Home)

	fmt.Println("(htpp://localhost:8877) - Server started on port ", port)
	http.ListenAndServe(port, nil)
}


