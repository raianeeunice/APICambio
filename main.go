package main

import (
	"api/src/config"
	"api/src/router"
	"log"
	"net/http"
)

func main() {
	config.Carregar()
	r := router.Gerar()

	log.Fatal(http.ListenAndServe(":5000", r))
}