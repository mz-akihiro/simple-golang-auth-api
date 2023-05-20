package main

import (
	"study-go-register-mysql/login"
	"study-go-register-mysql/logout"
	"study-go-register-mysql/signup"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/signup", signup.SignupHandler)
	http.HandleFunc("/login", login.LoginHandler)
	http.HandleFunc("/logout", logout.LogoutHandler)
	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
