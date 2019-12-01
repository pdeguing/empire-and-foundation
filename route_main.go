package main

import (
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		info("error when getting session", err)
		generateHTML(w, nil, "layout", "public.navbar", "index")
	} else {
		info("session is valid")
		http.Redirect(w, r, "/dashboard", 302)
	}
}
