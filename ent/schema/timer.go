package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/index"
)

// Timer holds the schema definition for the Timer entity.
type Timer struct {
	ent.Schema
}

// Fields of the Timer.
func (Timer) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("action").
			Values(
				"upgrade_metal_prod",
				"upgrade_hydrogen_prod",
				"upgrade_silica_prod",
				"upgrade_solar_prod",
				"upgrade_urbanism",
				"upgrade_metal_storage",
				"upgrade_hydrogen_storage",
				"upgrade_silica_storage",
				"upgrade_research_center",
				"upgrade_ship_factory",
				"build_caravel",
				"build_light_fighter",
				"build_corvette",
				"build_frigate",
				"build_probe",
				"build_small_cargo",
				"build_medium_cargo",
				"build_colonization_ark",
				"test", // Used in the tests
			).
			Immutable(),
		field.Enum("group").
			Values(
				"building",
				"ship",
			).
			Immutable(),
		field.Time("end_time").
			Immutable(),
	}
}

// Edges of the Timer.
func (Timer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("planet", Planet.Type).
			Ref("timers").
			Unique(),
	}
}

// Indexes of the Timer.
func (Timer) Indexes() []ent.Index {
	return []ent.Index{
		// There can only be one active command for each group on a planet.
		index.Edges("planet").
			Fields("group").
			Unique(),
	}
}
