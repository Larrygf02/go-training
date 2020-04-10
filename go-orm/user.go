package main

import (
	"fmt"
	"net/http"
)

func AllUseres(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All users endpoint hit")
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "New user endpint hit")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete user endpoint hit")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update user endpoint hit")
}
