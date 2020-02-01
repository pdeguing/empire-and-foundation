package data

import (
	"context"
	"github.com/pdeguing/empire-and-foundation/ent"
	"github.com/pdeguing/empire-and-foundation/ent/timer"
	"math/rand"
	"strconv"
	"syreclabs.com/go/faker"
	"time"
)

type planetFactory struct {
	ent.Planet
	timers           []*ent.Timer
	updatedResources bool
	client           *ent.Client
	ctx              context.Context
}

var randomGen = rand.New(rand.NewSource(time.Now().Unix()))

// NewPlanetFactory creates a factory initialized with random data.
func NewPlanetFactory() *planetFactory {
	f := planetFactory{}
	f.Name = faker.Lorem().Word() // TODO: replace with names generated by the planet name generator used in the seeder.
	f.PlanetType = randomPlanetType(randomGen)
	f.PlanetSkin = faker.RandomChoice([]string{"ako", "avalon", "earth", "farma", "jinx"})
	f.RegionCode = 999
	f.SystemCode, _ = strconv.Atoi(faker.Number().Between(0, 255))
	f.OrbitCode, _ = strconv.Atoi(faker.Number().Between(1, 15))
	f.SuborbitCode = 0
	f.PositionCode = getPositionCode(f.RegionCode, f.SystemCode, f.OrbitCode, f.SuborbitCode)
	f.LastResourceUpdate = time.Now().Add(-time.Hour)
	f.CreatedAt = faker.Time().Backward(time.Hour * 24 * 365)
	f.UpdatedAt = faker.Time().Backward(time.Since(f.CreatedAt))
	f.updatedResources = false
	f.client = Client
	f.ctx = context.Background()
	return &f
}

// WithOwner adds a owner to the planet.
func (f *planetFactory) WithOwner() *planetFactory {
	return f.ForOwner(NewUserFactory().Client(f.client).MustCreate())
}

// ForOwner sets u as the owner of the planet.
func (f *planetFactory) ForOwner(u *ent.User) *planetFactory {
	f.Edges.Owner = u
	return f
}

// WithBeginnerResources sets the planets resources as if the user has been playing
// for a short time.
func (f *planetFactory) WithBeginnerResources() *planetFactory {
	f.MetalProdLevel = 2
	f.MetalStorageLevel = 1
	f.Metal = getMaxStorage(f.MetalStorageLevel) / 2
	f.SilicaProdLevel = 2
	f.SilicaStorageLevel = 1
	f.Silica = getMaxStorage(f.SilicaStorageLevel) * 3 / 4
	f.HydrogenProdLevel = 1
	f.HydrogenStorageLevel = 1
	f.Hydrogen = getMaxStorage(f.HydrogenStorageLevel)
	f.PopulationProdLevel = 2
	f.PopulationStorageLevel = f.PopulationProdLevel
	f.Population = getMaxStorage(f.PopulationStorageLevel)
	f.SolarProdLevel = 5
	return f
}

func (f *planetFactory) WithTimer(a timer.Action, d time.Duration) *planetFactory {
	t, err := f.client.Timer.Create().
		SetAction(a).
		SetGroup(actions[a].Group).
		SetEndTime(timeNow().Add(d)).
		Save(f.ctx)
	if err != nil {
		panic(err)
	}
	f.timers = append(f.timers, t)
	return f
}

// WithUpdatedResources makes the factory return a planet with up-to-date
// resources (LastResourceUpdate == now)
func (f *planetFactory) WithUpdatedResources() *planetFactory {
	f.updatedResources = true
	return f
}

// Client uses c to create the entity.
func (f *planetFactory) Client(c *ent.Client) *planetFactory {
	f.client = c
	return f
}

// Tx uses c to create the entity.
func (f *planetFactory) Tx(tx *ent.Tx) *planetFactory {
	return f.Client(tx.Client())
}

// InContext executes the Create method in ctx.
func (f *planetFactory) InContext(ctx context.Context) *planetFactory {
	f.ctx = ctx
	return f
}

// Create returns the planet struct, which is saved to the database.
func (f *planetFactory) Create() (*ent.Planet, error) {
	q := f.client.Planet.Create().
		SetName(f.Name).
		SetPlanetType(f.PlanetType).
		SetPlanetSkin(f.PlanetSkin).
		SetRegionCode(f.RegionCode).
		SetSystemCode(f.SystemCode).
		SetOrbitCode(f.OrbitCode).
		SetSuborbitCode(f.SuborbitCode).
		SetPositionCode(f.PositionCode).
		SetMetal(f.Metal).
		SetMetalStorageLevel(f.MetalStorageLevel).
		SetSilica(f.Silica).
		SetSilicaStorageLevel(f.SilicaStorageLevel).
		SetHydrogen(f.Hydrogen).
		SetHydrogenStorageLevel(f.HydrogenStorageLevel).
		SetPopulation(f.Population).
		SetPopulationStorageLevel(f.PopulationStorageLevel).
		SetSolarProdLevel(f.SolarProdLevel).
		SetCreatedAt(f.CreatedAt).
		SetUpdatedAt(f.UpdatedAt).
		AddTimers(f.timers...)

	if f.Edges.Owner != nil {
		q = q.SetOwner(f.Edges.Owner)
	}

	p, err := q.Save(f.ctx)
	if err != nil {
		return nil, err
	}

	if f.updatedResources {
		UpdatePlanetResources(&f.Planet, timeNow())
	}

	return p, nil
}

// MustCreate returns the planet struct, which is saved to the database.
// If an error occurs a panic is raised.
func (f *planetFactory) MustCreate() *ent.Planet {
	p, err := f.Create()
	if err != nil {
		panic(err)
	}
	return p
}