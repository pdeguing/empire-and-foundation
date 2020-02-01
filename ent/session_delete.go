// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/pdeguing/empire-and-foundation/ent/predicate"
	"github.com/pdeguing/empire-and-foundation/ent/session"
)

// SessionDelete is the builder for deleting a Session entity.
type SessionDelete struct {
	config
	predicates []predicate.Session
}

// Where adds a new predicate to the delete builder.
func (sd *SessionDelete) Where(ps ...predicate.Session) *SessionDelete {
	sd.predicates = append(sd.predicates, ps...)
	return sd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sd *SessionDelete) Exec(ctx context.Context) (int, error) {
	return sd.sqlExec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (sd *SessionDelete) ExecX(ctx context.Context) int {
	n, err := sd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sd *SessionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: session.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: session.FieldID,
			},
		},
	}
	if ps := sd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, sd.driver, _spec)
}

// SessionDeleteOne is the builder for deleting a single Session entity.
type SessionDeleteOne struct {
	sd *SessionDelete
}

// Exec executes the deletion query.
func (sdo *SessionDeleteOne) Exec(ctx context.Context) error {
	n, err := sdo.sd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &ErrNotFound{session.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sdo *SessionDeleteOne) ExecX(ctx context.Context) {
	sdo.sd.ExecX(ctx)
}
