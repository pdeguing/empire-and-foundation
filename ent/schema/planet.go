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
	return []ent.Field{
		field.Int64(r.Type).
			NonNegative().
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
		field.Int("solar_prod_level").
			NonNegative().
			Default(0),
	}
}

type ShipFactoryMixin struct {}

func (ShipFactoryMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int("ship_factory_level").
			NonNegative().
			Default(0),
	}
}


type ResearchCenterMixin struct {}

func (ResearchCenterMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int("research_center_level").
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
		ShipFactoryMixin{},
		ResearchCenterMixin{},
		PositionMixin{},
		FleetMixin{},
	}
}

// Fields of the Planet.
func (Planet) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Default("Unnamed"),
		field.Enum("planet_type").
			Values("habitable", "mineral", "gas_giant", "ice_giant").
			Immutable(),
		field.String("planet_skin"),
		field.Time("last_resource_update").
			Default(time.Now),
	}
}

// Edges of the Planet.
func (Planet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("planets").
			Unique(),
		edge.To("timers", Timer.Type),
	}
}
