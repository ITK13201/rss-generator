// Code generated by ent, DO NOT EDIT.

package scrapingselector

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ITK13201/rss-generator/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldLTE(FieldID, id))
}

// Selector applies equality check predicate on the "selector" field. It's identical to SelectorEQ.
func Selector(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldEQ(FieldSelector, v))
}

// InnerSelector applies equality check predicate on the "inner_selector" field. It's identical to InnerSelectorEQ.
func InnerSelector(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldEQ(FieldInnerSelector, v))
}

// TitleSelector applies equality check predicate on the "title_selector" field. It's identical to TitleSelectorEQ.
func TitleSelector(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldEQ(FieldTitleSelector, v))
}

// DescriptionSelector applies equality check predicate on the "description_selector" field. It's identical to DescriptionSelectorEQ.
func DescriptionSelector(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldEQ(FieldDescriptionSelector, v))
}

// LinkSelector applies equality check predicate on the "link_selector" field. It's identical to LinkSelectorEQ.
func LinkSelector(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldEQ(FieldLinkSelector, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldEQ(FieldUpdatedAt, v))
}

// SelectorEQ applies the EQ predicate on the "selector" field.
func SelectorEQ(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldEQ(FieldSelector, v))
}

// SelectorNEQ applies the NEQ predicate on the "selector" field.
func SelectorNEQ(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldNEQ(FieldSelector, v))
}

// SelectorIn applies the In predicate on the "selector" field.
func SelectorIn(vs ...string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldIn(FieldSelector, vs...))
}

// SelectorNotIn applies the NotIn predicate on the "selector" field.
func SelectorNotIn(vs ...string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldNotIn(FieldSelector, vs...))
}

// SelectorGT applies the GT predicate on the "selector" field.
func SelectorGT(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldGT(FieldSelector, v))
}

// SelectorGTE applies the GTE predicate on the "selector" field.
func SelectorGTE(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldGTE(FieldSelector, v))
}

// SelectorLT applies the LT predicate on the "selector" field.
func SelectorLT(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldLT(FieldSelector, v))
}

// SelectorLTE applies the LTE predicate on the "selector" field.
func SelectorLTE(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldLTE(FieldSelector, v))
}

// SelectorContains applies the Contains predicate on the "selector" field.
func SelectorContains(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldContains(FieldSelector, v))
}

// SelectorHasPrefix applies the HasPrefix predicate on the "selector" field.
func SelectorHasPrefix(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldHasPrefix(FieldSelector, v))
}

// SelectorHasSuffix applies the HasSuffix predicate on the "selector" field.
func SelectorHasSuffix(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldHasSuffix(FieldSelector, v))
}

// SelectorEqualFold applies the EqualFold predicate on the "selector" field.
func SelectorEqualFold(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldEqualFold(FieldSelector, v))
}

// SelectorContainsFold applies the ContainsFold predicate on the "selector" field.
func SelectorContainsFold(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldContainsFold(FieldSelector, v))
}

// InnerSelectorEQ applies the EQ predicate on the "inner_selector" field.
func InnerSelectorEQ(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldEQ(FieldInnerSelector, v))
}

// InnerSelectorNEQ applies the NEQ predicate on the "inner_selector" field.
func InnerSelectorNEQ(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldNEQ(FieldInnerSelector, v))
}

// InnerSelectorIn applies the In predicate on the "inner_selector" field.
func InnerSelectorIn(vs ...string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldIn(FieldInnerSelector, vs...))
}

// InnerSelectorNotIn applies the NotIn predicate on the "inner_selector" field.
func InnerSelectorNotIn(vs ...string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldNotIn(FieldInnerSelector, vs...))
}

// InnerSelectorGT applies the GT predicate on the "inner_selector" field.
func InnerSelectorGT(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldGT(FieldInnerSelector, v))
}

// InnerSelectorGTE applies the GTE predicate on the "inner_selector" field.
func InnerSelectorGTE(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldGTE(FieldInnerSelector, v))
}

// InnerSelectorLT applies the LT predicate on the "inner_selector" field.
func InnerSelectorLT(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldLT(FieldInnerSelector, v))
}

// InnerSelectorLTE applies the LTE predicate on the "inner_selector" field.
func InnerSelectorLTE(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldLTE(FieldInnerSelector, v))
}

// InnerSelectorContains applies the Contains predicate on the "inner_selector" field.
func InnerSelectorContains(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldContains(FieldInnerSelector, v))
}

