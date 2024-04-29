package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

type TestFeed struct {
	ent.Schema
}

func (TestFeed) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique().Immutable(),
		field.String("title").NotEmpty().MaxLen(1023),
		field.String("description").NotEmpty().MaxLen(2047),
		field.String("link").NotEmpty().MaxLen(2047),
		field.Time("published_at"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (TestFeed) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("site", Site.Type).StorageKey(edge.Column("site_id")).Required().Unique(),
		edge.From("test_feed_items", TestFeedItem.Type).Ref("test_feed"),
	}
}

func (TestFeed) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "test_feeds"},
	}
}
