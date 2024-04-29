// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ITK13201/rss-generator/ent/predicate"
	"github.com/ITK13201/rss-generator/ent/testfeeditem"
)

// TestFeedItemDelete is the builder for deleting a TestFeedItem entity.
type TestFeedItemDelete struct {
	config
	hooks    []Hook
	mutation *TestFeedItemMutation
}

// Where appends a list predicates to the TestFeedItemDelete builder.
func (tfid *TestFeedItemDelete) Where(ps ...predicate.TestFeedItem) *TestFeedItemDelete {
	tfid.mutation.Where(ps...)
	return tfid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tfid *TestFeedItemDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, tfid.sqlExec, tfid.mutation, tfid.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (tfid *TestFeedItemDelete) ExecX(ctx context.Context) int {
	n, err := tfid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tfid *TestFeedItemDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(testfeeditem.Table, sqlgraph.NewFieldSpec(testfeeditem.FieldID, field.TypeInt))
	if ps := tfid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, tfid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	tfid.mutation.done = true
	return affected, err
}

// TestFeedItemDeleteOne is the builder for deleting a single TestFeedItem entity.
type TestFeedItemDeleteOne struct {
	tfid *TestFeedItemDelete
}

// Where appends a list predicates to the TestFeedItemDelete builder.
func (tfido *TestFeedItemDeleteOne) Where(ps ...predicate.TestFeedItem) *TestFeedItemDeleteOne {
	tfido.tfid.mutation.Where(ps...)
	return tfido
}

// Exec executes the deletion query.
func (tfido *TestFeedItemDeleteOne) Exec(ctx context.Context) error {
	n, err := tfido.tfid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{testfeeditem.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tfido *TestFeedItemDeleteOne) ExecX(ctx context.Context) {
	if err := tfido.Exec(ctx); err != nil {
		panic(err)
	}
}
