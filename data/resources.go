package data

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/pdeguing/empire-and-foundation/ent"
)

// getMaxStorage calculates the storage capacity for a resource with given storage level.
// TODO: Create separate methods for different resource types.
func getMaxStorage(storageLevel int) int64 {
	maxStorage := 100000 * int64(storageLevel) * int64(math.Pow(1.1, float64(storageLevel)))
	return maxStorage
}

// getProductivityBasedOnEnergyConsumption calculates the productivity of the mines based
// on the energy availability or shortage.
func getProductivityBasedOnEnergyConsumption(p *ent.Planet) (productivity float64) {
	prod := GetEnergyProd(p)
	cons := GetEnergyCons(p)
	if cons <= prod {
		return 1
	}
	return float64(prod) / float64(cons)
}

// GetEnergyProd calculates the current energy production
func GetEnergyProd(p *ent.Planet) int64 {
	solarProd := 1500 * int64(p.SolarProdLevel) * int64(math.Pow(1.1, float64(p.SolarProdLevel)))
	return solarProd
}

// GetEnergyCons calculates the current energy consumption
func GetEnergyCons(p *ent.Planet) int64 {
	metalCons := 500*int64(p.MetalProdLevel) * int64(math.Pow(1.1, float64(p.MetalProdLevel)))
	hydrogenCons := 1000*int64(p.HydrogenProdLevel) * int64(math.Pow(1.1, float64(p.HydrogenProdLevel)))
	silicaCons := 500*int64(p.SilicaProdLevel) * int64(math.Pow(1.1, float64(p.SilicaProdLevel)))
	populationCons := 250*int64(p.PopulationProdLevel) * int64(math.Pow(1.1, float64(p.PopulationProdLevel)))
	return metalCons + hydrogenCons + silicaCons + populationCons
}

// getMetalRate calculates the metal production and consumption per hour.
func getMetalRate(p *ent.Planet, productivity float64) int64 {
	return int64(60 * 12 * float64(p.MetalProdLevel) * math.Pow(1.1, float64(p.MetalProdLevel)) * productivity)
}

// getHydrogenRate calculates the hydrogen production and consumption per hour.
func getHydrogenRate(p *ent.Planet, productivity float64) int64 {
	return int64(60 * 12 * float64(p.HydrogenProdLevel) * math.Pow(1.1, float64(p.HydrogenProdLevel)) * productivity)
}

// getSilicaRate calculates the silica production and consumption per hour.
func getSilicaRate(p *ent.Planet, productivity float64) int64 {
	return int64(60 * 12 * float64(p.SilicaProdLevel) * math.Pow(1.1, float64(p.SilicaProdLevel)) * productivity)
}

// getPopulationRate calculates the population production and consumption per hour.
func getPopulationRate(p *ent.Planet, productivity float64) int64 {
	return int64(60 * 12 * float64(p.PopulationProdLevel) * math.Pow(1.1, float64(p.PopulationProdLevel)) * productivity)
}

// getNewStock calculates the current value in stock for a resource based on value and duration since last update.
func getNewStock(val int64, last time.Time, rate int64, storageLevel int, now time.Time) int64 {
	duration := int64(now.Sub(last) / time.Second)
	maxStorage := getMaxStorage(storageLevel)
	const secondsPerHour = 60 * 60
	current := val + duration*rate/secondsPerHour
	if current >= maxStorage {
		return maxStorage
	}
	return current
}

// UpdatePlanetResources updates the current planet struct to get up-to-date state
func UpdatePlanetResources(p *ent.Planet, now time.Time) {
	energyProductivity := getProductivityBasedOnEnergyConsumption(p)

	p.Metal = getNewStock(
		p.Metal,
		p.LastResourceUpdate,
		getMetalRate(p, energyProductivity),
		p.MetalStorageLevel,
		now,
	)
	p.Hydrogen = getNewStock(
		p.Hydrogen,
		p.LastResourceUpdate,
		getHydrogenRate(p, energyProductivity),
		p.HydrogenStorageLevel,
		now,
	)
	p.Silica = getNewStock(
		p.Silica,
		p.LastResourceUpdate,
		getSilicaRate(p, energyProductivity),
		p.SilicaStorageLevel,
		now,
	)
	p.Population = getNewStock(
		p.Population,
		p.LastResourceUpdate,
		getPopulationRate(p, energyProductivity),
		p.PopulationStorageLevel,
		now,
	)
	p.LastResourceUpdate = now
}

// SavePlanetResources saves all fields related to the resources to the database.
func SavePlanetResources(ctx context.Context, p *ent.Planet) (*ent.Planet, error) {
	p, err := p.Update().
		SetMetal(p.Metal).
		SetMetalProdLevel(p.MetalProdLevel).
		SetMetalStorageLevel(p.MetalStorageLevel).
		SetHydrogen(p.Hydrogen).
		SetHydrogenProdLevel(p.HydrogenProdLevel).
		SetHydrogenStorageLevel(p.HydrogenStorageLevel).
		SetSilica(p.Silica).
		SetSilicaProdLevel(p.SilicaProdLevel).
		SetSilicaStorageLevel(p.SilicaStorageLevel).
		SetPopulation(p.Population).
		SetPopulationProdLevel(p.PopulationProdLevel).
		SetPopulationStorageLevel(p.PopulationStorageLevel).
		SetSolarProdLevel(p.SolarProdLevel).
		SetLastResourceUpdate(p.LastResourceUpdate).
		Save(ctx)
	if err != nil {
		return p, fmt.Errorf("error while saving planet resource fields: %w", err)
	}
	return p, nil
}

// DebugPlanetResources returns a human-readable string with the planet's resource state.
func DebugPlanetResources(p *ent.Planet) string {
	energyProductivity := getProductivityBasedOnEnergyConsumption(p)
	return fmt.Sprintf("Resources of %s on %v:\nMetal       lvl %d: %6d + %4d u/h (max %4d u/h)\nHydrogen    lvl %d: %6d + %4d u/h (max %4d u/h)\nSilica      lvl %d: %6d + %4d u/h (max %4d u/h)\nPopulation  lvl %d: %6d + %4d u/h (max %4d u/h)\nSolar plant lvl %d\nProductivity: %.1f%%\n",
		p.Name,
		p.LastResourceUpdate,
		p.MetalProdLevel,
		p.Metal,
		getMetalRate(p, energyProductivity),
		getMetalRate(p, 1),
		p.HydrogenProdLevel,
		p.Hydrogen,
		getHydrogenRate(p, energyProductivity),
		getHydrogenRate(p, 1),
		p.SilicaProdLevel,
		p.Silica,
		getSilicaRate(p, energyProductivity),
		getSilicaRate(p, 1),
		p.PopulationProdLevel,
		p.Population,
		getPopulationRate(p, energyProductivity),
		getPopulationRate(p, 1),
		p.SolarProdLevel,
		energyProductivity,
	)
}
