// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/pdeguing/empire-and-foundation/ent/planet"
	"github.com/pdeguing/empire-and-foundation/ent/predicate"
	"github.com/pdeguing/empire-and-foundation/ent/timer"
)

// TimerUpdate is the builder for updating Timer entities.
type TimerUpdate struct {
	config
	hooks      []Hook
	mutation   *TimerMutation
	predicates []predicate.Timer
}

// Where adds a new predicate for the builder.
func (tu *TimerUpdate) Where(ps ...predicate.Timer) *TimerUpdate {
	tu.predicates = append(tu.predicates, ps...)
	return tu
}

// SetPlanetID sets the planet edge to Planet by id.
func (tu *TimerUpdate) SetPlanetID(id int) *TimerUpdate {
	tu.mutation.SetPlanetID(id)
	return tu
}

// SetNillablePlanetID sets the planet edge to Planet by id if the given value is not nil.
func (tu *TimerUpdate) SetNillablePlanetID(id *int) *TimerUpdate {
	if id != nil {
		tu = tu.SetPlanetID(*id)
	}
	return tu
}

// SetPlanet sets the planet edge to Planet.
func (tu *TimerUpdate) SetPlanet(p *Planet) *TimerUpdate {
	return tu.SetPlanetID(p.ID)
}

// ClearPlanet clears the planet edge to Planet.
func (tu *TimerUpdate) ClearPlanet() *TimerUpdate {
	tu.mutation.ClearPlanet()
	return tu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (tu *TimerUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(tu.hooks) == 0 {
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TimerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TimerUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TimerUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TimerUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tu *TimerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   timer.Table,
			Columns: timer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: timer.FieldID,
			},
		},
	}
	if ps := tu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if tu.mutation.PlanetCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.PlanetIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{timer.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// TimerUpdateOne is the builder for updating a single Timer entity.
type TimerUpdateOne struct {
	config
	hooks    []Hook
	mutation *TimerMutation
}

// SetPlanetID sets the planet edge to Planet by id.
func (tuo *TimerUpdateOne) SetPlanetID(id int) *TimerUpdateOne {
	tuo.mutation.SetPlanetID(id)
	return tuo
}

// SetNillablePlanetID sets the planet edge to Planet by id if the given value is not nil.
func (tuo *TimerUpdateOne) SetNillablePlanetID(id *int) *TimerUpdateOne {
	if id != nil {
		tuo = tuo.SetPlanetID(*id)
	}
	return tuo
}

// SetPlanet sets the planet edge to Planet.
func (tuo *TimerUpdateOne) SetPlanet(p *Planet) *TimerUpdateOne {
	return tuo.SetPlanetID(p.ID)
}

// ClearPlanet clears the planet edge to Planet.
func (tuo *TimerUpdateOne) ClearPlanet() *TimerUpdateOne {
	tuo.mutation.ClearPlanet()
	return tuo
}

// Save executes the query and returns the updated entity.
func (tuo *TimerUpdateOne) Save(ctx context.Context) (*Timer, error) {

	var (
		err  error
		node *Timer
	)
	if len(tuo.hooks) == 0 {
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TimerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			mut = tuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TimerUpdateOne) SaveX(ctx context.Context) *Timer {
	t, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return t
}

// Exec executes the query on the entity.
func (tuo *TimerUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TimerUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuo *TimerUpdateOne) sqlSave(ctx context.Context) (t *Timer, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   timer.Table,
			Columns: timer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: timer.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, fmt.Errorf("missing Timer.ID for update")
	}
	_spec.Node.ID.Value = id
	if tuo.mutation.PlanetCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.PlanetIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	t = &Timer{config: tuo.config}
	_spec.Assign = t.assignValues
	_spec.ScanValues = t.scanValues()
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{timer.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return t, nil
}
