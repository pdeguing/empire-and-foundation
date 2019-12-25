package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pdeguing/empire-and-foundation/data"
	"github.com/pdeguing/empire-and-foundation/ent"
	"github.com/pdeguing/empire-and-foundation/ent/planet"
	"github.com/pdeguing/empire-and-foundation/ent/timer"
	"github.com/pdeguing/empire-and-foundation/ent/user"
)

// Helper
func userPlanet(r *http.Request, tx *ent.Tx) (*ent.Planet, error) {
	n, err := strconv.Atoi(mux.Vars(r)["planetNumber"])
	if err != nil {
		return nil, newInternalServerError(fmt.Errorf("unable to parse url parameter \"planetNumber\": %v", err))
	}
	u := loggedInUser(r)

	p, err := tx.Planet.
		Query().
		Where(planet.HasOwnerWith(user.IDEQ(u.ID))).
		Offset(n - 1).
		First(r.Context())
	if _, ok := err.(*ent.ErrNotFound); ok {
		return nil, newNotFoundError(fmt.Errorf("unable to query planet #%d for user %d; it does not exist", n, u.ID))
	}
	if err != nil {
		return nil, newInternalServerError(fmt.Errorf("unable to retrieve planet for user: %v", err))
	}
	err = data.UpdateTimers(r.Context(), tx, p)
	if err != nil {
		return nil, newInternalServerError(fmt.Errorf("unable to update planet timers: %v", err))
	}
	return p, nil
}

