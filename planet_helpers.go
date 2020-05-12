package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/pdeguing/empire-and-foundation/data"
	"github.com/pdeguing/empire-and-foundation/ent"
	"github.com/pdeguing/empire-and-foundation/ent/planet"
	"github.com/pdeguing/empire-and-foundation/ent/timer"
	"github.com/pdeguing/empire-and-foundation/ent/user"
)

// userPlanet retrieves the planet requested for the logged in user
func userPlanet(r *http.Request, tx *ent.Tx) (*data.PlanetWithResourceInfo, error) {
	n, err := strconv.Atoi(mux.Vars(r)["planetNumber"])
	if err != nil {
		return nil, newInternalServerError(fmt.Errorf("unable to parse url parameter \"planetNumber\": %v", err))
	}
	u := loggedInUser(r)

	p, err := tx.Planet.
		Query().
		Where(
			planet.And(
				planet.HasOwnerWith(user.IDEQ(u.ID)),
				planet.ID(n),
			),
		).
		First(r.Context())
	if _, ok := err.(*ent.NotFoundError); ok {
		return nil, newNotFoundError(fmt.Errorf("unable to query planet #%d for user %d; it does not exist", n, u.ID))
	}
	if err != nil {
		return nil, newInternalServerError(fmt.Errorf("unable to retrieve planet for user: %v", err))
	}
	// UpdateTimers uses the old state of the planet to calculate the timer
	// durations. Therefore, this function must be called before the state is
	// updated.
	err = data.UpdateTimers(r.Context(), tx, p)
	if err != nil {
		return nil, newInternalServerError(fmt.Errorf("unable to update planet timers: %v", err))
	}
	pwr := data.NewPlanetWithResourceInfo(p)
	pwr.Update(time.Now())
	return pwr, nil
}

// userPlanets retrieves all planets for the logged in user
func userPlanets(r *http.Request, tx *ent.Tx) ([]*ent.Planet, error) {
	u := loggedInUser(r)

	p, err := tx.Planet.Query().
		Where(planet.HasOwnerWith(user.IDEQ(u.ID))).
		All(r.Context())
	if _, ok := err.(*ent.NotFoundError); ok {
		return nil, newNotFoundError(fmt.Errorf("unable to query planets for user %d; it does not exist", u.ID))
	}
	if err != nil {
		return nil, newInternalServerError(fmt.Errorf("unable to retrieve planets for user: %v", err))
	}
	return p, nil
}

// regionPlanets create slice of positions and planets in a system
func regionPlanets(w http.ResponseWriter, r *http.Request, region, system int) ([]*ent.Planet, error) {
	p, err := data.Client.Planet.
		Query().
		Where(planet.RegionCodeEQ(region)).
		Where(planet.SystemCodeEQ(system)).
		Order(ent.Asc(planet.FieldPositionCode)).
		WithOwner().
		All(r.Context())

	if err != nil {
		return nil, newInternalServerError(fmt.Errorf("could not retrieve user's planet from database: %v", err))
	}

	return p, err
}

// planetViewData contains the information for a page with upgradable
// items like constructions or research.
type planetViewData struct {
	UserPlanets []*ent.Planet
	Planet      *data.PlanetWithResourceInfo
	Timer       *data.Timer
}

// newPlanetViewData collects the data for the planet's construction, research and other
// views with upgrade mechanisms.
func newPlanetViewData(r *http.Request, g timer.Group) (*planetViewData, error) {
	var plist []*ent.Planet
	var p *data.PlanetWithResourceInfo
	var t *data.Timer
	err := data.WithTx(r.Context(), data.Client, func(tx *ent.Tx) error {
		var err error
		plist, err = userPlanets(r, tx)
		if err != nil {
			return err
		}
		p, err = userPlanet(r, tx)
		if err != nil {
			return err
		}
		// TODO: Remove wrapping if statement once all planet views have a group
		// they want to get the timers for.
		if g != "" {
			t, err = data.GetTimer(r.Context(), p.Planet, g)
			if err != nil {
				return newInternalServerError(err)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &planetViewData{
		UserPlanets: plist,
		Planet:      p,
		Timer:       t,
	}, nil
}
