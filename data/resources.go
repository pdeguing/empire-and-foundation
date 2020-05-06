package data

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/pdeguing/empire-and-foundation/ent"
)

type PlanetWithResourceInfo struct {
	*ent.Planet
	EnergyConsumption           int64
	EnergyProduction            int64
	PopulationSize              int64
	PopulationEmployment        int64
	MetalMineProduction         int64
	HydrogenExtractorProduction int64
	SilicaQuarryProduction      int64
	MetalStorageCapacity        int64
	HydrogenStorageCapacity     int64
	SilicaStorageCapacity       int64
	EnergyPenalty               bool
	PopulationPenalty           bool
	MetalStorageFull            bool
	HydrogenStorageFull         bool
	SilicaStorageFull           bool
}

func NewPlanetWithResourceInfo(p *ent.Planet) *PlanetWithResourceInfo {
	return &PlanetWithResourceInfo{
		Planet: p,
	}
}

func (p *PlanetWithResourceInfo) Update(now time.Time) {
	p.calcEnergyConsumption()
	p.calcEnergyProduction()
	p.calcPopulationSize()
	p.calcPopulationEmployment()

	p.calcMetalMineProduction()
	p.calcHydrogenExtractorProduction()
	p.calcSilicaQuarryProduction()

	p.applyProductionPenalties()
	p.calcNewResourcesAndProduction(now)
}

func (p *PlanetWithResourceInfo) calcEnergyConsumption() {
	metal := genericEnergyUsage(p.MetalProdLevel, 5, 1.25)
	hydro := genericEnergyUsage(p.HydrogenProdLevel, 15, 1.255)
	silic := genericEnergyUsage(p.SilicaProdLevel, 8, 1.25)
	urban := genericEnergyUsage(p.PopulationProdLevel, 10, 1.2)
	rsrch := genericEnergyUsage(p.ResearchCenterLevel, 40, 1.25)
	shipf := genericEnergyUsage(p.ShipFactoryLevel, 15, 1.25)
	p.EnergyConsumption = metal + hydro + silic + urban + rsrch + shipf
}

func (p *PlanetWithResourceInfo) calcEnergyProduction() {
	solar := SolarPlantEnergyProduction(p.SolarProdLevel)
	p.EnergyProduction = solar
}

func (p *PlanetWithResourceInfo) calcPopulationSize() {
	urban := PopulationStorageCapacity(p.PopulationProdLevel) // TODO: replace with actual population size instead of capacity
	p.PopulationSize = urban
}

func (p *PlanetWithResourceInfo) calcPopulationEmployment() {
	metal := genericPopulationEmployment(p.MetalProdLevel, 6, 1.75)
	hydro := genericPopulationEmployment(p.HydrogenProdLevel, 5, 1.6)
	silic := genericPopulationEmployment(p.SilicaProdLevel, 6, 1.75)
	solar := genericPopulationEmployment(p.SolarProdLevel, 2, 1.6)
	rsrch := genericPopulationEmployment(p.ResearchCenterLevel, 40, 1.7)
	shipf := genericPopulationEmployment(p.ShipFactoryLevel, 30, 1.7)
	p.PopulationEmployment = metal + hydro + silic + solar + rsrch + shipf
}

func (p *PlanetWithResourceInfo) calcMetalMineProduction() {
	p.MetalMineProduction = MetalHourlyProductionRate(p.MetalProdLevel)
}

func (p *PlanetWithResourceInfo) calcHydrogenExtractorProduction() {
	p.HydrogenExtractorProduction = HydrogenHourlyProductionRate(p.HydrogenProdLevel)
}

func (p *PlanetWithResourceInfo) calcSilicaQuarryProduction() {
	p.SilicaQuarryProduction = SilicaHourlyProductionRate(p.SilicaProdLevel)
}

func (p *PlanetWithResourceInfo) applyProductionPenalties() {
	p.EnergyPenalty = p.EnergyProduction < p.EnergyConsumption
	p.PopulationPenalty = p.PopulationSize < p.PopulationEmployment

	energyProductivity := productivity(p.EnergyProduction, p.EnergyConsumption)
	populationProductivity := productivity(p.PopulationSize, p.PopulationEmployment)
	productivity := math.Min(energyProductivity, populationProductivity)

	p.MetalMineProduction = int64(float64(p.MetalMineProduction) * productivity)
	p.HydrogenExtractorProduction = int64(float64(p.HydrogenExtractorProduction) * productivity)
	p.SilicaQuarryProduction = int64(float64(p.SilicaQuarryProduction) * productivity)

	// TODO: it would be nice that if there is a scarcity of hydrogen, productivity of
	// 		 hydrogen will be reduced last to save the population.
}

