// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ITK13201/rss-generator/ent/site"
	"github.com/ITK13201/rss-generator/ent/testfeed"
	"github.com/ITK13201/rss-generator/ent/testfeeditem"
	"github.com/google/uuid"
)

// TestFeedCreate is the builder for creating a TestFeed entity.
type TestFeedCreate struct {
	config
	mutation *TestFeedMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetTitle sets the "title" field.
func (tfc *TestFeedCreate) SetTitle(s string) *TestFeedCreate {
	tfc.mutation.SetTitle(s)
	return tfc
}

// SetDescription sets the "description" field.
func (tfc *TestFeedCreate) SetDescription(s string) *TestFeedCreate {
	tfc.mutation.SetDescription(s)
	return tfc
}

// SetLink sets the "link" field.
func (tfc *TestFeedCreate) SetLink(s string) *TestFeedCreate {
	tfc.mutation.SetLink(s)
	return tfc
}

// SetPublishedAt sets the "published_at" field.
func (tfc *TestFeedCreate) SetPublishedAt(t time.Time) *TestFeedCreate {
	tfc.mutation.SetPublishedAt(t)
	return tfc
}

// SetCreatedAt sets the "created_at" field.
func (tfc *TestFeedCreate) SetCreatedAt(t time.Time) *TestFeedCreate {
	tfc.mutation.SetCreatedAt(t)
	return tfc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tfc *TestFeedCreate) SetNillableCreatedAt(t *time.Time) *TestFeedCreate {
	if t != nil {
		tfc.SetCreatedAt(*t)
	}
	return tfc
}

// SetUpdatedAt sets the "updated_at" field.
func (tfc *TestFeedCreate) SetUpdatedAt(t time.Time) *TestFeedCreate {
	tfc.mutation.SetUpdatedAt(t)
	return tfc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tfc *TestFeedCreate) SetNillableUpdatedAt(t *time.Time) *TestFeedCreate {
	if t != nil {
		tfc.SetUpdatedAt(*t)
	}
	return tfc
}

// SetID sets the "id" field.
func (tfc *TestFeedCreate) SetID(u uuid.UUID) *TestFeedCreate {
	tfc.mutation.SetID(u)
	return tfc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (tfc *TestFeedCreate) SetNillableID(u *uuid.UUID) *TestFeedCreate {
	if u != nil {
		tfc.SetID(*u)
	}
	return tfc
}

// SetSiteID sets the "site" edge to the Site entity by ID.
func (tfc *TestFeedCreate) SetSiteID(id int) *TestFeedCreate {
	tfc.mutation.SetSiteID(id)
	return tfc
}

// SetSite sets the "site" edge to the Site entity.
func (tfc *TestFeedCreate) SetSite(s *Site) *TestFeedCreate {
	return tfc.SetSiteID(s.ID)
}

// AddTestFeedItemIDs adds the "test_feed_items" edge to the TestFeedItem entity by IDs.
func (tfc *TestFeedCreate) AddTestFeedItemIDs(ids ...int) *TestFeedCreate {
	tfc.mutation.AddTestFeedItemIDs(ids...)
	return tfc
}

// AddTestFeedItems adds the "test_feed_items" edges to the TestFeedItem entity.
func (tfc *TestFeedCreate) AddTestFeedItems(t ...*TestFeedItem) *TestFeedCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tfc.AddTestFeedItemIDs(ids...)
}

// Mutation returns the TestFeedMutation object of the builder.
func (tfc *TestFeedCreate) Mutation() *TestFeedMutation {
	return tfc.mutation
}

// Save creates the TestFeed in the database.
func (tfc *TestFeedCreate) Save(ctx context.Context) (*TestFeed, error) {
	tfc.defaults()
	return withHooks(ctx, tfc.sqlSave, tfc.mutation, tfc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tfc *TestFeedCreate) SaveX(ctx context.Context) *TestFeed {
	v, err := tfc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tfc *TestFeedCreate) Exec(ctx context.Context) error {
	_, err := tfc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tfc *TestFeedCreate) ExecX(ctx context.Context) {
	if err := tfc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tfc *TestFeedCreate) defaults() {
	if _, ok := tfc.mutation.CreatedAt(); !ok {
		v := testfeed.DefaultCreatedAt()
		tfc.mutation.SetCreatedAt(v)
	}
	if _, ok := tfc.mutation.UpdatedAt(); !ok {
		v := testfeed.DefaultUpdatedAt()
		tfc.mutation.SetUpdatedAt(v)
	}
	if _, ok := tfc.mutation.ID(); !ok {
		v := testfeed.DefaultID()
		tfc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tfc *TestFeedCreate) check() error {
	if _, ok := tfc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "TestFeed.title"`)}
	}
	if v, ok := tfc.mutation.Title(); ok {
		if err := testfeed.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "TestFeed.title": %w`, err)}
		}
	}
	if _, ok := tfc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "TestFeed.description"`)}
	}
	if v, ok := tfc.mutation.Description(); ok {
		if err := testfeed.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "TestFeed.description": %w`, err)}
		}
	}
	if _, ok := tfc.mutation.Link(); !ok {
		return &ValidationError{Name: "link", err: errors.New(`ent: missing required field "TestFeed.link"`)}
	}
	if v, ok := tfc.mutation.Link(); ok {
		if err := testfeed.LinkValidator(v); err != nil {
			return &ValidationError{Name: "link", err: fmt.Errorf(`ent: validator failed for field "TestFeed.link": %w`, err)}
		}
	}
	if _, ok := tfc.mutation.PublishedAt(); !ok {
		return &ValidationError{Name: "published_at", err: errors.New(`ent: missing required field "TestFeed.published_at"`)}
	}
	if _, ok := tfc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "TestFeed.created_at"`)}
	}
	if _, ok := tfc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "TestFeed.updated_at"`)}
	}
	if _, ok := tfc.mutation.SiteID(); !ok {
		return &ValidationError{Name: "site", err: errors.New(`ent: missing required edge "TestFeed.site"`)}
	}
	return nil
}

