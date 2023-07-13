package main

import (
	"log"
	"net/http"
	"simple-golang-auth-api/login"
	"simple-golang-auth-api/logout"
	"simple-golang-auth-api/signup"
)

func main() {
	http.HandleFunc("/signup", signup.SignupHandler)
	http.HandleFunc("/login", login.LoginHandler)
	http.HandleFunc("/logout", logout.LogoutHandler)
	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
