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

type constructionCost struct {
	data.Amounts
	EnergyConsumption    int64
	EnergyDelta          int64
	PopulationEmployment int64
	PopulationDelta      int64
}

type constructionCard struct {
	Level                int
	EnergyConsumption    int64
	PopulationEmployment int64
	Cost                 constructionCost
	Upgradable           bool
}

type constructionCards struct {
	MetalMine         constructionCard
	HydrogenExtractor constructionCard
	SilicaQuarry      constructionCard
	SolarPlant        constructionCard
	Urbanism          constructionCard
	MetalStorage      constructionCard
	HydrogenStorage   constructionCard
	SilicaStorage     constructionCard
	ResearchCenter    constructionCard
	ShipFactory       constructionCard
}

type constructionsViewData struct {
	planetViewData
	constructionCards
}

func getConstructionCards(planet *ent.Planet) constructionCards {
	metalMineCost := data.MetalMineCost(planet.MetalProdLevel + 1)
	hydrogenExtractorCost := data.HydrogenExtractorCost(planet.HydrogenProdLevel + 1)
	silicaQuarryCost := data.SilicaQuarryCost(planet.SilicaProdLevel + 1)
	solarPlantCost := data.SolarPlantCost(planet.SolarProdLevel + 1)
	urbanismCost := data.UrbanismCost(planet.PopulationProdLevel + 1)
	metalStorageCost := data.MetalStorageCost(planet.MetalStorageLevel + 1)
	hydrogenStorageCost := data.HydrogenStorageCost(planet.HydrogenStorageLevel + 1)
	silicaStorageCost := data.SilicaStorageCost(planet.SilicaStorageLevel + 1)
	researchCenterCost := data.ResearchCenterCost(planet.ResearchCenterLevel + 1)
	shipFactoryCost := data.ShipFactoryCost(planet.ShipFactoryLevel + 1)

	return constructionCards{
		MetalMine: constructionCard{
			Level: planet.MetalProdLevel,
			Cost: constructionCost{
				Amounts: metalMineCost,
			},
			Upgradable: data.HasResources(planet, metalMineCost),
		},
		HydrogenExtractor: constructionCard{
			Level: planet.HydrogenProdLevel,
			Cost: constructionCost{
				Amounts: hydrogenExtractorCost,
			},
			Upgradable: data.HasResources(planet, hydrogenExtractorCost),
		},
		SilicaQuarry: constructionCard{
			Level: planet.SilicaProdLevel,
			Cost: constructionCost{
				Amounts: silicaQuarryCost,
			},
			Upgradable: data.HasResources(planet, silicaQuarryCost),
		},
		SolarPlant: constructionCard{
			Level: planet.SolarProdLevel,
			Cost: constructionCost{
				Amounts: solarPlantCost,
			},
			Upgradable: data.HasResources(planet, solarPlantCost),
		},
		Urbanism: constructionCard{
			Level: planet.PopulationProdLevel,
			Cost: constructionCost{
				Amounts: urbanismCost,
			},
			Upgradable: data.HasResources(planet, urbanismCost),
		},
		MetalStorage: constructionCard{
			Level: planet.MetalStorageLevel,
			Cost: constructionCost{
				Amounts: metalStorageCost,
			},
			Upgradable: data.HasResources(planet, metalStorageCost),
		},
		HydrogenStorage: constructionCard{
			Level: planet.HydrogenStorageLevel,
			Cost: constructionCost{
				Amounts: hydrogenStorageCost,
			},
			Upgradable: data.HasResources(planet, hydrogenStorageCost),
		},
		SilicaStorage: constructionCard{
			Level: planet.SilicaStorageLevel,
			Cost: constructionCost{
				Amounts: silicaStorageCost,
			},
			Upgradable: data.HasResources(planet, silicaStorageCost),
		},
		ResearchCenter: constructionCard{
			Level: planet.ResearchCenterLevel,
			Cost: constructionCost{
				Amounts: researchCenterCost,
			},
			Upgradable: data.HasResources(planet, researchCenterCost),
		},
		ShipFactory: constructionCard{
			Level: planet.ShipFactoryLevel,
			Cost: constructionCost{
				Amounts: shipFactoryCost,
			},
			Upgradable: data.HasResources(planet, shipFactoryCost),
		},
	}
}

// GET /planet/{id}/constructions
// Show the constructions page for a planet
func serveConstructions(w http.ResponseWriter, r *http.Request) {
	p, err := newPlanetViewData(r, timer.GroupBuilding)
	if err != nil {
		serveError(w, r, err)
		return
	}
	b := getConstructionCards(p.Planet.Planet)
	vd := constructionsViewData{
		*p,
		b,
	}
	generateHTML(w, r, "planet-constructions", vd, "layout", "private.navbar", "dashboard", "leftbar", "planet.layout", "planet.header", "flash", "planet.constructions")
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
