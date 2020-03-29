// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/pdeguing/empire-and-foundation/ent/planet"
	"github.com/pdeguing/empire-and-foundation/ent/timer"
)

// TimerCreate is the builder for creating a Timer entity.
type TimerCreate struct {
	config
	mutation *TimerMutation
	hooks    []Hook
}

// SetAction sets the action field.
func (tc *TimerCreate) SetAction(t timer.Action) *TimerCreate {
	tc.mutation.SetAction(t)
	return tc
}

// SetGroup sets the group field.
func (tc *TimerCreate) SetGroup(t timer.Group) *TimerCreate {
	tc.mutation.SetGroup(t)
	return tc
}

// SetEndTime sets the end_time field.
func (tc *TimerCreate) SetEndTime(t time.Time) *TimerCreate {
	tc.mutation.SetEndTime(t)
	return tc
}

// SetPlanetID sets the planet edge to Planet by id.
func (tc *TimerCreate) SetPlanetID(id int) *TimerCreate {
	tc.mutation.SetPlanetID(id)
	return tc
}

// SetNillablePlanetID sets the planet edge to Planet by id if the given value is not nil.
func (tc *TimerCreate) SetNillablePlanetID(id *int) *TimerCreate {
	if id != nil {
		tc = tc.SetPlanetID(*id)
	}
	return tc
}

// SetPlanet sets the planet edge to Planet.
func (tc *TimerCreate) SetPlanet(p *Planet) *TimerCreate {
	return tc.SetPlanetID(p.ID)
}

// Save creates the Timer in the database.
func (tc *TimerCreate) Save(ctx context.Context) (*Timer, error) {
	if _, ok := tc.mutation.Action(); !ok {
		return nil, errors.New("ent: missing required field \"action\"")
	}
	if v, ok := tc.mutation.Action(); ok {
		if err := timer.ActionValidator(v); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"action\": %v", err)
		}
	}
	if _, ok := tc.mutation.Group(); !ok {
		return nil, errors.New("ent: missing required field \"group\"")
	}
	if v, ok := tc.mutation.Group(); ok {
		if err := timer.GroupValidator(v); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"group\": %v", err)
		}
	}
	if _, ok := tc.mutation.EndTime(); !ok {
		return nil, errors.New("ent: missing required field \"end_time\"")
	}
	var (
		err  error
		node *Timer
	)
	if len(tc.hooks) == 0 {
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TimerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tc.mutation = mutation
			node, err = tc.sqlSave(ctx)
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TimerCreate) SaveX(ctx context.Context) *Timer {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tc *TimerCreate) sqlSave(ctx context.Context) (*Timer, error) {
	var (
		t     = &Timer{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: timer.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: timer.FieldID,
			},
		}
	)
	if value, ok := tc.mutation.Action(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: timer.FieldAction,
		})
		t.Action = value
	}
	if value, ok := tc.mutation.Group(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: timer.FieldGroup,
		})
		t.Group = value
	}
	if value, ok := tc.mutation.EndTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: timer.FieldEndTime,
		})
		t.EndTime = value
	}
	if nodes := tc.mutation.PlanetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   timer.PlanetTable,
			Columns: []string{timer.PlanetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: planet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	t.ID = int(id)
	return t, nil
}
