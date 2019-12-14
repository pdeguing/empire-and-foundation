// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/pdeguing/empire-and-foundation/ent/planet"
)

// PlanetCreate is the builder for creating a Planet entity.
type PlanetCreate struct {
	config
	created_at             *time.Time
	updated_at             *time.Time
	metal                  *int64
	hydrogen               *int64
	population             *int64
	metal_last_update      *time.Time
	metal_prod_rate        *int
	hydrogen_last_update   *time.Time
	hydrogen_prod_rate     *int
	population_last_update *time.Time
	population_prod_rate   *int
	energy_cons            *int64
	energy_prod            *int64
	metal_prod_level       *int
	hydrogen_prod_level    *int
	energy_prod_level      *int
	population_prod_level  *int
	name                   *string
	owner                  map[int]struct{}
}

// SetCreatedAt sets the created_at field.
func (pc *PlanetCreate) SetCreatedAt(t time.Time) *PlanetCreate {
	pc.created_at = &t
	return pc
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (pc *PlanetCreate) SetNillableCreatedAt(t *time.Time) *PlanetCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUpdatedAt sets the updated_at field.
func (pc *PlanetCreate) SetUpdatedAt(t time.Time) *PlanetCreate {
	pc.updated_at = &t
	return pc
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (pc *PlanetCreate) SetNillableUpdatedAt(t *time.Time) *PlanetCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetMetal sets the metal field.
func (pc *PlanetCreate) SetMetal(i int64) *PlanetCreate {
	pc.metal = &i
	return pc
}

// SetNillableMetal sets the metal field if the given value is not nil.
func (pc *PlanetCreate) SetNillableMetal(i *int64) *PlanetCreate {
	if i != nil {
		pc.SetMetal(*i)
	}
	return pc
}

// SetHydrogen sets the hydrogen field.
func (pc *PlanetCreate) SetHydrogen(i int64) *PlanetCreate {
	pc.hydrogen = &i
	return pc
}

// SetNillableHydrogen sets the hydrogen field if the given value is not nil.
func (pc *PlanetCreate) SetNillableHydrogen(i *int64) *PlanetCreate {
	if i != nil {
		pc.SetHydrogen(*i)
	}
	return pc
}

// SetPopulation sets the population field.
func (pc *PlanetCreate) SetPopulation(i int64) *PlanetCreate {
	pc.population = &i
	return pc
}

// SetNillablePopulation sets the population field if the given value is not nil.
func (pc *PlanetCreate) SetNillablePopulation(i *int64) *PlanetCreate {
	if i != nil {
		pc.SetPopulation(*i)
	}
	return pc
}

// SetMetalLastUpdate sets the metal_last_update field.
func (pc *PlanetCreate) SetMetalLastUpdate(t time.Time) *PlanetCreate {
	pc.metal_last_update = &t
	return pc
}

// SetNillableMetalLastUpdate sets the metal_last_update field if the given value is not nil.
func (pc *PlanetCreate) SetNillableMetalLastUpdate(t *time.Time) *PlanetCreate {
	if t != nil {
		pc.SetMetalLastUpdate(*t)
	}
	return pc
}

// SetMetalProdRate sets the metal_prod_rate field.
func (pc *PlanetCreate) SetMetalProdRate(i int) *PlanetCreate {
	pc.metal_prod_rate = &i
	return pc
}

// SetNillableMetalProdRate sets the metal_prod_rate field if the given value is not nil.
func (pc *PlanetCreate) SetNillableMetalProdRate(i *int) *PlanetCreate {
	if i != nil {
		pc.SetMetalProdRate(*i)
	}
	return pc
}

// SetHydrogenLastUpdate sets the hydrogen_last_update field.
func (pc *PlanetCreate) SetHydrogenLastUpdate(t time.Time) *PlanetCreate {
	pc.hydrogen_last_update = &t
	return pc
}

// SetNillableHydrogenLastUpdate sets the hydrogen_last_update field if the given value is not nil.
func (pc *PlanetCreate) SetNillableHydrogenLastUpdate(t *time.Time) *PlanetCreate {
	if t != nil {
		pc.SetHydrogenLastUpdate(*t)
	}
	return pc
}

// SetHydrogenProdRate sets the hydrogen_prod_rate field.
func (pc *PlanetCreate) SetHydrogenProdRate(i int) *PlanetCreate {
	pc.hydrogen_prod_rate = &i
	return pc
}

// SetNillableHydrogenProdRate sets the hydrogen_prod_rate field if the given value is not nil.
func (pc *PlanetCreate) SetNillableHydrogenProdRate(i *int) *PlanetCreate {
	if i != nil {
		pc.SetHydrogenProdRate(*i)
	}
	return pc
}

// SetPopulationLastUpdate sets the population_last_update field.
func (pc *PlanetCreate) SetPopulationLastUpdate(t time.Time) *PlanetCreate {
	pc.population_last_update = &t
	return pc
}

// SetNillablePopulationLastUpdate sets the population_last_update field if the given value is not nil.
func (pc *PlanetCreate) SetNillablePopulationLastUpdate(t *time.Time) *PlanetCreate {
	if t != nil {
		pc.SetPopulationLastUpdate(*t)
	}
	return pc
}

// SetPopulationProdRate sets the population_prod_rate field.
func (pc *PlanetCreate) SetPopulationProdRate(i int) *PlanetCreate {
	pc.population_prod_rate = &i
	return pc
}

// SetNillablePopulationProdRate sets the population_prod_rate field if the given value is not nil.
func (pc *PlanetCreate) SetNillablePopulationProdRate(i *int) *PlanetCreate {
	if i != nil {
		pc.SetPopulationProdRate(*i)
	}
	return pc
}

// SetEnergyCons sets the energy_cons field.
func (pc *PlanetCreate) SetEnergyCons(i int64) *PlanetCreate {
	pc.energy_cons = &i
	return pc
}

// SetNillableEnergyCons sets the energy_cons field if the given value is not nil.
func (pc *PlanetCreate) SetNillableEnergyCons(i *int64) *PlanetCreate {
	if i != nil {
		pc.SetEnergyCons(*i)
	}
	return pc
}

// SetEnergyProd sets the energy_prod field.
func (pc *PlanetCreate) SetEnergyProd(i int64) *PlanetCreate {
	pc.energy_prod = &i
	return pc
}

// SetNillableEnergyProd sets the energy_prod field if the given value is not nil.
func (pc *PlanetCreate) SetNillableEnergyProd(i *int64) *PlanetCreate {
	if i != nil {
		pc.SetEnergyProd(*i)
	}
	return pc
}

// SetMetalProdLevel sets the metal_prod_level field.
func (pc *PlanetCreate) SetMetalProdLevel(i int) *PlanetCreate {
	pc.metal_prod_level = &i
	return pc
}

// SetNillableMetalProdLevel sets the metal_prod_level field if the given value is not nil.
func (pc *PlanetCreate) SetNillableMetalProdLevel(i *int) *PlanetCreate {
	if i != nil {
		pc.SetMetalProdLevel(*i)
	}
	return pc
}

// SetHydrogenProdLevel sets the hydrogen_prod_level field.
func (pc *PlanetCreate) SetHydrogenProdLevel(i int) *PlanetCreate {
	pc.hydrogen_prod_level = &i
	return pc
}

// SetNillableHydrogenProdLevel sets the hydrogen_prod_level field if the given value is not nil.
func (pc *PlanetCreate) SetNillableHydrogenProdLevel(i *int) *PlanetCreate {
	if i != nil {
		pc.SetHydrogenProdLevel(*i)
	}
	return pc
}

// SetEnergyProdLevel sets the energy_prod_level field.
func (pc *PlanetCreate) SetEnergyProdLevel(i int) *PlanetCreate {
	pc.energy_prod_level = &i
	return pc
}

// SetNillableEnergyProdLevel sets the energy_prod_level field if the given value is not nil.
func (pc *PlanetCreate) SetNillableEnergyProdLevel(i *int) *PlanetCreate {
	if i != nil {
		pc.SetEnergyProdLevel(*i)
	}
	return pc
}

// SetPopulationProdLevel sets the population_prod_level field.
func (pc *PlanetCreate) SetPopulationProdLevel(i int) *PlanetCreate {
	pc.population_prod_level = &i
	return pc
}

// SetNillablePopulationProdLevel sets the population_prod_level field if the given value is not nil.
func (pc *PlanetCreate) SetNillablePopulationProdLevel(i *int) *PlanetCreate {
	if i != nil {
		pc.SetPopulationProdLevel(*i)
	}
	return pc
}

// SetName sets the name field.
func (pc *PlanetCreate) SetName(s string) *PlanetCreate {
	pc.name = &s
	return pc
}

// SetNillableName sets the name field if the given value is not nil.
func (pc *PlanetCreate) SetNillableName(s *string) *PlanetCreate {
	if s != nil {
		pc.SetName(*s)
	}
	return pc
}

// SetOwnerID sets the owner edge to User by id.
func (pc *PlanetCreate) SetOwnerID(id int) *PlanetCreate {
	if pc.owner == nil {
		pc.owner = make(map[int]struct{})
	}
	pc.owner[id] = struct{}{}
	return pc
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (pc *PlanetCreate) SetNillableOwnerID(id *int) *PlanetCreate {
	if id != nil {
		pc = pc.SetOwnerID(*id)
	}
	return pc
}

// SetOwner sets the owner edge to User.
func (pc *PlanetCreate) SetOwner(u *User) *PlanetCreate {
	return pc.SetOwnerID(u.ID)
}

// Save creates the Planet in the database.
func (pc *PlanetCreate) Save(ctx context.Context) (*Planet, error) {
	if pc.created_at == nil {
		v := planet.DefaultCreatedAt()
		pc.created_at = &v
	}
	if pc.updated_at == nil {
		v := planet.DefaultUpdatedAt()
		pc.updated_at = &v
	}
	if pc.metal == nil {
		v := planet.DefaultMetal
		pc.metal = &v
	}
	if err := planet.MetalValidator(*pc.metal); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"metal\": %v", err)
	}
	if pc.hydrogen == nil {
		v := planet.DefaultHydrogen
		pc.hydrogen = &v
	}
	if err := planet.HydrogenValidator(*pc.hydrogen); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"hydrogen\": %v", err)
	}
	if pc.population == nil {
		v := planet.DefaultPopulation
		pc.population = &v
	}
	if err := planet.PopulationValidator(*pc.population); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"population\": %v", err)
	}
	if pc.metal_last_update == nil {
		v := planet.DefaultMetalLastUpdate()
		pc.metal_last_update = &v
	}
	if pc.metal_prod_rate == nil {
		v := planet.DefaultMetalProdRate
		pc.metal_prod_rate = &v
	}
	if pc.hydrogen_last_update == nil {
		v := planet.DefaultHydrogenLastUpdate()
		pc.hydrogen_last_update = &v
	}
	if pc.hydrogen_prod_rate == nil {
		v := planet.DefaultHydrogenProdRate
		pc.hydrogen_prod_rate = &v
	}
	if pc.population_last_update == nil {
		v := planet.DefaultPopulationLastUpdate()
		pc.population_last_update = &v
	}
	if pc.population_prod_rate == nil {
		v := planet.DefaultPopulationProdRate
		pc.population_prod_rate = &v
	}
	if pc.energy_cons == nil {
		v := planet.DefaultEnergyCons
		pc.energy_cons = &v
	}
	if err := planet.EnergyConsValidator(*pc.energy_cons); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"energy_cons\": %v", err)
	}
	if pc.energy_prod == nil {
		v := planet.DefaultEnergyProd
		pc.energy_prod = &v
	}
	if err := planet.EnergyProdValidator(*pc.energy_prod); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"energy_prod\": %v", err)
	}
	if pc.metal_prod_level == nil {
		v := planet.DefaultMetalProdLevel
		pc.metal_prod_level = &v
	}
	if err := planet.MetalProdLevelValidator(*pc.metal_prod_level); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"metal_prod_level\": %v", err)
	}
	if pc.hydrogen_prod_level == nil {
		v := planet.DefaultHydrogenProdLevel
		pc.hydrogen_prod_level = &v
	}
	if err := planet.HydrogenProdLevelValidator(*pc.hydrogen_prod_level); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"hydrogen_prod_level\": %v", err)
	}
	if pc.energy_prod_level == nil {
		v := planet.DefaultEnergyProdLevel
		pc.energy_prod_level = &v
	}
	if err := planet.EnergyProdLevelValidator(*pc.energy_prod_level); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"energy_prod_level\": %v", err)
	}
	if pc.population_prod_level == nil {
		v := planet.DefaultPopulationProdLevel
		pc.population_prod_level = &v
	}
	if err := planet.PopulationProdLevelValidator(*pc.population_prod_level); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"population_prod_level\": %v", err)
	}
	if pc.name == nil {
		v := planet.DefaultName
		pc.name = &v
	}
	if len(pc.owner) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"owner\"")
	}
	return pc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PlanetCreate) SaveX(ctx context.Context) *Planet {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pc *PlanetCreate) sqlSave(ctx context.Context) (*Planet, error) {
	var (
		res     sql.Result
		builder = sql.Dialect(pc.driver.Dialect())
		pl      = &Planet{config: pc.config}
	)
	tx, err := pc.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	insert := builder.Insert(planet.Table).Default()
	if value := pc.created_at; value != nil {
		insert.Set(planet.FieldCreatedAt, *value)
		pl.CreatedAt = *value
	}
	if value := pc.updated_at; value != nil {
		insert.Set(planet.FieldUpdatedAt, *value)
		pl.UpdatedAt = *value
	}
	if value := pc.metal; value != nil {
		insert.Set(planet.FieldMetal, *value)
		pl.Metal = *value
	}
	if value := pc.hydrogen; value != nil {
		insert.Set(planet.FieldHydrogen, *value)
		pl.Hydrogen = *value
	}
	if value := pc.population; value != nil {
		insert.Set(planet.FieldPopulation, *value)
		pl.Population = *value
	}
	if value := pc.metal_last_update; value != nil {
		insert.Set(planet.FieldMetalLastUpdate, *value)
		pl.MetalLastUpdate = *value
	}
	if value := pc.metal_prod_rate; value != nil {
		insert.Set(planet.FieldMetalProdRate, *value)
		pl.MetalProdRate = *value
	}
	if value := pc.hydrogen_last_update; value != nil {
		insert.Set(planet.FieldHydrogenLastUpdate, *value)
		pl.HydrogenLastUpdate = *value
	}
	if value := pc.hydrogen_prod_rate; value != nil {
		insert.Set(planet.FieldHydrogenProdRate, *value)
		pl.HydrogenProdRate = *value
	}
	if value := pc.population_last_update; value != nil {
		insert.Set(planet.FieldPopulationLastUpdate, *value)
		pl.PopulationLastUpdate = *value
	}
	if value := pc.population_prod_rate; value != nil {
		insert.Set(planet.FieldPopulationProdRate, *value)
		pl.PopulationProdRate = *value
	}
	if value := pc.energy_cons; value != nil {
		insert.Set(planet.FieldEnergyCons, *value)
		pl.EnergyCons = *value
	}
	if value := pc.energy_prod; value != nil {
		insert.Set(planet.FieldEnergyProd, *value)
		pl.EnergyProd = *value
	}
	if value := pc.metal_prod_level; value != nil {
		insert.Set(planet.FieldMetalProdLevel, *value)
		pl.MetalProdLevel = *value
	}
	if value := pc.hydrogen_prod_level; value != nil {
		insert.Set(planet.FieldHydrogenProdLevel, *value)
		pl.HydrogenProdLevel = *value
	}
	if value := pc.energy_prod_level; value != nil {
		insert.Set(planet.FieldEnergyProdLevel, *value)
		pl.EnergyProdLevel = *value
	}
	if value := pc.population_prod_level; value != nil {
		insert.Set(planet.FieldPopulationProdLevel, *value)
		pl.PopulationProdLevel = *value
	}
	if value := pc.name; value != nil {
		insert.Set(planet.FieldName, *value)
		pl.Name = *value
	}

	id, err := insertLastID(ctx, tx, insert.Returning(planet.FieldID))
	if err != nil {
		return nil, rollback(tx, err)
	}
	pl.ID = int(id)
	if len(pc.owner) > 0 {
		for eid := range pc.owner {
			query, args := builder.Update(planet.OwnerTable).
				Set(planet.OwnerColumn, eid).
				Where(sql.EQ(planet.FieldID, id)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
		}
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return pl, nil
}
