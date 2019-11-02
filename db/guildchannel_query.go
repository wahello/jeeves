// Code generated by entc, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/AlecAivazis/jeeves/db/guild"
	"github.com/AlecAivazis/jeeves/db/guildchannel"
	"github.com/AlecAivazis/jeeves/db/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
)

// GuildChannelQuery is the builder for querying GuildChannel entities.
type GuildChannelQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.GuildChannel
	// intermediate queries.
	sql *sql.Selector
}

// Where adds a new predicate for the builder.
func (gcq *GuildChannelQuery) Where(ps ...predicate.GuildChannel) *GuildChannelQuery {
	gcq.predicates = append(gcq.predicates, ps...)
	return gcq
}

// Limit adds a limit step to the query.
func (gcq *GuildChannelQuery) Limit(limit int) *GuildChannelQuery {
	gcq.limit = &limit
	return gcq
}

// Offset adds an offset step to the query.
func (gcq *GuildChannelQuery) Offset(offset int) *GuildChannelQuery {
	gcq.offset = &offset
	return gcq
}

// Order adds an order step to the query.
func (gcq *GuildChannelQuery) Order(o ...Order) *GuildChannelQuery {
	gcq.order = append(gcq.order, o...)
	return gcq
}

// QueryGuild chains the current query on the guild edge.
func (gcq *GuildChannelQuery) QueryGuild() *GuildQuery {
	query := &GuildQuery{config: gcq.config}

	builder := sql.Dialect(gcq.driver.Dialect())
	t1 := builder.Table(guild.Table)
	t2 := gcq.sqlQuery()
	t2.Select(t2.C(guildchannel.GuildColumn))
	query.sql = builder.Select(t1.Columns(guild.Columns...)...).
		From(t1).
		Join(t2).
		On(t1.C(guild.FieldID), t2.C(guildchannel.GuildColumn))
	return query
}

