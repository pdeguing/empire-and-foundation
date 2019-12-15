// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/pdeguing/empire-and-foundation/ent/commandplanet"
	"github.com/pdeguing/empire-and-foundation/ent/planet"
	"github.com/pdeguing/empire-and-foundation/ent/predicate"
)

// CommandPlanetQuery is the builder for querying CommandPlanet entities.
type CommandPlanetQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.CommandPlanet
	// intermediate queries.
	sql *sql.Selector
}

// Where adds a new predicate for the builder.
func (cpq *CommandPlanetQuery) Where(ps ...predicate.CommandPlanet) *CommandPlanetQuery {
	cpq.predicates = append(cpq.predicates, ps...)
	return cpq
}

// Limit adds a limit step to the query.
func (cpq *CommandPlanetQuery) Limit(limit int) *CommandPlanetQuery {
	cpq.limit = &limit
	return cpq
}

// Offset adds an offset step to the query.
func (cpq *CommandPlanetQuery) Offset(offset int) *CommandPlanetQuery {
	cpq.offset = &offset
	return cpq
}

// Order adds an order step to the query.
func (cpq *CommandPlanetQuery) Order(o ...Order) *CommandPlanetQuery {
	cpq.order = append(cpq.order, o...)
	return cpq
}

// QueryPlanet chains the current query on the planet edge.
func (cpq *CommandPlanetQuery) QueryPlanet() *PlanetQuery {
	query := &PlanetQuery{config: cpq.config}
	step := sql.NewStep(
		sql.From(commandplanet.Table, commandplanet.FieldID, cpq.sqlQuery()),
		sql.To(planet.Table, planet.FieldID),
		sql.Edge(sql.M2O, true, commandplanet.PlanetTable, commandplanet.PlanetColumn),
	)
	query.sql = sql.SetNeighbors(cpq.driver.Dialect(), step)
	return query
}

