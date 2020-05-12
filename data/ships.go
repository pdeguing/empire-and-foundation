package data

import (
	"github.com/pdeguing/empire-and-foundation/ent"
	"github.com/pdeguing/empire-and-foundation/ent/timer"
	"time"
)

var Ships = []Ship{
	Caravel{},
	LightFighter{},
	Corvette{},
	Frigate{},
	Probe{},
	SmallCargo{},
	MediumCargo{},
	ColonizationArk{},
}

type Ship interface {
	// BuildAction references the timer action that can be used to build or
	// cancel this ship.
	BuildAction() timer.Action

	// NumberOnPlanet returns the number of this type of ship on the specified
	// planet.
	NumberOnPlanet(p *ent.Planet) int64

	// IncreaseNumberOnPlanet increases the number of this type of ship on the
	// specified planet by 1.
	IncreaseNumberOnPlanet(p *ent.Planet)

	// Cost returns the cost to build this ship.
	Cost() Amounts

	// BuildDuration returns the time it takes to build a single ship of this
	// type.
	BuildDuration() time.Duration
}

type Caravel struct{}

func (c Caravel) BuildAction() timer.Action {
	return timer.ActionBuildCaravel
}

func (c Caravel) NumberOnPlanet(p *ent.Planet) int64 {
	return p.NumCaravel
}

func (c Caravel) IncreaseNumberOnPlanet(p *ent.Planet) {
	p.NumCaravel++
}

func (c Caravel) Cost() Amounts {
	return Amounts{
		Metal:    20,
		Hydrogen: 1,
		Silica:   10,
	}
}

func (c Caravel) BuildDuration() time.Duration {
	return time.Minute
}

type LightFighter struct{}

func (l LightFighter) BuildAction() timer.Action {
	return timer.ActionBuildLightFighter
}

func (l LightFighter) NumberOnPlanet(p *ent.Planet) int64 {
	return p.NumLightFighter
}

func (l LightFighter) IncreaseNumberOnPlanet(p *ent.Planet) {
	p.NumLightFighter++
}

func (l LightFighter) Cost() Amounts {
	return Amounts{
		Metal:    20,
		Hydrogen: 1,
		Silica:   10,
	}
}

func (l LightFighter) BuildDuration() time.Duration {
	return time.Minute
}

type Corvette struct{}

func (c Corvette) BuildAction() timer.Action {
	return timer.ActionBuildCorvette
}

func (c Corvette) NumberOnPlanet(p *ent.Planet) int64 {
	return p.NumCorvette
}

func (c Corvette) IncreaseNumberOnPlanet(p *ent.Planet) {
	p.NumCorvette++
}

func (c Corvette) Cost() Amounts {
	return Amounts{
		Metal:    20,
		Hydrogen: 1,
		Silica:   10,
	}
}

func (c Corvette) BuildDuration() time.Duration {
	return time.Minute
}

type Frigate struct{}

func (f Frigate) BuildAction() timer.Action {
	return timer.ActionBuildFrigate
}

func (f Frigate) NumberOnPlanet(p *ent.Planet) int64 {
	return p.NumFrigate
}

func (f Frigate) IncreaseNumberOnPlanet(p *ent.Planet) {
	p.NumFrigate++
}

func (f Frigate) Cost() Amounts {
	return Amounts{
		Metal:    20,
		Hydrogen: 1,
		Silica:   10,
	}
}

func (f Frigate) BuildDuration() time.Duration {
	return time.Minute
}

type Probe struct{}

func (pr Probe) BuildAction() timer.Action {
	return timer.ActionBuildProbe
}

func (pr Probe) NumberOnPlanet(p *ent.Planet) int64 {
	return p.NumProbe
}

func (pr Probe) IncreaseNumberOnPlanet(p *ent.Planet) {
	p.NumProbe++
}

func (pr Probe) Cost() Amounts {
	return Amounts{
		Metal:    20,
		Hydrogen: 1,
		Silica:   10,
	}
}

func (pr Probe) BuildDuration() time.Duration {
	return time.Minute
}

type SmallCargo struct{}

func (s SmallCargo) BuildAction() timer.Action {
	return timer.ActionBuildSmallCargo
}

func (s SmallCargo) NumberOnPlanet(p *ent.Planet) int64 {
	return p.NumSmallCargo
}

func (s SmallCargo) IncreaseNumberOnPlanet(p *ent.Planet) {
	p.NumSmallCargo++
}

func (s SmallCargo) Cost() Amounts {
	return Amounts{
		Metal:    20,
		Hydrogen: 1,
		Silica:   10,
	}
}

func (s SmallCargo) BuildDuration() time.Duration {
	return time.Minute
}

type MediumCargo struct{}

func (m MediumCargo) BuildAction() timer.Action {
	return timer.ActionBuildMediumCargo
}

func (m MediumCargo) NumberOnPlanet(p *ent.Planet) int64 {
	return p.NumMediumCargo
}

func (m MediumCargo) IncreaseNumberOnPlanet(p *ent.Planet) {
	p.NumMediumCargo++
}

func (m MediumCargo) Cost() Amounts {
	return Amounts{
		Metal:    20,
		Hydrogen: 1,
		Silica:   10,
	}
}

func (m MediumCargo) BuildDuration() time.Duration {
	return time.Minute
}

type ColonizationArk struct{}

func (c ColonizationArk) BuildAction() timer.Action {
	return timer.ActionBuildColonizationArk
}

func (c ColonizationArk) NumberOnPlanet(p *ent.Planet) int64 {
	return p.NumColonizationArk
}

func (c ColonizationArk) IncreaseNumberOnPlanet(p *ent.Planet) {
	p.NumColonizationArk++
}

func (c ColonizationArk) Cost() Amounts {
	return Amounts{
		Metal:    20,
		Hydrogen: 1,
		Silica:   10,
	}
}

func (c ColonizationArk) BuildDuration() time.Duration {
	return time.Minute
}
