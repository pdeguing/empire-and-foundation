package schema

import "github.com/facebookincubator/ent"

// Planet holds the schema definition for the Planet entity.
type Planet struct {
	ent.Schema
}

// Fields of the Planet.
func (Planet) Fields() []ent.Field {
	return []ent.Field{
		field.String("uuid").
		Unique(),
		field.Int("metal_stock").
		Positive(),
		field.Int("metal_mine").
		Positive(),
		field.Time("last_metal_update"),
	}
}

// Edges of the Planet.
func (Planet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
		Ref("planets").
		Unique(),
	}
}
