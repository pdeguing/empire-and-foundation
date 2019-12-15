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
	// TODO: Population???
}

func hasResources(p *ent.Planet, s Amounts) bool {
	return p.Metal >= s.Metal &&
		p.Hydrogen >= s.Hydrogen &&
		p.Silica >= s.Silica
}

func addStock(ctx context.Context, p *ent.Planet, s Amounts) error {
	_, err := p.Update().
		SetMetal(p.Metal + s.Metal).
		SetHydrogen(p.Hydrogen + s.Hydrogen).
		SetSilica(p.Silica + s.Silica).
		Save(ctx)
	return err
}

func subStock(ctx context.Context, p *ent.Planet, s Amounts) error {
	_, err := p.Update().
		SetMetal(p.Metal - s.Metal).
		SetHydrogen(p.Hydrogen - s.Hydrogen).
		SetSilica(p.Silica - s.Silica).
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
