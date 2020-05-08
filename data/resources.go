package data

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/pdeguing/empire-and-foundation/ent"
)

func HasResources(p *ent.Planet, s Amounts) bool {
	return p.Metal >= s.Metal &&
		p.Hydrogen >= s.Hydrogen &&
		p.Silica >= s.Silica
}

func addStock(p *ent.Planet, s Amounts) {
	p.Metal += s.Metal
	p.Hydrogen += s.Hydrogen
	p.Silica += s.Silica
}

func subStock(p *ent.Planet, s Amounts) {
	p.Metal -= s.Metal
	p.Hydrogen -= s.Hydrogen
	p.Silica -= s.Silica
}

type PlanetWithResourceInfo struct {
	*ent.Planet
	Buildings                   []Building
	EnergyConsumption           int64
	PopulationEmployment        int64
	EnergyProduction            int64
	PopulationSize              int64
	UnpenalizedHourlyProduction Amounts
	PenalizedHourlyProduction   Amounts
	Capacity                    Amounts
	MetalStorageFull            bool
	HydrogenStorageFull         bool
	SilicaStorageFull           bool
	EnergyPenalty               bool
	PopulationPenalty           bool
}

func NewPlanetWithResourceInfo(p *ent.Planet) *PlanetWithResourceInfo {
	return &PlanetWithResourceInfo{
		Planet: p,
	}
}

func (p *PlanetWithResourceInfo) Update(now time.Time) {
	p.calcTotals()
	p.applyProductionPenalties()
	p.calcNewResourcesAndProduction(now)
}

func (p *PlanetWithResourceInfo) calcTotals() {
	for _, b := range Buildings {
		b := b.ForPlanet(p.Planet)

		usage := b.Usage()
		p.EnergyConsumption += usage.Energy
		p.PopulationEmployment += usage.Population

		supply := b.Supply()
		p.EnergyProduction += supply.Energy
		p.PopulationSize += supply.Population // TODO: replace with actual population size instead of capacity

		production := b.Production()
		p.UnpenalizedHourlyProduction = p.UnpenalizedHourlyProduction.Add(production)

		capacity := b.Capacity()
		p.Capacity = p.Capacity.Add(capacity)
	}
}

func (p *PlanetWithResourceInfo) applyProductionPenalties() {
	p.EnergyPenalty = p.EnergyProduction < p.EnergyConsumption
	p.PopulationPenalty = p.PopulationSize < p.PopulationEmployment

	energyProductivity := productivity(p.EnergyProduction, p.EnergyConsumption)
	populationProductivity := productivity(p.PopulationSize, p.PopulationEmployment)
	productivity := math.Min(energyProductivity, populationProductivity)

	p.PenalizedHourlyProduction = p.UnpenalizedHourlyProduction.MulFloat64(productivity)

	// TODO: it would be nice that if there is a scarcity of hydrogen, productivity of
	// 		 hydrogen will be reduced last to save the population.
}

func (p *PlanetWithResourceInfo) calcNewResourcesAndProduction(now time.Time) {
	metalFreeSpace := MaxInt64(p.Capacity.Metal-p.Metal, 0)
	hydroFreeSpace := MaxInt64(p.Capacity.Hydrogen-p.Hydrogen, 0)
	silicFreeSpace := MaxInt64(p.Capacity.Silica-p.Silica, 0)

	duration := now.Sub(p.LastResourceUpdate)
	metalProduction := p.PenalizedHourlyProduction.Metal * int64(duration) / int64(time.Hour)
	hydroProduction := p.PenalizedHourlyProduction.Hydrogen * int64(duration) / int64(time.Hour)
	silicProduction := p.PenalizedHourlyProduction.Silica * int64(duration) / int64(time.Hour)

	metalProduction = MinInt64(metalProduction, metalFreeSpace)
	hydroProduction = MinInt64(hydroProduction, hydroFreeSpace)
	silicProduction = MinInt64(silicProduction, silicFreeSpace)

	p.Metal += metalProduction
	p.Hydrogen += hydroProduction
	p.Silica += silicProduction
	p.LastResourceUpdate = now

	// Recalculate to get the new hourly production (for display in the ui)
	metalFreeSpace = MaxInt64(p.Capacity.Metal-p.Metal, 0)
	hydroFreeSpace = MaxInt64(p.Capacity.Hydrogen-p.Hydrogen, 0)
	silicFreeSpace = MaxInt64(p.Capacity.Silica-p.Silica, 0)

	p.MetalStorageFull = metalFreeSpace == 0
	p.HydrogenStorageFull = hydroFreeSpace == 0
	p.SilicaStorageFull = silicFreeSpace == 0

	p.PenalizedHourlyProduction.Metal = MinInt64(p.PenalizedHourlyProduction.Metal, metalFreeSpace)
	p.PenalizedHourlyProduction.Hydrogen = MinInt64(p.PenalizedHourlyProduction.Hydrogen, hydroFreeSpace)
	p.PenalizedHourlyProduction.Silica = MinInt64(p.PenalizedHourlyProduction.Silica, silicFreeSpace)
}

func productivity(available, usage int64) float64 {
	if usage == 0 {
		return 1
	}
	return math.Min(float64(available)/float64(usage), 1)
}

func solarPlantEnergyProduction(level int) int64 {
	const initial = 30
	const base = 1.25
	l := float64(level + 1) // The player gets one production level for free
	return int64(initial * l * math.Pow(base, l))
}

