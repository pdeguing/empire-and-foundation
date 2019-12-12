package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Planet holds the schema definition for the Planet entity.
type Planet struct {
	ent.Schema
}

func (Planet) Mixin() []ent.Mixin {
    return []ent.Mixin{
        TimeMixin{},
    }
}

// Fields of the Planet.
func (Planet) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("metal_stock").
			NonNegative().
			Default(0),
		field.Int("metal_mine").
			NonNegative().
			Default(0),
		field.Time("last_metal_update").
			Default(time.Now),
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