// InnerSelectorHasPrefix applies the HasPrefix predicate on the "inner_selector" field.
func InnerSelectorHasPrefix(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldHasPrefix(FieldInnerSelector, v))
}

// InnerSelectorHasSuffix applies the HasSuffix predicate on the "inner_selector" field.
func InnerSelectorHasSuffix(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldHasSuffix(FieldInnerSelector, v))
}

// InnerSelectorEqualFold applies the EqualFold predicate on the "inner_selector" field.
func InnerSelectorEqualFold(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldEqualFold(FieldInnerSelector, v))
}

// InnerSelectorContainsFold applies the ContainsFold predicate on the "inner_selector" field.
func InnerSelectorContainsFold(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldContainsFold(FieldInnerSelector, v))
}

// TitleSelectorEQ applies the EQ predicate on the "title_selector" field.
func TitleSelectorEQ(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldEQ(FieldTitleSelector, v))
}

// TitleSelectorNEQ applies the NEQ predicate on the "title_selector" field.
func TitleSelectorNEQ(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldNEQ(FieldTitleSelector, v))
}

// TitleSelectorIn applies the In predicate on the "title_selector" field.
func TitleSelectorIn(vs ...string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldIn(FieldTitleSelector, vs...))
}

// TitleSelectorNotIn applies the NotIn predicate on the "title_selector" field.
func TitleSelectorNotIn(vs ...string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldNotIn(FieldTitleSelector, vs...))
}

// TitleSelectorGT applies the GT predicate on the "title_selector" field.
func TitleSelectorGT(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldGT(FieldTitleSelector, v))
}

// TitleSelectorGTE applies the GTE predicate on the "title_selector" field.
func TitleSelectorGTE(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldGTE(FieldTitleSelector, v))
}

// TitleSelectorLT applies the LT predicate on the "title_selector" field.
func TitleSelectorLT(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldLT(FieldTitleSelector, v))
}

// TitleSelectorLTE applies the LTE predicate on the "title_selector" field.
func TitleSelectorLTE(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldLTE(FieldTitleSelector, v))
}

// TitleSelectorContains applies the Contains predicate on the "title_selector" field.
func TitleSelectorContains(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldContains(FieldTitleSelector, v))
}

// TitleSelectorHasPrefix applies the HasPrefix predicate on the "title_selector" field.
func TitleSelectorHasPrefix(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldHasPrefix(FieldTitleSelector, v))
}

// TitleSelectorHasSuffix applies the HasSuffix predicate on the "title_selector" field.
func TitleSelectorHasSuffix(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldHasSuffix(FieldTitleSelector, v))
}

// TitleSelectorEqualFold applies the EqualFold predicate on the "title_selector" field.
func TitleSelectorEqualFold(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldEqualFold(FieldTitleSelector, v))
}

// TitleSelectorContainsFold applies the ContainsFold predicate on the "title_selector" field.
func TitleSelectorContainsFold(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldContainsFold(FieldTitleSelector, v))
}

// DescriptionSelectorEQ applies the EQ predicate on the "description_selector" field.
func DescriptionSelectorEQ(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldEQ(FieldDescriptionSelector, v))
}

// DescriptionSelectorNEQ applies the NEQ predicate on the "description_selector" field.
func DescriptionSelectorNEQ(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldNEQ(FieldDescriptionSelector, v))
}

// DescriptionSelectorIn applies the In predicate on the "description_selector" field.
func DescriptionSelectorIn(vs ...string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldIn(FieldDescriptionSelector, vs...))
}

// DescriptionSelectorNotIn applies the NotIn predicate on the "description_selector" field.
func DescriptionSelectorNotIn(vs ...string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldNotIn(FieldDescriptionSelector, vs...))
}

// DescriptionSelectorGT applies the GT predicate on the "description_selector" field.
func DescriptionSelectorGT(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldGT(FieldDescriptionSelector, v))
}

// DescriptionSelectorGTE applies the GTE predicate on the "description_selector" field.
func DescriptionSelectorGTE(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldGTE(FieldDescriptionSelector, v))
}

// DescriptionSelectorLT applies the LT predicate on the "description_selector" field.
func DescriptionSelectorLT(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldLT(FieldDescriptionSelector, v))
}

// DescriptionSelectorLTE applies the LTE predicate on the "description_selector" field.
func DescriptionSelectorLTE(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldLTE(FieldDescriptionSelector, v))
}

