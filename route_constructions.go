package main

import (
	"github.com/pdeguing/empire-and-foundation/data"
	"github.com/pdeguing/empire-and-foundation/ent"
	"github.com/pdeguing/empire-and-foundation/ent/timer"
	"net/http"
)

var constructionRouteActionMap = map[string]timer.Action{}

func init() {
	for _, building := range buildingInfos {
		constructionRouteActionMap[building.Uri] = building.Building.UpgradeAction()
	}
}

type buildingInfo struct {
	Building    data.Building
	Name        string
	Uri         string
	Image       string
	Description string
}

var buildingInfos = []buildingInfo{
	{
		Building:    data.MetalMine{},
		Name:        "Metal Mine",
		Uri:         "metal-mine",
		Image:       "metal-prod.png",
		Description: "Mining infrastructures extracting metal.",
	},
	{
		Building:    data.MetalStorage{},
		Name:        "Metal Storage",
		Uri:         "metal-storage",
		Image:       "metal-storage.png",
		Description: "Increases metal storage capacity.",
	},
	{
		Building:    data.HydrogenExtractor{},
		Name:        "Hydrogen Extractor",
		Uri:         "hydrogen-extractor",
		Image:       "hydrogen-prod.png",
		Description: "Enrichment infrastructures collecting hydrogen.",
	},
	{
		Building:    data.HydrogenStorage{},
		Name:        "Hydrogen Storage",
		Uri:         "hydrogen-storage",
		Image:       "hydrogen-storage.png",
		Description: "Increases hydrogen storage capacity.",
	},
	{
		Building:    data.SilicaQuarry{},
		Name:        "Silica Quarry",
		Uri:         "silica-quarry",
		Image:       "silica-prod.png",
		Description: "Mining infrastructures extracting silica.",
	},
	{
		Building:    data.SilicaStorage{},
		Name:        "Silica Storage",
		Uri:         "silica-storage",
		Image:       "silica-storage.png",
		Description: "Increases silica storage capacity.",
	},
	{
		Building:    data.SolarPlant{},
		Name:        "Solar Plant",
		Uri:         "solar-plant",
		Image:       "solar-prod.png",
		Description: "Looking up to the stars, the solar plant captures energy from its sun.",
	},
	{
		Building:    data.Urbanism{},
		Name:        "Urbanism",
		Uri:         "housing-facilities",
		Image:       "urbanism.png",
		Description: "Living accommodations for the population, allowing it to grow naturally.",
	},
	{
		Building:    data.ResearchCenter{},
		Name:        "Research Center",
		Uri:         "research-center",
		Image:       "solar-prod.png",
		Description: "Research new technologies.",
	},
	{
		Building:    data.ShipFactory{},
		Name:        "Ship Factory",
		Uri:         "ship-factory",
		Image:       "solar-prod.png",
		Description: "Build spaceships.",
	},
}

type constructionCard struct {
	Name         string
	Uri          string
	Image        string
	Description  string
	Level        int
	UpgradeCost  data.Amounts
	UpgradeUsage data.Levels
	DeltaUsage   data.Levels
	Upgradable   bool
	Timer        *data.Timer
}

func newConstructionCard(planet *ent.Planet, timer *data.Timer, info buildingInfo) constructionCard {
	building := info.Building.ForPlanet(planet)
	cost := building.NextLevel().Cost()
	usage := building.NextLevel().Usage()
	deltaUsage := usage.Sub(building.Usage())

	var buildingTimer *data.Timer
	if timer != nil && timer.Action == building.UpgradeAction() {
		buildingTimer = timer
	}

	return constructionCard{
		Name:         info.Name,
		Uri:          info.Uri,
		Image:        info.Image,
		Description:  info.Description,
		Level:        building.LevelOnPlanet(planet),
		UpgradeCost:  cost,
		UpgradeUsage: usage,
		DeltaUsage:   deltaUsage,
		Upgradable:   data.HasResources(planet, cost) && timer == nil,
		Timer:        buildingTimer,
	}
}

type constructionsViewData struct {
	planetViewData
	Cards []constructionCard
}

func getConstructionCards(planet *ent.Planet, timer *data.Timer) []constructionCard {
	var cards []constructionCard
	for _, info := range buildingInfos {
		cards = append(cards, newConstructionCard(planet, timer, info))
	}
	return cards
}

// GET /planet/{id}/constructions
// Show the constructions page for a planet
func serveConstructions(w http.ResponseWriter, r *http.Request) {
	p, err := newPlanetViewData(r, timer.GroupBuilding)
	if err != nil {
		serveError(w, r, err)
		return
	}
	vd := constructionsViewData{
		*p,
		getConstructionCards(p.Planet.Planet, p.Timer),
	}
	generateHTML(w, r, "planet-constructions", vd, "layout", "private.navbar", "dashboard", "leftbar", "planet.layout", "planet.header", "flash", "planet.constructions")
}

// POST /planet/{id}/construction/{action}/build
func servePlanetStartConstruction(w http.ResponseWriter, r *http.Request) {
	servePlanetStartAction(w, r, constructionRouteActionMap, "constructions")
}

// POST /planet/{id}/construction/{action}/cancel
func servePlanetCancelConstruction(w http.ResponseWriter, r *http.Request) {
	servePlanetCancelAction(w, r, constructionRouteActionMap, "constructions")
}
