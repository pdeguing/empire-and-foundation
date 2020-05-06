package data

import (
	"github.com/pdeguing/empire-and-foundation/ent"
	"github.com/pdeguing/empire-and-foundation/ent/timer"
	"time"
)

var Buildings = []Building{
	MetalMine{},
	HydrogenExtractor{},
	SilicaQuarry{},
	SolarPlant{},
	Urbanism{},
	MetalStorage{},
	HydrogenStorage{},
	SilicaStorage{},
	ResearchCenter{},
	ShipFactory{},
}

// Amounts stores the cost, capacity, or production of an in-game element.
type Amounts struct {
	Metal    int64
	Hydrogen int64
	Silica   int64
}

func (a Amounts) Add(b Amounts) Amounts {
	a.Metal += b.Metal
	a.Hydrogen += b.Hydrogen
	a.Silica += b.Silica
	return a
}

func (a Amounts) MulFloat64(m float64) Amounts {
	a.Metal = int64(float64(a.Metal) * m)
	a.Hydrogen = int64(float64(a.Hydrogen) * m)
	a.Silica = int64(float64(a.Silica) * m)
	return a
}

// Levels stores the usage or supply of energy and population on a planet.
type Levels struct {
	Energy     int64
	Population int64
}

func (l Levels) Sub(m Levels) Levels {
	l.Energy -= m.Energy
	l.Population -= m.Population
	return l
}

type Building interface {
	// UpgradeAction references the timer action that can be used to upgrade or
	// cancel the upgrade of this building.
	UpgradeAction() timer.Action

	// LevelOnPlanet returns the level of this building on the specified planet.
	LevelOnPlanet(p *ent.Planet) int

	// UpgradeBuildingOnPlanet upgrades this building on the specified planet.
	UpgradeBuildingOnPlanet(p *ent.Planet)

	// ForPlanet returns a copy of the building, but with the level set to that
	// of the building on the given planet. Successive calls to the cost/production/
	// capacity/duration functions will be calculated using this level.
	ForPlanet(p *ent.Planet) Building

	// NextLevel returns a copy of the building, but with a level that is 1 higher.
	NextLevel() Building

	// Cost returns the cost of the building at the current level.
	Cost() Amounts

	// Production returns the resources produced hourly by this building.
	Production() Amounts

	// Usage returns the continuous energy usage and population employment.
	Usage() Levels

	// Supply returns the continuous energy supply and population size.
	Supply() Levels

	// Capacity returns the maximum amount of resources that can be stored
	// before the facilities will halt production.
	Capacity() Amounts

	// BuildDuration returns the time it takes to get from the previous level
	// of this building to the current one.
	BuildDuration() time.Duration
}

type MetalMine struct {
	Level int
}

func (m MetalMine) UpgradeAction() timer.Action {
	return timer.ActionUpgradeMetalProd
}

func (m MetalMine) LevelOnPlanet(p *ent.Planet) int {
	return p.MetalProdLevel
}

func (m MetalMine) UpgradeBuildingOnPlanet(p *ent.Planet) {
	p.MetalProdLevel++
}

func (m MetalMine) ForPlanet(p *ent.Planet) Building {
	m.Level = p.MetalProdLevel
	return m
}

func (m MetalMine) NextLevel() Building {
	m.Level++
	return m
}

func (m MetalMine) Cost() Amounts {
	return Amounts{
		Metal:    genericConstructionCost(m.Level, 40, 1.5),
		Hydrogen: genericConstructionCost(m.Level, 30, 1.5),
		Silica:   genericConstructionCost(m.Level, 10, 1.5),
	}
}

func (m MetalMine) Production() Amounts {
	return Amounts{
		Metal: metalHourlyProductionRate(m.Level),
	}
}

func (m MetalMine) Usage() Levels {
	return Levels{
		Energy:     genericEnergyUsage(m.Level, 5, 1.25),
		Population: genericPopulationEmployment(m.Level, 6, 1.75),
	}
}

func (m MetalMine) Supply() Levels {
	return Levels{}
}

func (m MetalMine) Capacity() Amounts {
	return Amounts{}
}

func (m MetalMine) BuildDuration() time.Duration {
	return metalMineUpgradeDuration(m.Level)
}

type HydrogenExtractor struct {
	Level int
}

func (h HydrogenExtractor) UpgradeAction() timer.Action {
	return timer.ActionUpgradeHydrogenProd
}

func (h HydrogenExtractor) LevelOnPlanet(p *ent.Planet) int {
	return p.HydrogenProdLevel
}

