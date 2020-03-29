package data

import (
	"context"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pdeguing/empire-and-foundation/ent"
	"github.com/pdeguing/empire-and-foundation/ent/planet"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"

	"github.com/pdeguing/empire-and-foundation/ent/schema"
	"github.com/pdeguing/empire-and-foundation/ent/timer"
)

// freezeTime makes all future calls to timeNow() return the same time.
func freezeTime() {
	now := time.Now()
	timeNow = func() time.Time {
		return now
	}
}

func TestActions_ContainsAllSchemaDefinedActions(t *testing.T) {
	var actns []string
	for _, field := range (schema.Timer{}).Fields() {
		d := field.Descriptor()
		if d.Name == "action" {
			actns = d.Enums
			break
		}
	}
	for _, a := range actns {
		assert.Contains(t, actions, timer.Action(a), "Map \"actions\" does not contain the action %q", a)
	}
}

func TestIsBusy_ReturnsTrueIfTimerInGroupIsRunning(t *testing.T) {
	WithTestDatabase()
	p := NewPlanetFactory().WithTimer(timer.ActionUpgradeMetalProd, time.Hour).MustCreate()

	busy, err := IsBusy(context.Background(), p, timer.GroupBuilding)
	assert.NoError(t, err)
	assert.True(t, busy)
}

func TestIsBusy_ReturnsFalseIfNoTimerIsRunningInGroup(t *testing.T) {
	WithTestDatabase()
	p := NewPlanetFactory().MustCreate()

	busy, err := IsBusy(context.Background(), p, timer.GroupBuilding)
	assert.NoError(t, err)
	assert.False(t, busy)
}

func TestIsBusy_ReturnsFalseIfATimerIsRunningButInAnotherGroup(t *testing.T) {
	t.Skip("There are currently no other groups to make timers in.")
}

func TestIsBusy_ReturnsFalseIfTimerIsRunningOnAnotherPlanet(t *testing.T) {
	WithTestDatabase()
	p := NewPlanetFactory().MustCreate()
	NewPlanetFactory().WithTimer(timer.ActionUpgradeMetalProd, time.Hour).MustCreate()

	busy, err := IsBusy(context.Background(), p, timer.GroupBuilding)
	assert.NoError(t, err)
	assert.False(t, busy)
}

func TestGetTimer_ReturnsTimerInGroup(t *testing.T) {
	WithTestDatabase()
	freezeTime()
	p := NewPlanetFactory().WithTimer(timer.ActionUpgradeMetalProd, time.Hour).MustCreate()
	refTmr := p.QueryTimers().FirstX(context.Background())

	tmr, err := GetTimer(context.Background(), p, timer.GroupBuilding)
	assert.NoError(t, err)
	assert.NotNil(t, tmr)
	assert.Equal(t, timer.ActionUpgradeMetalProd, tmr.Action)
	assert.Equal(t, refTmr.EndTime, tmr.EndTime)
	assert.Equal(t, time.Hour, tmr.Duration())
}

func TestGetTimer_ReturnsNilIfNoTimerIsRunningInGroupForPlanet(t *testing.T) {
	WithTestDatabase()
	freezeTime()
	p := NewPlanetFactory().MustCreate()
	NewPlanetFactory().WithTimer(timer.ActionUpgradeMetalProd, time.Hour).MustCreate()

	tmr, err := GetTimer(context.Background(), p, timer.GroupBuilding)
	assert.NoError(t, err)
	assert.Nil(t, tmr)
}

func TestGetTimer_ReturnsNilIfATimerIsRunningButInAnotherGroup(t *testing.T) {
	t.Skip("There are currently no other groups to make timers in.")
}

func TestGetTimers_ReturnsAllTimers(t *testing.T) {
	WithTestDatabase()
	freezeTime()
	p := NewPlanetFactory().WithTimer(timer.ActionUpgradeMetalStorage, time.Hour).MustCreate()
	refTmr := p.QueryTimers().FirstX(context.Background())

	timers, err := GetTimers(context.Background(), p)
	assert.NoError(t, err)
	assert.Len(t, timers, 1)
	assert.NotNil(t, timers[timer.GroupBuilding])
	assert.Equal(t, timer.ActionUpgradeMetalStorage, timers[timer.GroupBuilding].Action)
	assert.Equal(t, refTmr.EndTime, timers[timer.GroupBuilding].EndTime)
	assert.Equal(t, time.Hour, timers[timer.GroupBuilding].Duration())

	// TODO: Test with more timers in other groups.
}