func newPlanetViewData(r *http.Request, g timer.Group) (*planetViewData, error) {
	var p *ent.Planet
	var t *data.Timer
	err := data.WithTx(r.Context(), data.Client, func(tx *ent.Tx) error {
		var err error
		p, err = userPlanet(r, tx)
		if err != nil {
			return err
		}
		// TODO: Remove wrapping if statement once all planet views have a group
		// they want to get the timers for.
		if g != "" {
			t, err = data.GetTimer(r.Context(), p, g)
			if err != nil {
				return newInternalServerError(err)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &planetViewData{
		Planet: p,
		Timer:  t,
	}, nil
}

type planetViewData struct {
	Planet *ent.Planet
	Timer  *data.Timer
}

// GET /planet/{id}
// Show the dashboard page for a planet
func servePlanet(w http.ResponseWriter, r *http.Request) {
	var p *ent.Planet
	var t map[timer.Group]*data.Timer
	err := data.WithTx(r.Context(), data.Client, func(tx *ent.Tx) error {
		var err error
		p, err = userPlanet(r, tx)
		if err != nil {
			return err
		}
		t, err = data.GetTimers(r.Context(), p)
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
	p, err := newPlanetViewData(r, timer.GroupBuilding)
	if err != nil {
		serveError(w, r, err)
		return
	}
	generateHTML(w, r, "planet-constructions", p, "layout", "private.navbar", "dashboard", "leftbar", "planet.layout", "planet.header", "planet.constructions")
}

// GET /planet/{id}/factories
// Show the factories page for a planet
func serveFactories(w http.ResponseWriter, r *http.Request) {
	p, err := newPlanetViewData(r, "")
	if err != nil {
		serveError(w, r, err)
		return
	}
	generateHTML(w, r, "planet-factories", p, "layout", "private.navbar", "dashboard", "leftbar", "planet.layout", "planet.header", "planet.factories")
}

// GET /planet/{id}/research
// Show the research page for a planet
func serveResearch(w http.ResponseWriter, r *http.Request) {
	p, err := newPlanetViewData(r, "")
	if err != nil {
		serveError(w, r, err)
		return
	}
	generateHTML(w, r, "planet-research", p, "layout", "private.navbar", "dashboard", "leftbar", "planet.layout", "planet.header", "planet.research")
}

// GET /planet/{id}/fleets
// Show the fleets page for a planet
func serveFleets(w http.ResponseWriter, r *http.Request) {
	p, err := newPlanetViewData(r, "")
	if err != nil {
		serveError(w, r, err)
		return
	}
	generateHTML(w, r, "planet-fleets", p, "layout", "private.navbar", "dashboard", "leftbar", "planet.layout", "planet.header", "planet.fleets")
}

// GET /planet/{id}/defenses
// Show the defenses page for a planet
func serveDefenses(w http.ResponseWriter, r *http.Request) {
	p, err := newPlanetViewData(r, "")
	if err != nil {
		serveError(w, r, err)
		return
	}
	generateHTML(w, r, "planet-defenses", p, "layout", "private.navbar", "dashboard", "leftbar", "planet.layout", "planet.header", "planet.defenses")
}

func serveUpgradeBuilding(w http.ResponseWriter, r *http.Request, a timer.Action) {
	var p *ent.Planet
	err := data.WithTx(r.Context(), data.Client, func(tx *ent.Tx) error {
		var err error
		p, err = userPlanet(r, tx)
		if err != nil {
			return err
		}
		err = data.StartTimer(r.Context(), tx, p, a)
		if err != nil {
			return newInternalServerError(fmt.Errorf("unable to start timer to upgrade building: %w", err))
		}
		return nil
	})
	if err != nil {
		if errors.Is(err, data.ErrActionPrerequisitesNotMet) {
			// TODO: Flash message
			http.Redirect(w, r, "/planet/"+strconv.Itoa(p.ID)+"/constructions", 302)
			return
		}
		if errors.Is(err, data.ErrTimerBussy) {
			// TODO: Flash message
			http.Redirect(w, r, "/planet/"+strconv.Itoa(p.ID)+"/constructions", 302)
			return
		}
		serveError(w, r, err)
		return
	}
	http.Redirect(w, r, "/planet/"+strconv.Itoa(p.ID)+"/constructions", 302)
}

func serveCancelBuilding(w http.ResponseWriter, r *http.Request, a timer.Action) {
	var p *ent.Planet
	err := data.WithTx(r.Context(), data.Client, func(tx *ent.Tx) error {
		var err error
		p, err = userPlanet(r, tx)
		if err != nil {
			return err
		}
		err = data.CancelTimer(r.Context(), tx, p, a)
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

// POST /planet/{id}/metal-mine/upgrade
// Upgrade the metal mine to the next level
func serveUpgradeMetalMine(w http.ResponseWriter, r *http.Request) {
	serveUpgradeBuilding(w, r, timer.ActionUpgradeMetalMine)
}

// POST /planet/{id}/metal-mine/cancel
// Cancel the upgrade of the metal mine
func serveCancelMetalMine(w http.ResponseWriter, r *http.Request) {
	serveCancelBuilding(w, r, timer.ActionUpgradeMetalMine)
}

// POST /planet/{id}/hydrogen-extractor/upgrade
// Upgrade the hydrogen extractor to the next level
func serveUpgradeHydrogenExtractor(w http.ResponseWriter, r *http.Request) {
	serveUpgradeBuilding(w, r, timer.ActionUpgradeHydrogenExtractor)
}

// POST /planet/{id}/hydrogen-extractor/cancel
// Cancel the upgrade of the hydrogen extractor
func serveCancelHydrogenExtractor(w http.ResponseWriter, r *http.Request) {
	serveCancelBuilding(w, r, timer.ActionUpgradeHydrogenExtractor)
}

// POST /planet/{id}/silica-quarry/upgrade
// Upgrade the silica quarry to the next level
func serveUpgradeSilicaQuarry(w http.ResponseWriter, r *http.Request) {
	serveUpgradeBuilding(w, r, timer.ActionUpgradeSilicaQuarry)
}

// POST /planet/{id}/silica-quarry/cancel
// Cancel the upgrade of the silica quarry
func serveCancelSilicaQuarry(w http.ResponseWriter, r *http.Request) {
	serveCancelBuilding(w, r, timer.ActionUpgradeSilicaQuarry)
}

// POST /planet/{id}/solar-plant/upgrade
// Upgrade the solar plant to the next level
func serveUpgradeSolarPlant(w http.ResponseWriter, r *http.Request) {
	serveUpgradeBuilding(w, r, timer.ActionUpgradeSolarPlant)
}

// POST /planet/{id}/solar-plant/cancel
// Cancel the upgrade of the solar plant
func serveCancelSolarPlant(w http.ResponseWriter, r *http.Request) {
	serveCancelBuilding(w, r, timer.ActionUpgradeSolarPlant)
}

// POST /planet/{id}/housing-facilities/upgrade
// Upgrade the housing facilities to the next level
func serveUpgradeHousingFacilities(w http.ResponseWriter, r *http.Request) {
	serveUpgradeBuilding(w, r, timer.ActionUpgradeHousingFacilities)
}

// POST /planet/{id}/housing-facilities/cancel
// Cancel the upgrade of the housing facilities
func serveCancelHousingFacilities(w http.ResponseWriter, r *http.Request) {
	serveCancelBuilding(w, r, timer.ActionUpgradeHousingFacilities)
}
