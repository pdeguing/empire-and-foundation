package main

import (
	"net/http"
)

func main() {
	info("Starting server...")
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)
	mux.HandleFunc("/err", err)

	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)

	mux.HandleFunc("/dashboard", dashboard)

	mux.HandleFunc("/planet/up_metal_mine", upMetalMine)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	info("Server started")
	server.ListenAndServe()
}
