package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"

	"time"
)

type TimeMixin struct{}

func (TimeMixin) Fields() []ent.Field {
    return []ent.Field{
        field.Time("created_at").
            Immutable().
            Default(time.Now),
        field.Time("updated_at").
            Default(time.Now).
            UpdateDefault(time.Now),
    }
}

type ResourcesMixin struct{}

func (ResourcesMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("metal").
			NonNegative().
			Default(0),
		field.Int64("hydrogen").
			NonNegative().
			Default(0),
		field.Int64("population").
			NonNegative().
			Default(0),
	}
}

type ProductionMixin struct{}

func (ProductionMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("metal_last_update").
			Default(time.Now),
		field.Int("metal_prod_rate").
			Default(0),
		field.Time("hydrogen_last_update").
			Default(time.Now),
		field.Int("hydrogen_prod_rate").
			Default(0),
		field.Time("population_last_update").
			Default(time.Now),
		field.Int("population_prod_rate").
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
	}
}

type BuildingsMixin struct{}

func (BuildingsMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int("metal_prod_level").
			NonNegative().
			Default(0),
		field.Int("hydrogen_prod_level").
			NonNegative().
			Default(0),
		field.Int("energy_prod_level").
			NonNegative().
			Default(0),
		field.Int("population_prod_level").
			NonNegative().
			Default(0),
	}
}
