package main

import (
	"log"
	"net/http"
	"social-network/src/router"
)

func main() {
	r := router.Generate()
	log.Fatal(http.ListenAndServe(":5000", r))
}