// First returns the first GuildChannel entity in the query. Returns *ErrNotFound when no guildchannel was found.
func (gcq *GuildChannelQuery) First(ctx context.Context) (*GuildChannel, error) {
	gcs, err := gcq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(gcs) == 0 {
		return nil, &ErrNotFound{guildchannel.Label}
	}
	return gcs[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gcq *GuildChannelQuery) FirstX(ctx context.Context) *GuildChannel {
	gc, err := gcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return gc
}

// FirstID returns the first GuildChannel id in the query. Returns *ErrNotFound when no id was found.
func (gcq *GuildChannelQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gcq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &ErrNotFound{guildchannel.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (gcq *GuildChannelQuery) FirstXID(ctx context.Context) int {
	id, err := gcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only GuildChannel entity in the query, returns an error if not exactly one entity was returned.
func (gcq *GuildChannelQuery) Only(ctx context.Context) (*GuildChannel, error) {
	gcs, err := gcq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(gcs) {
	case 1:
		return gcs[0], nil
	case 0:
		return nil, &ErrNotFound{guildchannel.Label}
	default:
		return nil, &ErrNotSingular{guildchannel.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gcq *GuildChannelQuery) OnlyX(ctx context.Context) *GuildChannel {
	gc, err := gcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return gc
}

// OnlyID returns the only GuildChannel id in the query, returns an error if not exactly one id was returned.
func (gcq *GuildChannelQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gcq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &ErrNotFound{guildchannel.Label}
	default:
		err = &ErrNotSingular{guildchannel.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (gcq *GuildChannelQuery) OnlyXID(ctx context.Context) int {
	id, err := gcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GuildChannels.
func (gcq *GuildChannelQuery) All(ctx context.Context) ([]*GuildChannel, error) {
	return gcq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (gcq *GuildChannelQuery) AllX(ctx context.Context) []*GuildChannel {
	gcs, err := gcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return gcs
}

// IDs executes the query and returns a list of GuildChannel ids.
func (gcq *GuildChannelQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := gcq.Select(guildchannel.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gcq *GuildChannelQuery) IDsX(ctx context.Context) []int {
	ids, err := gcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gcq *GuildChannelQuery) Count(ctx context.Context) (int, error) {
	return gcq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (gcq *GuildChannelQuery) CountX(ctx context.Context) int {
	count, err := gcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gcq *GuildChannelQuery) Exist(ctx context.Context) (bool, error) {
	return gcq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (gcq *GuildChannelQuery) ExistX(ctx context.Context) bool {
	exist, err := gcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gcq *GuildChannelQuery) Clone() *GuildChannelQuery {
	return &GuildChannelQuery{
		config:     gcq.config,
		limit:      gcq.limit,
		offset:     gcq.offset,
		order:      append([]Order{}, gcq.order...),
		unique:     append([]string{}, gcq.unique...),
		predicates: append([]predicate.GuildChannel{}, gcq.predicates...),
		// clone intermediate queries.
		sql: gcq.sql.Clone(),
	}
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Channel string `json:"channel,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.GuildChannel.Query().
//		GroupBy(guildchannel.FieldChannel).
//		Aggregate(db.Count()).
//		Scan(ctx, &v)
//
func (gcq *GuildChannelQuery) GroupBy(field string, fields ...string) *GuildChannelGroupBy {
	group := &GuildChannelGroupBy{config: gcq.config}
	group.fields = append([]string{field}, fields...)
	group.sql = gcq.sqlQuery()
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Channel string `json:"channel,omitempty"`
//	}
//
//	client.GuildChannel.Query().
//		Select(guildchannel.FieldChannel).
//		Scan(ctx, &v)
//
func (gcq *GuildChannelQuery) Select(field string, fields ...string) *GuildChannelSelect {
	selector := &GuildChannelSelect{config: gcq.config}
	selector.fields = append([]string{field}, fields...)
	selector.sql = gcq.sqlQuery()
	return selector
}

func (gcq *GuildChannelQuery) sqlAll(ctx context.Context) ([]*GuildChannel, error) {
	rows := &sql.Rows{}
	selector := gcq.sqlQuery()
	if unique := gcq.unique; len(unique) == 0 {
		selector.Distinct()
	}
	query, args := selector.Query()
	if err := gcq.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var gcs GuildChannels
	if err := gcs.FromRows(rows); err != nil {
		return nil, err
	}
	gcs.config(gcq.config)
	return gcs, nil
}

func (gcq *GuildChannelQuery) sqlCount(ctx context.Context) (int, error) {
	rows := &sql.Rows{}
	selector := gcq.sqlQuery()
	unique := []string{guildchannel.FieldID}
	if len(gcq.unique) > 0 {
		unique = gcq.unique
	}
	selector.Count(sql.Distinct(selector.Columns(unique...)...))
	query, args := selector.Query()
	if err := gcq.driver.Query(ctx, query, args, rows); err != nil {
		return 0, err
	}
	defer rows.Close()
	if !rows.Next() {
		return 0, errors.New("db: no rows found")
	}
	var n int
	if err := rows.Scan(&n); err != nil {
		return 0, fmt.Errorf("db: failed reading count: %v", err)
	}
	return n, nil
}

func (gcq *GuildChannelQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := gcq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("db: check existence: %v", err)
	}
	return n > 0, nil
}

func (gcq *GuildChannelQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(gcq.driver.Dialect())
	t1 := builder.Table(guildchannel.Table)
	selector := builder.Select(t1.Columns(guildchannel.Columns...)...).From(t1)
	if gcq.sql != nil {
		selector = gcq.sql
		selector.Select(selector.Columns(guildchannel.Columns...)...)
	}
	for _, p := range gcq.predicates {
		p(selector)
	}
	for _, p := range gcq.order {
		p(selector)
	}
	if offset := gcq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gcq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GuildChannelGroupBy is the builder for group-by GuildChannel entities.
type GuildChannelGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate queries.
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gcgb *GuildChannelGroupBy) Aggregate(fns ...Aggregate) *GuildChannelGroupBy {
	gcgb.fns = append(gcgb.fns, fns...)
	return gcgb
}

// Scan applies the group-by query and scan the result into the given value.
func (gcgb *GuildChannelGroupBy) Scan(ctx context.Context, v interface{}) error {
	return gcgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (gcgb *GuildChannelGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := gcgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (gcgb *GuildChannelGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(gcgb.fields) > 1 {
		return nil, errors.New("db: GuildChannelGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := gcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (gcgb *GuildChannelGroupBy) StringsX(ctx context.Context) []string {
	v, err := gcgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (gcgb *GuildChannelGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(gcgb.fields) > 1 {
		return nil, errors.New("db: GuildChannelGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := gcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (gcgb *GuildChannelGroupBy) IntsX(ctx context.Context) []int {
	v, err := gcgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (gcgb *GuildChannelGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(gcgb.fields) > 1 {
		return nil, errors.New("db: GuildChannelGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := gcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (gcgb *GuildChannelGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := gcgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (gcgb *GuildChannelGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(gcgb.fields) > 1 {
		return nil, errors.New("db: GuildChannelGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := gcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (gcgb *GuildChannelGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := gcgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gcgb *GuildChannelGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := gcgb.sqlQuery().Query()
	if err := gcgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (gcgb *GuildChannelGroupBy) sqlQuery() *sql.Selector {
	selector := gcgb.sql
	columns := make([]string, 0, len(gcgb.fields)+len(gcgb.fns))
	columns = append(columns, gcgb.fields...)
	for _, fn := range gcgb.fns {
		columns = append(columns, fn.SQL(selector))
	}
	return selector.Select(columns...).GroupBy(gcgb.fields...)
}

// GuildChannelSelect is the builder for select fields of GuildChannel entities.
type GuildChannelSelect struct {
	config
	fields []string
	// intermediate queries.
	sql *sql.Selector
}

// Scan applies the selector query and scan the result into the given value.
func (gcs *GuildChannelSelect) Scan(ctx context.Context, v interface{}) error {
	return gcs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (gcs *GuildChannelSelect) ScanX(ctx context.Context, v interface{}) {
	if err := gcs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (gcs *GuildChannelSelect) Strings(ctx context.Context) ([]string, error) {
	if len(gcs.fields) > 1 {
		return nil, errors.New("db: GuildChannelSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := gcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (gcs *GuildChannelSelect) StringsX(ctx context.Context) []string {
	v, err := gcs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (gcs *GuildChannelSelect) Ints(ctx context.Context) ([]int, error) {
	if len(gcs.fields) > 1 {
		return nil, errors.New("db: GuildChannelSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := gcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (gcs *GuildChannelSelect) IntsX(ctx context.Context) []int {
	v, err := gcs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (gcs *GuildChannelSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(gcs.fields) > 1 {
		return nil, errors.New("db: GuildChannelSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := gcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (gcs *GuildChannelSelect) Float64sX(ctx context.Context) []float64 {
	v, err := gcs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (gcs *GuildChannelSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(gcs.fields) > 1 {
		return nil, errors.New("db: GuildChannelSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := gcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (gcs *GuildChannelSelect) BoolsX(ctx context.Context) []bool {
	v, err := gcs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gcs *GuildChannelSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := gcs.sqlQuery().Query()
	if err := gcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (gcs *GuildChannelSelect) sqlQuery() sql.Querier {
	view := "guildchannel_view"
	return sql.Dialect(gcs.driver.Dialect()).
		Select(gcs.fields...).From(gcs.sql.As(view))
}
