package main

import (
	"net/http"

	"github.com/pdeguing/empire-and-foundation/data"
	"github.com/pdeguing/empire-and-foundation/ent"
)

type dashboardViewData struct {
	UserPlanets	[]*ent.Planet
}

// GET /dashboard
// Show player main dashboard page
func serveDashboard(w http.ResponseWriter, r *http.Request) {
	var p []*ent.Planet
	err := data.WithTx(r.Context(), data.Client, func(tx *ent.Tx) error {
		var err error
		p, err = userPlanets(r, tx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		serveError(w, r, err)
		return
	}
	d := dashboardViewData{
		UserPlanets: p,
	}
	generateHTML(w, r, "dashboard", d, "layout", "private.navbar", "dashboard", "leftbar", "empire")
}

// GET /dashboard/cartography
// Show map page
func serveCartography(w http.ResponseWriter, r *http.Request) {
	p, err := regionPlanets(w, r)
	if err != nil {
		serveError(w, r, err)
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