func (p *PlanetWithResourceInfo) calcNewResourcesAndProduction(now time.Time) {
	p.MetalStorageCapacity = MetalStorageCapacity(p.MetalStorageLevel)
	p.HydrogenStorageCapacity = HydrogenStorageCapacity(p.HydrogenStorageLevel)
	p.SilicaStorageCapacity = SilicaStorageCapacity(p.SilicaStorageLevel)

	metalFreeSpace := MaxInt64(p.MetalStorageCapacity-p.Metal, 0)
	hydroFreeSpace := MaxInt64(p.HydrogenStorageCapacity-p.Hydrogen, 0)
	silicFreeSpace := MaxInt64(p.SilicaStorageCapacity-p.Silica, 0)

	duration := now.Sub(p.LastResourceUpdate)
	metalProduction := p.MetalMineProduction * int64(duration) / int64(time.Hour)
	hydroProduction := p.HydrogenExtractorProduction * int64(duration) / int64(time.Hour)
	silicProduction := p.SilicaQuarryProduction * int64(duration) / int64(time.Hour)

	metalProduction = MinInt64(metalProduction, metalFreeSpace)
	hydroProduction = MinInt64(hydroProduction, hydroFreeSpace)
	silicProduction = MinInt64(silicProduction, silicFreeSpace)

	p.Metal += metalProduction
	p.Hydrogen += hydroProduction
	p.Silica += silicProduction
	p.LastResourceUpdate = now

	// Recalculate to get the new hourly production (for display in the ui)
	metalFreeSpace = MaxInt64(p.MetalStorageCapacity-p.Metal, 0)
	hydroFreeSpace = MaxInt64(p.HydrogenStorageCapacity-p.Hydrogen, 0)
	silicFreeSpace = MaxInt64(p.SilicaStorageCapacity-p.Silica, 0)

	p.MetalStorageFull = metalFreeSpace == 0
	p.HydrogenStorageFull = hydroFreeSpace == 0
	p.SilicaStorageFull = silicFreeSpace == 0

	p.MetalMineProduction = MinInt64(p.MetalMineProduction, metalFreeSpace)
	p.HydrogenExtractorProduction = MinInt64(p.HydrogenExtractorProduction, hydroFreeSpace)
	p.SilicaQuarryProduction = MinInt64(p.SilicaQuarryProduction, silicFreeSpace)
}

// Amounts stores the cost or capacity of something.
type Amounts struct {
	Metal    int64
	Hydrogen int64
	Silica   int64
}

func productivity(available, usage int64) float64 {
	if usage == 0 {
		return 1
	}
	return math.Min(float64(available)/float64(usage), 1)
}

func SolarPlantEnergyProduction(level int) int64 {
	const initial = 30
	const base = 1.25
	l := float64(level + 1) // The player gets one production level for free
	return int64(initial * l * math.Pow(base, l))
}

func PopulationStorageCapacity(level int) int64 {
	const initial = 10
	const base = 1.7
	l := float64(level + 1) // The player gets one capacity level for free
	return int64(initial * l * math.Pow(base, l))
}

func genericHourlyProductionRate(level int, initial, base float64) int64 {
	l := float64(level)
	return int64(l * initial * math.Pow(base, l))
}

func MetalHourlyProductionRate(level int) int64 {
	return genericHourlyProductionRate(level, 120, 1.05)
}

func HydrogenHourlyProductionRate(level int) int64 {
	return genericHourlyProductionRate(level, 160, 1.08)
}

func SilicaHourlyProductionRate(level int) int64 {
	return genericHourlyProductionRate(level, 160, 1.05)
}

func genericStorageCapacity(level int, initial, base float64) int64 {
	const expBase = 1.5
	l := float64(level)
	return int64(initial + base*l*l*math.Pow(expBase, l))
}

func MetalStorageCapacity(level int) int64 {
	return genericStorageCapacity(level, 1000, 1000)
}

func HydrogenStorageCapacity(level int) int64 {
	return genericStorageCapacity(level, 1000, 1200)
}

