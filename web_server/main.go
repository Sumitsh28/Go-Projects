package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

 if r.Method != "GET" {
	http.Error(w, "Not found", http.StatusNotFound)
	return
 }

 fmt.Fprintf(w, "HELLO!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w,"ParseForm() err: %v",err)
		return
	}

	fmt.Fprintf(w,"Post successful!")
	name:= r.FormValue("name")
	fmt.Fprintf(w, "NAME: %s\n",name)
}


func main(){

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Printf("Server starting on PORT 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil{
		log.Fatal(err)
	}
}