func TestGetTimers_ReturnsEmptyMapIfNoTimersExistsForPlanet(t *testing.T) {
	WithTestDatabase()
	freezeTime()
	p := NewPlanetFactory().MustCreate()
	NewPlanetFactory().WithTimer(timer.ActionUpgradeMetalProd, time.Hour).MustCreate()

	timers, err := GetTimers(context.Background(), p)
	assert.NoError(t, err)
	assert.Len(t, timers, 0)
}

func TestStartTimer_CanNotStartTimerIfItIsBusy(t *testing.T) {
	WithTestDatabase()
	p := NewPlanetFactory().WithTimer(timer.ActionUpgradeMetalProd, time.Hour).MustCreate()
	tx, err := Client.Tx(context.Background())
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()

	err = StartTimer(context.Background(), tx, p, timer.ActionUpgradeHydrogenStorage)
	assert.Equal(t, ErrTimerBusy, err)
}

func TestStartTimer_CanNotStartTimerIfPrerequisitesAreNotMet(t *testing.T) {
	WithTestDatabase()
	p := NewPlanetFactory().MustCreate()
	tx, err := Client.Tx(context.Background())
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()
	actions[timer.ActionTest] = action{
		Valid: func(p *ent.Planet) bool {
			return false
		},
	}

	err = StartTimer(context.Background(), tx, p, timer.ActionTest)
	assert.Equal(t, ErrActionPrerequisitesNotMet, err)
}

func TestStartTimer_PersistsResourcesToDatabase(t *testing.T) {
	WithTestDatabase()
	freezeTime()
	p := NewPlanetFactory().WithUpdatedResources().MustCreate()
	tx, err := Client.Tx(context.Background())
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()

	actions[timer.ActionTest] = action{
		Group: timer.GroupBuilding,
		Duration: func(p *ent.Planet) time.Duration {
			return time.Hour
		},
		Valid: func(p *ent.Planet) bool {
			return true
		},
		Start: func(p *ent.Planet) error {
			p.Metal = 1
			p.MetalProdLevel = 2
			p.MetalStorageLevel = 3
			p.Silica = 4
			p.SilicaProdLevel = 5
			p.SilicaStorageLevel = 6
			p.Hydrogen = 7
			p.HydrogenProdLevel = 8
			p.HydrogenStorageLevel = 9
			p.Population = 10
			p.PopulationProdLevel = 11
			p.PopulationStorageLevel = 12
			p.SolarProdLevel = 13
			return nil
		},
	}

	err = StartTimer(context.Background(), tx, p, timer.ActionTest)
	assert.NoError(t, err)

	pDatabase := tx.Planet.GetX(context.Background(), p.ID)
	assert.EqualValues(t, 1, pDatabase.Metal)
	assert.EqualValues(t, 2, pDatabase.MetalProdLevel)
	assert.EqualValues(t, 3, pDatabase.MetalStorageLevel)
	assert.EqualValues(t, 4, pDatabase.Silica)
	assert.EqualValues(t, 5, pDatabase.SilicaProdLevel)
	assert.EqualValues(t, 6, pDatabase.SilicaStorageLevel)
	assert.EqualValues(t, 7, pDatabase.Hydrogen)
	assert.EqualValues(t, 8, pDatabase.HydrogenProdLevel)
	assert.EqualValues(t, 9, pDatabase.HydrogenStorageLevel)
	assert.EqualValues(t, 10, pDatabase.Population)
	assert.EqualValues(t, 11, pDatabase.PopulationProdLevel)
	assert.EqualValues(t, 12, pDatabase.PopulationStorageLevel)
	assert.EqualValues(t, 13, pDatabase.SolarProdLevel)
	assert.WithinDuration(t, timeNow(), pDatabase.LastResourceUpdate, time.Millisecond)
}

