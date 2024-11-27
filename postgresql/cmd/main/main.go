package main

import (
	"fmt"
	"go_learn/postgresql/pkg/router"
	"log"
	"net/http"
)

func main() {

	r:= router.PostRouter()
	fmt.Printf("Starting server on PORT: 8080")

	log.Fatal(http.ListenAndServe(":8080",r))
}