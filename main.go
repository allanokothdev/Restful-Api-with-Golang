package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type Article struct {
	Title string `json:"Title"`
	Desc string	`json:"desc"`
	Content string	`json:"content"`
}

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Homepage!")
	fmt.Println("Endpoint Hit: homepage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}