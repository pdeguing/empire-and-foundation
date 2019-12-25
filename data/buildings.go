package data

import (
	"context"
	"math"
	"time"

	"github.com/pdeguing/empire-and-foundation/ent"
)

// Amounts holds the amount of each resource that is necessary or stored by something.
type Amounts struct {
	Metal    int64
	Hydrogen int64
	Silica   int64
}

func hasResources(p *ent.Planet, s Amounts) bool {
	return p.Metal >= s.Metal &&
		p.Hydrogen >= s.Hydrogen &&
		p.Silica >= s.Silica
}

func addStock(ctx context.Context, p *ent.Planet, s Amounts) error {
	p.Metal += s.Metal
	p.Hydrogen += s.Hydrogen
	p.Silica += s.Silica
	_, err := p.Update().
		SetMetal(p.Metal).
		SetHydrogen(p.Hydrogen).
		SetSilica(p.Silica).
		Save(ctx)
	return err
}

func subStock(ctx context.Context, p *ent.Planet, s Amounts) error {
	p.Metal -= s.Metal
	p.Hydrogen -= s.Hydrogen
	p.Silica -= s.Silica
	_, err := p.Update().
		SetMetal(p.Metal).
		SetHydrogen(p.Hydrogen).
		SetSilica(p.Silica).
		Save(ctx)
	return err
}

func getMetalMineUpgradeCost(newLevel int) Amounts {
	return Amounts{
		Metal:    60,
		Hydrogen: 1,
		Silica:   12,
	}
}

func getMetalMineUpgradeDuration(newLevel int) time.Duration {
	return time.Second * time.Duration(42*math.Pow(1.5, float64(newLevel)))
}

func getHydrogenExtractorUpgradeCost(newLevel int) Amounts {
	return Amounts{
		Metal:    11,
		Hydrogen: 0,
		Silica:   30,
	}
}

func getHydrogenExtractorUpgradeDuration(newLevel int) time.Duration {
	return time.Second * time.Duration(42*math.Pow(1.5, float64(newLevel)))
}

func getSilicaQuarryUpgradeCost(newLevel int) Amounts {
	return Amounts{
		Metal:    60,
		Hydrogen: 0,
		Silica:   5,
	}
}

func getSilicaQuarryUpgradeDuration(newLevel int) time.Duration {
	return time.Second * time.Duration(42*math.Pow(1.5, float64(newLevel)))
}

func getSolarPlantUpgradeCost(newLevel int) Amounts {
	return Amounts{
		Metal:    45,
		Hydrogen: 0,
		Silica:   40,
	}
}

func getSolarPlantUpgradeDuration(newLevel int) time.Duration {
	return time.Second * time.Duration(42*math.Pow(1.5, float64(newLevel)))
}

func getHousingFacilitiesUpgradeCost(newLevel int) Amounts {
	return Amounts{
		Metal:    1,
		Hydrogen: 0,
		Silica:   20,
	}
}

func getHousingFacilitiesUpgradeDuration(newLevel int) time.Duration {
	return time.Second * time.Duration(42*math.Pow(1.5, float64(newLevel)))
}
