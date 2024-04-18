//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	conf := &gen.Config{
		Features: []gen.Feature{
			gen.FeatureUpsert,
			gen.FeatureExecQuery,
		},
	}
	opts := []entc.Option{
		entc.TemplateDir("./ent/template"),
	}
	if err := entc.Generate("./ent/schema", conf, opts...); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
