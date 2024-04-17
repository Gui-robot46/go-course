package main

import (
	"fmt"
	"net/http"
)

const portNumber string = ":8080"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	fmt.Printf(fmt.Sprintf("Aplicação escutando na porta: %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
