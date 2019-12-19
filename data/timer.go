package data

import (
	"context"
	"errors"
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

	// Complete is triggered when the timer is done.
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
			_, err := p.Update().
				SetMetalProdLevel(p.MetalProdLevel + 1).
				Save(ctx)
			return err
		},
		Cancel: func(ctx context.Context, tx *ent.Tx, p *ent.Planet) error {
			c := getMetalMineUpgradeCost(p.MetalProdLevel + 1)
			return addStock(ctx, p, c)
		},
	},
}

// IsBussy checks if there is currently a timer in progress for the group.
func IsBussy(ctx context.Context, p *ent.Planet, g timer.Group) (bool, error) {
	return p.QueryTimers().
		Where(timer.GroupEQ(g)).
		Exist(ctx)
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
		return nil, err
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
		return nil, err
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
func StartTimer(ctx context.Context, p *ent.Planet, action timer.Action) error {
	return WithTx(ctx, Client, func(tx *ent.Tx) error {
		// Retrieve the planet *again*, but from within the transaction.
		pTx, err := tx.Planet.Get(ctx, p.ID)
		if err != nil {
			return err
		}
		a := actions[action]
		bussy, err := IsBussy(ctx, pTx, a.Group)
		if err != nil {
			return err
		}
		if bussy {
			return ErrTimerBussy
		}
		if !a.Valid(pTx) {
			return ErrActionPrerequisitesNotMet
		}
		d := a.Duration(pTx)
		err = a.Start(ctx, tx, pTx)
		if err != nil {
			return err
		}
		_, err = tx.Timer.
			Create().
			SetPlanet(pTx).
			SetAction(action).
			SetGroup(a.Group).
			SetEndTime(time.Now().Add(d)).
			Save(ctx)
		return err
	})
}

// CancelTimer aborts the timer and triggers the action's Cancel function immediately.
func CancelTimer(ctx context.Context, p *ent.Planet, a timer.Action) error {
	return WithTx(ctx, Client, func(tx *ent.Tx) error {
		// Retrieve the planet *again*, but from within the transaction.
		pTx, err := tx.Planet.Get(ctx, p.ID)
		if err != nil {
			return err
		}
		err = actions[a].Cancel(ctx, tx, pTx)
		if err != nil {
			return err
		}
		_, err = tx.Timer.
			Delete().
			Where(timer.HasPlanetWith(planet.IDEQ(pTx.ID))).
			Where(timer.ActionEQ(a)).
			Exec(ctx)
		return err
	})
}

// UpdateTimers checks if timers have completed, and if so, triggers the action's
// Complete function and cleans up the timers. This function must be called before
// any information manipulated by the timers/actions is queried.
func UpdateTimers(ctx context.Context, p *ent.Planet) error {
	return WithTx(ctx, Client, func(tx *ent.Tx) error {
		// Retrieve the planet *again*, but from within the transaction.
		pTx, err := tx.Planet.Get(ctx, p.ID)
		if err != nil {
			return err
		}
		timers, err := pTx.QueryTimers().
			Where(timer.EndTimeLTE(time.Now())).
			Order(ent.Asc(timer.FieldEndTime)).
			All(ctx)
		if err != nil {
			return err
		}
		for _, t := range timers {
			err = actions[t.Action].Complete(ctx, tx, pTx)
			if err != nil {
				return err
			}
		}
		_, err = tx.Timer.
			Delete().
			Where(timer.EndTimeLTE(time.Now())).
			Where(timer.HasPlanetWith(planet.IDEQ(pTx.ID))).
			Exec(ctx)
		return err
	})
}