// First returns the first CommandPlanet entity in the query. Returns *ErrNotFound when no commandplanet was found.
func (cpq *CommandPlanetQuery) First(ctx context.Context) (*CommandPlanet, error) {
	cps, err := cpq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(cps) == 0 {
		return nil, &ErrNotFound{commandplanet.Label}
	}
	return cps[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cpq *CommandPlanetQuery) FirstX(ctx context.Context) *CommandPlanet {
	cp, err := cpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return cp
}

// FirstID returns the first CommandPlanet id in the query. Returns *ErrNotFound when no id was found.
func (cpq *CommandPlanetQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cpq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &ErrNotFound{commandplanet.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (cpq *CommandPlanetQuery) FirstXID(ctx context.Context) int {
	id, err := cpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only CommandPlanet entity in the query, returns an error if not exactly one entity was returned.
func (cpq *CommandPlanetQuery) Only(ctx context.Context) (*CommandPlanet, error) {
	cps, err := cpq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(cps) {
	case 1:
		return cps[0], nil
	case 0:
		return nil, &ErrNotFound{commandplanet.Label}
	default:
		return nil, &ErrNotSingular{commandplanet.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cpq *CommandPlanetQuery) OnlyX(ctx context.Context) *CommandPlanet {
	cp, err := cpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return cp
}

// OnlyID returns the only CommandPlanet id in the query, returns an error if not exactly one id was returned.
func (cpq *CommandPlanetQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cpq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &ErrNotFound{commandplanet.Label}
	default:
		err = &ErrNotSingular{commandplanet.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (cpq *CommandPlanetQuery) OnlyXID(ctx context.Context) int {
	id, err := cpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CommandPlanets.
func (cpq *CommandPlanetQuery) All(ctx context.Context) ([]*CommandPlanet, error) {
	return cpq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (cpq *CommandPlanetQuery) AllX(ctx context.Context) []*CommandPlanet {
	cps, err := cpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return cps
}

// IDs executes the query and returns a list of CommandPlanet ids.
func (cpq *CommandPlanetQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := cpq.Select(commandplanet.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cpq *CommandPlanetQuery) IDsX(ctx context.Context) []int {
	ids, err := cpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cpq *CommandPlanetQuery) Count(ctx context.Context) (int, error) {
	return cpq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (cpq *CommandPlanetQuery) CountX(ctx context.Context) int {
	count, err := cpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cpq *CommandPlanetQuery) Exist(ctx context.Context) (bool, error) {
	return cpq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (cpq *CommandPlanetQuery) ExistX(ctx context.Context) bool {
	exist, err := cpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cpq *CommandPlanetQuery) Clone() *CommandPlanetQuery {
	return &CommandPlanetQuery{
		config:     cpq.config,
		limit:      cpq.limit,
		offset:     cpq.offset,
		order:      append([]Order{}, cpq.order...),
		unique:     append([]string{}, cpq.unique...),
		predicates: append([]predicate.CommandPlanet{}, cpq.predicates...),
		// clone intermediate queries.
		sql: cpq.sql.Clone(),
	}
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Typ commandplanet.Typ `json:"typ,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CommandPlanet.Query().
//		GroupBy(commandplanet.FieldTyp).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (cpq *CommandPlanetQuery) GroupBy(field string, fields ...string) *CommandPlanetGroupBy {
	group := &CommandPlanetGroupBy{config: cpq.config}
	group.fields = append([]string{field}, fields...)
	group.sql = cpq.sqlQuery()
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Typ commandplanet.Typ `json:"typ,omitempty"`
//	}
//
//	client.CommandPlanet.Query().
//		Select(commandplanet.FieldTyp).
//		Scan(ctx, &v)
//
func (cpq *CommandPlanetQuery) Select(field string, fields ...string) *CommandPlanetSelect {
	selector := &CommandPlanetSelect{config: cpq.config}
	selector.fields = append([]string{field}, fields...)
	selector.sql = cpq.sqlQuery()
	return selector
}

func (cpq *CommandPlanetQuery) sqlAll(ctx context.Context) ([]*CommandPlanet, error) {
	rows := &sql.Rows{}
	selector := cpq.sqlQuery()
	if unique := cpq.unique; len(unique) == 0 {
		selector.Distinct()
	}
	query, args := selector.Query()
	if err := cpq.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var cps CommandPlanets
	if err := cps.FromRows(rows); err != nil {
		return nil, err
	}
	cps.config(cpq.config)
	return cps, nil
}

func (cpq *CommandPlanetQuery) sqlCount(ctx context.Context) (int, error) {
	rows := &sql.Rows{}
	selector := cpq.sqlQuery()
	unique := []string{commandplanet.FieldID}
	if len(cpq.unique) > 0 {
		unique = cpq.unique
	}
	selector.Count(sql.Distinct(selector.Columns(unique...)...))
	query, args := selector.Query()
	if err := cpq.driver.Query(ctx, query, args, rows); err != nil {
		return 0, err
	}
	defer rows.Close()
	if !rows.Next() {
		return 0, errors.New("ent: no rows found")
	}
	var n int
	if err := rows.Scan(&n); err != nil {
		return 0, fmt.Errorf("ent: failed reading count: %v", err)
	}
	return n, nil
}

func (cpq *CommandPlanetQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := cpq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (cpq *CommandPlanetQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(cpq.driver.Dialect())
	t1 := builder.Table(commandplanet.Table)
	selector := builder.Select(t1.Columns(commandplanet.Columns...)...).From(t1)
	if cpq.sql != nil {
		selector = cpq.sql
		selector.Select(selector.Columns(commandplanet.Columns...)...)
	}
	for _, p := range cpq.predicates {
		p(selector)
	}
	for _, p := range cpq.order {
		p(selector)
	}
	if offset := cpq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cpq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CommandPlanetGroupBy is the builder for group-by CommandPlanet entities.
type CommandPlanetGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate queries.
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cpgb *CommandPlanetGroupBy) Aggregate(fns ...Aggregate) *CommandPlanetGroupBy {
	cpgb.fns = append(cpgb.fns, fns...)
	return cpgb
}

// Scan applies the group-by query and scan the result into the given value.
func (cpgb *CommandPlanetGroupBy) Scan(ctx context.Context, v interface{}) error {
	return cpgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cpgb *CommandPlanetGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := cpgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (cpgb *CommandPlanetGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(cpgb.fields) > 1 {
		return nil, errors.New("ent: CommandPlanetGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := cpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cpgb *CommandPlanetGroupBy) StringsX(ctx context.Context) []string {
	v, err := cpgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (cpgb *CommandPlanetGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(cpgb.fields) > 1 {
		return nil, errors.New("ent: CommandPlanetGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := cpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cpgb *CommandPlanetGroupBy) IntsX(ctx context.Context) []int {
	v, err := cpgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (cpgb *CommandPlanetGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(cpgb.fields) > 1 {
		return nil, errors.New("ent: CommandPlanetGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := cpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cpgb *CommandPlanetGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := cpgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (cpgb *CommandPlanetGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(cpgb.fields) > 1 {
		return nil, errors.New("ent: CommandPlanetGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := cpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cpgb *CommandPlanetGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := cpgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cpgb *CommandPlanetGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := cpgb.sqlQuery().Query()
	if err := cpgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cpgb *CommandPlanetGroupBy) sqlQuery() *sql.Selector {
	selector := cpgb.sql
	columns := make([]string, 0, len(cpgb.fields)+len(cpgb.fns))
	columns = append(columns, cpgb.fields...)
	for _, fn := range cpgb.fns {
		columns = append(columns, fn.SQL(selector))
	}
	return selector.Select(columns...).GroupBy(cpgb.fields...)
}

// CommandPlanetSelect is the builder for select fields of CommandPlanet entities.
type CommandPlanetSelect struct {
	config
	fields []string
	// intermediate queries.
	sql *sql.Selector
}

// Scan applies the selector query and scan the result into the given value.
func (cps *CommandPlanetSelect) Scan(ctx context.Context, v interface{}) error {
	return cps.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cps *CommandPlanetSelect) ScanX(ctx context.Context, v interface{}) {
	if err := cps.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (cps *CommandPlanetSelect) Strings(ctx context.Context) ([]string, error) {
	if len(cps.fields) > 1 {
		return nil, errors.New("ent: CommandPlanetSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := cps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cps *CommandPlanetSelect) StringsX(ctx context.Context) []string {
	v, err := cps.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (cps *CommandPlanetSelect) Ints(ctx context.Context) ([]int, error) {
	if len(cps.fields) > 1 {
		return nil, errors.New("ent: CommandPlanetSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := cps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cps *CommandPlanetSelect) IntsX(ctx context.Context) []int {
	v, err := cps.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (cps *CommandPlanetSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(cps.fields) > 1 {
		return nil, errors.New("ent: CommandPlanetSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := cps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cps *CommandPlanetSelect) Float64sX(ctx context.Context) []float64 {
	v, err := cps.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (cps *CommandPlanetSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(cps.fields) > 1 {
		return nil, errors.New("ent: CommandPlanetSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := cps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cps *CommandPlanetSelect) BoolsX(ctx context.Context) []bool {
	v, err := cps.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cps *CommandPlanetSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := cps.sqlQuery().Query()
	if err := cps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cps *CommandPlanetSelect) sqlQuery() sql.Querier {
	view := "commandplanet_view"
	return sql.Dialect(cps.driver.Dialect()).
		Select(cps.fields...).From(cps.sql.As(view))
}