func TestStartTimer_CreatesTimerInDatabase(t *testing.T) {
	WithTestDatabase()
	freezeTime()
	p := NewPlanetFactory().WithBeginnerResources().MustCreate()
	tx, err := Client.Tx(context.Background())
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()

	err = StartTimer(context.Background(), tx, p, timer.ActionUpgradeMetalProd)
	assert.NoError(t, err)

	tmr, err := tx.Timer.Query().Where(timer.HasPlanetWith(planet.IDEQ(p.ID))).First(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, timer.ActionUpgradeMetalProd, tmr.Action)
	assert.Equal(t, timer.GroupBuilding, tmr.Group)
	assert.WithinDuration(t, timeNow().Add(actions[timer.ActionUpgradeMetalProd].Duration(p)), tmr.EndTime, 0)
}

func TestStartTimer_ReturnsErrorWhenActionStartFunctionReturnsError(t *testing.T) {
	WithTestDatabase()
	freezeTime()
	p := NewPlanetFactory().MustCreate()
	tx, err := Client.Tx(context.Background())
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()

	expectedErr := errors.New("expected error")
	actions[timer.ActionTest] = action{
		Group: timer.GroupBuilding,
		Duration: func(p *ent.Planet) time.Duration {
			return time.Hour
		},
		Valid: func(p *ent.Planet) bool {
			return true
		},
		Start: func(p *ent.Planet) error {
			return expectedErr
		},
	}

	err = StartTimer(context.Background(), tx, p, timer.ActionTest)
	assert.True(t, errors.Is(err, expectedErr))
}

func TestCancelTimer_RemovesTimerFromDatabase(t *testing.T) {
	WithTestDatabase()
	tx, err := Client.Tx(context.Background())
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()
	p := NewPlanetFactory().Tx(tx).WithTimer(timer.ActionUpgradeMetalProd, time.Hour).MustCreate()

	err = CancelTimer(context.Background(), tx, p, timer.ActionUpgradeMetalProd)
	assert.NoError(t, err)

	n := tx.Timer.Query().Where(timer.HasPlanetWith(planet.IDEQ(p.ID))).CountX(context.Background())
	assert.Equal(t, 0, n)
}

func TestCancelTimer_ReturnsErrorIfNoTimerIsRunning(t *testing.T) {
	WithTestDatabase()
	p := NewPlanetFactory().MustCreate()
	tx, err := Client.Tx(context.Background())
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()

	err = CancelTimer(context.Background(), tx, p, timer.ActionUpgradeMetalProd)
	assert.Equal(t, ErrTimerNotRunning, err)
}

func TestCancelTimer_TimersInOtherGroupsAreLeftRunning(t *testing.T) {
	t.Skip("There are no other groups yet.")
}

func TestCancelTimer_TimersForOtherPlanetsAreLeftRunning(t *testing.T) {
	WithTestDatabase()
	tx, err := Client.Tx(context.Background())
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()
	p1 := NewPlanetFactory().Tx(tx).WithTimer(timer.ActionUpgradeMetalProd, time.Hour).MustCreate()
	p2 := NewPlanetFactory().Tx(tx).WithTimer(timer.ActionUpgradeMetalProd, time.Hour).MustCreate()

	err = CancelTimer(context.Background(), tx, p1, timer.ActionUpgradeMetalProd)
	assert.NoError(t, err)

	n := tx.Timer.Query().Where(timer.HasPlanetWith(planet.IDEQ(p2.ID))).CountX(context.Background())
	assert.Equal(t, 1, n)
}

func TestCancelTimer_ReturnsErrorWhenActionCancelFunctionReturnsError(t *testing.T) {
	WithTestDatabase()
	tx, err := Client.Tx(context.Background())
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()

	expectedErr := errors.New("expected error")
	actions[timer.ActionTest] = action{
		Group: timer.GroupBuilding,
		Cancel: func(p *ent.Planet) error {
			return expectedErr
		},
	}

	p := NewPlanetFactory().Tx(tx).WithTimer(timer.ActionTest, time.Hour).MustCreate()

	err = CancelTimer(context.Background(), tx, p, timer.ActionTest)
	assert.True(t, errors.Is(err, expectedErr))
}

