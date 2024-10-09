package main

import (
	"log"
	"net/http"

	"github.com/IsralTeo/api-go-jwt/route"
)

func main() {

	r := route.InitRouter()
	log.Fatal(http.ListenAndServe(":8080", r))

}
