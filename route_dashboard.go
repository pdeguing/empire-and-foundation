package main

import (
	"net/http"

	"github.com/pdeguing/empire-and-foundation/data"
	"github.com/pdeguing/empire-and-foundation/ent"
)

type dashboardViewData struct {
	UserPlanets	[]*ent.Planet
}

func getUserPlanets(r *http.Request) ([]*ent.Planet, error) {
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
		return nil, err
	}
	return p, nil
}

// GET /dashboard
// Show player main dashboard page
func serveDashboard(w http.ResponseWriter, r *http.Request) {
	userPlanets, err := getUserPlanets(r)
	if err != nil {
		serveError(w, r, err)
		return
	}

	d := dashboardViewData{
		UserPlanets: userPlanets,
	}
	generateHTML(w, r, "dashboard", d, "layout", "private.navbar", "dashboard", "leftbar", "empire")
}

// GET /dashboard/cartography
// Show map page
func serveCartography(w http.ResponseWriter, r *http.Request) {
	regionPlanets, err := regionPlanets(w, r)
	if err != nil {
		serveError(w, r, err)
		return
	}

	userPlanets, err := getUserPlanets(r)
	if err != nil {
		serveError(w, r, err)
		return
	}

	d := struct {
		RegionPlanets	[]*ent.Planet
		UserPlanets	[]*ent.Planet
	}{
		regionPlanets,
		userPlanets,
	}
	generateHTML(w, r, "cartography", d, "layout", "private.navbar", "dashboard", "leftbar", "cartography")
}

// GET /dashboard/fleetcontrol
func serveFleetControl(w http.ResponseWriter, r *http.Request) {
	userPlanets, err := getUserPlanets(r)
	if err != nil {
		serveError(w, r, err)
		return
	}
	d := struct {
		UserPlanets	[]*ent.Planet
	}{
		userPlanets,
	}
	generateHTML(w, r, "fleet", d, "layout", "private.navbar", "dashboard", "leftbar", "fleetcontrol")
}

// GET /dashboard/technology
func serveTechnology(w http.ResponseWriter, r *http.Request) {
	userPlanets, err := getUserPlanets(r)
	if err != nil {
		serveError(w, r, err)
		return
	}
	d := struct {
		UserPlanets	[]*ent.Planet
	}{
		userPlanets,
	}
	generateHTML(w, r, "technology", d, "layout", "private.navbar", "dashboard", "leftbar", "technology")
}

// GET /dashboard/diplomacy
func serveDiplomacy(w http.ResponseWriter, r *http.Request) {
	userPlanets, err := getUserPlanets(r)
	if err != nil {
		serveError(w, r, err)
		return
	}
	d := struct {
		UserPlanets	[]*ent.Planet
	}{
		userPlanets,
	}
	generateHTML(w, r, "diplomacy", d, "layout", "private.navbar", "dashboard", "leftbar", "diplomacy")
}

// GET /dashboard/story
func serveStory(w http.ResponseWriter, r *http.Request) {
	userPlanets, err := getUserPlanets(r)
	if err != nil {
		serveError(w, r, err)
		return
	}
	d := struct {
		UserPlanets	[]*ent.Planet
	}{
		userPlanets,
	}
	generateHTML(w, r, "story", d, "layout", "private.navbar", "dashboard", "leftbar", "story")
}

// GET /dashboard/wiki
func serveWiki(w http.ResponseWriter, r *http.Request) {
	userPlanets, err := getUserPlanets(r)
	if err != nil {
		serveError(w, r, err)
		return
	}
	d := struct {
		UserPlanets	[]*ent.Planet
	}{
		userPlanets,
	}
	generateHTML(w, r, "wiki", d, "layout", "private.navbar", "dashboard", "leftbar", "wiki")
}

// GET /dashboard/news
func serveNews(w http.ResponseWriter, r *http.Request) {
	userPlanets, err := getUserPlanets(r)
	if err != nil {
		serveError(w, r, err)
		return
	}
	d := struct {
		UserPlanets	[]*ent.Planet
	}{
		userPlanets,
	}
	generateHTML(w, r, "news", d, "layout", "private.navbar", "dashboard", "leftbar", "news")
}
