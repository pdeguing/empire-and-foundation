package main

import (
	"net/http"
)

func err(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		generateHTML(w, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}

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