func SilicaStorageCapacity(level int) int64 {
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

func MetalMineCost(level int) Amounts {
	return Amounts{
		Metal:    genericConstructionCost(level, 40, 1.5),
		Hydrogen: genericConstructionCost(level, 30, 1.5),
		Silica:   genericConstructionCost(level, 10, 1.5),
	}
}

func HydrogenExtractorCost(level int) Amounts {
	return Amounts{
		Metal:    genericConstructionCost(level, 50, 1.5),
		Hydrogen: genericConstructionCost(level, 10, 1.5),
		Silica:   genericConstructionCost(level, 40, 1.5),
	}
}

func SilicaQuarryCost(level int) Amounts {
	return Amounts{
		Metal:    genericConstructionCost(level, 60, 1.5),
		Hydrogen: genericConstructionCost(level, 35, 1.5),
		Silica:   genericConstructionCost(level, 5, 1.5),
	}
}

func UrbanismCost(level int) Amounts {
	return Amounts{
		Metal:    genericConstructionCost(level, 10, 1.5),
		Hydrogen: genericConstructionCost(level, 30, 1.5),
		Silica:   genericConstructionCost(level, 60, 1.5),
	}
}

func SolarPlantCost(level int) Amounts {
	return Amounts{
		Metal:    genericConstructionCost(level, 30, 1.5),
		Hydrogen: genericConstructionCost(level, 5, 1.5),
		Silica:   genericConstructionCost(level, 40, 1.5),
	}
}

func genericStorageCost(level int, initial, base float64) int64 {
	l := float64(level)
	return int64(initial * l * l * math.Pow(base, l))
}

func MetalStorageCost(level int) Amounts {
	return Amounts{
		Metal:    genericStorageCost(level, 30, 1.5),
		Hydrogen: genericStorageCost(level, 5, 1.5),
		Silica:   genericStorageCost(level, 40, 1.5),
	}
}

func HydrogenStorageCost(level int) Amounts {
	return Amounts{
		Metal:    genericStorageCost(level, 35, 1.5),
		Hydrogen: 0,
		Silica:   genericStorageCost(level, 35, 1.5),
	}
}

func SilicaStorageCost(level int) Amounts {
	return Amounts{
		Metal:    genericStorageCost(level, 40, 1.5),
		Hydrogen: genericStorageCost(level, 5, 1.5),
		Silica:   genericStorageCost(level, 20, 1.5),
	}
}

func ResearchCenterCost(level int) Amounts {
	return Amounts{
		Metal:    genericConstructionCost(level, 115, 1.5),
		Hydrogen: genericConstructionCost(level, 450, 1.5),
		Silica:   genericConstructionCost(level, 700, 1.5),
	}
}

func ShipFactoryCost(level int) Amounts {
	return Amounts{
		Metal:    genericConstructionCost(level, 600, 1.5),
		Hydrogen: genericConstructionCost(level, 100, 1.5),
		Silica:   genericConstructionCost(level, 350, 1.5),
	}
}

func genericUpgradeDuration(level int, initial time.Duration, base float64) time.Duration {
	l := float64(level)
	i := float64(initial)
	return time.Duration(i * math.Pow(base, l))
}

func MetalMineUpgradeDuration(level int) time.Duration {
	return genericUpgradeDuration(level, 104*time.Second, 1.4)
}

func HydrogenExtractorUpgradeDuration(level int) time.Duration {
	return genericUpgradeDuration(level, 86*time.Second, 1.4)
}

func SilicaQuarryUpgradeDuration(level int) time.Duration {
	return genericUpgradeDuration(level, 78*time.Second, 1.4)
}

func UrbanismUpgradeDuration(level int) time.Duration {
	return genericUpgradeDuration(level, 52*time.Second, 1.4)
}

func SolarPlantUpgradeDuration(level int) time.Duration {
	return genericUpgradeDuration(level, 35*time.Second, 1.6)
}

func MetalStorageUpgradeDuration(level int) time.Duration {
	return genericUpgradeDuration(level, 43*time.Second, 2.0)
}

func HydrogenStorageUpgradeDuration(level int) time.Duration {
	return genericUpgradeDuration(level, 52*time.Second, 2.0)
}

func SilicaStorageUpgradeDuration(level int) time.Duration {
	return genericUpgradeDuration(level, 35*time.Second, 2.0)
}

func ResearchCenterUpgradeDuration(level int) time.Duration {
	return genericUpgradeDuration(level, 21*time.Minute+36*time.Second, 1.6)
}

func ShipFactoryUpgradeDuration(level int) time.Duration {
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
