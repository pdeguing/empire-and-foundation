package main

import (
	"net/http"

	"github.com/pdeguing/empire-and-foundation/data"
)

func upMetalMine(w http.ResponseWriter, r *http.Request) {
	user := user(r)
	planet, err := data.PlanetByUserId(user.Id)
	if err != nil {
		info("could not get user planet: ", err)
	}
	planet.UpgradeMine()
	http.Redirect(w, r, "/dashboard/planet/constructions", 302)
	info("up metal mine")
}
