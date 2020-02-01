package data

import (
	"math"
	"time"

	"github.com/pdeguing/empire-and-foundation/ent"
)

// Amounts holds the amount of each resource that is necessary or stored by something.
type Amounts struct {
	Metal    int64
	Hydrogen int64
	Silica   int64
	Population	int64
}

func hasResources(p *ent.Planet, s Amounts) bool {
	return p.Metal >= s.Metal &&
		p.Hydrogen >= s.Hydrogen &&
		p.Silica >= s.Silica
}

func addStock(p *ent.Planet, s Amounts) {
	p.Metal += s.Metal
	p.Hydrogen += s.Hydrogen
	p.Silica += s.Silica
	p.Population += s.Population
}

func subStock(p *ent.Planet, s Amounts) {
	p.Metal -= s.Metal
	p.Hydrogen -= s.Hydrogen
	p.Silica -= s.Silica
	p.Population -= s.Population
}

func GetMetalProdUpgradeCost(newLevel int) Amounts {
	if newLevel == 1 {
		return Amounts{
			Metal:	0,
			Hydrogen: 0,
			Silica: 0,
		}
	}
	return Amounts{
		Metal:    60,
		Hydrogen: 1,
		Silica:   12,
	}
}

func getMetalProdUpgradeDuration(newLevel int) time.Duration {
	return time.Second * time.Duration(42*math.Pow(1.5, float64(newLevel)))
}

func GetHydrogenProdUpgradeCost(newLevel int) Amounts {
	if newLevel == 1 {
		return Amounts{
			Metal:	0,
			Hydrogen: 0,
			Silica: 0,
		}
	}
	return Amounts{
		Metal:    11,
		Hydrogen: 0,
		Silica:   30,
	}
}

func getHydrogenProdUpgradeDuration(newLevel int) time.Duration {
	return time.Second * time.Duration(42*math.Pow(1.5, float64(newLevel)))
}

func GetSilicaProdUpgradeCost(newLevel int) Amounts {
	if newLevel == 1 {
		return Amounts{
			Metal:	0,
			Hydrogen: 0,
			Silica: 0,
		}
	}
	return Amounts{
		Metal:    60,
		Hydrogen: 0,
		Silica:   5,
	}
}

func getSilicaProdUpgradeDuration(newLevel int) time.Duration {
	return time.Second * time.Duration(42*math.Pow(1.5, float64(newLevel)))
}

func GetSolarProdUpgradeCost(newLevel int) Amounts {
	if newLevel == 1 {
		return Amounts{
			Metal:	0,
			Hydrogen: 0,
			Silica: 0,
		}
	}
	return Amounts{
		Metal:    45,
		Hydrogen: 0,
		Silica:   40,
	}
}

func getSolarProdUpgradeDuration(newLevel int) time.Duration {
	return time.Second * time.Duration(42*math.Pow(1.5, float64(newLevel)))
}

func GetUrbanismUpgradeCost(newLevel int) Amounts {
	if newLevel == 1 {
		return Amounts{
			Metal:	0,
			Hydrogen: 0,
			Silica: 0,
		}
	}
	return Amounts{
		Metal:    1,
		Hydrogen: 0,
		Silica:   20,
	}
}

func getUrbanismUpgradeDuration(newLevel int) time.Duration {
	return time.Second * time.Duration(42*math.Pow(1.5, float64(newLevel)))
}

func GetMetalStorageUpgradeCost(newLevel int) Amounts {
	if newLevel == 1 {
		return Amounts{
			Metal:	0,
			Hydrogen: 0,
			Silica: 0,
		}
	}
	return Amounts{
		Metal:    10,
		Hydrogen: 0,
		Silica:   5,
	}
}

func getMetalStorageUpgradeDuration(newLevel int) time.Duration {
	return time.Second * time.Duration(42*math.Pow(1.5, float64(newLevel)))
}

func GetHydrogenStorageUpgradeCost(newLevel int) Amounts {
	if newLevel == 1 {
		return Amounts{
			Metal:	0,
			Hydrogen: 0,
			Silica: 0,
		}
	}
	return Amounts{
		Metal:    10,
		Hydrogen: 2,
		Silica:   3,
	}
}

func getHydrogenStorageUpgradeDuration(newLevel int) time.Duration {
	return time.Second * time.Duration(42*math.Pow(1.5, float64(newLevel)))
}

func GetSilicaStorageUpgradeCost(newLevel int) Amounts {
	if newLevel == 1 {
		return Amounts{
			Metal:	0,
			Hydrogen: 0,
			Silica: 0,
		}
	}
	return Amounts{
		Metal:    8,
		Hydrogen: 0,
		Silica:   10,
	}
}

func getSilicaStorageUpgradeDuration(newLevel int) time.Duration {
	return time.Second * time.Duration(42*math.Pow(1.5, float64(newLevel)))
}
