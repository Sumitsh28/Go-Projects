package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"strings"
)

func hello(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hello %q \n",strings.Split(html.EscapeString(r.URL.Path), "/")[1])

}

func main(){

	http.HandleFunc("/",hello)
	log.Fatal(http.ListenAndServe(":8080",nil))
}