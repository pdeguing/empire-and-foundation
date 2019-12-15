package data

import (
	"context"
	"errors"
	"time"

	"github.com/pdeguing/empire-and-foundation/ent"
	"github.com/pdeguing/empire-and-foundation/ent/commandplanet"
	"github.com/pdeguing/empire-and-foundation/ent/planet"
)

// ErrCommandPrerequisitesNotMet is returned when a command's Valid()
// method returns false. The command cannot be started.
var ErrCommandPrerequisitesNotMet = errors.New("Cannot start the command because its prerequisites (`! cmd.Valid()`) were not met")

// ErrCommandBussy is returned when another command for the same planet
// and in the same group is already running. Only one command can be
// running at a time.
var ErrCommandBussy = errors.New("Another command is already running for this planet and group")

// TODO: Rename to timer
type delayedCommand struct {
	// Group specifies the group the command belongs to. For each planet
	// there can only be one running command in each group at any given time.
	Group commandplanet.Group

	// Duration returns the time to wait between the start of the command and the completion.
	Duration func(p *ent.Planet) time.Duration

	// Valid checks if the prerequisites of the command are satisfied.
	Valid func(p *ent.Planet) bool

	// Start is triggered when the command is scheduled.
	Start func(ctx context.Context, tx *ent.Tx, p *ent.Planet) error

	// Complete is triggered when the time has elapsed.
	Complete func(ctx context.Context, tx *ent.Tx, p *ent.Planet) error

	// Cancel is triggered when the command is canceled before
	// the time has elapsed.
	Cancel func(ctx context.Context, tx *ent.Tx, p *ent.Planet) error
}

// Timer contains information about a single running timer.
type Timer struct {
	Typ     commandplanet.Typ
	EndTime time.Time
}

// Duration returns the time left until the timer completes.
func (t *Timer) Duration() time.Duration {
	return time.Until(t.EndTime)
}

var delayedCommands = map[commandplanet.Typ]delayedCommand{
	commandplanet.TypUpgradeMetalMine: delayedCommand{
		Group: commandplanet.GroupBuilding,
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

// IsBussy checks if there is currently a command in progress in the group.
func IsBussy(ctx context.Context, p *ent.Planet, g commandplanet.Group) (bool, error) {
	return p.QueryCommands().
		Where(commandplanet.GroupEQ(g)).
		Exist(ctx)
}

// GetCommandInGroup returns the type of command in progress in the group.
func GetCommandInGroup(ctx context.Context, p *ent.Planet, g commandplanet.Group) (*Timer, error) {
	cmd, err := p.QueryCommands().
		Where(commandplanet.GroupEQ(g)).
		Only(ctx)
	if _, ok := err.(*ent.ErrNotFound); ok {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &Timer{
		Typ:     cmd.Typ,
		EndTime: cmd.EndTime,
	}, nil
}

// StartCommand executes the Complete function for command of type t after a calculated duration has passed.
func StartCommand(ctx context.Context, p *ent.Planet, t commandplanet.Typ) error {
	return WithTx(ctx, Client, func(tx *ent.Tx) error {
		// Retrieve the planet *again*, but from within the transaction.
		pTx, err := tx.Planet.Get(ctx, p.ID)
		if err != nil {
			return err
		}
		cmd := delayedCommands[t]
		bussy, err := IsBussy(ctx, pTx, cmd.Group)
		if err != nil {
			return err
		}
		if bussy {
			return ErrCommandBussy
		}
		if !cmd.Valid(pTx) {
			return ErrCommandPrerequisitesNotMet
		}
		d := cmd.Duration(pTx)
		err = cmd.Start(ctx, tx, pTx)
		if err != nil {
			return err
		}
		_, err = tx.CommandPlanet.
			Create().
			SetPlanet(pTx).
			SetTyp(t).
			SetGroup(cmd.Group).
			SetEndTime(time.Now().Add(d)).
			Save(ctx)
		return err
	})
}

// CancelCommand aborts the command from being executed and triggers the Cancel function directly.
func CancelCommand(ctx context.Context, p *ent.Planet, t commandplanet.Typ) error {
	return WithTx(ctx, Client, func(tx *ent.Tx) error {
		// Retrieve the planet *again*, but from within the transaction.
		pTx, err := tx.Planet.Get(ctx, p.ID)
		if err != nil {
			return err
		}
		err = delayedCommands[t].Cancel(ctx, tx, pTx)
		if err != nil {
			return err
		}
		_, err = tx.CommandPlanet.
			Delete().
			Where(commandplanet.HasPlanetWith(planet.IDEQ(pTx.ID))).
			Where(commandplanet.TypEQ(t)).
			Exec(ctx)
		return err
	})
}

// UpdateCommands runs the Complete triggers for all commands for a planet
// for which the duration has elapsed.
func UpdateCommands(ctx context.Context, p *ent.Planet) error {
	return WithTx(ctx, Client, func(tx *ent.Tx) error {
		// Retrieve the planet *again*, but from within the transaction.
		pTx, err := tx.Planet.Get(ctx, p.ID)
		if err != nil {
			return err
		}
		cmds, err := pTx.QueryCommands().
			Where(commandplanet.EndTimeLTE(time.Now())).
			Order(ent.Asc(commandplanet.FieldEndTime)).
			All(ctx)
		if err != nil {
			return err
		}
		for _, cmd := range cmds {
			err = delayedCommands[cmd.Typ].Complete(ctx, tx, pTx)
			if err != nil {
				return err
			}
		}
		_, err = tx.CommandPlanet.
			Delete().
			Where(commandplanet.EndTimeLTE(time.Now())).
			Where(commandplanet.HasPlanetWith(planet.IDEQ(pTx.ID))).
			Exec(ctx)
		return err
	})
}
