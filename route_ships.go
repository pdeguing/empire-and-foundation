package main

import (
	"github.com/pdeguing/empire-and-foundation/data"
	"github.com/pdeguing/empire-and-foundation/ent"
	"github.com/pdeguing/empire-and-foundation/ent/timer"
	"net/http"
)

var shipRouteActionMap = map[string]timer.Action{}

func init() {
	for _, ship := range shipInfos {
		shipRouteActionMap[ship.Uri] = ship.Ship.BuildAction()
	}
}

type shipInfo struct {
	Ship        data.Ship
	Name        string
	Uri         string
	Image       string
	Description string
}

var shipInfos = []shipInfo{
	{
		Ship:        data.Caravel{},
		Name:        "Caravel",
		Uri:         "caravel",
		Image:       "ship-caravel.png",
		Description: "A well-rounded ship. Fast, reliable, armed, with decent cargo. The perfect tool for exploring and setting early presence between the stars.",
	},
	{
		Ship:        data.LightFighter{},
		Name:        "Light Fighter",
		Uri:         "light-fighter",
		Image:       "ship-light-fighter.png",
		Description: "Fast and maneuvrable. When grouped they represent a high threat to any fleet overhelmed by their numbers. The most agile attack force.",
	},
	{
		Ship:        data.Corvette{},
		Name:        "Corvette",
		Uri:         "corvette",
		Image:       "ship-corvette.png",
		Description: "Lorem ipsum dolor sit amet",
	},
	{
		Ship:        data.Frigate{},
		Name:        "Frigate",
		Uri:         "frigate",
		Image:       "ship-frigate.png",
		Description: "Lorem ipsum dolor sit amet",
	},
	{
		Ship:        data.Probe{},
		Name:        "Probe",
		Uri:         "probe",
		Image:       "ship-probe.png",
		Description: "Lorem ipsum dolor sit amet",
	},
	{
		Ship:        data.SmallCargo{},
		Name:        "Small Cargo",
		Uri:         "small-cargo",
		Image:       "ship-small-cargo.png",
		Description: "Lorem ipsum dolor sit amet",
	},
	{
		Ship:        data.MediumCargo{},
		Name:        "Medium Cargo",
		Uri:         "medium-cargo",
		Image:       "ship-medium-cargo.png",
		Description: "Lorem ipsum dolor sit amet",
	},
	{
		Ship:        data.ColonizationArk{},
		Name:        "Colonization Ark",
		Uri:         "colonization-ark",
		Image:       "ship-colonization-ark.png",
		Description: "Lorem ipsum dolor sit amet",
	},
}

type shipCard struct {
	Name        string
	Uri         string
	Image       string
	Description string
	Count       int64
	BuildCost   data.Amounts
	Buildable   bool
	Timer       *data.Timer
}

func newShipCard(planet *ent.Planet, timer *data.Timer, info shipInfo) shipCard {
	ship := info.Ship
	cost := ship.Cost()

	var shipTimer *data.Timer
	if timer != nil && timer.Action == ship.BuildAction() {
		shipTimer = timer
	}

	return shipCard{
		Name:        info.Name,
		Uri:         info.Uri,
		Image:       info.Image,
		Description: info.Description,
		Count:       ship.NumberOnPlanet(planet),
		BuildCost:   cost,
		Buildable:   data.HasResources(planet, cost) && timer == nil,
		Timer:       shipTimer,
	}
}

type factoriesViewData struct {
	planetViewData
	Cards []shipCard
}

func getShipCards(planet *ent.Planet, timer *data.Timer) []shipCard {
	var cards []shipCard
	for _, info := range shipInfos {
		cards = append(cards, newShipCard(planet, timer, info))
	}
	return cards
}

// GET /planet/{id}/factories
// Show the factories page for a planet
func serveFactories(w http.ResponseWriter, r *http.Request) {
	p, err := newPlanetViewData(r, timer.GroupShip)
	if err != nil {
		serveError(w, r, err)
		return
	}
	vd := factoriesViewData{
		*p,
		getShipCards(p.Planet.Planet, p.Timer),
	}
	generateHTML(w, r, "planet-factories", vd, "layout", "private.navbar", "dashboard", "leftbar", "planet.layout", "planet.header", "flash", "planet.factories")
}

// POST /planet/{id}/factory/{action}/build
func servePlanetStartFactory(w http.ResponseWriter, r *http.Request) {
	servePlanetStartAction(w, r, shipRouteActionMap, "factories")
}

// POST /planet/{id}/factory/{action}/cancel
func servePlanetCancelFactory(w http.ResponseWriter, r *http.Request) {
	servePlanetCancelAction(w, r, shipRouteActionMap, "factories")
}
