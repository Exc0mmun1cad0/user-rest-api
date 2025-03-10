package main

import (
	"net/http"
)

const (
	addr = ":8080" // TODO: move to config
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /users", CreateUserHandler)
	mux.HandleFunc("GET /user/{id}", GetUserHandler)
	mux.HandleFunc("PATCH /user/{id}", UpdateUserHandler)
	mux.HandleFunc("DELETE /user/{id}", DeleteUserHandler)

	srv := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	srv.ListenAndServe()
}
