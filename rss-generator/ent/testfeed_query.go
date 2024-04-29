// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ITK13201/rss-generator/ent/predicate"
	"github.com/ITK13201/rss-generator/ent/site"
	"github.com/ITK13201/rss-generator/ent/testfeed"
	"github.com/ITK13201/rss-generator/ent/testfeeditem"
	"github.com/google/uuid"
)

// TestFeedQuery is the builder for querying TestFeed entities.
type TestFeedQuery struct {
	config
	ctx               *QueryContext
	order             []testfeed.OrderOption
	inters            []Interceptor
	predicates        []predicate.TestFeed
	withSite          *SiteQuery
	withTestFeedItems *TestFeedItemQuery
	withFKs           bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TestFeedQuery builder.
func (tfq *TestFeedQuery) Where(ps ...predicate.TestFeed) *TestFeedQuery {
	tfq.predicates = append(tfq.predicates, ps...)
	return tfq
}

// Limit the number of records to be returned by this query.
func (tfq *TestFeedQuery) Limit(limit int) *TestFeedQuery {
	tfq.ctx.Limit = &limit
	return tfq
}

// Offset to start from.
func (tfq *TestFeedQuery) Offset(offset int) *TestFeedQuery {
	tfq.ctx.Offset = &offset
	return tfq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tfq *TestFeedQuery) Unique(unique bool) *TestFeedQuery {
	tfq.ctx.Unique = &unique
	return tfq
}

// Order specifies how the records should be ordered.
func (tfq *TestFeedQuery) Order(o ...testfeed.OrderOption) *TestFeedQuery {
	tfq.order = append(tfq.order, o...)
	return tfq
}

// QuerySite chains the current query on the "site" edge.
func (tfq *TestFeedQuery) QuerySite() *SiteQuery {
	query := (&SiteClient{config: tfq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tfq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(testfeed.Table, testfeed.FieldID, selector),
			sqlgraph.To(site.Table, site.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, testfeed.SiteTable, testfeed.SiteColumn),
		)
		fromU = sqlgraph.SetNeighbors(tfq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTestFeedItems chains the current query on the "test_feed_items" edge.
func (tfq *TestFeedQuery) QueryTestFeedItems() *TestFeedItemQuery {
	query := (&TestFeedItemClient{config: tfq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tfq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(testfeed.Table, testfeed.FieldID, selector),
			sqlgraph.To(testfeeditem.Table, testfeeditem.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, testfeed.TestFeedItemsTable, testfeed.TestFeedItemsColumn),
		)
		fromU = sqlgraph.SetNeighbors(tfq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TestFeed entity from the query.
// Returns a *NotFoundError when no TestFeed was found.
func (tfq *TestFeedQuery) First(ctx context.Context) (*TestFeed, error) {
	nodes, err := tfq.Limit(1).All(setContextOp(ctx, tfq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{testfeed.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tfq *TestFeedQuery) FirstX(ctx context.Context) *TestFeed {
	node, err := tfq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TestFeed ID from the query.
// Returns a *NotFoundError when no TestFeed ID was found.
func (tfq *TestFeedQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = tfq.Limit(1).IDs(setContextOp(ctx, tfq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{testfeed.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tfq *TestFeedQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := tfq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TestFeed entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TestFeed entity is found.
// Returns a *NotFoundError when no TestFeed entities are found.
func (tfq *TestFeedQuery) Only(ctx context.Context) (*TestFeed, error) {
	nodes, err := tfq.Limit(2).All(setContextOp(ctx, tfq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{testfeed.Label}
	default:
		return nil, &NotSingularError{testfeed.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tfq *TestFeedQuery) OnlyX(ctx context.Context) *TestFeed {
	node, err := tfq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TestFeed ID in the query.
// Returns a *NotSingularError when more than one TestFeed ID is found.
// Returns a *NotFoundError when no entities are found.
func (tfq *TestFeedQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = tfq.Limit(2).IDs(setContextOp(ctx, tfq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{testfeed.Label}
	default:
		err = &NotSingularError{testfeed.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tfq *TestFeedQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := tfq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TestFeeds.
func (tfq *TestFeedQuery) All(ctx context.Context) ([]*TestFeed, error) {
	ctx = setContextOp(ctx, tfq.ctx, "All")
	if err := tfq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*TestFeed, *TestFeedQuery]()
	return withInterceptors[[]*TestFeed](ctx, tfq, qr, tfq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tfq *TestFeedQuery) AllX(ctx context.Context) []*TestFeed {
	nodes, err := tfq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TestFeed IDs.
func (tfq *TestFeedQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if tfq.ctx.Unique == nil && tfq.path != nil {
		tfq.Unique(true)
	}
	ctx = setContextOp(ctx, tfq.ctx, "IDs")
	if err = tfq.Select(testfeed.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tfq *TestFeedQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := tfq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tfq *TestFeedQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tfq.ctx, "Count")
	if err := tfq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tfq, querierCount[*TestFeedQuery](), tfq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tfq *TestFeedQuery) CountX(ctx context.Context) int {
	count, err := tfq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tfq *TestFeedQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tfq.ctx, "Exist")
	switch _, err := tfq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tfq *TestFeedQuery) ExistX(ctx context.Context) bool {
	exist, err := tfq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TestFeedQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tfq *TestFeedQuery) Clone() *TestFeedQuery {
	if tfq == nil {
		return nil
	}
	return &TestFeedQuery{
		config:            tfq.config,
		ctx:               tfq.ctx.Clone(),
		order:             append([]testfeed.OrderOption{}, tfq.order...),
		inters:            append([]Interceptor{}, tfq.inters...),
		predicates:        append([]predicate.TestFeed{}, tfq.predicates...),
		withSite:          tfq.withSite.Clone(),
		withTestFeedItems: tfq.withTestFeedItems.Clone(),
		// clone intermediate query.
		sql:  tfq.sql.Clone(),
		path: tfq.path,
	}
}

// WithSite tells the query-builder to eager-load the nodes that are connected to
// the "site" edge. The optional arguments are used to configure the query builder of the edge.
func (tfq *TestFeedQuery) WithSite(opts ...func(*SiteQuery)) *TestFeedQuery {
	query := (&SiteClient{config: tfq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tfq.withSite = query
	return tfq
}

// WithTestFeedItems tells the query-builder to eager-load the nodes that are connected to
// the "test_feed_items" edge. The optional arguments are used to configure the query builder of the edge.
func (tfq *TestFeedQuery) WithTestFeedItems(opts ...func(*TestFeedItemQuery)) *TestFeedQuery {
	query := (&TestFeedItemClient{config: tfq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tfq.withTestFeedItems = query
	return tfq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TestFeed.Query().
//		GroupBy(testfeed.FieldTitle).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tfq *TestFeedQuery) GroupBy(field string, fields ...string) *TestFeedGroupBy {
	tfq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TestFeedGroupBy{build: tfq}
	grbuild.flds = &tfq.ctx.Fields
	grbuild.label = testfeed.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//	}
//
//	client.TestFeed.Query().
//		Select(testfeed.FieldTitle).
//		Scan(ctx, &v)
func (tfq *TestFeedQuery) Select(fields ...string) *TestFeedSelect {
	tfq.ctx.Fields = append(tfq.ctx.Fields, fields...)
	sbuild := &TestFeedSelect{TestFeedQuery: tfq}
	sbuild.label = testfeed.Label
	sbuild.flds, sbuild.scan = &tfq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TestFeedSelect configured with the given aggregations.
func (tfq *TestFeedQuery) Aggregate(fns ...AggregateFunc) *TestFeedSelect {
	return tfq.Select().Aggregate(fns...)
}

func (tfq *TestFeedQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tfq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tfq); err != nil {
				return err
			}
		}
	}
	for _, f := range tfq.ctx.Fields {
		if !testfeed.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tfq.path != nil {
		prev, err := tfq.path(ctx)
		if err != nil {
			return err
		}
		tfq.sql = prev
	}
	return nil
}

func (tfq *TestFeedQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TestFeed, error) {
	var (
		nodes       = []*TestFeed{}
		withFKs     = tfq.withFKs
		_spec       = tfq.querySpec()
		loadedTypes = [2]bool{
			tfq.withSite != nil,
			tfq.withTestFeedItems != nil,
		}
	)
	if tfq.withSite != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, testfeed.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TestFeed).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TestFeed{config: tfq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tfq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tfq.withSite; query != nil {
		if err := tfq.loadSite(ctx, query, nodes, nil,
			func(n *TestFeed, e *Site) { n.Edges.Site = e }); err != nil {
			return nil, err
		}
	}
	if query := tfq.withTestFeedItems; query != nil {
		if err := tfq.loadTestFeedItems(ctx, query, nodes,
			func(n *TestFeed) { n.Edges.TestFeedItems = []*TestFeedItem{} },
			func(n *TestFeed, e *TestFeedItem) { n.Edges.TestFeedItems = append(n.Edges.TestFeedItems, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tfq *TestFeedQuery) loadSite(ctx context.Context, query *SiteQuery, nodes []*TestFeed, init func(*TestFeed), assign func(*TestFeed, *Site)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*TestFeed)
	for i := range nodes {
		if nodes[i].site_id == nil {
			continue
		}
		fk := *nodes[i].site_id
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(site.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "site_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (tfq *TestFeedQuery) loadTestFeedItems(ctx context.Context, query *TestFeedItemQuery, nodes []*TestFeed, init func(*TestFeed), assign func(*TestFeed, *TestFeedItem)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*TestFeed)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.TestFeedItem(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(testfeed.TestFeedItemsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.test_feed_id
		if fk == nil {
			return fmt.Errorf(`foreign-key "test_feed_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "test_feed_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (tfq *TestFeedQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tfq.querySpec()
	_spec.Node.Columns = tfq.ctx.Fields
	if len(tfq.ctx.Fields) > 0 {
		_spec.Unique = tfq.ctx.Unique != nil && *tfq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tfq.driver, _spec)
}

func (tfq *TestFeedQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(testfeed.Table, testfeed.Columns, sqlgraph.NewFieldSpec(testfeed.FieldID, field.TypeUUID))
	_spec.From = tfq.sql
	if unique := tfq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tfq.path != nil {
		_spec.Unique = true
	}
	if fields := tfq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, testfeed.FieldID)
		for i := range fields {
			if fields[i] != testfeed.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tfq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tfq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tfq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tfq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tfq *TestFeedQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tfq.driver.Dialect())
	t1 := builder.Table(testfeed.Table)
	columns := tfq.ctx.Fields
	if len(columns) == 0 {
		columns = testfeed.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tfq.sql != nil {
		selector = tfq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tfq.ctx.Unique != nil && *tfq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range tfq.predicates {
		p(selector)
	}
	for _, p := range tfq.order {
		p(selector)
	}
	if offset := tfq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tfq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TestFeedGroupBy is the group-by builder for TestFeed entities.
type TestFeedGroupBy struct {
	selector
	build *TestFeedQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tfgb *TestFeedGroupBy) Aggregate(fns ...AggregateFunc) *TestFeedGroupBy {
	tfgb.fns = append(tfgb.fns, fns...)
	return tfgb
}

// Scan applies the selector query and scans the result into the given value.
func (tfgb *TestFeedGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tfgb.build.ctx, "GroupBy")
	if err := tfgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TestFeedQuery, *TestFeedGroupBy](ctx, tfgb.build, tfgb, tfgb.build.inters, v)
}

func (tfgb *TestFeedGroupBy) sqlScan(ctx context.Context, root *TestFeedQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tfgb.fns))
	for _, fn := range tfgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tfgb.flds)+len(tfgb.fns))
		for _, f := range *tfgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tfgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tfgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TestFeedSelect is the builder for selecting fields of TestFeed entities.
type TestFeedSelect struct {
	*TestFeedQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (tfs *TestFeedSelect) Aggregate(fns ...AggregateFunc) *TestFeedSelect {
	tfs.fns = append(tfs.fns, fns...)
	return tfs
}

// Scan applies the selector query and scans the result into the given value.
func (tfs *TestFeedSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tfs.ctx, "Select")
	if err := tfs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TestFeedQuery, *TestFeedSelect](ctx, tfs.TestFeedQuery, tfs, tfs.inters, v)
}

func (tfs *TestFeedSelect) sqlScan(ctx context.Context, root *TestFeedQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(tfs.fns))
	for _, fn := range tfs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*tfs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tfs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
