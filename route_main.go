package main

import (
	"net/http"
)

// GET /
// Show the frontpage or redirect the user to their dashboard
func serveIndex(w http.ResponseWriter, r *http.Request) {
	if isAuthenticated(r) {
		http.Redirect(w, r, "/dashboard", 302)
		return
	}
	generateHTML(w, r, "frontpage", nil, "layout", "public.navbar", "index", "flash")
}
