package main

import "net/http"

type shipInfo struct {
	Name        string
	Uri         string
	Image       string
	Description string
}

var shipInfos = []shipInfo{
	{
		Name:        "Caravel",
		Uri:         "caravel",
		Image:       "ship-caravel.png",
		Description: "A well-rounded ship. Fast, reliable, armed, with decent cargo. The perfect tool for exploring and setting early presence between the stars.",
	},
	{
		Name:        "Light Fighter",
		Uri:         "light-fighter",
		Image:       "ship-light-fighter.png",
		Description: "Fast and maneuvrable. When grouped they represent a high threat to any fleet overhelmed by their numbers. The most agile attack force.",
	},
	{
		Name:        "Corvette",
		Uri:         "corvette",
		Image:       "ship-corvette.png",
		Description: "Lorem ipsum dolor sit amet",
	},
	{
		Name:        "Frigate",
		Uri:         "frigate",
		Image:       "ship-frigate.png",
		Description: "Lorem ipsum dolor sit amet",
	},
	{
		Name:        "Probe",
		Uri:         "probe",
		Image:       "ship-probe.png",
		Description: "Lorem ipsum dolor sit amet",
	},
	{
		Name:        "Small cargo",
		Uri:         "small-cargo",
		Image:       "ship-small-cargo.png",
		Description: "Lorem ipsum dolor sit amet",
	},
	{
		Name:        "Medium Cargo",
		Uri:         "medium-cargo",
		Image:       "ship-medium-cargo.png",
		Description: "Lorem ipsum dolor sit amet",
	},
	{
		Name:        "Colonization Ark",
		Uri:         "colonization-ark",
		Image:       "ship-colonization-ark.png",
		Description: "Lorem ipsum dolor sit amet",
	},
}

type factoriesViewData struct {
	planetViewData
	Cards []shipInfo
}

// GET /planet/{id}/factories
// Show the factories page for a planet
func serveFactories(w http.ResponseWriter, r *http.Request) {
	p, err := newPlanetViewData(r, "")
	if err != nil {
		serveError(w, r, err)
		return
	}
	vd := factoriesViewData{
		*p,
		shipInfos,
	}
	generateHTML(w, r, "planet-factories", vd, "layout", "private.navbar", "dashboard", "leftbar", "planet.layout", "planet.header", "flash", "planet.factories")
}
