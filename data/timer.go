package data

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/pdeguing/empire-and-foundation/ent"
	"github.com/pdeguing/empire-and-foundation/ent/planet"
	"github.com/pdeguing/empire-and-foundation/ent/timer"
)

// ErrActionPrerequisitesNotMet is returned when a action's Valid()
// method returns false. The timer for the action cannot be started.
var ErrActionPrerequisitesNotMet = errors.New("Cannot start the timer because the action's prerequisites (`! t.Valid()`) were not met")

// ErrTimerBussy is returned when another timer for the same planet
// and in the same group is already running. Only one timer can be
// running at a time.
var ErrTimerBussy = errors.New("Another timer is already running for this planet and group")

// ErrTimerNotRunning is returned when trying to cancel a timer that
// is not currently running.
var ErrTimerNotRunning = errors.New("The timer cannot be cancelled because it is not running")

type action struct {
	// Group specifies the group the action belongs to. For each planet
	// there can only be one running action in each group at any given time.
	Group timer.Group

	// Duration returns the time to wait between the start of the action and the completion.
	Duration func(p *ent.Planet) time.Duration

	// Valid checks if the prerequisites of the action are satisfied.
	Valid func(p *ent.Planet) bool

	// Start is triggered when the action is scheduled using a timer.
	Start func(ctx context.Context, tx *ent.Tx, p *ent.Planet) error

	// Complete is triggered when the timer is done. Make sure that any update
	// to the planet is also updated in the passed in model and not only in
	// the database. If not, the update won't be visible in the view until
	// after a reload.
	Complete func(ctx context.Context, tx *ent.Tx, p *ent.Planet) error

	// Cancel is triggered when the timer is canceled before
	// it finished.
	Cancel func(ctx context.Context, tx *ent.Tx, p *ent.Planet) error
}

// Timer contains information about a single running timer.
type Timer struct {
	Action  timer.Action
	EndTime time.Time
}

// Duration returns the time left until the timer completes.
func (t *Timer) Duration() time.Duration {
	return time.Until(t.EndTime)
}

