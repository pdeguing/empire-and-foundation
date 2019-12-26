package main

import (
	"errors"
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
func userPlanet(r *http.Request, tx *ent.Tx) (*ent.Planet, error) {
	n, err := strconv.Atoi(mux.Vars(r)["planetNumber"])
	if err != nil {
		return nil, newInternalServerError(fmt.Errorf("unable to parse url parameter \"planetNumber\": %v", err))
	}
	u := loggedInUser(r)

	p, err := tx.Planet.
		Query().
		Where(planet.HasOwnerWith(user.IDEQ(u.ID))).
		Offset(n - 1).
		First(r.Context())
	if _, ok := err.(*ent.ErrNotFound); ok {
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
	data.UpdatePlanetState(p, time.Now())
	return p, nil
}

// newPlanetViewData collects the data for the planet's construction, research and other
// views with upgrade mechanisms.
func newPlanetViewData(r *http.Request, g timer.Group) (*planetViewData, error) {
	var p *ent.Planet
	var t *data.Timer
	err := data.WithTx(r.Context(), data.Client, func(tx *ent.Tx) error {
		var err error
		p, err = userPlanet(r, tx)
		if err != nil {
			return err
		}
		// TODO: Remove wrapping if statement once all planet views have a group
		// they want to get the timers for.
		if g != "" {
			t, err = data.GetTimer(r.Context(), p, g)
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
		Planet: p,
		Timer:  t,
	}, nil
}

// planetViewData contains the information for a page with upgradable
// items like constructions or research.
type planetViewData struct {
	Planet *ent.Planet
	Timer  *data.Timer
}

// serveUpgradeBuilding progresses the request to start an upgrade timer.
func serveUpgradeBuilding(w http.ResponseWriter, r *http.Request, a timer.Action) {
	var p *ent.Planet
	err := data.WithTx(r.Context(), data.Client, func(tx *ent.Tx) error {
		var err error
		p, err = userPlanet(r, tx)
		if err != nil {
			return err
		}
		err = data.StartTimer(r.Context(), tx, p, a)
		if err != nil {
			return newInternalServerError(fmt.Errorf("unable to start timer to upgrade building: %w", err))
		}
		return nil
	})
	if err != nil {
		if errors.Is(err, data.ErrActionPrerequisitesNotMet) {
			flash(r, flashDanger, "There are not enough resources on this planet.")
			http.Redirect(w, r, "/planet/"+strconv.Itoa(p.ID)+"/constructions", 302)
			return
		}
		if errors.Is(err, data.ErrTimerBussy) {
			flash(r, flashWarning, "Something is already being upgraded.")
			http.Redirect(w, r, "/planet/"+strconv.Itoa(p.ID)+"/constructions", 302)
			return
		}
		serveError(w, r, err)
		return
	}
	http.Redirect(w, r, "/planet/"+strconv.Itoa(p.ID)+"/constructions", 302)
}

// serveCancelBuilding progresses the request to cancel a timer.
func serveCancelBuilding(w http.ResponseWriter, r *http.Request, a timer.Action) {
	var p *ent.Planet
	err := data.WithTx(r.Context(), data.Client, func(tx *ent.Tx) error {
		var err error
		p, err = userPlanet(r, tx)
		if err != nil {
			return err
		}
		err = data.CancelTimer(r.Context(), tx, p, a)
		if err != nil {
			return newInternalServerError(fmt.Errorf("unable to cancel timer to upgrade building: %w", err))
		}
		return nil
	})
	if err != nil && !errors.Is(err, data.ErrTimerNotRunning) {
		serveError(w, r, err)
		return
	}
	http.Redirect(w, r, "/planet/"+strconv.Itoa(p.ID)+"/constructions", 302)
}
