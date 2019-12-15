package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

type ResourceMixin struct {
    Type string
}

func (r ResourceMixin) Fields() []ent.Field {
    return []ent.Field {
        field.Int64(r.Type).
            NonNegative().
            Default(0),
        field.Time(r.Type + "_last_update").
            Default(time.Now),
        field.Int(r.Type + "_rate").
            Default(0),
        field.Int(r.Type + "_prod_level").
            NonNegative().
            Default(0),
        field.Int(r.Type + "_storage_level").
            NonNegative().
            Default(0),
    }
}

type EnergyMixin struct{}

func (EnergyMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("energy_cons").
			NonNegative().
			Default(0),
		field.Int64("energy_prod").
			NonNegative().
			Default(0),
		field.Int("solar_prod_level").
			NonNegative().
			Default(0),
	}
}

// Planet holds the schema definition for the Planet entity.
type Planet struct {
	ent.Schema
}

func (Planet) Mixin() []ent.Mixin {
    return []ent.Mixin{
        TimeMixin{},
	ResourceMixin{Type: "metal"},
	ResourceMixin{Type: "hydrogen"},
	ResourceMixin{Type: "silica"},
	ResourceMixin{Type: "population"},
	EnergyMixin{},
    }
}

// Fields of the Planet.
func (Planet) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Default("Unnamed"),
	}
}

// Edges of the Planet.
func (Planet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("planets").
			Unique(),
		edge.To("commands", CommandPlanet.Type),
	}
}
