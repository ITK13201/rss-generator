package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// FeedItem holds the schema definition for the FeedItems entity.
type FeedItem struct {
	ent.Schema
}

// Fields of the FeedItem.
func (FeedItem) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.String("title").NotEmpty(),
		field.String("description").NotEmpty(),
		field.String("link").Optional(),
		field.Time("published_at"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the FeedItem.
func (FeedItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("feed", Feed.Type).StorageKey(edge.Column("feed_id")).Required().Unique(),
	}
}

func (FeedItem) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "feed_items"},
	}
}
