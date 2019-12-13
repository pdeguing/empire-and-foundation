package main

import (
	"net/http"
)

// GET /dashboard
// Show player main dashboard page
func dashboard(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, r, "dashboard", nil, "layout", "private.navbar", "dashboard", "rightbar", "empire")
}

// GET /dashboard/cartography
// Show map page
func cartography(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, r, "cartography", nil, "layout", "private.navbar", "dashboard", "rightbar", "cartography")
}

// GET /dashboard/fleetcontrol
func fleetcontrol(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, r, "fleet", nil, "layout", "private.navbar", "dashboard", "rightbar", "fleetcontrol")
}

// GET /dashboard/technology
func technology(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, r, "technology", nil, "layout", "private.navbar", "dashboard", "rightbar", "technology")
}

// GET /dashboard/diplomacy
func diplomacy(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, r, "diplomacy", nil, "layout", "private.navbar", "dashboard", "rightbar", "diplomacy")
}

// GET /dashboard/story
func story(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, r, "story", nil, "layout", "private.navbar", "dashboard", "rightbar", "story")
}

// GET /dashboard/wiki
func wiki(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, r, "wiki", nil, "layout", "private.navbar", "dashboard", "rightbar", "wiki")
}

// GET /dashboard/news
func news(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, r, "news", nil, "layout", "private.navbar", "dashboard", "rightbar", "news")
}

// GET /dashboard/planet
// Show the constructions page for a planet
func planet(w http.ResponseWriter, r *http.Request) {
	u := loggedInUser(r)
	p, err := u.QueryPlanets().
		First(r.Context())
	if err != nil {
		internalServerError(w, r, err, "Could not retrieve user's planet from database")
		return
	}
	generateHTML(w, r, "planet-dashboard", p, "layout", "private.navbar", "dashboard", "rightbar", "planet.layout", "planet.header", "planet.overview")
}

// GET /dashboard/planet/constructions
// Show the constructions page for a planet
func constructions(w http.ResponseWriter, r *http.Request) {
	u := loggedInUser(r)
	p, err := u.QueryPlanets().
		First(r.Context())
	if err != nil {
		internalServerError(w, r, err, "Could not retrieve user's planet from database")
		return
	}
	generateHTML(w, r, "planet-constructions", p, "layout", "private.navbar", "dashboard", "rightbar", "planet.layout", "planet.header", "planet.constructions")
}

// GET /dashboard/planet/factories
// Show the factories page for a planet
func factories(w http.ResponseWriter, r *http.Request) {
	u := loggedInUser(r)
	p, err := u.QueryPlanets().
		First(r.Context())
	if err != nil {
		internalServerError(w, r, err, "Could not retrieve user's planet from database")
		return
	}
	generateHTML(w, r, "planet-factories", p, "layout", "private.navbar", "dashboard", "rightbar", "planet.layout", "planet.header", "planet.factories")
}

// GET /dashboard/planet/research
// Show the research page for a planet
func research(w http.ResponseWriter, r *http.Request) {
	u := loggedInUser(r)
	p, err := u.QueryPlanets().
		First(r.Context())
	if err != nil {
		internalServerError(w, r, err, "Could not retrieve user's planet from database")
		return
	}
	generateHTML(w, r, "planet-research", p, "layout", "private.navbar", "dashboard", "rightbar", "planet.layout", "planet.header", "planet.research")
}

// GET /dashboard/planet/fleets
// Show the fleets page for a planet
func fleets(w http.ResponseWriter, r *http.Request) {
	u := loggedInUser(r)
	p, err := u.QueryPlanets().
		First(r.Context())
	if err != nil {
		internalServerError(w, r, err, "Could not retrieve user's planet from database")
		return
	}
	generateHTML(w, r, "planet-fleets", p, "layout", "private.navbar", "dashboard", "rightbar", "planet.layout", "planet.header", "planet.fleets")
}

// GET /dashboard/planet/defenses
// Show the defenses page for a planet
func defenses(w http.ResponseWriter, r *http.Request) {
	u := loggedInUser(r)
	p, err := u.QueryPlanets().
		First(r.Context())
	if err != nil {
		internalServerError(w, r, err, "Could not retrieve user's planet from database")
		return
	}
	generateHTML(w, r, "planet-defenses", p, "layout", "private.navbar", "dashboard", "rightbar", "planet.layout", "planet.header", "planet.defenses")
}
