package main

import (
	"net/http"
)

// GET /dashboard
// Show player main dashboard page
func serveDashboard(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, r, "dashboard", nil, "layout", "private.navbar", "dashboard", "leftbar", "empire")
}

// GET /dashboard/cartography
// Show map page
func serveCartography(w http.ResponseWriter, r *http.Request) {
	p, err := regionPlanets(w, r)
	if err != nil {
		serverError(w, r, err)
	}
	generateHTML(w, r, "cartography", p, "layout", "private.navbar", "dashboard", "leftbar", "cartography")
}

// GET /dashboard/fleetcontrol
func serveFleetControl(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, r, "fleet", nil, "layout", "private.navbar", "dashboard", "leftbar", "fleetcontrol")
}

// GET /dashboard/technology
func serveTechnology(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, r, "technology", nil, "layout", "private.navbar", "dashboard", "leftbar", "technology")
}

// GET /dashboard/diplomacy
func serveDiplomacy(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, r, "diplomacy", nil, "layout", "private.navbar", "dashboard", "leftbar", "diplomacy")
}

// GET /dashboard/story
func serveStory(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, r, "story", nil, "layout", "private.navbar", "dashboard", "leftbar", "story")
}

// GET /dashboard/wiki
func serveWiki(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, r, "wiki", nil, "layout", "private.navbar", "dashboard", "leftbar", "wiki")
}

// GET /dashboard/news
func serveNews(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, r, "news", nil, "layout", "private.navbar", "dashboard", "leftbar", "news")
}
