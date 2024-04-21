package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Site holds the schema definition for the Site entity.
type Site struct {
	ent.Schema
}

// Fields of the Site.
func (Site) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable(),
		field.String("slug").NotEmpty().Unique(),
		field.String("title").NotEmpty(),
		field.String("description").Optional(),
		field.String("url").NotEmpty(),
		field.Bool("enable_js_rendering").Default(false),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Site.
func (Site) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("scraping_selector", ScrapingSelector.Type).
			Ref("site").
			Unique(),
	}
}

func (Site) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sites"},
	}
}
