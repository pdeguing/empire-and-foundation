package data

import (
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

// getNewMetalRate calculates the metal production and consumption per hour.
func getNewMetalRate(p *ent.Planet) int {
	return int(60 * 12 * float64(p.MetalProdLevel) * math.Pow(1.1, float64(p.MetalProdLevel)) * float64(p.EnergyProd / p.EnergyCons))
}

// getNewHydrogenRate calculates the hydrogen production and consumption per hour.
func getNewHydrogenRate(p *ent.Planet) int {
	return int(60 * 12 * float64(p.HydrogenProdLevel) * math.Pow(1.1, float64(p.HydrogenProdLevel)) * float64(p.EnergyProd / p.EnergyCons))
}

// getNewSilicaRate calculates the silica production and consumption per hour.
func getNewSilicaRate(p *ent.Planet) int {
	return int(60 * 12 * float64(p.SilicaProdLevel) * math.Pow(1.1, float64(p.SilicaProdLevel)) * float64(p.EnergyProd / p.EnergyCons))
}

// getNewPopulationRate calculates the population production and consumption per hour.
func getNewPopulationRate(p *ent.Planet) int {
	return int(60 * 12 * float64(p.SilicaProdLevel) * math.Pow(1.1, float64(p.SilicaProdLevel)) * float64(p.EnergyProd / p.EnergyCons))
}

// getEnergyCons calculates the current energy consumption
func getEnergyCons(p *ent.Planet) int64 {
	consumption := int64(500 * int64(p.MetalProdLevel) + int64(math.Pow(1.1, float64(p.MetalProdLevel))))
	consumption += int64(1000 * int64(p.HydrogenProdLevel) + int64(math.Pow(1.1, float64(p.HydrogenProdLevel))))
	consumption += int64(500 * int64(p.SilicaProdLevel) + int64(math.Pow(1.1, float64(p.SilicaProdLevel))))
	consumption += int64(250 * int64(p.PopulationProdLevel) + int64(math.Pow(1.1, float64(p.PopulationProdLevel))))
	return consumption
}

// getEnergyProd calculates the current energy production
func getEnergyProd(solarProdLevel int) int64 {
	return int64(1500 * int64(solarProdLevel) * int64(math.Pow(1.1, float64(solarProdLevel))))
}

// getNewStock calculates the current value in stock for a resource based on value and duration since last update.
func getNewStock(val int64, last time.Time, rate int, storageLevel int, now time.Time) int64 {
	duration := int64(now.Sub(last) / time.Second)
	maxStorage := getMaxStorage(storageLevel)
	const secondsPerHour = 60 * 60
	current := val + duration*int64(rate)/secondsPerHour
	if current >= maxStorage {
		return maxStorage
	}
	return current
}

// UpdatePlanetState updates the current planet struct to get up-to-date state
func UpdatePlanetState(p *ent.Planet, now time.Time) {
	p.Metal = getNewStock(
		p.Metal,
		p.MetalLastUpdate,
		p.MetalRate,
		p.MetalStorageLevel,
		now,
	)
	p.MetalLastUpdate = now
	p.Hydrogen = getNewStock(
		p.Hydrogen,
		p.HydrogenLastUpdate,
		p.HydrogenRate,
		p.HydrogenStorageLevel,
		now,
	)
	p.HydrogenLastUpdate = now
	p.Silica = getNewStock(
		p.Silica,
		p.SilicaLastUpdate,
		p.SilicaRate,
		p.SilicaStorageLevel,
		now,
	)
	p.SilicaLastUpdate = now
	p.Population = getNewStock(
		p.Population,
		p.PopulationLastUpdate,
		p.PopulationRate,
		p.PopulationStorageLevel,
		now,
	)
	p.PopulationLastUpdate = now
	p.EnergyCons = getEnergyCons(p)
	p.EnergyProd = getEnergyProd(p.SolarProdLevel)
}
