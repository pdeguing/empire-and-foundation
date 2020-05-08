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

// timeNow aliases the time.Now() function but can be overwritten to allow
// testing of time sensitive methods.
var timeNow = time.Now

// ErrActionPrerequisitesNotMet is returned when a action's Valid()
// method returns false. The timer for the action cannot be started.
var ErrActionPrerequisitesNotMet = errors.New("cannot start the timer because the action's prerequisites (`! t.Valid()`) were not met")

// ErrTimerBusy is returned when another timer for the same planet
// and in the same group is already running. Only one timer can be
// running at a time.
var ErrTimerBusy = errors.New("another timer is already running for this planet and group")

// ErrTimerNotRunning is returned when trying to cancel a timer that
// is not currently running.
var ErrTimerNotRunning = errors.New("the timer cannot be cancelled because it is not running")

type action struct {
	// Group specifies the group the action belongs to. For each planet
	// there can only be one running action in each group at any given time.
	Group timer.Group

	// Duration returns the time to wait between the start of the action and the completion.
	Duration func(p *ent.Planet) time.Duration

	// Valid checks if the prerequisites of the action are satisfied.
	Valid func(p *ent.Planet) bool

	// Start is triggered when the action is scheduled using a timer.
	Start func(p *ent.Planet) error

	// Complete is triggered when the timer is done.
	Complete func(p *ent.Planet) error

	// Cancel is triggered when the timer is canceled before
	// it finished.
	Cancel func(p *ent.Planet) error
}

// Timer contains information about a single running timer.
type Timer struct {
	Action  timer.Action
	EndTime time.Time
}

// Duration returns the time left until the timer completes.
func (t *Timer) Duration() time.Duration {
	return t.EndTime.Sub(timeNow())
}

// actions contains a map of actions that can be executed using a timer.
// All types should be defined using the enum fields in the Timer ent schema
// and, vice versa, all enum values should exist in this map.
var actions = map[timer.Action]action{
	timer.ActionTest: {}, // Overwritten in some tests
}

func init() {
	for _, b := range Buildings {
		actions[b.UpgradeAction()] = newActionFromBuilding(b)
	}
}

func newActionFromBuilding(b Building) action {
	return action{
		Group: timer.GroupBuilding,
		Duration: func(p *ent.Planet) time.Duration {
			return b.ForPlanet(p).NextLevel().BuildDuration()
		},
		Valid: func(p *ent.Planet) bool {
			c := b.ForPlanet(p).NextLevel().Cost()
			return HasResources(p, c)
		},
		Start: func(p *ent.Planet) error {
			c := b.ForPlanet(p).NextLevel().Cost()
			subStock(p, c)
			return nil
		},
		Complete: func(p *ent.Planet) error {
			b.UpgradeBuildingOnPlanet(p)
			return nil
		},
		Cancel: func(p *ent.Planet) error {
			c := b.ForPlanet(p).NextLevel().Cost()
			const refundFraction = 0.7
			refund := c.MulFloat64(refundFraction)
			addStock(p, refund)
			return nil
		},
	}
}

// IsBusy checks if there is currently a timer in progress for the group.
func IsBusy(ctx context.Context, p *ent.Planet, g timer.Group) (bool, error) {
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
	if _, ok := err.(*ent.NotFoundError); ok {
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
	if _, ok := err.(*ent.NotFoundError); ok {
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
	busy, err := IsBusy(ctx, p, a.Group)
	if err != nil {
		return err
	}
	if busy {
		return ErrTimerBusy
	}
	if !a.Valid(p) {
		return ErrActionPrerequisitesNotMet
	}
	d := a.Duration(p)
	err = a.Start(p)
	if err != nil {
		return fmt.Errorf("error while calling \"Start\" function for action %q: %w", action, err)
	}
	_, err = SavePlanetResources(ctx, p)
	if err != nil {
		return err
	}
	_, err = tx.Timer.
		Create().
		SetPlanet(p).
		SetAction(action).
		SetGroup(a.Group).
		SetEndTime(timeNow().Add(d)).
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
	err = actions[a].Cancel(p)
	if err != nil {
		return fmt.Errorf("error while calling \"Cancel\" function for action %q: %w", a, err)
	}
	_, err = SavePlanetResources(ctx, p)
	return err
}

// UpdateTimers checks if timers have completed, and if so, triggers the action's
// Complete function and cleans up the timers. This function must be called before
// any information manipulated by the timers/actions is queried.
// In contrast to StartTimer and CancelTimer, UpdateTimers expects the planet
// state *NOT* to be updated. UpdateTimers makes use of the old state of the
// planet to calculate durations and update the state in steps.
func UpdateTimers(ctx context.Context, tx *ent.Tx, p *ent.Planet) error {
	now := timeNow()
	timers, err := p.QueryTimers().
		Where(timer.EndTimeLTE(now)).
		Order(ent.Asc(timer.FieldEndTime)).
		All(ctx)
	if err != nil {
		return fmt.Errorf("unable to retrieve running timers: %v", err)
	}
	if len(timers) == 0 {
		return nil // Fast path
	}
	pwr := NewPlanetWithResourceInfo(p)
	for _, t := range timers {
		pwr.Update(t.EndTime)
		err = actions[t.Action].Complete(pwr.Planet)
		if err != nil {
			return fmt.Errorf("error while calling \"Complete\" function for action %q: %w", t.Action, err)
		}
	}
	p, err = SavePlanetResources(ctx, p)
	if err != nil {
		return err
	}
	_, err = tx.Timer.
		Delete().
		Where(timer.EndTimeLTE(now)).
		Where(timer.HasPlanetWith(planet.IDEQ(p.ID))).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("unable to delete finished timers: %v", err)
	}
	return nil
}
