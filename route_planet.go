package main

import (
	"net/http"
	"github.com/pdeguing/empire-and-foundation/data"
)

func upMetalMine(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
		return
	}
	user, err := sess.User()
	if err != nil {
		info("cannot get user from session: ", err)
	}
	planet, err := data.PlanetByUserId(user.Id)
	if err != nil {
		info("could not get user planet: ", err)
	}
	planet.UpgradeMine()
	http.Redirect(w, r, "/dashboard", 302)
	info("up metal mine")
}