func (h HydrogenExtractor) UpgradeBuildingOnPlanet(p *ent.Planet) {
	p.HydrogenProdLevel++
}

func (h HydrogenExtractor) ForPlanet(p *ent.Planet) Building {
	h.Level = p.HydrogenProdLevel
	return h
}

func (h HydrogenExtractor) NextLevel() Building {
	h.Level++
	return h
}

func (h HydrogenExtractor) Cost() Amounts {
	return Amounts{
		Metal:    genericConstructionCost(h.Level, 50, 1.5),
		Hydrogen: genericConstructionCost(h.Level, 10, 1.5),
		Silica:   genericConstructionCost(h.Level, 40, 1.5),
	}
}

func (h HydrogenExtractor) Production() Amounts {
	return Amounts{
		Hydrogen: hydrogenHourlyProductionRate(h.Level),
	}
}

func (h HydrogenExtractor) Usage() Levels {
	return Levels{
		Energy:     genericEnergyUsage(h.Level, 15, 1.255),
		Population: genericPopulationEmployment(h.Level, 5, 1.6),
	}
}

func (h HydrogenExtractor) Supply() Levels {
	return Levels{}
}

func (h HydrogenExtractor) Capacity() Amounts {
	return Amounts{}
}

func (h HydrogenExtractor) BuildDuration() time.Duration {
	return hydrogenExtractorUpgradeDuration(h.Level)
}

type SilicaQuarry struct {
	Level int
}

func (s SilicaQuarry) UpgradeAction() timer.Action {
	return timer.ActionUpgradeSilicaProd
}

func (s SilicaQuarry) LevelOnPlanet(p *ent.Planet) int {
	return p.SilicaProdLevel
}

func (s SilicaQuarry) UpgradeBuildingOnPlanet(p *ent.Planet) {
	p.SilicaProdLevel++
}

func (s SilicaQuarry) ForPlanet(p *ent.Planet) Building {
	s.Level = p.SilicaProdLevel
	return s
}

func (s SilicaQuarry) NextLevel() Building {
	s.Level++
	return s
}

func (s SilicaQuarry) Cost() Amounts {
	return Amounts{
		Metal:    genericConstructionCost(s.Level, 60, 1.5),
		Hydrogen: genericConstructionCost(s.Level, 35, 1.5),
		Silica:   genericConstructionCost(s.Level, 5, 1.5),
	}
}

func (s SilicaQuarry) Production() Amounts {
	return Amounts{
		Silica: silicaHourlyProductionRate(s.Level),
	}
}

func (s SilicaQuarry) Usage() Levels {
	return Levels{
		Energy:     genericEnergyUsage(s.Level, 8, 1.25),
		Population: genericPopulationEmployment(s.Level, 6, 1.75),
	}
}

func (s SilicaQuarry) Supply() Levels {
	return Levels{}
}

func (s SilicaQuarry) Capacity() Amounts {
	return Amounts{}
}

func (s SilicaQuarry) BuildDuration() time.Duration {
	return silicaQuarryUpgradeDuration(s.Level)
}

type Urbanism struct {
	Level int
}

func (u Urbanism) UpgradeAction() timer.Action {
	return timer.ActionUpgradeUrbanism
}

func (u Urbanism) LevelOnPlanet(p *ent.Planet) int {
	return p.PopulationProdLevel
}

func (u Urbanism) UpgradeBuildingOnPlanet(p *ent.Planet) {
	p.PopulationProdLevel++
}

func (u Urbanism) ForPlanet(p *ent.Planet) Building {
	u.Level = p.PopulationProdLevel
	return u
}

func (u Urbanism) NextLevel() Building {
	u.Level++
	return u
}

func (u Urbanism) Cost() Amounts {
	return Amounts{
		Metal:    genericConstructionCost(u.Level, 10, 1.5),
		Hydrogen: genericConstructionCost(u.Level, 30, 1.5),
		Silica:   genericConstructionCost(u.Level, 60, 1.5),
	}
}

func (u Urbanism) Production() Amounts {
	return Amounts{}
}

func (u Urbanism) Usage() Levels {
	return Levels{
		Energy: genericEnergyUsage(u.Level, 10, 1.2),
	}
}

func (u Urbanism) Supply() Levels {
	return Levels{
		Population: populationStorageCapacity(u.Level),
	}
}

func (u Urbanism) Capacity() Amounts {
	return Amounts{}
}

func (u Urbanism) BuildDuration() time.Duration {
	return urbanismUpgradeDuration(u.Level)
}

