package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type ScrapingSetting struct {
	ent.Schema
}

func (ScrapingSetting) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.String("selector").NotEmpty(),
		field.String("inner_selector").NotEmpty(),
		field.String("title_selector").NotEmpty(),
		field.String("description_selector").Optional(),
		field.String("link_selector").Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the ScrapingSelector.
func (ScrapingSetting) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("site", Site.Type).StorageKey(edge.Column("site_id")).Required().Unique(),
	}
}

func (ScrapingSetting) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "scraping_settings"},
	}
}
