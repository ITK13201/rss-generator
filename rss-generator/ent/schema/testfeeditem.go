package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

type TestFeedItem struct {
	ent.Schema
}

func (TestFeedItem) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.String("title").NotEmpty().MaxLen(1023),
		field.String("description").NotEmpty().MaxLen(2047),
		field.String("link").Optional().MaxLen(2047),
		field.Time("published_at"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (TestFeedItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("test_feed", TestFeed.Type).StorageKey(edge.Column("test_feed_id")).Required().Unique(),
	}
}

func (TestFeedItem) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "test_feed_items"},
	}
}
