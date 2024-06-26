// Code generated by ent, DO NOT EDIT.

package feed

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ITK13201/rss-generator/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Feed {
	return predicate.Feed(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Feed {
	return predicate.Feed(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Feed {
	return predicate.Feed(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Feed {
	return predicate.Feed(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Feed {
	return predicate.Feed(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Feed {
	return predicate.Feed(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Feed {
	return predicate.Feed(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Feed {
	return predicate.Feed(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Feed {
	return predicate.Feed(sql.FieldLTE(FieldID, id))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Feed {
	return predicate.Feed(sql.FieldEQ(FieldTitle, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Feed {
	return predicate.Feed(sql.FieldEQ(FieldDescription, v))
}

// Link applies equality check predicate on the "link" field. It's identical to LinkEQ.
func Link(v string) predicate.Feed {
	return predicate.Feed(sql.FieldEQ(FieldLink, v))
}

// PublishedAt applies equality check predicate on the "published_at" field. It's identical to PublishedAtEQ.
func PublishedAt(v time.Time) predicate.Feed {
	return predicate.Feed(sql.FieldEQ(FieldPublishedAt, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Feed {
	return predicate.Feed(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Feed {
	return predicate.Feed(sql.FieldEQ(FieldUpdatedAt, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Feed {
	return predicate.Feed(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Feed {
	return predicate.Feed(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Feed {
	return predicate.Feed(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Feed {
	return predicate.Feed(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Feed {
	return predicate.Feed(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Feed {
	return predicate.Feed(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Feed {
	return predicate.Feed(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Feed {
	return predicate.Feed(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Feed {
	return predicate.Feed(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Feed {
	return predicate.Feed(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Feed {
	return predicate.Feed(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Feed {
	return predicate.Feed(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Feed {
	return predicate.Feed(sql.FieldContainsFold(FieldTitle, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Feed {
	return predicate.Feed(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Feed {
	return predicate.Feed(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Feed {
	return predicate.Feed(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Feed {
	return predicate.Feed(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Feed {
	return predicate.Feed(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Feed {
	return predicate.Feed(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Feed {
	return predicate.Feed(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Feed {
	return predicate.Feed(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Feed {
	return predicate.Feed(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Feed {
	return predicate.Feed(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Feed {
	return predicate.Feed(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Feed {
	return predicate.Feed(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Feed {
	return predicate.Feed(sql.FieldContainsFold(FieldDescription, v))
}

// LinkEQ applies the EQ predicate on the "link" field.
func LinkEQ(v string) predicate.Feed {
	return predicate.Feed(sql.FieldEQ(FieldLink, v))
}

// LinkNEQ applies the NEQ predicate on the "link" field.
func LinkNEQ(v string) predicate.Feed {
	return predicate.Feed(sql.FieldNEQ(FieldLink, v))
}

// LinkIn applies the In predicate on the "link" field.
func LinkIn(vs ...string) predicate.Feed {
	return predicate.Feed(sql.FieldIn(FieldLink, vs...))
}

// LinkNotIn applies the NotIn predicate on the "link" field.
func LinkNotIn(vs ...string) predicate.Feed {
	return predicate.Feed(sql.FieldNotIn(FieldLink, vs...))
}

// LinkGT applies the GT predicate on the "link" field.
func LinkGT(v string) predicate.Feed {
	return predicate.Feed(sql.FieldGT(FieldLink, v))
}

// LinkGTE applies the GTE predicate on the "link" field.
func LinkGTE(v string) predicate.Feed {
	return predicate.Feed(sql.FieldGTE(FieldLink, v))
}

// LinkLT applies the LT predicate on the "link" field.
func LinkLT(v string) predicate.Feed {
	return predicate.Feed(sql.FieldLT(FieldLink, v))
}

// LinkLTE applies the LTE predicate on the "link" field.
func LinkLTE(v string) predicate.Feed {
	return predicate.Feed(sql.FieldLTE(FieldLink, v))
}

// LinkContains applies the Contains predicate on the "link" field.
func LinkContains(v string) predicate.Feed {
	return predicate.Feed(sql.FieldContains(FieldLink, v))
}

// LinkHasPrefix applies the HasPrefix predicate on the "link" field.
func LinkHasPrefix(v string) predicate.Feed {
	return predicate.Feed(sql.FieldHasPrefix(FieldLink, v))
}

// LinkHasSuffix applies the HasSuffix predicate on the "link" field.
func LinkHasSuffix(v string) predicate.Feed {
	return predicate.Feed(sql.FieldHasSuffix(FieldLink, v))
}

// LinkEqualFold applies the EqualFold predicate on the "link" field.
func LinkEqualFold(v string) predicate.Feed {
	return predicate.Feed(sql.FieldEqualFold(FieldLink, v))
}

// LinkContainsFold applies the ContainsFold predicate on the "link" field.
func LinkContainsFold(v string) predicate.Feed {
	return predicate.Feed(sql.FieldContainsFold(FieldLink, v))
}

// PublishedAtEQ applies the EQ predicate on the "published_at" field.
func PublishedAtEQ(v time.Time) predicate.Feed {
	return predicate.Feed(sql.FieldEQ(FieldPublishedAt, v))
}

// PublishedAtNEQ applies the NEQ predicate on the "published_at" field.
func PublishedAtNEQ(v time.Time) predicate.Feed {
	return predicate.Feed(sql.FieldNEQ(FieldPublishedAt, v))
}

// PublishedAtIn applies the In predicate on the "published_at" field.
func PublishedAtIn(vs ...time.Time) predicate.Feed {
	return predicate.Feed(sql.FieldIn(FieldPublishedAt, vs...))
}

// PublishedAtNotIn applies the NotIn predicate on the "published_at" field.
func PublishedAtNotIn(vs ...time.Time) predicate.Feed {
	return predicate.Feed(sql.FieldNotIn(FieldPublishedAt, vs...))
}

// PublishedAtGT applies the GT predicate on the "published_at" field.
func PublishedAtGT(v time.Time) predicate.Feed {
	return predicate.Feed(sql.FieldGT(FieldPublishedAt, v))
}

// PublishedAtGTE applies the GTE predicate on the "published_at" field.
func PublishedAtGTE(v time.Time) predicate.Feed {
	return predicate.Feed(sql.FieldGTE(FieldPublishedAt, v))
}

// PublishedAtLT applies the LT predicate on the "published_at" field.
func PublishedAtLT(v time.Time) predicate.Feed {
	return predicate.Feed(sql.FieldLT(FieldPublishedAt, v))
}

// PublishedAtLTE applies the LTE predicate on the "published_at" field.
func PublishedAtLTE(v time.Time) predicate.Feed {
	return predicate.Feed(sql.FieldLTE(FieldPublishedAt, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Feed {
	return predicate.Feed(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Feed {
	return predicate.Feed(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Feed {
	return predicate.Feed(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Feed {
	return predicate.Feed(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Feed {
	return predicate.Feed(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Feed {
	return predicate.Feed(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Feed {
	return predicate.Feed(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Feed {
	return predicate.Feed(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Feed {
	return predicate.Feed(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Feed {
	return predicate.Feed(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Feed {
	return predicate.Feed(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Feed {
	return predicate.Feed(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Feed {
	return predicate.Feed(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Feed {
	return predicate.Feed(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Feed {
	return predicate.Feed(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Feed {
	return predicate.Feed(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasSite applies the HasEdge predicate on the "site" edge.
func HasSite() predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, SiteTable, SiteColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSiteWith applies the HasEdge predicate on the "site" edge with a given conditions (other predicates).
func HasSiteWith(preds ...predicate.Site) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		step := newSiteStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFeedItems applies the HasEdge predicate on the "feed_items" edge.
func HasFeedItems() predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, FeedItemsTable, FeedItemsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFeedItemsWith applies the HasEdge predicate on the "feed_items" edge with a given conditions (other predicates).
func HasFeedItemsWith(preds ...predicate.FeedItem) predicate.Feed {
	return predicate.Feed(func(s *sql.Selector) {
		step := newFeedItemsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Feed) predicate.Feed {
	return predicate.Feed(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Feed) predicate.Feed {
	return predicate.Feed(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Feed) predicate.Feed {
	return predicate.Feed(sql.NotPredicates(p))
}
