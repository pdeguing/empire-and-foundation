package main

import "net/http"

type researchInfo struct {
	Name        string
	Uri         string
	Image       string
	Description string
}

var researchInfos = []researchInfo{
	{
		Name:        "Metallurgy",
		Uri:         "metallurgy",
		Image:       "tech-metallurgy.png",
		Description: "Unlocks more ships and structures.\nIncreases armor X% per level.",
	},
	{
		Name:        "Exploitation",
		Uri:         "exploitation",
		Image:       "tech-exploitation.png",
		Description: "Increases max level for resources buildings. Increases production X% per level.",
	},
	{
		Name:        "Energy",
		Uri:         "energy",
		Image:       "tech-energy.png",
		Description: "Increases solar plant max level. Unlocks atomic physics research. Increases energy production X% per level. ",
	},
	{
		Name:        "Biotechnology",
		Uri:         "biotechnology",
		Image:       "tech-biotechnology.png",
		Description: "Enables colonizing and bio-stations. Increases population growth X% per level.",
	},
	{
		Name:        "Propulsion",
		Uri:         "propulsion",
		Image:       "tech-propulsion.png",
		Description: "Unlocks various ships. Increases ships speed X% per level. Unlocks advanced propulsion technologies.",
	},
	{
		Name:        "Weaponry",
		Uri:         "weaponry",
		Image:       "tech-weaponry.png",
		Description: "Unlocks various ships. Increases ships attack X% per level. Unlocks advanced weaponry technologies.",
	},
	{
		Name:        "Mathematics",
		Uri:         "mathematics",
		Image:       "tech-mathematics.png",
		Description: "Increases max number of colonies and stations per X. Increases max laboratory level.",
	},
	{
		Name:        "Cybernetics",
		Uri:         "cybernetics",
		Image:       "tech-cybernetics.png",
		Description: "Increases max robotic factory level. Unlocks advanced systems. Increases construction speed X% per level.",
	},
}

type researchViewData struct {
	planetViewData
	Cards []researchInfo
}

// GET /planet/{id}/research
// Show the research page for a planet
func serveResearch(w http.ResponseWriter, r *http.Request) {
	p, err := newPlanetViewData(r, "")
	if err != nil {
		serveError(w, r, err)
		return
	}
	vd := researchViewData{
		*p,
		researchInfos,
	}
	generateHTML(w, r, "planet-research", vd, "layout", "private.navbar", "dashboard", "leftbar", "planet.layout", "planet.header", "flash", "planet.research")
}
