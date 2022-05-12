//go:build ignore
// +build ignore

package main

import (
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"log"
)

func main() {
	opts := []entc.Option{
		entc.TemplateDir("./template"),
	}

	if err := entc.Generate("./schema", &gen.Config{
		Features: []gen.Feature{
			gen.FeatureUpsert,
		},
	}, opts...); err != nil {
		log.Fatalf("Error: failed running ent codegen: %v", err)
	}
}
