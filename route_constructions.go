package main

import (
	"github.com/pdeguing/empire-and-foundation/data"
	"github.com/pdeguing/empire-and-foundation/ent"
	"github.com/pdeguing/empire-and-foundation/ent/timer"
	"net/http"
)

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
	b := getConstructionCards(p.Planet.Planet, p.Timer)
	vd := constructionsViewData{
		*p,
		b,
	}
	generateHTML(w, r, "planet-constructions", vd, "layout", "private.navbar", "dashboard", "leftbar", "planet.layout", "planet.header", "flash", "planet.constructions")
}

// POST /planet/{id}/metal-mine/upgrade
// Upgrade the metal mine to the next level
func serveUpgradeMetalMine(w http.ResponseWriter, r *http.Request) {
	serveUpgradeBuilding(w, r, timer.ActionUpgradeMetalProd)
}

// POST /planet/{id}/metal-mine/cancel
// Cancel the upgrade of the metal mine
func serveCancelMetalMine(w http.ResponseWriter, r *http.Request) {
	serveCancelBuilding(w, r, timer.ActionUpgradeMetalProd)
}

// POST /planet/{id}/hydrogen-extractor/upgrade
// Upgrade the hydrogen extractor to the next level
func serveUpgradeHydrogenExtractor(w http.ResponseWriter, r *http.Request) {
	serveUpgradeBuilding(w, r, timer.ActionUpgradeHydrogenProd)
}

// POST /planet/{id}/hydrogen-extractor/cancel
// Cancel the upgrade of the hydrogen extractor
func serveCancelHydrogenExtractor(w http.ResponseWriter, r *http.Request) {
	serveCancelBuilding(w, r, timer.ActionUpgradeHydrogenProd)
}

// POST /planet/{id}/silica-quarry/upgrade
// Upgrade the silica quarry to the next level
func serveUpgradeSilicaQuarry(w http.ResponseWriter, r *http.Request) {
	serveUpgradeBuilding(w, r, timer.ActionUpgradeSilicaProd)
}

// POST /planet/{id}/silica-quarry/cancel
// Cancel the upgrade of the silica quarry
func serveCancelSilicaQuarry(w http.ResponseWriter, r *http.Request) {
	serveCancelBuilding(w, r, timer.ActionUpgradeSilicaProd)
}

// POST /planet/{id}/solar-plant/upgrade
// Upgrade the solar plant to the next level
func serveUpgradeSolarPlant(w http.ResponseWriter, r *http.Request) {
	serveUpgradeBuilding(w, r, timer.ActionUpgradeSolarProd)
}

// POST /planet/{id}/solar-plant/cancel
// Cancel the upgrade of the solar plant
func serveCancelSolarPlant(w http.ResponseWriter, r *http.Request) {
	serveCancelBuilding(w, r, timer.ActionUpgradeSolarProd)
}

// POST /planet/{id}/housing-facilities/upgrade
// Upgrade the housing facilities to the next level
func serveUpgradeUrbanism(w http.ResponseWriter, r *http.Request) {
	serveUpgradeBuilding(w, r, timer.ActionUpgradeUrbanism)
}

// POST /planet/{id}/housing-facilities/cancel
// Cancel the upgrade of the housing facilities
func serveCancelUrbanism(w http.ResponseWriter, r *http.Request) {
	serveCancelBuilding(w, r, timer.ActionUpgradeUrbanism)
}

// POST /planet/{id}/metal-storage/upgrade
// Upgrade the metal storage to the next level
func serveUpgradeMetalStorage(w http.ResponseWriter, r *http.Request) {
	serveUpgradeBuilding(w, r, timer.ActionUpgradeMetalStorage)
}

// POST /planet/{id}/metal-storage/cancel
// Cancel the upgrade of the metal storage
func serveCancelMetalStorage(w http.ResponseWriter, r *http.Request) {
	serveCancelBuilding(w, r, timer.ActionUpgradeMetalStorage)
}

// POST /planet/{id}/hydrogen-storage/upgrade
// Upgrade the hydrogen storage to the next level
func serveUpgradeHydrogenStorage(w http.ResponseWriter, r *http.Request) {
	serveUpgradeBuilding(w, r, timer.ActionUpgradeHydrogenStorage)
}

// POST /planet/{id}/hydrogen-storage/cancel
// Cancel the upgrade of the hydrogen storage
func serveCancelHydrogenStorage(w http.ResponseWriter, r *http.Request) {
	serveCancelBuilding(w, r, timer.ActionUpgradeHydrogenStorage)
}

// POST /planet/{id}/silica-storage/upgrade
// Upgrade the silica storage to the next level
func serveUpgradeSilicaStorage(w http.ResponseWriter, r *http.Request) {
	serveUpgradeBuilding(w, r, timer.ActionUpgradeSilicaStorage)
}

// POST /planet/{id}/silica-storage/cancel
// Cancel the upgrade of the silica storage
func serveCancelSilicaStorage(w http.ResponseWriter, r *http.Request) {
	serveCancelBuilding(w, r, timer.ActionUpgradeSilicaStorage)
}

// POST /planet/{id}/research-center/upgrade
// Upgrade the research center to the next level
func serveUpgradeResearchCenter(w http.ResponseWriter, r *http.Request) {
	serveUpgradeBuilding(w, r, timer.ActionUpgradeResearchCenter)
}

// POST /planet/{id}/research-center/cancel
// Cancel the upgrade of the research center
func serveCancelResearchCenter(w http.ResponseWriter, r *http.Request) {
	serveCancelBuilding(w, r, timer.ActionUpgradeResearchCenter)
}

// POST /planet/{id}/ship-factory/upgrade
// Upgrade the ship factory to the next level
func serveUpgradeShipFactory(w http.ResponseWriter, r *http.Request) {
	serveUpgradeBuilding(w, r, timer.ActionUpgradeShipFactory)
}

// POST /planet/{id}/ship-factory/cancel
// Cancel the upgrade of the ship factory
func serveCancelShipFactory(w http.ResponseWriter, r *http.Request) {
	serveCancelBuilding(w, r, timer.ActionUpgradeShipFactory)
}