func TestCancelTimer_PersistsResourcesToDatabase(t *testing.T) {
	WithTestDatabase()
	freezeTime()
	tx, err := Client.Tx(context.Background())
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()

	actions[timer.ActionTest] = action{
		Group: timer.GroupBuilding,
		Cancel: func(p *ent.Planet) error {
			p.Metal = 1
			p.MetalProdLevel = 2
			p.MetalStorageLevel = 3
			p.Silica = 4
			p.SilicaProdLevel = 5
			p.SilicaStorageLevel = 6
			p.Hydrogen = 7
			p.HydrogenProdLevel = 8
			p.HydrogenStorageLevel = 9
			p.Population = 10
			p.PopulationProdLevel = 11
			p.PopulationStorageLevel = 12
			p.SolarProdLevel = 13
			return nil
		},
	}

	p := NewPlanetFactory().Tx(tx).WithUpdatedResources().WithTimer(timer.ActionTest, time.Hour).MustCreate()

	err = CancelTimer(context.Background(), tx, p, timer.ActionTest)
	assert.NoError(t, err)

	pDatabase := tx.Planet.GetX(context.Background(), p.ID)
	assert.EqualValues(t, 1, pDatabase.Metal)
	assert.EqualValues(t, 2, pDatabase.MetalProdLevel)
	assert.EqualValues(t, 3, pDatabase.MetalStorageLevel)
	assert.EqualValues(t, 4, pDatabase.Silica)
	assert.EqualValues(t, 5, pDatabase.SilicaProdLevel)
	assert.EqualValues(t, 6, pDatabase.SilicaStorageLevel)
	assert.EqualValues(t, 7, pDatabase.Hydrogen)
	assert.EqualValues(t, 8, pDatabase.HydrogenProdLevel)
	assert.EqualValues(t, 9, pDatabase.HydrogenStorageLevel)
	assert.EqualValues(t, 10, pDatabase.Population)
	assert.EqualValues(t, 11, pDatabase.PopulationProdLevel)
	assert.EqualValues(t, 12, pDatabase.PopulationStorageLevel)
	assert.EqualValues(t, 13, pDatabase.SolarProdLevel)
	assert.WithinDuration(t, timeNow(), pDatabase.LastResourceUpdate, time.Millisecond)
}

func TestUpdateTimers_SucceedsWhenNoTimersArePresent(t *testing.T) {
	WithTestDatabase()
	p := NewPlanetFactory().MustCreate()
	tx, err := Client.Tx(context.Background())
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()

	err = UpdateTimers(context.Background(), tx, p)
	assert.NoError(t, err)
}

func TestUpdateTimers_SucceedsWhenMultipleTimersCompleteSimultaneously(t *testing.T) {
	t.Skip("Currently there is only a single group, so there can only be a single timer.")

	WithTestDatabase()
	freezeTime()
	tx, err := Client.Tx(context.Background())
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()

	i := 0
	actions[timer.ActionTest] = action{
		Group: timer.GroupBuilding,
		Complete: func(p *ent.Planet) error {
			assert.Equal(t, timeNow().Add(time.Hour), p.LastResourceUpdate)
			i++
			return nil
		},
	}

	p := NewPlanetFactory().
		WithTimer(timer.ActionTest, time.Hour).
		WithTimer(timer.ActionTest, time.Hour).
		MustCreate()

	err = UpdateTimers(context.Background(), tx, p)
	assert.NoError(t, err)
	assert.Equal(t, 2, i)
}

func TestUpdateTimers_SucceedsWhenMultipleTimersCompleteSequentially(t *testing.T) {
	t.Skip("Currently there is only a single group, so there can only be a single timer.")
}

func TestUpdateTimers_UpdatesPlanetsResourcesIncrementally(t *testing.T) {
	t.Skip("Currently there is only a single group, so there can only be a single timer.")
}

