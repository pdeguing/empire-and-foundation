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
				"upgrade_metal_mine",
				"upgrade_hydrogen_extractor",
				"upgrade_silica_quarry",
				"upgrade_solar_plant",
				"upgrade_housing_facilities",
			).
			Immutable(),
		field.Enum("group").
			Values("building").
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