// DescriptionSelectorContains applies the Contains predicate on the "description_selector" field.
func DescriptionSelectorContains(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldContains(FieldDescriptionSelector, v))
}

// DescriptionSelectorHasPrefix applies the HasPrefix predicate on the "description_selector" field.
func DescriptionSelectorHasPrefix(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldHasPrefix(FieldDescriptionSelector, v))
}

// DescriptionSelectorHasSuffix applies the HasSuffix predicate on the "description_selector" field.
func DescriptionSelectorHasSuffix(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldHasSuffix(FieldDescriptionSelector, v))
}

// DescriptionSelectorIsNil applies the IsNil predicate on the "description_selector" field.
func DescriptionSelectorIsNil() predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldIsNull(FieldDescriptionSelector))
}

// DescriptionSelectorNotNil applies the NotNil predicate on the "description_selector" field.
func DescriptionSelectorNotNil() predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldNotNull(FieldDescriptionSelector))
}

// DescriptionSelectorEqualFold applies the EqualFold predicate on the "description_selector" field.
func DescriptionSelectorEqualFold(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldEqualFold(FieldDescriptionSelector, v))
}

// DescriptionSelectorContainsFold applies the ContainsFold predicate on the "description_selector" field.
func DescriptionSelectorContainsFold(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldContainsFold(FieldDescriptionSelector, v))
}

// LinkSelectorEQ applies the EQ predicate on the "link_selector" field.
func LinkSelectorEQ(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldEQ(FieldLinkSelector, v))
}

// LinkSelectorNEQ applies the NEQ predicate on the "link_selector" field.
func LinkSelectorNEQ(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldNEQ(FieldLinkSelector, v))
}

// LinkSelectorIn applies the In predicate on the "link_selector" field.
func LinkSelectorIn(vs ...string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldIn(FieldLinkSelector, vs...))
}

// LinkSelectorNotIn applies the NotIn predicate on the "link_selector" field.
func LinkSelectorNotIn(vs ...string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldNotIn(FieldLinkSelector, vs...))
}

// LinkSelectorGT applies the GT predicate on the "link_selector" field.
func LinkSelectorGT(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldGT(FieldLinkSelector, v))
}

// LinkSelectorGTE applies the GTE predicate on the "link_selector" field.
func LinkSelectorGTE(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldGTE(FieldLinkSelector, v))
}

// LinkSelectorLT applies the LT predicate on the "link_selector" field.
func LinkSelectorLT(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldLT(FieldLinkSelector, v))
}

// LinkSelectorLTE applies the LTE predicate on the "link_selector" field.
func LinkSelectorLTE(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldLTE(FieldLinkSelector, v))
}

// LinkSelectorContains applies the Contains predicate on the "link_selector" field.
func LinkSelectorContains(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldContains(FieldLinkSelector, v))
}

// LinkSelectorHasPrefix applies the HasPrefix predicate on the "link_selector" field.
func LinkSelectorHasPrefix(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldHasPrefix(FieldLinkSelector, v))
}

// LinkSelectorHasSuffix applies the HasSuffix predicate on the "link_selector" field.
func LinkSelectorHasSuffix(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldHasSuffix(FieldLinkSelector, v))
}

// LinkSelectorIsNil applies the IsNil predicate on the "link_selector" field.
func LinkSelectorIsNil() predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldIsNull(FieldLinkSelector))
}

// LinkSelectorNotNil applies the NotNil predicate on the "link_selector" field.
func LinkSelectorNotNil() predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldNotNull(FieldLinkSelector))
}

// LinkSelectorEqualFold applies the EqualFold predicate on the "link_selector" field.
func LinkSelectorEqualFold(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldEqualFold(FieldLinkSelector, v))
}

// LinkSelectorContainsFold applies the ContainsFold predicate on the "link_selector" field.
func LinkSelectorContainsFold(v string) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldContainsFold(FieldLinkSelector, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasSite applies the HasEdge predicate on the "site" edge.
func HasSite() predicate.ScrapingSelector {
	return predicate.ScrapingSelector(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, SiteTable, SiteColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSiteWith applies the HasEdge predicate on the "site" edge with a given conditions (other predicates).
func HasSiteWith(preds ...predicate.Site) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(func(s *sql.Selector) {
		step := newSiteStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ScrapingSelector) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ScrapingSelector) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ScrapingSelector) predicate.ScrapingSelector {
	return predicate.ScrapingSelector(sql.NotPredicates(p))
}
