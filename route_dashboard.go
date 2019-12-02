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
	generateHTML(w, planet, "layout", "private.navbar", "dashboard", "rightbar", "empire")
}

// GET /dashboard/cartography
// Show map page
func cartography(w http.ResponseWriter, r *http.Request) {
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
	generateHTML(w, planet, "layout", "private.navbar", "dashboard", "rightbar", "cartography")
}

// GET /dashboard/fleetcontrol
func fleetcontrol(w http.ResponseWriter, r *http.Request) {
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
	generateHTML(w, planet, "layout", "private.navbar", "dashboard", "rightbar", "fleetcontrol")
}

// GET /dashboard/technology
func technology(w http.ResponseWriter, r *http.Request) {
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
	generateHTML(w, planet, "layout", "private.navbar", "dashboard", "rightbar", "technology")
}

// GET /dashboard/diplomacy
func diplomacy(w http.ResponseWriter, r *http.Request) {
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
	generateHTML(w, planet, "layout", "private.navbar", "dashboard", "rightbar", "diplomacy")
}

// GET /dashboard/story
func story(w http.ResponseWriter, r *http.Request) {
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
	generateHTML(w, planet, "layout", "private.navbar", "dashboard", "rightbar", "story")
}

// GET /dashboard/wiki
func wiki(w http.ResponseWriter, r *http.Request) {
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
	generateHTML(w, planet, "layout", "private.navbar", "dashboard", "rightbar", "wiki")
}

// GET /dashboard/news
func news(w http.ResponseWriter, r *http.Request) {
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
	generateHTML(w, planet, "layout", "private.navbar", "dashboard", "rightbar", "news")
}

// GET /dashboard/planet
// Show the constructions page for a planet
func planet(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
		return
	}
	user, err := sess.User()
	if err != nil {
		danger("cannot get user from session")
	}
	planet, err := data.PlanetByUserId(user.Id)
	if err != nil {
		info("could not get user planet", err)
	}
	generateHTML(w, planet, "layout", "private.navbar", "dashboard", "rightbar", "planet.layout", "planet.header", "planet.overview")
}

// GET /dashboard/planet/constructions
// Show the constructions page for a planet
func constructions(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
		return
	}
	user, err := sess.User()
	if err != nil {
		danger("cannot get user from session")
	}
	planet, err := data.PlanetByUserId(user.Id)
	if err != nil {
		info("could not get user planet", err)
	}
	generateHTML(w, planet, "layout", "private.navbar", "dashboard", "rightbar", "planet.layout", "planet.header", "planet.constructions")
}

// GET /dashboard/planet/factories
// Show the factories page for a planet
func factories(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
		return
	}
	user, err := sess.User()
	if err != nil {
		danger("cannot get user from session")
	}
	planet, err := data.PlanetByUserId(user.Id)
	if err != nil {
		info("could not get user planet", err)
	}
	generateHTML(w, planet, "layout", "private.navbar", "dashboard", "rightbar", "planet.layout", "planet.header", "planet.factories")
}

// GET /dashboard/planet/research
// Show the research page for a planet
func research(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
		return
	}
	user, err := sess.User()
	if err != nil {
		danger("cannot get user from session")
	}
	planet, err := data.PlanetByUserId(user.Id)
	if err != nil {
		info("could not get user planet", err)
	}
	generateHTML(w, planet, "layout", "private.navbar", "dashboard", "rightbar", "planet.layout", "planet.header", "planet.research")
}
// GET /dashboard/planet/fleets
// Show the fleets page for a planet
func fleets(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
		return
	}
	user, err := sess.User()
	if err != nil {
		danger("cannot get user from session")
	}
	planet, err := data.PlanetByUserId(user.Id)
	if err != nil {
		info("could not get user planet", err)
	}
	generateHTML(w, planet, "layout", "private.navbar", "dashboard", "rightbar", "planet.layout", "planet.header", "planet.fleets")
}

// GET /dashboard/planet/defenses
// Show the defenses page for a planet
func defenses(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
		return
	}
	user, err := sess.User()
	if err != nil {
		danger("cannot get user from session")
	}
	planet, err := data.PlanetByUserId(user.Id)
	if err != nil {
		info("could not get user planet", err)
	}
	generateHTML(w, planet, "layout", "private.navbar", "dashboard", "rightbar", "planet.layout", "planet.header", "planet.defenses")
}
