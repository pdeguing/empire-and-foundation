// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/pdeguing/empire-and-foundation/ent/planet"
	"github.com/pdeguing/empire-and-foundation/ent/predicate"
	"github.com/pdeguing/empire-and-foundation/ent/timer"
	"github.com/pdeguing/empire-and-foundation/ent/user"
)

// PlanetQuery is the builder for querying Planet entities.
type PlanetQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.Planet
	// eager-loading edges.
	withOwner  *UserQuery
	withTimers *TimerQuery
	withFKs    bool
	// intermediate query.
	sql *sql.Selector
}

// Where adds a new predicate for the builder.
func (pq *PlanetQuery) Where(ps ...predicate.Planet) *PlanetQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit adds a limit step to the query.
func (pq *PlanetQuery) Limit(limit int) *PlanetQuery {
	pq.limit = &limit
	return pq
}

// Offset adds an offset step to the query.
func (pq *PlanetQuery) Offset(offset int) *PlanetQuery {
	pq.offset = &offset
	return pq
}

// Order adds an order step to the query.
func (pq *PlanetQuery) Order(o ...Order) *PlanetQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryOwner chains the current query on the owner edge.
func (pq *PlanetQuery) QueryOwner() *UserQuery {
	query := &UserQuery{config: pq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(planet.Table, planet.FieldID, pq.sqlQuery()),
		sqlgraph.To(user.Table, user.FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, planet.OwnerTable, planet.OwnerColumn),
	)
	query.sql = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
	return query
}

// QueryTimers chains the current query on the timers edge.
func (pq *PlanetQuery) QueryTimers() *TimerQuery {
	query := &TimerQuery{config: pq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(planet.Table, planet.FieldID, pq.sqlQuery()),
		sqlgraph.To(timer.Table, timer.FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, planet.TimersTable, planet.TimersColumn),
	)
	query.sql = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
	return query
}