func TestUpdateTimers_ReturnsErrorWhenActionCompleteFunctionReturnsError(t *testing.T) {
	WithTestDatabase()
	freezeTime()
	tx, err := Client.Tx(context.Background())
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()

	expectedErr := errors.New("expected error")
	actions[timer.ActionTest] = action{
		Group: timer.GroupBuilding,
		Complete: func(p *ent.Planet) error {
			return expectedErr
		},
	}

	p := NewPlanetFactory().WithTimer(timer.ActionTest, -time.Minute).MustCreate()

	err = UpdateTimers(context.Background(), tx, p)
	assert.True(t, errors.Is(err, expectedErr))
}

func TestUpdateTimers_PersistsResourcesToDatabase(t *testing.T) {
	WithTestDatabase()
	freezeTime()
	tx, err := Client.Tx(context.Background())
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()

	actions[timer.ActionTest] = action{
		Group: timer.GroupBuilding,
		Complete: func(p *ent.Planet) error {
			p.Metal = 1
			p.MetalProdLevel = 2
			p.MetalStorageLevel = 3
			p.Silica = 4
			p.SilicaProdLevel = 5
			p.SilicaStorageLevel = 6
			p.Hydrogen = 7
			p.HydrogenProdLevel = 8
			p.HydrogenStorageLevel = 9
			p.Population = 10
			p.PopulationProdLevel = 11
			p.PopulationStorageLevel = 12
			p.SolarProdLevel = 13
			return nil
		},
	}

	p := NewPlanetFactory().WithTimer(timer.ActionTest, -time.Minute).MustCreate()

	err = UpdateTimers(context.Background(), tx, p)
	assert.NoError(t, err)

	pDatabase := tx.Planet.GetX(context.Background(), p.ID)
	assert.EqualValues(t, 1, pDatabase.Metal)
	assert.EqualValues(t, 2, pDatabase.MetalProdLevel)
	assert.EqualValues(t, 3, pDatabase.MetalStorageLevel)
	assert.EqualValues(t, 4, pDatabase.Silica)
	assert.EqualValues(t, 5, pDatabase.SilicaProdLevel)
	assert.EqualValues(t, 6, pDatabase.SilicaStorageLevel)
	assert.EqualValues(t, 7, pDatabase.Hydrogen)
	assert.EqualValues(t, 8, pDatabase.HydrogenProdLevel)
	assert.EqualValues(t, 9, pDatabase.HydrogenStorageLevel)
	assert.EqualValues(t, 10, pDatabase.Population)
	assert.EqualValues(t, 11, pDatabase.PopulationProdLevel)
	assert.EqualValues(t, 12, pDatabase.PopulationStorageLevel)
	assert.EqualValues(t, 13, pDatabase.SolarProdLevel)
	assert.WithinDuration(t, timeNow().Add(-time.Minute), pDatabase.LastResourceUpdate, time.Millisecond)
}

func TestUpdateTimers_DeletesTimersThanAreCompleted(t *testing.T) {
	WithTestDatabase()
	tx, err := Client.Tx(context.Background())
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()
	p := NewPlanetFactory().Tx(tx).WithTimer(timer.ActionUpgradeMetalProd, -time.Minute).MustCreate()

	err = UpdateTimers(context.Background(), tx, p)
	assert.NoError(t, err)

	n := tx.Timer.Query().Where(timer.HasPlanetWith(planet.IDEQ(p.ID))).CountX(context.Background())
	assert.Equal(t, 0, n)
}

func TestUpdateTimers_DoesntTouchOtherPlanetTimers(t *testing.T) {
	WithTestDatabase()
	tx, err := Client.Tx(context.Background())
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()
	p1 := NewPlanetFactory().Tx(tx).WithTimer(timer.ActionUpgradeMetalProd, -time.Minute).MustCreate()
	p2 := NewPlanetFactory().Tx(tx).WithTimer(timer.ActionUpgradeMetalProd, -time.Minute).MustCreate()

	err = UpdateTimers(context.Background(), tx, p1)
	assert.NoError(t, err)

	n := tx.Timer.Query().Where(timer.HasPlanetWith(planet.IDEQ(p2.ID))).CountX(context.Background())
	assert.Equal(t, 1, n)
}