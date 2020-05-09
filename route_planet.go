package main

import (
	"net/http"

	"github.com/pdeguing/empire-and-foundation/data"
	"github.com/pdeguing/empire-and-foundation/ent"
	"github.com/pdeguing/empire-and-foundation/ent/timer"
)

type planetOverviewViewData struct {
	UserPlanets []*ent.Planet
	Planet      *data.PlanetWithResourceInfo
	EnergyProd  int64
	EnergyCons  int64
	Timers      map[timer.Group]*data.Timer
}

// GET /planet/{id}
// Show the dashboard page for a planet
func servePlanet(w http.ResponseWriter, r *http.Request) {
	var plist []*ent.Planet
	var p *data.PlanetWithResourceInfo
	var t map[timer.Group]*data.Timer
	err := data.WithTx(r.Context(), data.Client, func(tx *ent.Tx) error {
		var err error
		plist, err = userPlanets(r, tx)
		if err != nil {
			return err
		}
		p, err = userPlanet(r, tx)
		if err != nil {
			return err
		}
		t, err = data.GetTimers(r.Context(), p.Planet)
		if err != nil {
			return newInternalServerError(err)
		}
		return nil
	})
	if err != nil {
		serveError(w, r, err)
		return
	}
	pv := planetOverviewViewData{
		UserPlanets: plist,
		Planet:      p,
		Timers:      t,
	}
	generateHTML(w, r, "planet-dashboard", pv, "layout", "private.navbar", "dashboard", "leftbar", "planet.layout", "planet.header", "flash", "planet.overview")
}

// GET /planet/{id}/fleets
// Show the fleets page for a planet
func serveFleets(w http.ResponseWriter, r *http.Request) {
	p, err := newPlanetViewData(r, "")
	if err != nil {
		serveError(w, r, err)
		return
	}
	generateHTML(w, r, "planet-fleets", p, "layout", "private.navbar", "dashboard", "leftbar", "planet.layout", "planet.header", "flash", "planet.fleets")
}

// GET /planet/{id}/defenses
// Show the defenses page for a planet
func serveDefenses(w http.ResponseWriter, r *http.Request) {
	p, err := newPlanetViewData(r, "")
	if err != nil {
		serveError(w, r, err)
		return
	}
	generateHTML(w, r, "planet-defenses", p, "layout", "private.navbar", "dashboard", "leftbar", "planet.layout", "planet.header", "flash", "planet.defenses")
}
