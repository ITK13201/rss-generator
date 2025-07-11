// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ITK13201/rss-generator/ent/feed"
	"github.com/ITK13201/rss-generator/ent/feeditem"
	"github.com/ITK13201/rss-generator/ent/predicate"
	"github.com/google/uuid"
)

// FeedItemUpdate is the builder for updating FeedItem entities.
type FeedItemUpdate struct {
	config
	hooks    []Hook
	mutation *FeedItemMutation
}

// Where appends a list predicates to the FeedItemUpdate builder.
func (fiu *FeedItemUpdate) Where(ps ...predicate.FeedItem) *FeedItemUpdate {
	fiu.mutation.Where(ps...)
	return fiu
}

// SetTitle sets the "title" field.
func (fiu *FeedItemUpdate) SetTitle(s string) *FeedItemUpdate {
	fiu.mutation.SetTitle(s)
	return fiu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (fiu *FeedItemUpdate) SetNillableTitle(s *string) *FeedItemUpdate {
	if s != nil {
		fiu.SetTitle(*s)
	}
	return fiu
}

// SetDescription sets the "description" field.
func (fiu *FeedItemUpdate) SetDescription(s string) *FeedItemUpdate {
	fiu.mutation.SetDescription(s)
	return fiu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (fiu *FeedItemUpdate) SetNillableDescription(s *string) *FeedItemUpdate {
	if s != nil {
		fiu.SetDescription(*s)
	}
	return fiu
}

// SetLink sets the "link" field.
func (fiu *FeedItemUpdate) SetLink(s string) *FeedItemUpdate {
	fiu.mutation.SetLink(s)
	return fiu
}

// SetNillableLink sets the "link" field if the given value is not nil.
func (fiu *FeedItemUpdate) SetNillableLink(s *string) *FeedItemUpdate {
	if s != nil {
		fiu.SetLink(*s)
	}
	return fiu
}

// ClearLink clears the value of the "link" field.
func (fiu *FeedItemUpdate) ClearLink() *FeedItemUpdate {
	fiu.mutation.ClearLink()
	return fiu
}

// SetPublishedAt sets the "published_at" field.
func (fiu *FeedItemUpdate) SetPublishedAt(t time.Time) *FeedItemUpdate {
	fiu.mutation.SetPublishedAt(t)
	return fiu
}

// SetNillablePublishedAt sets the "published_at" field if the given value is not nil.
func (fiu *FeedItemUpdate) SetNillablePublishedAt(t *time.Time) *FeedItemUpdate {
	if t != nil {
		fiu.SetPublishedAt(*t)
	}
	return fiu
}

// SetCreatedAt sets the "created_at" field.
func (fiu *FeedItemUpdate) SetCreatedAt(t time.Time) *FeedItemUpdate {
	fiu.mutation.SetCreatedAt(t)
	return fiu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fiu *FeedItemUpdate) SetNillableCreatedAt(t *time.Time) *FeedItemUpdate {
	if t != nil {
		fiu.SetCreatedAt(*t)
	}
	return fiu
}

// SetUpdatedAt sets the "updated_at" field.
func (fiu *FeedItemUpdate) SetUpdatedAt(t time.Time) *FeedItemUpdate {
	fiu.mutation.SetUpdatedAt(t)
	return fiu
}

// SetFeedID sets the "feed" edge to the Feed entity by ID.
func (fiu *FeedItemUpdate) SetFeedID(id uuid.UUID) *FeedItemUpdate {
	fiu.mutation.SetFeedID(id)
	return fiu
}

// SetFeed sets the "feed" edge to the Feed entity.
func (fiu *FeedItemUpdate) SetFeed(f *Feed) *FeedItemUpdate {
	return fiu.SetFeedID(f.ID)
}

// Mutation returns the FeedItemMutation object of the builder.
func (fiu *FeedItemUpdate) Mutation() *FeedItemMutation {
	return fiu.mutation
}

// ClearFeed clears the "feed" edge to the Feed entity.
func (fiu *FeedItemUpdate) ClearFeed() *FeedItemUpdate {
	fiu.mutation.ClearFeed()
	return fiu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fiu *FeedItemUpdate) Save(ctx context.Context) (int, error) {
	fiu.defaults()
	return withHooks(ctx, fiu.sqlSave, fiu.mutation, fiu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fiu *FeedItemUpdate) SaveX(ctx context.Context) int {
	affected, err := fiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fiu *FeedItemUpdate) Exec(ctx context.Context) error {
	_, err := fiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fiu *FeedItemUpdate) ExecX(ctx context.Context) {
	if err := fiu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fiu *FeedItemUpdate) defaults() {
	if _, ok := fiu.mutation.UpdatedAt(); !ok {
		v := feeditem.UpdateDefaultUpdatedAt()
		fiu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fiu *FeedItemUpdate) check() error {
	if v, ok := fiu.mutation.Title(); ok {
		if err := feeditem.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "FeedItem.title": %w`, err)}
		}
	}
	if _, ok := fiu.mutation.FeedID(); fiu.mutation.FeedCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "FeedItem.feed"`)
	}
	return nil
}

func (fiu *FeedItemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := fiu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(feeditem.Table, feeditem.Columns, sqlgraph.NewFieldSpec(feeditem.FieldID, field.TypeInt))
	if ps := fiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fiu.mutation.Title(); ok {
		_spec.SetField(feeditem.FieldTitle, field.TypeString, value)
	}
	if value, ok := fiu.mutation.Description(); ok {
		_spec.SetField(feeditem.FieldDescription, field.TypeString, value)
	}
	if value, ok := fiu.mutation.Link(); ok {
		_spec.SetField(feeditem.FieldLink, field.TypeString, value)
	}
	if fiu.mutation.LinkCleared() {
		_spec.ClearField(feeditem.FieldLink, field.TypeString)
	}
	if value, ok := fiu.mutation.PublishedAt(); ok {
		_spec.SetField(feeditem.FieldPublishedAt, field.TypeTime, value)
	}
	if value, ok := fiu.mutation.CreatedAt(); ok {
		_spec.SetField(feeditem.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := fiu.mutation.UpdatedAt(); ok {
		_spec.SetField(feeditem.FieldUpdatedAt, field.TypeTime, value)
	}
	if fiu.mutation.FeedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   feeditem.FeedTable,
			Columns: []string{feeditem.FeedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(feed.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fiu.mutation.FeedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   feeditem.FeedTable,
			Columns: []string{feeditem.FeedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(feed.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{feeditem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	fiu.mutation.done = true
	return n, nil
}

// FeedItemUpdateOne is the builder for updating a single FeedItem entity.
type FeedItemUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FeedItemMutation
}

// SetTitle sets the "title" field.
func (fiuo *FeedItemUpdateOne) SetTitle(s string) *FeedItemUpdateOne {
	fiuo.mutation.SetTitle(s)
	return fiuo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (fiuo *FeedItemUpdateOne) SetNillableTitle(s *string) *FeedItemUpdateOne {
	if s != nil {
		fiuo.SetTitle(*s)
	}
	return fiuo
}

// SetDescription sets the "description" field.
func (fiuo *FeedItemUpdateOne) SetDescription(s string) *FeedItemUpdateOne {
	fiuo.mutation.SetDescription(s)
	return fiuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (fiuo *FeedItemUpdateOne) SetNillableDescription(s *string) *FeedItemUpdateOne {
	if s != nil {
		fiuo.SetDescription(*s)
	}
	return fiuo
}

// SetLink sets the "link" field.
func (fiuo *FeedItemUpdateOne) SetLink(s string) *FeedItemUpdateOne {
	fiuo.mutation.SetLink(s)
	return fiuo
}

// SetNillableLink sets the "link" field if the given value is not nil.
func (fiuo *FeedItemUpdateOne) SetNillableLink(s *string) *FeedItemUpdateOne {
	if s != nil {
		fiuo.SetLink(*s)
	}
	return fiuo
}

// ClearLink clears the value of the "link" field.
func (fiuo *FeedItemUpdateOne) ClearLink() *FeedItemUpdateOne {
	fiuo.mutation.ClearLink()
	return fiuo
}

// SetPublishedAt sets the "published_at" field.
func (fiuo *FeedItemUpdateOne) SetPublishedAt(t time.Time) *FeedItemUpdateOne {
	fiuo.mutation.SetPublishedAt(t)
	return fiuo
}

// SetNillablePublishedAt sets the "published_at" field if the given value is not nil.
func (fiuo *FeedItemUpdateOne) SetNillablePublishedAt(t *time.Time) *FeedItemUpdateOne {
	if t != nil {
		fiuo.SetPublishedAt(*t)
	}
	return fiuo
}

// SetCreatedAt sets the "created_at" field.
func (fiuo *FeedItemUpdateOne) SetCreatedAt(t time.Time) *FeedItemUpdateOne {
	fiuo.mutation.SetCreatedAt(t)
	return fiuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fiuo *FeedItemUpdateOne) SetNillableCreatedAt(t *time.Time) *FeedItemUpdateOne {
	if t != nil {
		fiuo.SetCreatedAt(*t)
	}
	return fiuo
}

// SetUpdatedAt sets the "updated_at" field.
func (fiuo *FeedItemUpdateOne) SetUpdatedAt(t time.Time) *FeedItemUpdateOne {
	fiuo.mutation.SetUpdatedAt(t)
	return fiuo
}

// SetFeedID sets the "feed" edge to the Feed entity by ID.
func (fiuo *FeedItemUpdateOne) SetFeedID(id uuid.UUID) *FeedItemUpdateOne {
	fiuo.mutation.SetFeedID(id)
	return fiuo
}

// SetFeed sets the "feed" edge to the Feed entity.
func (fiuo *FeedItemUpdateOne) SetFeed(f *Feed) *FeedItemUpdateOne {
	return fiuo.SetFeedID(f.ID)
}

// Mutation returns the FeedItemMutation object of the builder.
func (fiuo *FeedItemUpdateOne) Mutation() *FeedItemMutation {
	return fiuo.mutation
}

// ClearFeed clears the "feed" edge to the Feed entity.
func (fiuo *FeedItemUpdateOne) ClearFeed() *FeedItemUpdateOne {
	fiuo.mutation.ClearFeed()
	return fiuo
}

// Where appends a list predicates to the FeedItemUpdate builder.
func (fiuo *FeedItemUpdateOne) Where(ps ...predicate.FeedItem) *FeedItemUpdateOne {
	fiuo.mutation.Where(ps...)
	return fiuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fiuo *FeedItemUpdateOne) Select(field string, fields ...string) *FeedItemUpdateOne {
	fiuo.fields = append([]string{field}, fields...)
	return fiuo
}

// Save executes the query and returns the updated FeedItem entity.
func (fiuo *FeedItemUpdateOne) Save(ctx context.Context) (*FeedItem, error) {
	fiuo.defaults()
	return withHooks(ctx, fiuo.sqlSave, fiuo.mutation, fiuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fiuo *FeedItemUpdateOne) SaveX(ctx context.Context) *FeedItem {
	node, err := fiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fiuo *FeedItemUpdateOne) Exec(ctx context.Context) error {
	_, err := fiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fiuo *FeedItemUpdateOne) ExecX(ctx context.Context) {
	if err := fiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fiuo *FeedItemUpdateOne) defaults() {
	if _, ok := fiuo.mutation.UpdatedAt(); !ok {
		v := feeditem.UpdateDefaultUpdatedAt()
		fiuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fiuo *FeedItemUpdateOne) check() error {
	if v, ok := fiuo.mutation.Title(); ok {
		if err := feeditem.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "FeedItem.title": %w`, err)}
		}
	}
	if _, ok := fiuo.mutation.FeedID(); fiuo.mutation.FeedCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "FeedItem.feed"`)
	}
	return nil
}

func (fiuo *FeedItemUpdateOne) sqlSave(ctx context.Context) (_node *FeedItem, err error) {
	if err := fiuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(feeditem.Table, feeditem.Columns, sqlgraph.NewFieldSpec(feeditem.FieldID, field.TypeInt))
	id, ok := fiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "FeedItem.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, feeditem.FieldID)
		for _, f := range fields {
			if !feeditem.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != feeditem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fiuo.mutation.Title(); ok {
		_spec.SetField(feeditem.FieldTitle, field.TypeString, value)
	}
	if value, ok := fiuo.mutation.Description(); ok {
		_spec.SetField(feeditem.FieldDescription, field.TypeString, value)
	}
	if value, ok := fiuo.mutation.Link(); ok {
		_spec.SetField(feeditem.FieldLink, field.TypeString, value)
	}
	if fiuo.mutation.LinkCleared() {
		_spec.ClearField(feeditem.FieldLink, field.TypeString)
	}
	if value, ok := fiuo.mutation.PublishedAt(); ok {
		_spec.SetField(feeditem.FieldPublishedAt, field.TypeTime, value)
	}
	if value, ok := fiuo.mutation.CreatedAt(); ok {
		_spec.SetField(feeditem.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := fiuo.mutation.UpdatedAt(); ok {
		_spec.SetField(feeditem.FieldUpdatedAt, field.TypeTime, value)
	}
	if fiuo.mutation.FeedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   feeditem.FeedTable,
			Columns: []string{feeditem.FeedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(feed.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fiuo.mutation.FeedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   feeditem.FeedTable,
			Columns: []string{feeditem.FeedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(feed.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &FeedItem{config: fiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{feeditem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	fiuo.mutation.done = true
	return _node, nil
}
