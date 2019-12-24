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

type PositionMixin struct{}

func (PositionMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int("region_code").
			NonNegative().
			Immutable(),
		field.Int("system_code").
			NonNegative().
			Immutable(),
		field.Int("orbit_code").
			NonNegative().
			Immutable(),
		field.Int("suborbit_code").
			NonNegative().
			Immutable(),
		field.Int("position_code").
			NonNegative().
			Immutable().
			Unique(),
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
	PositionMixin{},
    }
}

// Fields of the Planet.
func (Planet) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Default("Unnamed"),
		field.Enum("planet_type").
			Values("habitable", "mineral", "gas giant", "ice giant").
			Immutable(),
		field.String("planet_skin"),
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
