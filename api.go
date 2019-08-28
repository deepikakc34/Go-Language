package main

import (
    "fmt"
	"net/http"
)

func page(w http.ResponseWriter, r *http.Request){
	fmt.Println("This is boring")
}

func handlerequest() {
	net/http.Handlefunc("/", page)
}

func main() {
	handlerequest()
}