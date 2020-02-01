// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/pdeguing/empire-and-foundation/ent/planet"
	"github.com/pdeguing/empire-and-foundation/ent/timer"
	"github.com/pdeguing/empire-and-foundation/ent/user"
)

// PlanetCreate is the builder for creating a Planet entity.
type PlanetCreate struct {
	config
	created_at               *time.Time
	updated_at               *time.Time
	metal                    *int64
	metal_prod_level         *int
	metal_storage_level      *int
	hydrogen                 *int64
	hydrogen_prod_level      *int
	hydrogen_storage_level   *int
	silica                   *int64
	silica_prod_level        *int
	silica_storage_level     *int
	population               *int64
	population_prod_level    *int
	population_storage_level *int
	solar_prod_level         *int
	region_code              *int
	system_code              *int
	orbit_code               *int
	suborbit_code            *int
	position_code            *int
	name                     *string
	planet_type              *planet.PlanetType
	planet_skin              *string
	last_resource_update     *time.Time
	owner                    map[int]struct{}
	timers                   map[int]struct{}
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

// SetMetalStorageLevel sets the metal_storage_level field.
func (pc *PlanetCreate) SetMetalStorageLevel(i int) *PlanetCreate {
	pc.metal_storage_level = &i
	return pc
}

// SetNillableMetalStorageLevel sets the metal_storage_level field if the given value is not nil.
func (pc *PlanetCreate) SetNillableMetalStorageLevel(i *int) *PlanetCreate {
	if i != nil {
		pc.SetMetalStorageLevel(*i)
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

// SetHydrogenStorageLevel sets the hydrogen_storage_level field.
func (pc *PlanetCreate) SetHydrogenStorageLevel(i int) *PlanetCreate {
	pc.hydrogen_storage_level = &i
	return pc
}

// SetNillableHydrogenStorageLevel sets the hydrogen_storage_level field if the given value is not nil.
func (pc *PlanetCreate) SetNillableHydrogenStorageLevel(i *int) *PlanetCreate {
	if i != nil {
		pc.SetHydrogenStorageLevel(*i)
	}
	return pc
}

// SetSilica sets the silica field.
func (pc *PlanetCreate) SetSilica(i int64) *PlanetCreate {
	pc.silica = &i
	return pc
}

// SetNillableSilica sets the silica field if the given value is not nil.
func (pc *PlanetCreate) SetNillableSilica(i *int64) *PlanetCreate {
	if i != nil {
		pc.SetSilica(*i)
	}
	return pc
}

// SetSilicaProdLevel sets the silica_prod_level field.
func (pc *PlanetCreate) SetSilicaProdLevel(i int) *PlanetCreate {
	pc.silica_prod_level = &i
	return pc
}

// SetNillableSilicaProdLevel sets the silica_prod_level field if the given value is not nil.
func (pc *PlanetCreate) SetNillableSilicaProdLevel(i *int) *PlanetCreate {
	if i != nil {
		pc.SetSilicaProdLevel(*i)
	}
	return pc
}

// SetSilicaStorageLevel sets the silica_storage_level field.
func (pc *PlanetCreate) SetSilicaStorageLevel(i int) *PlanetCreate {
	pc.silica_storage_level = &i
	return pc
}

// SetNillableSilicaStorageLevel sets the silica_storage_level field if the given value is not nil.
func (pc *PlanetCreate) SetNillableSilicaStorageLevel(i *int) *PlanetCreate {
	if i != nil {
		pc.SetSilicaStorageLevel(*i)
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

// SetPopulationStorageLevel sets the population_storage_level field.
func (pc *PlanetCreate) SetPopulationStorageLevel(i int) *PlanetCreate {
	pc.population_storage_level = &i
	return pc
}

// SetNillablePopulationStorageLevel sets the population_storage_level field if the given value is not nil.
func (pc *PlanetCreate) SetNillablePopulationStorageLevel(i *int) *PlanetCreate {
	if i != nil {
		pc.SetPopulationStorageLevel(*i)
	}
	return pc
}

// SetSolarProdLevel sets the solar_prod_level field.
func (pc *PlanetCreate) SetSolarProdLevel(i int) *PlanetCreate {
	pc.solar_prod_level = &i
	return pc
}

// SetNillableSolarProdLevel sets the solar_prod_level field if the given value is not nil.
func (pc *PlanetCreate) SetNillableSolarProdLevel(i *int) *PlanetCreate {
	if i != nil {
		pc.SetSolarProdLevel(*i)
	}
	return pc
}

// SetRegionCode sets the region_code field.
func (pc *PlanetCreate) SetRegionCode(i int) *PlanetCreate {
	pc.region_code = &i
	return pc
}

// SetSystemCode sets the system_code field.
func (pc *PlanetCreate) SetSystemCode(i int) *PlanetCreate {
	pc.system_code = &i
	return pc
}

// SetOrbitCode sets the orbit_code field.
func (pc *PlanetCreate) SetOrbitCode(i int) *PlanetCreate {
	pc.orbit_code = &i
	return pc
}

// SetSuborbitCode sets the suborbit_code field.
func (pc *PlanetCreate) SetSuborbitCode(i int) *PlanetCreate {
	pc.suborbit_code = &i
	return pc
}

// SetPositionCode sets the position_code field.
func (pc *PlanetCreate) SetPositionCode(i int) *PlanetCreate {
	pc.position_code = &i
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

// SetPlanetType sets the planet_type field.
func (pc *PlanetCreate) SetPlanetType(pt planet.PlanetType) *PlanetCreate {
	pc.planet_type = &pt
	return pc
}

// SetPlanetSkin sets the planet_skin field.
func (pc *PlanetCreate) SetPlanetSkin(s string) *PlanetCreate {
	pc.planet_skin = &s
	return pc
}

// SetLastResourceUpdate sets the last_resource_update field.
func (pc *PlanetCreate) SetLastResourceUpdate(t time.Time) *PlanetCreate {
	pc.last_resource_update = &t
	return pc
}

// SetNillableLastResourceUpdate sets the last_resource_update field if the given value is not nil.
func (pc *PlanetCreate) SetNillableLastResourceUpdate(t *time.Time) *PlanetCreate {
	if t != nil {
		pc.SetLastResourceUpdate(*t)
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

// AddTimerIDs adds the timers edge to Timer by ids.
func (pc *PlanetCreate) AddTimerIDs(ids ...int) *PlanetCreate {
	if pc.timers == nil {
		pc.timers = make(map[int]struct{})
	}
	for i := range ids {
		pc.timers[ids[i]] = struct{}{}
	}
	return pc
}

// AddTimers adds the timers edges to Timer.
func (pc *PlanetCreate) AddTimers(t ...*Timer) *PlanetCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return pc.AddTimerIDs(ids...)
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
	if pc.metal_prod_level == nil {
		v := planet.DefaultMetalProdLevel
		pc.metal_prod_level = &v
	}
	if err := planet.MetalProdLevelValidator(*pc.metal_prod_level); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"metal_prod_level\": %v", err)
	}
	if pc.metal_storage_level == nil {
		v := planet.DefaultMetalStorageLevel
		pc.metal_storage_level = &v
	}
	if err := planet.MetalStorageLevelValidator(*pc.metal_storage_level); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"metal_storage_level\": %v", err)
	}
	if pc.hydrogen == nil {
		v := planet.DefaultHydrogen
		pc.hydrogen = &v
	}
	if err := planet.HydrogenValidator(*pc.hydrogen); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"hydrogen\": %v", err)
	}
	if pc.hydrogen_prod_level == nil {
		v := planet.DefaultHydrogenProdLevel
		pc.hydrogen_prod_level = &v
	}
	if err := planet.HydrogenProdLevelValidator(*pc.hydrogen_prod_level); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"hydrogen_prod_level\": %v", err)
	}
	if pc.hydrogen_storage_level == nil {
		v := planet.DefaultHydrogenStorageLevel
		pc.hydrogen_storage_level = &v
	}
	if err := planet.HydrogenStorageLevelValidator(*pc.hydrogen_storage_level); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"hydrogen_storage_level\": %v", err)
	}
	if pc.silica == nil {
		v := planet.DefaultSilica
		pc.silica = &v
	}
	if err := planet.SilicaValidator(*pc.silica); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"silica\": %v", err)
	}
	if pc.silica_prod_level == nil {
		v := planet.DefaultSilicaProdLevel
		pc.silica_prod_level = &v
	}
	if err := planet.SilicaProdLevelValidator(*pc.silica_prod_level); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"silica_prod_level\": %v", err)
	}
	if pc.silica_storage_level == nil {
		v := planet.DefaultSilicaStorageLevel
		pc.silica_storage_level = &v
	}
	if err := planet.SilicaStorageLevelValidator(*pc.silica_storage_level); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"silica_storage_level\": %v", err)
	}
	if pc.population == nil {
		v := planet.DefaultPopulation
		pc.population = &v
	}
	if err := planet.PopulationValidator(*pc.population); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"population\": %v", err)
	}
	if pc.population_prod_level == nil {
		v := planet.DefaultPopulationProdLevel
		pc.population_prod_level = &v
	}
	if err := planet.PopulationProdLevelValidator(*pc.population_prod_level); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"population_prod_level\": %v", err)
	}
	if pc.population_storage_level == nil {
		v := planet.DefaultPopulationStorageLevel
		pc.population_storage_level = &v
	}
	if err := planet.PopulationStorageLevelValidator(*pc.population_storage_level); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"population_storage_level\": %v", err)
	}
	if pc.solar_prod_level == nil {
		v := planet.DefaultSolarProdLevel
		pc.solar_prod_level = &v
	}
	if err := planet.SolarProdLevelValidator(*pc.solar_prod_level); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"solar_prod_level\": %v", err)
	}
	if pc.region_code == nil {
		return nil, errors.New("ent: missing required field \"region_code\"")
	}
	if err := planet.RegionCodeValidator(*pc.region_code); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"region_code\": %v", err)
	}
	if pc.system_code == nil {
		return nil, errors.New("ent: missing required field \"system_code\"")
	}
	if err := planet.SystemCodeValidator(*pc.system_code); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"system_code\": %v", err)
	}
	if pc.orbit_code == nil {
		return nil, errors.New("ent: missing required field \"orbit_code\"")
	}
	if err := planet.OrbitCodeValidator(*pc.orbit_code); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"orbit_code\": %v", err)
	}
	if pc.suborbit_code == nil {
		return nil, errors.New("ent: missing required field \"suborbit_code\"")
	}
	if err := planet.SuborbitCodeValidator(*pc.suborbit_code); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"suborbit_code\": %v", err)
	}
	if pc.position_code == nil {
		return nil, errors.New("ent: missing required field \"position_code\"")
	}
	if err := planet.PositionCodeValidator(*pc.position_code); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"position_code\": %v", err)
	}
	if pc.name == nil {
		v := planet.DefaultName
		pc.name = &v
	}
	if pc.planet_type == nil {
		return nil, errors.New("ent: missing required field \"planet_type\"")
	}
	if err := planet.PlanetTypeValidator(*pc.planet_type); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"planet_type\": %v", err)
	}
	if pc.planet_skin == nil {
		return nil, errors.New("ent: missing required field \"planet_skin\"")
	}
	if pc.last_resource_update == nil {
		v := planet.DefaultLastResourceUpdate()
		pc.last_resource_update = &v
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
		pl    = &Planet{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: planet.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: planet.FieldID,
			},
		}
	)
	if value := pc.created_at; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: planet.FieldCreatedAt,
		})
		pl.CreatedAt = *value
	}
	if value := pc.updated_at; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: planet.FieldUpdatedAt,
		})
		pl.UpdatedAt = *value
	}
	if value := pc.metal; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  *value,
			Column: planet.FieldMetal,
		})
		pl.Metal = *value
	}
	if value := pc.metal_prod_level; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: planet.FieldMetalProdLevel,
		})
		pl.MetalProdLevel = *value
	}
	if value := pc.metal_storage_level; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: planet.FieldMetalStorageLevel,
		})
		pl.MetalStorageLevel = *value
	}
	if value := pc.hydrogen; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  *value,
			Column: planet.FieldHydrogen,
		})
		pl.Hydrogen = *value
	}
	if value := pc.hydrogen_prod_level; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: planet.FieldHydrogenProdLevel,
		})
		pl.HydrogenProdLevel = *value
	}
	if value := pc.hydrogen_storage_level; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: planet.FieldHydrogenStorageLevel,
		})
		pl.HydrogenStorageLevel = *value
	}
	if value := pc.silica; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  *value,
			Column: planet.FieldSilica,
		})
		pl.Silica = *value
	}
	if value := pc.silica_prod_level; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: planet.FieldSilicaProdLevel,
		})
		pl.SilicaProdLevel = *value
	}
	if value := pc.silica_storage_level; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: planet.FieldSilicaStorageLevel,
		})
		pl.SilicaStorageLevel = *value
	}
	if value := pc.population; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  *value,
			Column: planet.FieldPopulation,
		})
		pl.Population = *value
	}
	if value := pc.population_prod_level; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: planet.FieldPopulationProdLevel,
		})
		pl.PopulationProdLevel = *value
	}
	if value := pc.population_storage_level; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: planet.FieldPopulationStorageLevel,
		})
		pl.PopulationStorageLevel = *value
	}
	if value := pc.solar_prod_level; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: planet.FieldSolarProdLevel,
		})
		pl.SolarProdLevel = *value
	}
	if value := pc.region_code; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: planet.FieldRegionCode,
		})
		pl.RegionCode = *value
	}
	if value := pc.system_code; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: planet.FieldSystemCode,
		})
		pl.SystemCode = *value
	}
	if value := pc.orbit_code; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: planet.FieldOrbitCode,
		})
		pl.OrbitCode = *value
	}
	if value := pc.suborbit_code; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: planet.FieldSuborbitCode,
		})
		pl.SuborbitCode = *value
	}
	if value := pc.position_code; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: planet.FieldPositionCode,
		})
		pl.PositionCode = *value
	}
	if value := pc.name; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: planet.FieldName,
		})
		pl.Name = *value
	}
	if value := pc.planet_type; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  *value,
			Column: planet.FieldPlanetType,
		})
		pl.PlanetType = *value
	}
	if value := pc.planet_skin; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: planet.FieldPlanetSkin,
		})
		pl.PlanetSkin = *value
	}
	if value := pc.last_resource_update; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: planet.FieldLastResourceUpdate,
		})
		pl.LastResourceUpdate = *value
	}
	if nodes := pc.owner; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   planet.OwnerTable,
			Columns: []string{planet.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.timers; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   planet.TimersTable,
			Columns: []string{planet.TimersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: timer.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	pl.ID = int(id)
	return pl, nil
}
