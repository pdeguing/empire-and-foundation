package main

import (
	"net/http"
	"strconv"
	"time"
	"math"

	"github.com/gorilla/mux"
	"github.com/pdeguing/empire-and-foundation/ent"
)

// Calculates the storage capacity for a resource with given storage level
func getMaxStorage(storageLevel int) int64 {
	maxStorage := 100000 *  int64(storageLevel) * int64(math.Pow(1.1, float64(storageLevel)))
	return maxStorage
}

// Calculates the current value in stock for a resource based on value and duration since last update
func getStock(val int64, last time.Time, rate int, storageLevel int) int64 {
	duration := int64(time.Since(last) / time.Second)
	maxStorage := getMaxStorage(storageLevel)
	current := val + duration * int64(rate)
	if current >= maxStorage {
		return maxStorage
	}
	return current
}

// Updates the current planet struct to get up-to-date state
func getPlanetState(p *ent.Planet) *ent.Planet {
	p.Metal = getStock(
		p.Metal,
		p.MetalLastUpdate,
		p.MetalRate,
		p.MetalStorageLevel,
	)
	p.Hydrogen = getStock(
		p.Hydrogen,
		p.HydrogenLastUpdate,
		p.HydrogenRate,
		p.HydrogenStorageLevel,
	)
	p.Silica = getStock(
		p.Silica,
		p.SilicaLastUpdate,
		p.SilicaRate,
		p.SilicaStorageLevel,
	)
	p.Population = getStock(
		p.Population,
		p.PopulationLastUpdate,
		p.PopulationRate,
		p.PopulationStorageLevel,
	)
	return p
}

// Retrieves the planet requested for the logged in user
func userPlanet(w http.ResponseWriter, r *http.Request) (*ent.Planet, bool) {
	n, err := strconv.Atoi(mux.Vars(r)["planetNumber"])
	if err != nil {
		serveInternalServerError(w, r, err, "Cannot parse url variable 'planetNumber'")
		return nil, false
	}
	u := loggedInUser(r)
	p, err := u.QueryPlanets().
		Offset(n - 1).
		First(r.Context())
	if _, ok := err.(*ent.ErrNotFound); ok {
		serveNotFoundError(w, r)
		return nil, false
	}
	if err != nil {
		serveInternalServerError(w, r, err, "Could not retrieve user's planet from database")
		return nil, false
	}
	p = getPlanetState(p)
	return p, true
}
