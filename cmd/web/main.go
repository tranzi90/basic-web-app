package main

import (
	"fmt"
	"log"
	"net/http"
	"webApp/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)

	fmt.Println(fmt.Sprintf("Starting app on port %s", portNumber[1:]))

	log.Fatal(http.ListenAndServe(portNumber, nil))
}
