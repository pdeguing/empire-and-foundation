package main

import (
	"net/http"
)

func upMetalMine(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/dashboard/planet/constructions", 302)
	info("up metal mine")
}
