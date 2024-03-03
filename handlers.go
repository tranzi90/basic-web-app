package main

import (
	"fmt"
	"net/http"
)

func Home(writer http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(writer, "WUBBA LUBBA DUBDUB !!!")
}
