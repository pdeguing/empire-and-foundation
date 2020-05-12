package main

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	"github.com/pdeguing/empire-and-foundation/data"
	"github.com/pdeguing/empire-and-foundation/ent"
	"github.com/pdeguing/empire-and-foundation/ent/timer"
)

var planetRouteActionMap = map[string]timer.Action{}

func init() {
	for _, building := range buildingInfos {
		planetRouteActionMap[building.Uri] = building.Building.UpgradeAction()
	}
}

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

// POST /planet/{id}/{action}/build
// servePlanetStartAction progresses the request to start an upgrade or build timer.
func servePlanetStartAction(w http.ResponseWriter, r *http.Request) {
	uri := mux.Vars(r)["action"]
	a, ok := planetRouteActionMap[uri]
	if !ok {
		http.NotFound(w, r)
		return
	}

	var p *data.PlanetWithResourceInfo
	err := data.WithTx(r.Context(), data.Client, func(tx *ent.Tx) error {
		var err error
		p, err = userPlanet(r, tx)
		if err != nil {
			return err
		}
		err = data.StartTimer(r.Context(), tx, p.Planet, a)
		if err != nil {
			return newInternalServerError(fmt.Errorf("unable to start timer to upgrade building: %w", err))
		}
		return nil
	})
	if err != nil {
		if errors.Is(err, data.ErrActionPrerequisitesNotMet) {
			flash(r, flashDanger, "There are not enough resources on this planet.")
			http.Redirect(w, r, "/planet/"+strconv.Itoa(p.ID)+"/constructions", 302)
			return
		}
		if errors.Is(err, data.ErrTimerBusy) {
			flash(r, flashWarning, "Something is already being upgraded.")
			http.Redirect(w, r, "/planet/"+strconv.Itoa(p.ID)+"/constructions", 302)
			return
		}
		serveError(w, r, err)
		return
	}
	http.Redirect(w, r, "/planet/"+strconv.Itoa(p.ID)+"/constructions", 302)
}

// POST /planet/{id}/{action}/cancel
// servePlanetCancelAction progresses the request to cancel a timer.
func servePlanetCancelAction(w http.ResponseWriter, r *http.Request) {
	uri := mux.Vars(r)["action"]
	a, ok := planetRouteActionMap[uri]
	if !ok {
		http.NotFound(w, r)
		return
	}

	var p *data.PlanetWithResourceInfo
	err := data.WithTx(r.Context(), data.Client, func(tx *ent.Tx) error {
		var err error
		p, err = userPlanet(r, tx)
		if err != nil {
			return err
		}
		err = data.CancelTimer(r.Context(), tx, p.Planet, a)
		if err != nil {
			return newInternalServerError(fmt.Errorf("unable to cancel timer to upgrade building: %w", err))
		}
		return nil
	})
	if err != nil && !errors.Is(err, data.ErrTimerNotRunning) {
		serveError(w, r, err)
		return
	}
	http.Redirect(w, r, "/planet/"+strconv.Itoa(p.ID)+"/constructions", 302)
}
