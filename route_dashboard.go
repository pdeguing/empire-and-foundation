package main

import (
	"github.com/pdeguing/empire-and-foundation/data"
	"net/http"
)

// GET /dashboard
// Show player main dashboard page
func dashboard(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
		return
	}
	user, err := sess.User()
	if err != nil {
		info("cannot get user from session")
	}
	planet, err := data.PlanetByUserId(user.Id)
	if err != nil {
		info("could not get user planet", err)
		if planet, err = user.CreatePlanet(); err != nil {
			danger("could not create user planet:", err)
		}
	}
	newStock := planet.GetMetalStock()
	info("updated metal stock to :", newStock)
	generateHTML(w, planet, "layout", "private.navbar", "dashboard", "rightbar")
}
