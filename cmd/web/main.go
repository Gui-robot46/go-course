package main

import (
	"fmt"
	"net/http"

	handler "hello-world/pkg/handlers"
)

const portNumber string = ":8080"

func main() {
	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handler.About)
	fmt.Printf(fmt.Sprintf("Aplicação escutando na porta: %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
