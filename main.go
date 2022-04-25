package main

import (
	"fmt"
	"log"
	"net/http"
	"social-network/src/config"
	"social-network/src/router"
)

func main() {
	config.Load()
	r := router.Generate()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
