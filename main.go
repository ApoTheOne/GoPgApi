package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	requestHandler()
}

func requestHandler() {
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login Endpoint!")
	fmt.Println("Hit it... Login!")
}

func register(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Register Endpoint!")
	fmt.Println("Hit it... Register!")
}
