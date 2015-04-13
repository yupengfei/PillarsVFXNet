package main

import (
	"log"
	"net/http"
)

func main() {
	RouterBinding()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
