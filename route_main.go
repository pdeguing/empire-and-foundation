package main

import (
	"net/http"
)

// GET /
// Show the frontpage or redirect the user to their dashboard
func index(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		info("error when getting session", err)
		generateHTML(w, r, nil, "layout", "public.navbar", "index")
	} else {
		info("session is valid")
		http.Redirect(w, r, "/dashboard", 302)
	}
}
