package main

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pdeguing/empire-and-foundation/ent"
)

// Helper
func userPlanet(w http.ResponseWriter, r *http.Request) (*ent.Planet, bool) {
	n, err := strconv.Atoi(mux.Vars(r)["planetNumber"])
	if err != nil {
		serveInternalServerError(w, r, err, "Cannot parse url variable 'planetNumber'")
		return nil, false
	}
	u := loggedInUser(r)
	p, err := u.QueryPlanets().
		Offset(n - 1).
		First(r.Context())
	if _, ok := err.(*ent.ErrNotFound); ok {
		serveNotFoundError(w, r)
		return nil, false
	}
	if err != nil {
		serveInternalServerError(w, r, err, "Could not retrieve user's planet from database")
		return nil, false
	}
	return p, true
}

// GET /dashboard/planet
// Show the dashboard page for a planet
func servePlanet(w http.ResponseWriter, r *http.Request) {
	if p, ok := userPlanet(w, r); ok {
		generateHTML(w, r, "planet-dashboard", p, "layout", "private.navbar", "dashboard", "leftbar", "planet.layout", "planet.header", "planet.overview")
	}
}

// GET /dashboard/planet/constructions
// Show the constructions page for a planet
func serveConstructions(w http.ResponseWriter, r *http.Request) {
	if p, ok := userPlanet(w, r); ok {
		generateHTML(w, r, "planet-constructions", p, "layout", "private.navbar", "dashboard", "leftbar", "planet.layout", "planet.header", "planet.constructions")
	}
}

// GET /dashboard/planet/factories
// Show the factories page for a planet
func serveFactories(w http.ResponseWriter, r *http.Request) {
	if p, ok := userPlanet(w, r); ok {
		generateHTML(w, r, "planet-factories", p, "layout", "private.navbar", "dashboard", "leftbar", "planet.layout", "planet.header", "planet.factories")
	}
}

// GET /dashboard/planet/research
// Show the research page for a planet
func serveResearch(w http.ResponseWriter, r *http.Request) {
	if p, ok := userPlanet(w, r); ok {
		generateHTML(w, r, "planet-research", p, "layout", "private.navbar", "dashboard", "leftbar", "planet.layout", "planet.header", "planet.research")
	}
}

// GET /dashboard/planet/fleets
// Show the fleets page for a planet
func serveFleets(w http.ResponseWriter, r *http.Request) {
	if p, ok := userPlanet(w, r); ok {
		generateHTML(w, r, "planet-fleets", p, "layout", "private.navbar", "dashboard", "leftbar", "planet.layout", "planet.header", "planet.fleets")
	}
}

// GET /dashboard/planet/defenses
// Show the defenses page for a planet
func serveDefenses(w http.ResponseWriter, r *http.Request) {
	if p, ok := userPlanet(w, r); ok {
		generateHTML(w, r, "planet-defenses", p, "layout", "private.navbar", "dashboard", "leftbar", "planet.layout", "planet.header", "planet.defenses")
	}
}

// Temporarily deprecated?
func serveUpMetalMine(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/dashboard/planet/constructions", 302)
	info("up metal mine")
}
