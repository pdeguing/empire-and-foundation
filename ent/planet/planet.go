// Code generated by entc, DO NOT EDIT.

package planet

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the planet type in the database.
	Label = "planet"
	// FieldID holds the string denoting the id field in the database.
	FieldID                     = "id"                       // FieldCreatedAt holds the string denoting the created_at vertex property in the database.
	FieldCreatedAt              = "created_at"               // FieldUpdatedAt holds the string denoting the updated_at vertex property in the database.
	FieldUpdatedAt              = "updated_at"               // FieldMetal holds the string denoting the metal vertex property in the database.
	FieldMetal                  = "metal"                    // FieldMetalProdLevel holds the string denoting the metal_prod_level vertex property in the database.
	FieldMetalProdLevel         = "metal_prod_level"         // FieldMetalStorageLevel holds the string denoting the metal_storage_level vertex property in the database.
	FieldMetalStorageLevel      = "metal_storage_level"      // FieldHydrogen holds the string denoting the hydrogen vertex property in the database.
	FieldHydrogen               = "hydrogen"                 // FieldHydrogenProdLevel holds the string denoting the hydrogen_prod_level vertex property in the database.
	FieldHydrogenProdLevel      = "hydrogen_prod_level"      // FieldHydrogenStorageLevel holds the string denoting the hydrogen_storage_level vertex property in the database.
	FieldHydrogenStorageLevel   = "hydrogen_storage_level"   // FieldSilica holds the string denoting the silica vertex property in the database.
	FieldSilica                 = "silica"                   // FieldSilicaProdLevel holds the string denoting the silica_prod_level vertex property in the database.
	FieldSilicaProdLevel        = "silica_prod_level"        // FieldSilicaStorageLevel holds the string denoting the silica_storage_level vertex property in the database.
	FieldSilicaStorageLevel     = "silica_storage_level"     // FieldPopulation holds the string denoting the population vertex property in the database.
	FieldPopulation             = "population"               // FieldPopulationProdLevel holds the string denoting the population_prod_level vertex property in the database.
	FieldPopulationProdLevel    = "population_prod_level"    // FieldPopulationStorageLevel holds the string denoting the population_storage_level vertex property in the database.
	FieldPopulationStorageLevel = "population_storage_level" // FieldSolarProdLevel holds the string denoting the solar_prod_level vertex property in the database.
	FieldSolarProdLevel         = "solar_prod_level"         // FieldShipFactoryLevel holds the string denoting the ship_factory_level vertex property in the database.
	FieldShipFactoryLevel       = "ship_factory_level"       // FieldResearchCenterLevel holds the string denoting the research_center_level vertex property in the database.
	FieldResearchCenterLevel    = "research_center_level"    // FieldRegionCode holds the string denoting the region_code vertex property in the database.
	FieldRegionCode             = "region_code"              // FieldSystemCode holds the string denoting the system_code vertex property in the database.
	FieldSystemCode             = "system_code"              // FieldOrbitCode holds the string denoting the orbit_code vertex property in the database.
	FieldOrbitCode              = "orbit_code"               // FieldSuborbitCode holds the string denoting the suborbit_code vertex property in the database.
	FieldSuborbitCode           = "suborbit_code"            // FieldPositionCode holds the string denoting the position_code vertex property in the database.
	FieldPositionCode           = "position_code"            // FieldName holds the string denoting the name vertex property in the database.
	FieldName                   = "name"                     // FieldPlanetType holds the string denoting the planet_type vertex property in the database.
	FieldPlanetType             = "planet_type"              // FieldPlanetSkin holds the string denoting the planet_skin vertex property in the database.
	FieldPlanetSkin             = "planet_skin"              // FieldLastResourceUpdate holds the string denoting the last_resource_update vertex property in the database.
	FieldLastResourceUpdate     = "last_resource_update"

	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeTimers holds the string denoting the timers edge name in mutations.
	EdgeTimers = "timers"

	// Table holds the table name of the planet in the database.
	Table = "planets"
	// OwnerTable is the table the holds the owner relation/edge.
	OwnerTable = "planets"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_planets"
	// TimersTable is the table the holds the timers relation/edge.
	TimersTable = "timers"
	// TimersInverseTable is the table name for the Timer entity.
	// It exists in this package in order to avoid circular dependency with the "timer" package.
	TimersInverseTable = "timers"
	// TimersColumn is the table column denoting the timers relation/edge.
	TimersColumn = "planet_timers"
)

// Columns holds all SQL columns for planet fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldMetal,
	FieldMetalProdLevel,
	FieldMetalStorageLevel,
	FieldHydrogen,
	FieldHydrogenProdLevel,
	FieldHydrogenStorageLevel,
	FieldSilica,
	FieldSilicaProdLevel,
	FieldSilicaStorageLevel,
	FieldPopulation,
	FieldPopulationProdLevel,
	FieldPopulationStorageLevel,
	FieldSolarProdLevel,
	FieldShipFactoryLevel,
	FieldResearchCenterLevel,
	FieldRegionCode,
	FieldSystemCode,
	FieldOrbitCode,
	FieldSuborbitCode,
	FieldPositionCode,
	FieldName,
	FieldPlanetType,
	FieldPlanetSkin,
	FieldLastResourceUpdate,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Planet type.
var ForeignKeys = []string{
	"user_planets",
}

