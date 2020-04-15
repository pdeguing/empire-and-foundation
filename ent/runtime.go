// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/pdeguing/empire-and-foundation/ent/planet"
	"github.com/pdeguing/empire-and-foundation/ent/schema"
	"github.com/pdeguing/empire-and-foundation/ent/user"

	"github.com/facebookincubator/ent"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	planetMixin := schema.Planet{}.Mixin()
	planetMixinFields := [...][]ent.Field{
		planetMixin[0].Fields(),
		planetMixin[1].Fields(),
		planetMixin[2].Fields(),
		planetMixin[3].Fields(),
		planetMixin[4].Fields(),
		planetMixin[5].Fields(),
		planetMixin[6].Fields(),
	}
	planetFields := schema.Planet{}.Fields()
	_ = planetFields
	// planetDescCreatedAt is the schema descriptor for created_at field.
	planetDescCreatedAt := planetMixinFields[0][0].Descriptor()
	// planet.DefaultCreatedAt holds the default value on creation for the created_at field.
	planet.DefaultCreatedAt = planetDescCreatedAt.Default.(func() time.Time)
	// planetDescUpdatedAt is the schema descriptor for updated_at field.
	planetDescUpdatedAt := planetMixinFields[0][1].Descriptor()
	// planet.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	planet.DefaultUpdatedAt = planetDescUpdatedAt.Default.(func() time.Time)
	// planet.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	planet.UpdateDefaultUpdatedAt = planetDescUpdatedAt.UpdateDefault.(func() time.Time)
	// planetDescMetal is the schema descriptor for metal field.
	planetDescMetal := planetMixinFields[1][0].Descriptor()
	// planet.DefaultMetal holds the default value on creation for the metal field.
	planet.DefaultMetal = planetDescMetal.Default.(int64)
	// planet.MetalValidator is a validator for the "metal" field. It is called by the builders before save.
	planet.MetalValidator = planetDescMetal.Validators[0].(func(int64) error)
	// planetDescMetalProdLevel is the schema descriptor for metal_prod_level field.
	planetDescMetalProdLevel := planetMixinFields[1][1].Descriptor()
	// planet.DefaultMetalProdLevel holds the default value on creation for the metal_prod_level field.
	planet.DefaultMetalProdLevel = planetDescMetalProdLevel.Default.(int)
	// planet.MetalProdLevelValidator is a validator for the "metal_prod_level" field. It is called by the builders before save.
	planet.MetalProdLevelValidator = planetDescMetalProdLevel.Validators[0].(func(int) error)
	// planetDescMetalStorageLevel is the schema descriptor for metal_storage_level field.
	planetDescMetalStorageLevel := planetMixinFields[1][2].Descriptor()
	// planet.DefaultMetalStorageLevel holds the default value on creation for the metal_storage_level field.
	planet.DefaultMetalStorageLevel = planetDescMetalStorageLevel.Default.(int)
	// planet.MetalStorageLevelValidator is a validator for the "metal_storage_level" field. It is called by the builders before save.
	planet.MetalStorageLevelValidator = planetDescMetalStorageLevel.Validators[0].(func(int) error)
	// planetDescHydrogen is the schema descriptor for hydrogen field.
	planetDescHydrogen := planetMixinFields[2][0].Descriptor()
	// planet.DefaultHydrogen holds the default value on creation for the hydrogen field.
	planet.DefaultHydrogen = planetDescHydrogen.Default.(int64)
	// planet.HydrogenValidator is a validator for the "hydrogen" field. It is called by the builders before save.
	planet.HydrogenValidator = planetDescHydrogen.Validators[0].(func(int64) error)
	// planetDescHydrogenProdLevel is the schema descriptor for hydrogen_prod_level field.
	planetDescHydrogenProdLevel := planetMixinFields[2][1].Descriptor()
	// planet.DefaultHydrogenProdLevel holds the default value on creation for the hydrogen_prod_level field.
	planet.DefaultHydrogenProdLevel = planetDescHydrogenProdLevel.Default.(int)
	// planet.HydrogenProdLevelValidator is a validator for the "hydrogen_prod_level" field. It is called by the builders before save.
	planet.HydrogenProdLevelValidator = planetDescHydrogenProdLevel.Validators[0].(func(int) error)
	// planetDescHydrogenStorageLevel is the schema descriptor for hydrogen_storage_level field.
	planetDescHydrogenStorageLevel := planetMixinFields[2][2].Descriptor()
	// planet.DefaultHydrogenStorageLevel holds the default value on creation for the hydrogen_storage_level field.
	planet.DefaultHydrogenStorageLevel = planetDescHydrogenStorageLevel.Default.(int)
	// planet.HydrogenStorageLevelValidator is a validator for the "hydrogen_storage_level" field. It is called by the builders before save.
	planet.HydrogenStorageLevelValidator = planetDescHydrogenStorageLevel.Validators[0].(func(int) error)
	// planetDescSilica is the schema descriptor for silica field.
	planetDescSilica := planetMixinFields[3][0].Descriptor()
	// planet.DefaultSilica holds the default value on creation for the silica field.
	planet.DefaultSilica = planetDescSilica.Default.(int64)
	// planet.SilicaValidator is a validator for the "silica" field. It is called by the builders before save.
	planet.SilicaValidator = planetDescSilica.Validators[0].(func(int64) error)
	// planetDescSilicaProdLevel is the schema descriptor for silica_prod_level field.
	planetDescSilicaProdLevel := planetMixinFields[3][1].Descriptor()
	// planet.DefaultSilicaProdLevel holds the default value on creation for the silica_prod_level field.
	planet.DefaultSilicaProdLevel = planetDescSilicaProdLevel.Default.(int)
	// planet.SilicaProdLevelValidator is a validator for the "silica_prod_level" field. It is called by the builders before save.
	planet.SilicaProdLevelValidator = planetDescSilicaProdLevel.Validators[0].(func(int) error)
	// planetDescSilicaStorageLevel is the schema descriptor for silica_storage_level field.
	planetDescSilicaStorageLevel := planetMixinFields[3][2].Descriptor()
	// planet.DefaultSilicaStorageLevel holds the default value on creation for the silica_storage_level field.
	planet.DefaultSilicaStorageLevel = planetDescSilicaStorageLevel.Default.(int)
	// planet.SilicaStorageLevelValidator is a validator for the "silica_storage_level" field. It is called by the builders before save.
	planet.SilicaStorageLevelValidator = planetDescSilicaStorageLevel.Validators[0].(func(int) error)
	// planetDescPopulation is the schema descriptor for population field.
	planetDescPopulation := planetMixinFields[4][0].Descriptor()
	// planet.DefaultPopulation holds the default value on creation for the population field.
	planet.DefaultPopulation = planetDescPopulation.Default.(int64)
	// planet.PopulationValidator is a validator for the "population" field. It is called by the builders before save.
	planet.PopulationValidator = planetDescPopulation.Validators[0].(func(int64) error)
	// planetDescPopulationProdLevel is the schema descriptor for population_prod_level field.
	planetDescPopulationProdLevel := planetMixinFields[4][1].Descriptor()
	// planet.DefaultPopulationProdLevel holds the default value on creation for the population_prod_level field.
	planet.DefaultPopulationProdLevel = planetDescPopulationProdLevel.Default.(int)
	// planet.PopulationProdLevelValidator is a validator for the "population_prod_level" field. It is called by the builders before save.
	planet.PopulationProdLevelValidator = planetDescPopulationProdLevel.Validators[0].(func(int) error)
	// planetDescPopulationStorageLevel is the schema descriptor for population_storage_level field.
	planetDescPopulationStorageLevel := planetMixinFields[4][2].Descriptor()
	// planet.DefaultPopulationStorageLevel holds the default value on creation for the population_storage_level field.
	planet.DefaultPopulationStorageLevel = planetDescPopulationStorageLevel.Default.(int)
	// planet.PopulationStorageLevelValidator is a validator for the "population_storage_level" field. It is called by the builders before save.
	planet.PopulationStorageLevelValidator = planetDescPopulationStorageLevel.Validators[0].(func(int) error)
	// planetDescSolarProdLevel is the schema descriptor for solar_prod_level field.
	planetDescSolarProdLevel := planetMixinFields[5][0].Descriptor()
	// planet.DefaultSolarProdLevel holds the default value on creation for the solar_prod_level field.
	planet.DefaultSolarProdLevel = planetDescSolarProdLevel.Default.(int)
	// planet.SolarProdLevelValidator is a validator for the "solar_prod_level" field. It is called by the builders before save.
	planet.SolarProdLevelValidator = planetDescSolarProdLevel.Validators[0].(func(int) error)
	// planetDescRegionCode is the schema descriptor for region_code field.
	planetDescRegionCode := planetMixinFields[6][0].Descriptor()
	// planet.RegionCodeValidator is a validator for the "region_code" field. It is called by the builders before save.
	planet.RegionCodeValidator = planetDescRegionCode.Validators[0].(func(int) error)
	// planetDescSystemCode is the schema descriptor for system_code field.
	planetDescSystemCode := planetMixinFields[6][1].Descriptor()
	// planet.SystemCodeValidator is a validator for the "system_code" field. It is called by the builders before save.
	planet.SystemCodeValidator = planetDescSystemCode.Validators[0].(func(int) error)
	// planetDescOrbitCode is the schema descriptor for orbit_code field.
	planetDescOrbitCode := planetMixinFields[6][2].Descriptor()
	// planet.OrbitCodeValidator is a validator for the "orbit_code" field. It is called by the builders before save.
	planet.OrbitCodeValidator = planetDescOrbitCode.Validators[0].(func(int) error)
	// planetDescSuborbitCode is the schema descriptor for suborbit_code field.
	planetDescSuborbitCode := planetMixinFields[6][3].Descriptor()
	// planet.SuborbitCodeValidator is a validator for the "suborbit_code" field. It is called by the builders before save.
	planet.SuborbitCodeValidator = planetDescSuborbitCode.Validators[0].(func(int) error)
	// planetDescPositionCode is the schema descriptor for position_code field.
	planetDescPositionCode := planetMixinFields[6][4].Descriptor()
	// planet.PositionCodeValidator is a validator for the "position_code" field. It is called by the builders before save.
	planet.PositionCodeValidator = planetDescPositionCode.Validators[0].(func(int) error)
	// planetDescName is the schema descriptor for name field.
	planetDescName := planetFields[0].Descriptor()
	// planet.DefaultName holds the default value on creation for the name field.
	planet.DefaultName = planetDescName.Default.(string)
	// planetDescLastResourceUpdate is the schema descriptor for last_resource_update field.
	planetDescLastResourceUpdate := planetFields[3].Descriptor()
	// planet.DefaultLastResourceUpdate holds the default value on creation for the last_resource_update field.
	planet.DefaultLastResourceUpdate = planetDescLastResourceUpdate.Default.(func() time.Time)
	userMixin := schema.User{}.Mixin()
	userMixinFields := [...][]ent.Field{
		userMixin[0].Fields(),
	}
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields[0][0].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields[0][1].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
}