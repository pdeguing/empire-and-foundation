package main

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pdeguing/empire-and-foundation/data"
	"github.com/pdeguing/empire-and-foundation/ent"
	"github.com/pdeguing/empire-and-foundation/ent/timer"
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
	err = data.UpdateTimers(r.Context(), p)
	if err != nil {
		serveInternalServerError(w, r, err, "Cannot update the planet timers")
		return nil, false
	}
	return p, true
}

func newPlanetViewData(w http.ResponseWriter, r *http.Request, g timer.Group) (*planetViewData, bool) {
	p, ok := userPlanet(w, r)
	if !ok {
		return nil, false
	}
	var err error
	var t *data.Timer
	if g != "" {
		t, err = data.GetTimer(r.Context(), p, g)
		if err != nil {
			serveInternalServerError(w, r, err, "Cannot get timer in group for planet")
			return nil, false
		}
	}
	return &planetViewData{
		Planet: p,
		Timer:  t,
	}, true
}

type planetViewData struct {
	Planet *ent.Planet
	Timer  *data.Timer
}

// GET /planet/{id}
// Show the dashboard page for a planet
func servePlanet(w http.ResponseWriter, r *http.Request) {
	p, ok := userPlanet(w, r)
	if !ok {
		return
	}
	t, err := data.GetTimers(r.Context(), p)
	if err != nil {
		serveInternalServerError(w, r, err, "Cannot get timers for planet")
		return
	}
	pv := planetOverviewViewData{
		Planet: p,
		Timers: t,
	}
	generateHTML(w, r, "planet-dashboard", pv, "layout", "private.navbar", "dashboard", "leftbar", "planet.layout", "planet.header", "planet.overview")
}

type planetOverviewViewData struct {
	Planet *ent.Planet
	Timers map[timer.Group]*data.Timer
}

// GET /planet/{id}/constructions
// Show the constructions page for a planet
func serveConstructions(w http.ResponseWriter, r *http.Request) {
	if p, ok := newPlanetViewData(w, r, timer.GroupBuilding); ok {
		generateHTML(w, r, "planet-constructions", p, "layout", "private.navbar", "dashboard", "leftbar", "planet.layout", "planet.header", "planet.constructions")
	}
}

// GET /planet/{id}/factories
// Show the factories page for a planet
func serveFactories(w http.ResponseWriter, r *http.Request) {
	if p, ok := newPlanetViewData(w, r, ""); ok {
		generateHTML(w, r, "planet-factories", p, "layout", "private.navbar", "dashboard", "leftbar", "planet.layout", "planet.header", "planet.factories")
	}
}

// GET /planet/{id}/research
// Show the research page for a planet
func serveResearch(w http.ResponseWriter, r *http.Request) {
	if p, ok := newPlanetViewData(w, r, ""); ok {
		generateHTML(w, r, "planet-research", p, "layout", "private.navbar", "dashboard", "leftbar", "planet.layout", "planet.header", "planet.research")
	}
}

// GET /planet/{id}/fleets
// Show the fleets page for a planet
func serveFleets(w http.ResponseWriter, r *http.Request) {
	if p, ok := newPlanetViewData(w, r, ""); ok {
		generateHTML(w, r, "planet-fleets", p, "layout", "private.navbar", "dashboard", "leftbar", "planet.layout", "planet.header", "planet.fleets")
	}
}

// GET /planet/{id}/defenses
// Show the defenses page for a planet
func serveDefenses(w http.ResponseWriter, r *http.Request) {
	if p, ok := newPlanetViewData(w, r, ""); ok {
		generateHTML(w, r, "planet-defenses", p, "layout", "private.navbar", "dashboard", "leftbar", "planet.layout", "planet.header", "planet.defenses")
	}
}

// POST /planet/{id}/up_metal_mine
// Upgrade the metal mine to the next level
func serveUpMetalMine(w http.ResponseWriter, r *http.Request) {
	p, ok := userPlanet(w, r)
	if !ok {
		return
	}
	err := data.StartTimer(r.Context(), p, timer.ActionUpgradeMetalMine)
	if err == data.ErrActionPrerequisitesNotMet {
		// TODO: Flash message
		http.Redirect(w, r, "/planet/"+strconv.Itoa(p.ID)+"/constructions", 302)
		return
	}
	if err == data.ErrTimerBussy {
		// TODO: Flash message
		http.Redirect(w, r, "/planet/"+strconv.Itoa(p.ID)+"/constructions", 302)
		return
	}
	if err != nil {
		serveInternalServerError(w, r, err, "cannot start command")
		return
	}
	http.Redirect(w, r, "/planet/"+strconv.Itoa(p.ID)+"/constructions", 302)
}