// First returns the first Planet entity in the query. Returns *NotFoundError when no planet was found.
func (pq *PlanetQuery) First(ctx context.Context) (*Planet, error) {
	pls, err := pq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(pls) == 0 {
		return nil, &NotFoundError{planet.Label}
	}
	return pls[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PlanetQuery) FirstX(ctx context.Context) *Planet {
	pl, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return pl
}

// FirstID returns the first Planet id in the query. Returns *NotFoundError when no id was found.
func (pq *PlanetQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{planet.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (pq *PlanetQuery) FirstXID(ctx context.Context) int {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Planet entity in the query, returns an error if not exactly one entity was returned.
func (pq *PlanetQuery) Only(ctx context.Context) (*Planet, error) {
	pls, err := pq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(pls) {
	case 1:
		return pls[0], nil
	case 0:
		return nil, &NotFoundError{planet.Label}
	default:
		return nil, &NotSingularError{planet.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PlanetQuery) OnlyX(ctx context.Context) *Planet {
	pl, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return pl
}

// OnlyID returns the only Planet id in the query, returns an error if not exactly one id was returned.
func (pq *PlanetQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{planet.Label}
	default:
		err = &NotSingularError{planet.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (pq *PlanetQuery) OnlyXID(ctx context.Context) int {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Planets.
func (pq *PlanetQuery) All(ctx context.Context) ([]*Planet, error) {
	return pq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (pq *PlanetQuery) AllX(ctx context.Context) []*Planet {
	pls, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return pls
}

// IDs executes the query and returns a list of Planet ids.
func (pq *PlanetQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := pq.Select(planet.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PlanetQuery) IDsX(ctx context.Context) []int {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PlanetQuery) Count(ctx context.Context) (int, error) {
	return pq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PlanetQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PlanetQuery) Exist(ctx context.Context) (bool, error) {
	return pq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PlanetQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PlanetQuery) Clone() *PlanetQuery {
	return &PlanetQuery{
		config:     pq.config,
		limit:      pq.limit,
		offset:     pq.offset,
		order:      append([]Order{}, pq.order...),
		unique:     append([]string{}, pq.unique...),
		predicates: append([]predicate.Planet{}, pq.predicates...),
		// clone intermediate query.
		sql: pq.sql.Clone(),
	}
}

//  WithOwner tells the query-builder to eager-loads the nodes that are connected to
// the "owner" edge. The optional arguments used to configure the query builder of the edge.
func (pq *PlanetQuery) WithOwner(opts ...func(*UserQuery)) *PlanetQuery {
	query := &UserQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withOwner = query
	return pq
}

//  WithTimers tells the query-builder to eager-loads the nodes that are connected to
// the "timers" edge. The optional arguments used to configure the query builder of the edge.
func (pq *PlanetQuery) WithTimers(opts ...func(*TimerQuery)) *PlanetQuery {
	query := &TimerQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withTimers = query
	return pq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Planet.Query().
//		GroupBy(planet.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (pq *PlanetQuery) GroupBy(field string, fields ...string) *PlanetGroupBy {
	group := &PlanetGroupBy{config: pq.config}
	group.fields = append([]string{field}, fields...)
	group.sql = pq.sqlQuery()
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Planet.Query().
//		Select(planet.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (pq *PlanetQuery) Select(field string, fields ...string) *PlanetSelect {
	selector := &PlanetSelect{config: pq.config}
	selector.fields = append([]string{field}, fields...)
	selector.sql = pq.sqlQuery()
	return selector
}

func (pq *PlanetQuery) sqlAll(ctx context.Context) ([]*Planet, error) {
	var (
		nodes       = []*Planet{}
		withFKs     = pq.withFKs
		_spec       = pq.querySpec()
		loadedTypes = [2]bool{
			pq.withOwner != nil,
			pq.withTimers != nil,
		}
	)
	if pq.withOwner != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, planet.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &Planet{config: pq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := pq.withOwner; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Planet)
		for i := range nodes {
			if fk := nodes[i].owner_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "owner_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Owner = n
			}
		}
	}

	if query := pq.withTimers; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Planet)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Timer(func(s *sql.Selector) {
			s.Where(sql.InValues(planet.TimersColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.planet_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "planet_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "planet_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Timers = append(node.Edges.Timers, n)
		}
	}

	return nodes, nil
}

func (pq *PlanetQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PlanetQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := pq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (pq *PlanetQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   planet.Table,
			Columns: planet.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: planet.FieldID,
			},
		},
		From:   pq.sql,
		Unique: true,
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *PlanetQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(planet.Table)
	selector := builder.Select(t1.Columns(planet.Columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(planet.Columns...)...)
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PlanetGroupBy is the builder for group-by Planet entities.
type PlanetGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate query.
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PlanetGroupBy) Aggregate(fns ...Aggregate) *PlanetGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the group-by query and scan the result into the given value.
func (pgb *PlanetGroupBy) Scan(ctx context.Context, v interface{}) error {
	return pgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pgb *PlanetGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := pgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (pgb *PlanetGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: PlanetGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pgb *PlanetGroupBy) StringsX(ctx context.Context) []string {
	v, err := pgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (pgb *PlanetGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: PlanetGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pgb *PlanetGroupBy) IntsX(ctx context.Context) []int {
	v, err := pgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (pgb *PlanetGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: PlanetGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pgb *PlanetGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := pgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (pgb *PlanetGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: PlanetGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pgb *PlanetGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := pgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pgb *PlanetGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pgb.sqlQuery().Query()
	if err := pgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pgb *PlanetGroupBy) sqlQuery() *sql.Selector {
	selector := pgb.sql
	columns := make([]string, 0, len(pgb.fields)+len(pgb.fns))
	columns = append(columns, pgb.fields...)
	for _, fn := range pgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(pgb.fields...)
}

// PlanetSelect is the builder for select fields of Planet entities.
type PlanetSelect struct {
	config
	fields []string
	// intermediate queries.
	sql *sql.Selector
}

// Scan applies the selector query and scan the result into the given value.
func (ps *PlanetSelect) Scan(ctx context.Context, v interface{}) error {
	return ps.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ps *PlanetSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ps.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (ps *PlanetSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: PlanetSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ps *PlanetSelect) StringsX(ctx context.Context) []string {
	v, err := ps.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (ps *PlanetSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: PlanetSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ps *PlanetSelect) IntsX(ctx context.Context) []int {
	v, err := ps.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (ps *PlanetSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: PlanetSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ps *PlanetSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ps.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (ps *PlanetSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: PlanetSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ps *PlanetSelect) BoolsX(ctx context.Context) []bool {
	v, err := ps.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ps *PlanetSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ps.sqlQuery().Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ps *PlanetSelect) sqlQuery() sql.Querier {
	selector := ps.sql
	selector.Select(selector.Columns(ps.fields...)...)
	return selector
}
