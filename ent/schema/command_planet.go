package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/index"
)

// CommandPlanet holds the schema definition for the CommandPlanet entity.
type CommandPlanet struct {
	ent.Schema
}

// Fields of the CommandPlanet.
func (CommandPlanet) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("typ").
			Values("upgrade_metal_mine").
			Immutable(),
		field.Enum("group").
			Values("building").
			Immutable(),
		field.Time("end_time").
			Immutable(),
	}
}

// Edges of the CommandPlanet.
func (CommandPlanet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("planet", Planet.Type).
			Ref("commands").
			Unique(),
	}
}

// Indexes of the CommandPlanet.
func (CommandPlanet) Indexes() []ent.Index {
    return []ent.Index{
		// There can only be one active command for each group on a planet.
		index.Edges("planet").
			Fields("group").
            Unique(),
    }
}