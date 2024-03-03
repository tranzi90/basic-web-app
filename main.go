package main

import (
	"fmt"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", Home)

	fmt.Println(fmt.Sprintf("Starting app on port %s", portNumber[1:]))

	log.Fatal(http.ListenAndServe(portNumber, nil))
}