func populationStorageCapacity(level int) int64 {
	const initial = 10
	const base = 1.7
	l := float64(level + 1) // The player gets one capacity level for free
	return int64(initial * l * math.Pow(base, l))
}

func genericHourlyProductionRate(level int, initial, base float64) int64 {
	l := float64(level)
	return int64(l * initial * math.Pow(base, l))
}

func metalHourlyProductionRate(level int) int64 {
	return genericHourlyProductionRate(level, 120, 1.05)
}

func hydrogenHourlyProductionRate(level int) int64 {
	return genericHourlyProductionRate(level, 160, 1.08)
}

func silicaHourlyProductionRate(level int) int64 {
	return genericHourlyProductionRate(level, 160, 1.05)
}

func genericStorageCapacity(level int, initial, base float64) int64 {
	const expBase = 1.5
	l := float64(level)
	return int64(initial + base*l*l*math.Pow(expBase, l))
}

func metalStorageCapacity(level int) int64 {
	return genericStorageCapacity(level, 1000, 1000)
}

func hydrogenStorageCapacity(level int) int64 {
	return genericStorageCapacity(level, 1000, 1200)
}

func silicaStorageCapacity(level int) int64 {
	return genericStorageCapacity(level, 1000, 1100)
}

func genericConstructionCost(level int, initial, base float64) int64 {
	l := float64(level)
	prevL := float64(level - 1)
	upgradeCost := prevL * initial * (math.Pow(base, l) - math.Pow(base, prevL))
	newCost := initial * math.Pow(base, l)
	return int64(upgradeCost + newCost)
}

func genericPopulationEmployment(level int, initial, base float64) int64 {
	l := float64(level)
	return int64(initial * l * math.Pow(base, l))
}

func genericEnergyUsage(level int, initial, base float64) int64 {
	l := float64(level)
	return int64(initial * l * math.Pow(base, l))
}

func genericStorageCost(level int, initial, base float64) int64 {
	l := float64(level)
	return int64(initial * l * l * math.Pow(base, l))
}

func genericUpgradeDuration(level int, initial time.Duration, base float64) time.Duration {
	l := float64(level)
	i := float64(initial)
	return time.Duration(i * math.Pow(base, l))
}

func metalMineUpgradeDuration(level int) time.Duration {
	return genericUpgradeDuration(level, 104*time.Second, 1.4)
}

func hydrogenExtractorUpgradeDuration(level int) time.Duration {
	return genericUpgradeDuration(level, 86*time.Second, 1.4)
}

func silicaQuarryUpgradeDuration(level int) time.Duration {
	return genericUpgradeDuration(level, 78*time.Second, 1.4)
}

func urbanismUpgradeDuration(level int) time.Duration {
	return genericUpgradeDuration(level, 52*time.Second, 1.4)
}

func solarPlantUpgradeDuration(level int) time.Duration {
	return genericUpgradeDuration(level, 35*time.Second, 1.6)
}

func metalStorageUpgradeDuration(level int) time.Duration {
	return genericUpgradeDuration(level, 43*time.Second, 2.0)
}

func hydrogenStorageUpgradeDuration(level int) time.Duration {
	return genericUpgradeDuration(level, 52*time.Second, 2.0)
}

func silicaStorageUpgradeDuration(level int) time.Duration {
	return genericUpgradeDuration(level, 35*time.Second, 2.0)
}

func researchCenterUpgradeDuration(level int) time.Duration {
	return genericUpgradeDuration(level, 21*time.Minute+36*time.Second, 1.6)
}

func shipFactoryUpgradeDuration(level int) time.Duration {
	return genericUpgradeDuration(level, 7*time.Minute+12*time.Second, 1.5)
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
		SetResearchCenterLevel(p.ResearchCenterLevel).
		SetShipFactoryLevel(p.ShipFactoryLevel).
		SetLastResourceUpdate(p.LastResourceUpdate).
		Save(ctx)
	if err != nil {
		return p, fmt.Errorf("error while saving planet resource fields: %w", err)
	}
	return p, nil
}

// DebugPlanetResources returns a human-readable string with the planet's resource state.
//func DebugPlanetResources(p *ent.Planet) string {
//	pwr := NewPlanetWithResourceInfo(p, time.Now())
//
//	//energyProductivity := getProductivityBasedOnEnergyConsumption(p)
//	return fmt.Sprintf("Resources of %s on %v:\nMetal       lvl %d: %6d + %4d u/h (max %4d u/h)\nHydrogen    lvl %d: %6d + %4d u/h (max %4d u/h)\nSilica      lvl %d: %6d + %4d u/h (max %4d u/h)\nPopulation  lvl %d: %6d + %4d u/h (max %4d u/h)\nSolar plant lvl %d\nProductivity: %.1f%%\n",
//		p.Name,
//		p.LastResourceUpdate,
//		p.MetalProdLevel,
//		p.Metal,
//		getMetalRate(p, energyProductivity),
//		getMetalRate(p, 1),
//		p.HydrogenProdLevel,
//		p.Hydrogen,
//		getHydrogenRate(p, energyProductivity),
//		getHydrogenRate(p, 1),
//		p.SilicaProdLevel,
//		p.Silica,
//		getSilicaRate(p, energyProductivity),
//		getSilicaRate(p, 1),
//		p.PopulationProdLevel,
//		p.Population,
//		getPopulationRate(p, energyProductivity),
//		getPopulationRate(p, 1),
//		p.SolarProdLevel,
//		energyProductivity,
//	)
//}
