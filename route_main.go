package main

import (
	"net/http"
)

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

func invalidCsrfToken(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, r, "It's not possible to do this right now. Please go back, reload, and try again.", 403)
}
