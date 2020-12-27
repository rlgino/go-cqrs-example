package controller

import (
	"fmt"
	"net/http"
)

// Run http server
func Run() {
	http.HandleFunc("/user", createUser)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User created")
}