func (tfc *TestFeedCreate) sqlSave(ctx context.Context) (*TestFeed, error) {
	if err := tfc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tfc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tfc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	tfc.mutation.id = &_node.ID
	tfc.mutation.done = true
	return _node, nil
}

func (tfc *TestFeedCreate) createSpec() (*TestFeed, *sqlgraph.CreateSpec) {
	var (
		_node = &TestFeed{config: tfc.config}
		_spec = sqlgraph.NewCreateSpec(testfeed.Table, sqlgraph.NewFieldSpec(testfeed.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = tfc.conflict
	if id, ok := tfc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := tfc.mutation.Title(); ok {
		_spec.SetField(testfeed.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := tfc.mutation.Description(); ok {
		_spec.SetField(testfeed.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := tfc.mutation.Link(); ok {
		_spec.SetField(testfeed.FieldLink, field.TypeString, value)
		_node.Link = value
	}
	if value, ok := tfc.mutation.PublishedAt(); ok {
		_spec.SetField(testfeed.FieldPublishedAt, field.TypeTime, value)
		_node.PublishedAt = value
	}
	if value, ok := tfc.mutation.CreatedAt(); ok {
		_spec.SetField(testfeed.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tfc.mutation.UpdatedAt(); ok {
		_spec.SetField(testfeed.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := tfc.mutation.SiteIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   testfeed.SiteTable,
			Columns: []string{testfeed.SiteColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(site.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.site_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tfc.mutation.TestFeedItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   testfeed.TestFeedItemsTable,
			Columns: []string{testfeed.TestFeedItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(testfeeditem.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.TestFeed.Create().
//		SetTitle(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TestFeedUpsert) {
//			SetTitle(v+v).
//		}).
//		Exec(ctx)
func (tfc *TestFeedCreate) OnConflict(opts ...sql.ConflictOption) *TestFeedUpsertOne {
	tfc.conflict = opts
	return &TestFeedUpsertOne{
		create: tfc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.TestFeed.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tfc *TestFeedCreate) OnConflictColumns(columns ...string) *TestFeedUpsertOne {
	tfc.conflict = append(tfc.conflict, sql.ConflictColumns(columns...))
	return &TestFeedUpsertOne{
		create: tfc,
	}
}

type (
	// TestFeedUpsertOne is the builder for "upsert"-ing
	//  one TestFeed node.
	TestFeedUpsertOne struct {
		create *TestFeedCreate
	}

	// TestFeedUpsert is the "OnConflict" setter.
	TestFeedUpsert struct {
		*sql.UpdateSet
	}
)

// SetTitle sets the "title" field.
func (u *TestFeedUpsert) SetTitle(v string) *TestFeedUpsert {
	u.Set(testfeed.FieldTitle, v)
	return u
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *TestFeedUpsert) UpdateTitle() *TestFeedUpsert {
	u.SetExcluded(testfeed.FieldTitle)
	return u
}

// SetDescription sets the "description" field.
func (u *TestFeedUpsert) SetDescription(v string) *TestFeedUpsert {
	u.Set(testfeed.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *TestFeedUpsert) UpdateDescription() *TestFeedUpsert {
	u.SetExcluded(testfeed.FieldDescription)
	return u
}

// SetLink sets the "link" field.
func (u *TestFeedUpsert) SetLink(v string) *TestFeedUpsert {
	u.Set(testfeed.FieldLink, v)
	return u
}

// UpdateLink sets the "link" field to the value that was provided on create.
func (u *TestFeedUpsert) UpdateLink() *TestFeedUpsert {
	u.SetExcluded(testfeed.FieldLink)
	return u
}

// SetPublishedAt sets the "published_at" field.
func (u *TestFeedUpsert) SetPublishedAt(v time.Time) *TestFeedUpsert {
	u.Set(testfeed.FieldPublishedAt, v)
	return u
}

// UpdatePublishedAt sets the "published_at" field to the value that was provided on create.
func (u *TestFeedUpsert) UpdatePublishedAt() *TestFeedUpsert {
	u.SetExcluded(testfeed.FieldPublishedAt)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *TestFeedUpsert) SetCreatedAt(v time.Time) *TestFeedUpsert {
	u.Set(testfeed.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TestFeedUpsert) UpdateCreatedAt() *TestFeedUpsert {
	u.SetExcluded(testfeed.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TestFeedUpsert) SetUpdatedAt(v time.Time) *TestFeedUpsert {
	u.Set(testfeed.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TestFeedUpsert) UpdateUpdatedAt() *TestFeedUpsert {
	u.SetExcluded(testfeed.FieldUpdatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.TestFeed.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(testfeed.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *TestFeedUpsertOne) UpdateNewValues() *TestFeedUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(testfeed.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.TestFeed.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *TestFeedUpsertOne) Ignore() *TestFeedUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TestFeedUpsertOne) DoNothing() *TestFeedUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TestFeedCreate.OnConflict
// documentation for more info.
func (u *TestFeedUpsertOne) Update(set func(*TestFeedUpsert)) *TestFeedUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TestFeedUpsert{UpdateSet: update})
	}))
	return u
}

// SetTitle sets the "title" field.
func (u *TestFeedUpsertOne) SetTitle(v string) *TestFeedUpsertOne {
	return u.Update(func(s *TestFeedUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *TestFeedUpsertOne) UpdateTitle() *TestFeedUpsertOne {
	return u.Update(func(s *TestFeedUpsert) {
		s.UpdateTitle()
	})
}

// SetDescription sets the "description" field.
func (u *TestFeedUpsertOne) SetDescription(v string) *TestFeedUpsertOne {
	return u.Update(func(s *TestFeedUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *TestFeedUpsertOne) UpdateDescription() *TestFeedUpsertOne {
	return u.Update(func(s *TestFeedUpsert) {
		s.UpdateDescription()
	})
}

// SetLink sets the "link" field.
func (u *TestFeedUpsertOne) SetLink(v string) *TestFeedUpsertOne {
	return u.Update(func(s *TestFeedUpsert) {
		s.SetLink(v)
	})
}

// UpdateLink sets the "link" field to the value that was provided on create.
func (u *TestFeedUpsertOne) UpdateLink() *TestFeedUpsertOne {
	return u.Update(func(s *TestFeedUpsert) {
		s.UpdateLink()
	})
}

// SetPublishedAt sets the "published_at" field.
func (u *TestFeedUpsertOne) SetPublishedAt(v time.Time) *TestFeedUpsertOne {
	return u.Update(func(s *TestFeedUpsert) {
		s.SetPublishedAt(v)
	})
}

// UpdatePublishedAt sets the "published_at" field to the value that was provided on create.
func (u *TestFeedUpsertOne) UpdatePublishedAt() *TestFeedUpsertOne {
	return u.Update(func(s *TestFeedUpsert) {
		s.UpdatePublishedAt()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *TestFeedUpsertOne) SetCreatedAt(v time.Time) *TestFeedUpsertOne {
	return u.Update(func(s *TestFeedUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TestFeedUpsertOne) UpdateCreatedAt() *TestFeedUpsertOne {
	return u.Update(func(s *TestFeedUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TestFeedUpsertOne) SetUpdatedAt(v time.Time) *TestFeedUpsertOne {
	return u.Update(func(s *TestFeedUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TestFeedUpsertOne) UpdateUpdatedAt() *TestFeedUpsertOne {
	return u.Update(func(s *TestFeedUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *TestFeedUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TestFeedCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TestFeedUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TestFeedUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: TestFeedUpsertOne.ID is not supported by MySQL driver. Use TestFeedUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TestFeedUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TestFeedCreateBulk is the builder for creating many TestFeed entities in bulk.
type TestFeedCreateBulk struct {
	config
	err      error
	builders []*TestFeedCreate
	conflict []sql.ConflictOption
}

// Save creates the TestFeed entities in the database.
func (tfcb *TestFeedCreateBulk) Save(ctx context.Context) ([]*TestFeed, error) {
	if tfcb.err != nil {
		return nil, tfcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tfcb.builders))
	nodes := make([]*TestFeed, len(tfcb.builders))
	mutators := make([]Mutator, len(tfcb.builders))
	for i := range tfcb.builders {
		func(i int, root context.Context) {
			builder := tfcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TestFeedMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tfcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = tfcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tfcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tfcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tfcb *TestFeedCreateBulk) SaveX(ctx context.Context) []*TestFeed {
	v, err := tfcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tfcb *TestFeedCreateBulk) Exec(ctx context.Context) error {
	_, err := tfcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tfcb *TestFeedCreateBulk) ExecX(ctx context.Context) {
	if err := tfcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.TestFeed.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TestFeedUpsert) {
//			SetTitle(v+v).
//		}).
//		Exec(ctx)
func (tfcb *TestFeedCreateBulk) OnConflict(opts ...sql.ConflictOption) *TestFeedUpsertBulk {
	tfcb.conflict = opts
	return &TestFeedUpsertBulk{
		create: tfcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.TestFeed.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tfcb *TestFeedCreateBulk) OnConflictColumns(columns ...string) *TestFeedUpsertBulk {
	tfcb.conflict = append(tfcb.conflict, sql.ConflictColumns(columns...))
	return &TestFeedUpsertBulk{
		create: tfcb,
	}
}

// TestFeedUpsertBulk is the builder for "upsert"-ing
// a bulk of TestFeed nodes.
type TestFeedUpsertBulk struct {
	create *TestFeedCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.TestFeed.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(testfeed.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *TestFeedUpsertBulk) UpdateNewValues() *TestFeedUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(testfeed.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.TestFeed.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *TestFeedUpsertBulk) Ignore() *TestFeedUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TestFeedUpsertBulk) DoNothing() *TestFeedUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TestFeedCreateBulk.OnConflict
// documentation for more info.
func (u *TestFeedUpsertBulk) Update(set func(*TestFeedUpsert)) *TestFeedUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TestFeedUpsert{UpdateSet: update})
	}))
	return u
}

// SetTitle sets the "title" field.
func (u *TestFeedUpsertBulk) SetTitle(v string) *TestFeedUpsertBulk {
	return u.Update(func(s *TestFeedUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *TestFeedUpsertBulk) UpdateTitle() *TestFeedUpsertBulk {
	return u.Update(func(s *TestFeedUpsert) {
		s.UpdateTitle()
	})
}

// SetDescription sets the "description" field.
func (u *TestFeedUpsertBulk) SetDescription(v string) *TestFeedUpsertBulk {
	return u.Update(func(s *TestFeedUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *TestFeedUpsertBulk) UpdateDescription() *TestFeedUpsertBulk {
	return u.Update(func(s *TestFeedUpsert) {
		s.UpdateDescription()
	})
}

// SetLink sets the "link" field.
func (u *TestFeedUpsertBulk) SetLink(v string) *TestFeedUpsertBulk {
	return u.Update(func(s *TestFeedUpsert) {
		s.SetLink(v)
	})
}

// UpdateLink sets the "link" field to the value that was provided on create.
func (u *TestFeedUpsertBulk) UpdateLink() *TestFeedUpsertBulk {
	return u.Update(func(s *TestFeedUpsert) {
		s.UpdateLink()
	})
}

// SetPublishedAt sets the "published_at" field.
func (u *TestFeedUpsertBulk) SetPublishedAt(v time.Time) *TestFeedUpsertBulk {
	return u.Update(func(s *TestFeedUpsert) {
		s.SetPublishedAt(v)
	})
}

// UpdatePublishedAt sets the "published_at" field to the value that was provided on create.
func (u *TestFeedUpsertBulk) UpdatePublishedAt() *TestFeedUpsertBulk {
	return u.Update(func(s *TestFeedUpsert) {
		s.UpdatePublishedAt()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *TestFeedUpsertBulk) SetCreatedAt(v time.Time) *TestFeedUpsertBulk {
	return u.Update(func(s *TestFeedUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TestFeedUpsertBulk) UpdateCreatedAt() *TestFeedUpsertBulk {
	return u.Update(func(s *TestFeedUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TestFeedUpsertBulk) SetUpdatedAt(v time.Time) *TestFeedUpsertBulk {
	return u.Update(func(s *TestFeedUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TestFeedUpsertBulk) UpdateUpdatedAt() *TestFeedUpsertBulk {
	return u.Update(func(s *TestFeedUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *TestFeedUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the TestFeedCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TestFeedCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TestFeedUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}