type SolarPlant struct {
	Level int
}

func (s SolarPlant) UpgradeAction() timer.Action {
	return timer.ActionUpgradeSolarProd
}

func (s SolarPlant) LevelOnPlanet(p *ent.Planet) int {
	return p.SolarProdLevel
}

func (s SolarPlant) UpgradeBuildingOnPlanet(p *ent.Planet) {
	p.SolarProdLevel++
}

func (s SolarPlant) ForPlanet(p *ent.Planet) Building {
	s.Level = p.SolarProdLevel
	return s
}

func (s SolarPlant) NextLevel() Building {
	s.Level++
	return s
}

func (s SolarPlant) Cost() Amounts {
	return Amounts{
		Metal:    genericConstructionCost(s.Level, 30, 1.5),
		Hydrogen: genericConstructionCost(s.Level, 5, 1.5),
		Silica:   genericConstructionCost(s.Level, 40, 1.5),
	}
}

func (s SolarPlant) Production() Amounts {
	return Amounts{}
}

func (s SolarPlant) Usage() Levels {
	return Levels{
		Population: genericPopulationEmployment(s.Level, 2, 1.6),
	}
}

func (s SolarPlant) Supply() Levels {
	return Levels{
		Energy: solarPlantEnergyProduction(s.Level),
	}
}

func (s SolarPlant) Capacity() Amounts {
	return Amounts{}
}

func (s SolarPlant) BuildDuration() time.Duration {
	return solarPlantUpgradeDuration(s.Level)
}

type MetalStorage struct {
	Level int
}

func (m MetalStorage) UpgradeAction() timer.Action {
	return timer.ActionUpgradeMetalStorage
}

func (m MetalStorage) LevelOnPlanet(p *ent.Planet) int {
	return p.MetalStorageLevel
}

func (m MetalStorage) UpgradeBuildingOnPlanet(p *ent.Planet) {
	p.MetalStorageLevel++
}

func (m MetalStorage) ForPlanet(p *ent.Planet) Building {
	m.Level = p.MetalStorageLevel
	return m
}

func (m MetalStorage) NextLevel() Building {
	m.Level++
	return m
}

func (m MetalStorage) Cost() Amounts {
	return Amounts{
		Metal:    genericStorageCost(m.Level, 30, 1.5),
		Hydrogen: genericStorageCost(m.Level, 5, 1.5),
		Silica:   genericStorageCost(m.Level, 40, 1.5),
	}
}

func (m MetalStorage) Production() Amounts {
	return Amounts{}
}

func (m MetalStorage) Usage() Levels {
	return Levels{}
}

func (m MetalStorage) Supply() Levels {
	return Levels{}
}

func (m MetalStorage) Capacity() Amounts {
	return Amounts{
		Metal: metalStorageCapacity(m.Level),
	}
}

func (m MetalStorage) BuildDuration() time.Duration {
	return metalStorageUpgradeDuration(m.Level)
}

type HydrogenStorage struct {
	Level int
}

func (h HydrogenStorage) UpgradeAction() timer.Action {
	return timer.ActionUpgradeHydrogenStorage
}

func (h HydrogenStorage) LevelOnPlanet(p *ent.Planet) int {
	return p.HydrogenStorageLevel
}

func (h HydrogenStorage) UpgradeBuildingOnPlanet(p *ent.Planet) {
	p.HydrogenStorageLevel++
}

func (h HydrogenStorage) ForPlanet(p *ent.Planet) Building {
	h.Level = p.HydrogenStorageLevel
	return h
}

func (h HydrogenStorage) NextLevel() Building {
	h.Level++
	return h
}

func (h HydrogenStorage) Cost() Amounts {
	return Amounts{
		Metal:    genericStorageCost(h.Level, 35, 1.5),
		Hydrogen: 0,
		Silica:   genericStorageCost(h.Level, 35, 1.5),
	}
}

func (h HydrogenStorage) Production() Amounts {
	return Amounts{}
}

func (h HydrogenStorage) Usage() Levels {
	return Levels{}
}

func (h HydrogenStorage) Supply() Levels {
	return Levels{}
}

func (h HydrogenStorage) Capacity() Amounts {
	return Amounts{
		Hydrogen: hydrogenStorageCapacity(h.Level),
	}
}

func (h HydrogenStorage) BuildDuration() time.Duration {
	return hydrogenStorageUpgradeDuration(h.Level)
}

type SilicaStorage struct {
	Level int
}