// actions contains a map of acctions that can be executed using a timer.
// All types should be defined using the enum fields in the Timer ent schema
// and, vice versa, all enum values should exist in this map.
var actions = map[timer.Action]action{
	timer.ActionUpgradeMetalMine: action{
		Group: timer.GroupBuilding,
		Duration: func(p *ent.Planet) time.Duration {
			return getMetalMineUpgradeDuration(p.MetalProdLevel + 1)
		},
		Valid: func(p *ent.Planet) bool {
			c := getMetalMineUpgradeCost(p.MetalProdLevel + 1)
			return hasResources(p, c)
		},
		Start: func(ctx context.Context, tx *ent.Tx, p *ent.Planet) error {
			c := getMetalMineUpgradeCost(p.MetalProdLevel + 1)
			return subStock(ctx, p, c)
		},
		Complete: func(ctx context.Context, tx *ent.Tx, p *ent.Planet) error {
			p.MetalProdLevel++
			p.MetalRate = getNewMetalRate(p)
			_, err := p.Update().
				SetMetalProdLevel(p.MetalProdLevel).
				SetMetalRate(p.MetalRate).
				SetMetalLastUpdate(p.MetalLastUpdate).
				Save(ctx)
			return err
		},
		Cancel: func(ctx context.Context, tx *ent.Tx, p *ent.Planet) error {
			c := getMetalMineUpgradeCost(p.MetalProdLevel + 1)
			return addStock(ctx, p, c)
		},
	},
	timer.ActionUpgradeHydrogenExtractor: action{
		Group: timer.GroupBuilding,
		Duration: func(p *ent.Planet) time.Duration {
			return getHydrogenExtractorUpgradeDuration(p.HydrogenProdLevel + 1)
		},
		Valid: func(p *ent.Planet) bool {
			c := getHydrogenExtractorUpgradeCost(p.HydrogenProdLevel + 1)
			return hasResources(p, c)
		},
		Start: func(ctx context.Context, tx *ent.Tx, p *ent.Planet) error {
			c := getHydrogenExtractorUpgradeCost(p.HydrogenProdLevel + 1)
			return subStock(ctx, p, c)
		},
		Complete: func(ctx context.Context, tx *ent.Tx, p *ent.Planet) error {
			p.HydrogenProdLevel++
			p.HydrogenRate = getNewHydrogenRate(p)
			_, err := p.Update().
				SetHydrogenProdLevel(p.HydrogenProdLevel).
				SetHydrogenRate(p.HydrogenRate).
				SetHydrogenLastUpdate(p.HydrogenLastUpdate).
				Save(ctx)
			return err
		},
		Cancel: func(ctx context.Context, tx *ent.Tx, p *ent.Planet) error {
			c := getHydrogenExtractorUpgradeCost(p.HydrogenProdLevel + 1)
			return addStock(ctx, p, c)
		},
	},
	timer.ActionUpgradeSilicaQuarry: action{
		Group: timer.GroupBuilding,
		Duration: func(p *ent.Planet) time.Duration {
			return getSilicaQuarryUpgradeDuration(p.SilicaProdLevel + 1)
		},
		Valid: func(p *ent.Planet) bool {
			c := getSilicaQuarryUpgradeCost(p.SilicaProdLevel + 1)
			return hasResources(p, c)
		},
		Start: func(ctx context.Context, tx *ent.Tx, p *ent.Planet) error {
			c := getSilicaQuarryUpgradeCost(p.SilicaProdLevel + 1)
			return subStock(ctx, p, c)
		},
		Complete: func(ctx context.Context, tx *ent.Tx, p *ent.Planet) error {
			p.SilicaProdLevel++
			p.SilicaRate = getNewSilicaRate(p)
			_, err := p.Update().
				SetSilicaProdLevel(p.SilicaProdLevel).
				SetSilicaRate(p.SilicaRate).
				SetSilicaLastUpdate(p.SilicaLastUpdate).
				Save(ctx)
			return err
		},
		Cancel: func(ctx context.Context, tx *ent.Tx, p *ent.Planet) error {
			c := getSilicaQuarryUpgradeCost(p.SilicaProdLevel + 1)
			return addStock(ctx, p, c)
		},
	},
	timer.ActionUpgradeSolarPlant: action{
		Group: timer.GroupBuilding,
		Duration: func(p *ent.Planet) time.Duration {
			return getSolarPlantUpgradeDuration(p.SolarProdLevel + 1)
		},
		Valid: func(p *ent.Planet) bool {
			c := getSolarPlantUpgradeCost(p.SolarProdLevel + 1)
			return hasResources(p, c)
		},
		Start: func(ctx context.Context, tx *ent.Tx, p *ent.Planet) error {
			c := getSolarPlantUpgradeCost(p.SolarProdLevel + 1)
			return subStock(ctx, p, c)
		},
		Complete: func(ctx context.Context, tx *ent.Tx, p *ent.Planet) error {
			p.SolarProdLevel++
			_, err := p.Update().
				SetSolarProdLevel(p.SolarProdLevel).
				Save(ctx)
			return err
		},
		Cancel: func(ctx context.Context, tx *ent.Tx, p *ent.Planet) error {
			c := getSolarPlantUpgradeCost(p.SolarProdLevel + 1)
			return addStock(ctx, p, c)
		},
	},
	timer.ActionUpgradeHousingFacilities: action{
		Group: timer.GroupBuilding,
		Duration: func(p *ent.Planet) time.Duration {
			return getHousingFacilitiesUpgradeDuration(p.PopulationStorageLevel + 1)
		},
		Valid: func(p *ent.Planet) bool {
			c := getHousingFacilitiesUpgradeCost(p.PopulationStorageLevel + 1)
			return hasResources(p, c)
		},
		Start: func(ctx context.Context, tx *ent.Tx, p *ent.Planet) error {
			c := getHousingFacilitiesUpgradeCost(p.PopulationStorageLevel + 1)
			return subStock(ctx, p, c)
		},
		Complete: func(ctx context.Context, tx *ent.Tx, p *ent.Planet) error {
			p.PopulationStorageLevel++
			_, err := p.Update().
				SetPopulationStorageLevel(p.PopulationStorageLevel).
				Save(ctx)
			return err
		},
		Cancel: func(ctx context.Context, tx *ent.Tx, p *ent.Planet) error {
			c := getHousingFacilitiesUpgradeCost(p.PopulationStorageLevel + 1)
			return addStock(ctx, p, c)
		},
	},
}

// IsBussy checks if there is currently a timer in progress for the group.
func IsBussy(ctx context.Context, p *ent.Planet, g timer.Group) (bool, error) {
	b, err := p.QueryTimers().
		Where(timer.GroupEQ(g)).
		Exist(ctx)
	if err != nil {
		return true, fmt.Errorf("unable to query existence of running timer: %v", err)
	}
	return b, nil
}

