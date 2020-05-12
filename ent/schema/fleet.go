package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

type FleetMixin struct{}

func (FleetMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("num_caravel").
			NonNegative().
			Default(0),
		field.Int64("num_light_fighter").
			NonNegative().
			Default(0),
		field.Int64("num_corvette").
			NonNegative().
			Default(0),
		field.Int64("num_frigate").
			NonNegative().
			Default(0),
		field.Int64("num_probe").
			NonNegative().
			Default(0),
		field.Int64("num_small_cargo").
			NonNegative().
			Default(0),
		field.Int64("num_medium_cargo").
			NonNegative().
			Default(0),
		field.Int64("num_colonization_ark").
			NonNegative().
			Default(0),
	}
}
