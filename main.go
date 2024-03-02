package main

import (
	"fmt"
	"log"
	"net/http"
)

func Home(writer http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintf(writer, "WUBBA LUBBA DUBDUB")
	if err != nil {
		println(err)
	}
}

func main() {
	http.HandleFunc("/", Home)

	//err := http.ListenAndServe(":8080", nil)
	//
	//println(err.Error())

	log.Fatal(http.ListenAndServe(":8080", nil))
}