var (
	// DefaultCreatedAt holds the default value on creation for the created_at field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the updated_at field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultMetal holds the default value on creation for the metal field.
	DefaultMetal int64
	// MetalValidator is a validator for the "metal" field. It is called by the builders before save.
	MetalValidator func(int64) error
	// DefaultMetalProdLevel holds the default value on creation for the metal_prod_level field.
	DefaultMetalProdLevel int
	// MetalProdLevelValidator is a validator for the "metal_prod_level" field. It is called by the builders before save.
	MetalProdLevelValidator func(int) error
	// DefaultMetalStorageLevel holds the default value on creation for the metal_storage_level field.
	DefaultMetalStorageLevel int
	// MetalStorageLevelValidator is a validator for the "metal_storage_level" field. It is called by the builders before save.
	MetalStorageLevelValidator func(int) error
	// DefaultHydrogen holds the default value on creation for the hydrogen field.
	DefaultHydrogen int64
	// HydrogenValidator is a validator for the "hydrogen" field. It is called by the builders before save.
	HydrogenValidator func(int64) error
	// DefaultHydrogenProdLevel holds the default value on creation for the hydrogen_prod_level field.
	DefaultHydrogenProdLevel int
	// HydrogenProdLevelValidator is a validator for the "hydrogen_prod_level" field. It is called by the builders before save.
	HydrogenProdLevelValidator func(int) error
	// DefaultHydrogenStorageLevel holds the default value on creation for the hydrogen_storage_level field.
	DefaultHydrogenStorageLevel int
	// HydrogenStorageLevelValidator is a validator for the "hydrogen_storage_level" field. It is called by the builders before save.
	HydrogenStorageLevelValidator func(int) error
	// DefaultSilica holds the default value on creation for the silica field.
	DefaultSilica int64
	// SilicaValidator is a validator for the "silica" field. It is called by the builders before save.
	SilicaValidator func(int64) error
	// DefaultSilicaProdLevel holds the default value on creation for the silica_prod_level field.
	DefaultSilicaProdLevel int
	// SilicaProdLevelValidator is a validator for the "silica_prod_level" field. It is called by the builders before save.
	SilicaProdLevelValidator func(int) error
	// DefaultSilicaStorageLevel holds the default value on creation for the silica_storage_level field.
	DefaultSilicaStorageLevel int
	// SilicaStorageLevelValidator is a validator for the "silica_storage_level" field. It is called by the builders before save.
	SilicaStorageLevelValidator func(int) error
	// DefaultPopulation holds the default value on creation for the population field.
	DefaultPopulation int64
	// PopulationValidator is a validator for the "population" field. It is called by the builders before save.
	PopulationValidator func(int64) error
	// DefaultPopulationProdLevel holds the default value on creation for the population_prod_level field.
	DefaultPopulationProdLevel int
	// PopulationProdLevelValidator is a validator for the "population_prod_level" field. It is called by the builders before save.
	PopulationProdLevelValidator func(int) error
	// DefaultPopulationStorageLevel holds the default value on creation for the population_storage_level field.
	DefaultPopulationStorageLevel int
	// PopulationStorageLevelValidator is a validator for the "population_storage_level" field. It is called by the builders before save.
	PopulationStorageLevelValidator func(int) error
	// DefaultSolarProdLevel holds the default value on creation for the solar_prod_level field.
	DefaultSolarProdLevel int
	// SolarProdLevelValidator is a validator for the "solar_prod_level" field. It is called by the builders before save.
	SolarProdLevelValidator func(int) error
	// DefaultShipFactoryLevel holds the default value on creation for the ship_factory_level field.
	DefaultShipFactoryLevel int
	// ShipFactoryLevelValidator is a validator for the "ship_factory_level" field. It is called by the builders before save.
	ShipFactoryLevelValidator func(int) error
	// DefaultResearchCenterLevel holds the default value on creation for the research_center_level field.
	DefaultResearchCenterLevel int
	// ResearchCenterLevelValidator is a validator for the "research_center_level" field. It is called by the builders before save.
	ResearchCenterLevelValidator func(int) error
	// RegionCodeValidator is a validator for the "region_code" field. It is called by the builders before save.
	RegionCodeValidator func(int) error
	// SystemCodeValidator is a validator for the "system_code" field. It is called by the builders before save.
	SystemCodeValidator func(int) error
	// OrbitCodeValidator is a validator for the "orbit_code" field. It is called by the builders before save.
	OrbitCodeValidator func(int) error
	// SuborbitCodeValidator is a validator for the "suborbit_code" field. It is called by the builders before save.
	SuborbitCodeValidator func(int) error
	// PositionCodeValidator is a validator for the "position_code" field. It is called by the builders before save.
	PositionCodeValidator func(int) error
	// DefaultName holds the default value on creation for the name field.
	DefaultName string
	// DefaultLastResourceUpdate holds the default value on creation for the last_resource_update field.
	DefaultLastResourceUpdate func() time.Time
)

// PlanetType defines the type for the planet_type enum field.
type PlanetType string

// PlanetType values.
const (
	PlanetTypeHabitable PlanetType = "habitable"
	PlanetTypeMineral   PlanetType = "mineral"
	PlanetTypeGasGiant  PlanetType = "gas_giant"
	PlanetTypeIceGiant  PlanetType = "ice_giant"
)

func (s PlanetType) String() string {
	return string(s)
}

// PlanetTypeValidator is a validator for the "pt" field enum values. It is called by the builders before save.
func PlanetTypeValidator(pt PlanetType) error {
	switch pt {
	case PlanetTypeHabitable, PlanetTypeMineral, PlanetTypeGasGiant, PlanetTypeIceGiant:
		return nil
	default:
		return fmt.Errorf("planet: invalid enum value for planet_type field: %q", pt)
	}
}
