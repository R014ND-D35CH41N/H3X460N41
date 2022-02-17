package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("[MESSAGE]: Starting server on port 8080...")
	router := rest.InitHandlers()
	log.Fatal(http.ListenAndServe(":8080", router))
}
