package main

import (
	"net/http"

	"github.com/pdeguing/empire-and-foundation/data"
	"github.com/pdeguing/empire-and-foundation/ent"
	"github.com/pdeguing/empire-and-foundation/ent/timer"
)

type planetOverviewViewData struct {
	UserPlanets	[]*ent.Planet
	Planet *ent.Planet
	Timers map[timer.Group]*data.Timer
}

// GET /planet/{id}
// Show the dashboard page for a planet
func servePlanet(w http.ResponseWriter, r *http.Request) {
	var plist []*ent.Planet
	var p *ent.Planet
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
		UserPlanets: plist,
		Planet: p,
		Timers: t,
	}
	generateHTML(w, r, "planet-dashboard", pv, "layout", "private.navbar", "dashboard", "leftbar", "planet.layout", "planet.header", "flash", "planet.overview")
}

// GET /planet/{id}/constructions
// Show the constructions page for a planet
func serveConstructions(w http.ResponseWriter, r *http.Request) {
	p, err := newPlanetViewData(r, timer.GroupBuilding)
	if err != nil {
		serveError(w, r, err)
		return
	}
	generateHTML(w, r, "planet-constructions", p, "layout", "private.navbar", "dashboard", "leftbar", "planet.layout", "planet.header", "flash", "planet.constructions")
}

// GET /planet/{id}/factories
// Show the factories page for a planet
func serveFactories(w http.ResponseWriter, r *http.Request) {
	p, err := newPlanetViewData(r, "")
	if err != nil {
		serveError(w, r, err)
		return
	}
	generateHTML(w, r, "planet-factories", p, "layout", "private.navbar", "dashboard", "leftbar", "planet.layout", "planet.header", "flash", "planet.factories")
}

// GET /planet/{id}/research
// Show the research page for a planet
func serveResearch(w http.ResponseWriter, r *http.Request) {
	p, err := newPlanetViewData(r, "")
	if err != nil {
		serveError(w, r, err)
		return
	}
	generateHTML(w, r, "planet-research", p, "layout", "private.navbar", "dashboard", "leftbar", "planet.layout", "planet.header", "flash", "planet.research")
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