func (s SilicaStorage) UpgradeAction() timer.Action {
	return timer.ActionUpgradeSilicaStorage
}

func (s SilicaStorage) LevelOnPlanet(p *ent.Planet) int {
	return p.SilicaStorageLevel
}

func (s SilicaStorage) UpgradeBuildingOnPlanet(p *ent.Planet) {
	p.SilicaStorageLevel++
}

func (s SilicaStorage) ForPlanet(p *ent.Planet) Building {
	s.Level = p.SilicaStorageLevel
	return s
}

func (s SilicaStorage) NextLevel() Building {
	s.Level++
	return s
}

func (s SilicaStorage) Cost() Amounts {
	return Amounts{
		Metal:    genericStorageCost(s.Level, 40, 1.5),
		Hydrogen: genericStorageCost(s.Level, 5, 1.5),
		Silica:   genericStorageCost(s.Level, 20, 1.5),
	}
}

func (s SilicaStorage) Production() Amounts {
	return Amounts{}
}

func (s SilicaStorage) Usage() Levels {
	return Levels{}
}

func (s SilicaStorage) Supply() Levels {
	return Levels{}
}

func (s SilicaStorage) Capacity() Amounts {
	return Amounts{
		Silica: silicaStorageCapacity(s.Level),
	}
}

func (s SilicaStorage) BuildDuration() time.Duration {
	return silicaStorageUpgradeDuration(s.Level)
}

type ResearchCenter struct {
	Level int
}

func (r ResearchCenter) UpgradeAction() timer.Action {
	return timer.ActionUpgradeResearchCenter
}

func (r ResearchCenter) LevelOnPlanet(p *ent.Planet) int {
	return p.ResearchCenterLevel
}

func (r ResearchCenter) UpgradeBuildingOnPlanet(p *ent.Planet) {
	p.ResearchCenterLevel++
}

func (r ResearchCenter) ForPlanet(p *ent.Planet) Building {
	r.Level = p.ResearchCenterLevel
	return r
}

func (r ResearchCenter) NextLevel() Building {
	r.Level++
	return r
}

func (r ResearchCenter) Cost() Amounts {
	return Amounts{
		Metal:    genericConstructionCost(r.Level, 115, 1.5),
		Hydrogen: genericConstructionCost(r.Level, 450, 1.5),
		Silica:   genericConstructionCost(r.Level, 700, 1.5),
	}
}

func (r ResearchCenter) Production() Amounts {
	return Amounts{}
}

func (r ResearchCenter) Usage() Levels {
	return Levels{
		Energy:     genericEnergyUsage(r.Level, 40, 1.25),
		Population: genericPopulationEmployment(r.Level, 40, 1.7),
	}
}

func (r ResearchCenter) Supply() Levels {
	return Levels{}
}

func (r ResearchCenter) Capacity() Amounts {
	return Amounts{}
}

func (r ResearchCenter) BuildDuration() time.Duration {
	return researchCenterUpgradeDuration(r.Level)
}

type ShipFactory struct {
	Level int
}

func (s ShipFactory) UpgradeAction() timer.Action {
	return timer.ActionUpgradeShipFactory
}

func (s ShipFactory) LevelOnPlanet(p *ent.Planet) int {
	return p.ShipFactoryLevel
}

func (s ShipFactory) UpgradeBuildingOnPlanet(p *ent.Planet) {
	p.ShipFactoryLevel++
}

func (s ShipFactory) ForPlanet(p *ent.Planet) Building {
	s.Level = p.ShipFactoryLevel
	return s
}

func (s ShipFactory) NextLevel() Building {
	s.Level++
	return s
}

func (s ShipFactory) Cost() Amounts {
	return Amounts{
		Metal:    genericConstructionCost(s.Level, 600, 1.5),
		Hydrogen: genericConstructionCost(s.Level, 100, 1.5),
		Silica:   genericConstructionCost(s.Level, 350, 1.5),
	}
}

func (s ShipFactory) Production() Amounts {
	return Amounts{}
}

func (s ShipFactory) Usage() Levels {
	return Levels{
		Energy:     genericEnergyUsage(s.Level, 15, 1.25),
		Population: genericPopulationEmployment(s.Level, 30, 1.7),
	}
}

func (s ShipFactory) Supply() Levels {
	return Levels{}
}

func (s ShipFactory) Capacity() Amounts {
	return Amounts{}
}

func (s ShipFactory) BuildDuration() time.Duration {
	return shipFactoryUpgradeDuration(s.Level)
}