// GetTimer returns information about the in progress timer in the group.
func GetTimer(ctx context.Context, p *ent.Planet, g timer.Group) (*Timer, error) {
	t, err := p.QueryTimers().
		Where(timer.GroupEQ(g)).
		Only(ctx)
	if _, ok := err.(*ent.ErrNotFound); ok {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve timer in group %q for planet: %v", g, err)
	}
	return &Timer{
		Action:  t.Action,
		EndTime: t.EndTime,
	}, nil
}

// GetTimers returns a map with information about all active timers for the planet.
func GetTimers(ctx context.Context, p *ent.Planet) (map[timer.Group]*Timer, error) {
	timers, err := p.QueryTimers().
		All(ctx)
	if _, ok := err.(*ent.ErrNotFound); ok {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve timers for planet: %v", err)
	}
	tm := make(map[timer.Group]*Timer)
	for _, t := range timers {
		tm[t.Group] = &Timer{
			Action:  t.Action,
			EndTime: t.EndTime,
		}
	}
	return tm, nil
}

// StartTimer starts a timer for the action a if all prerequisites are met.
// After the duration defined by the action, the timer completes.
// StartTimer expects the planet state to be up-to-date.
func StartTimer(ctx context.Context, tx *ent.Tx, p *ent.Planet, action timer.Action) error {
	a, ok := actions[action]
	if !ok {
		// A test exists to check that all actions defined in the schema
		// are also defined in the action map. This error should never
		// occur in production if the test is used.
		return fmt.Errorf("action %q for timer is not yet defined", action)
	}
	bussy, err := IsBussy(ctx, p, a.Group)
	if err != nil {
		return err
	}
	if bussy {
		return ErrTimerBussy
	}
	if !a.Valid(p) {
		return ErrActionPrerequisitesNotMet
	}
	d := a.Duration(p)
	err = a.Start(ctx, tx, p)
	if err != nil {
		return fmt.Errorf("error while calling \"Start\" function for action %q: %v", action, err)
	}
	_, err = tx.Timer.
		Create().
		SetPlanet(p).
		SetAction(action).
		SetGroup(a.Group).
		SetEndTime(time.Now().Add(d)).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("unable to create timer entry: %v", err)
	}
	return nil
}

// CancelTimer aborts the timer and triggers the action's Cancel function immediately.
// CancelTimer expects the planet state to be up-to-date.
func CancelTimer(ctx context.Context, tx *ent.Tx, p *ent.Planet, a timer.Action) error {
	n, err := tx.Timer.
		Delete().
		Where(timer.HasPlanetWith(planet.IDEQ(p.ID))).
		Where(timer.ActionEQ(a)).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("unable to delete timer to cancel it: %v", err)
	}
	if n == 0 {
		return ErrTimerNotRunning
	}
	err = actions[a].Cancel(ctx, tx, p)
	return err
}

// UpdateTimers checks if timers have completed, and if so, triggers the action's
// Complete function and cleans up the timers. This function must be called before
// any information manipulated by the timers/actions is queried.
// In contrast to StartTimer and CancelTimer, UpdateTimers expects the planet
// state *NOT* to be updated. UpdateTimers makes use of the old state of the
// planet to calculate durations and update the state in steps.
func UpdateTimers(ctx context.Context, tx *ent.Tx, p *ent.Planet) error {
	timers, err := p.QueryTimers().
		Where(timer.EndTimeLTE(time.Now())).
		Order(ent.Asc(timer.FieldEndTime)).
		All(ctx)
	if err != nil {
		return fmt.Errorf("unable to retrieve running timers: %v", err)
	}
	if len(timers) == 0 {
		return nil // Fast path
	}
	for _, t := range timers {
		UpdatePlanetState(p, t.EndTime)
		err = actions[t.Action].Complete(ctx, tx, p)
		if err != nil {
			return fmt.Errorf("error while calling \"Complete\" function for action %q: %v", t.Action, err)
		}
	}
	_, err = tx.Timer.
		Delete().
		Where(timer.EndTimeLTE(time.Now())).
		Where(timer.HasPlanetWith(planet.IDEQ(p.ID))).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("unable to delete finished timers: %v", err)
	}
	return nil
}
