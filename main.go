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

	// Those routes are temporary and should be adapted to handle multiple planets per user.
	mux.HandleFunc("/dashboard", dashboard)
	mux.HandleFunc("/dashboard/cartography", cartography)
	mux.HandleFunc("/dashboard/fleetcontrol", fleetcontrol)
	mux.HandleFunc("/dashboard/technology", technology)
	mux.HandleFunc("/dashboard/diplomacy", diplomacy)
	mux.HandleFunc("/dashboard/story", story)
	mux.HandleFunc("/dashboard/wiki", wiki)
	mux.HandleFunc("/dashboard/news", news)

	mux.HandleFunc("/dashboard/planet", planet)
	mux.HandleFunc("/dashboard/planet/constructions", constructions)
	mux.HandleFunc("/dashboard/planet/factories", factories)
	mux.HandleFunc("/dashboard/planet/research", research)
	mux.HandleFunc("/dashboard/planet/fleets", fleets)
	mux.HandleFunc("/dashboard/planet/defenses", defenses)

	mux.HandleFunc("/planet/up_metal_mine", upMetalMine)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	info("Server started")
	server.ListenAndServe()
}
