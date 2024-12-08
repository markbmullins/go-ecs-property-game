package game

import (
	"slices"
	"time"

	"github.com/markbmullins/city-developer/pkg/components"
	"github.com/markbmullins/city-developer/pkg/ecs"
	"github.com/markbmullins/city-developer/pkg/neighborhoods"
	"github.com/markbmullins/city-developer/pkg/systems"
)

func InitializeGame() *ecs.World {
	world := ecs.NewWorld()
	initialDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

	initializeGameTime(world, initialDate)
	initializePlayer(world)
	initializeProperties(world)
	initializeNeighborhoods(world)
	initializeSystems(world)

	return world
}

var allNeighborhoods = []*components.Neighborhood{
	neighborhoods.GetDowntownDistrictNeighborhood(),
	neighborhoods.GetHistoricHeightsNeighborhood(),
	neighborhoods.GetTechValleyNeighborhood(),
	neighborhoods.GetCedarGroveNeighborhood(),
	neighborhoods.GetWillowFlatsNeighborhood(),
}

var allProperties = slices.Concat(
	neighborhoods.GetDowntownProperties(),
	neighborhoods.GetHistoricHeightsProperties(),
	neighborhoods.GetTechValleyProperties(),
	neighborhoods.GetCedarGroveProperties(),
	neighborhoods.GetWillowFlatsProperties(),
)

func initializeGameTime(world *ecs.World, initialDate time.Time) {
	timeEntity := ecs.NewEntity(0)
	timeEntity.AddComponent("GameTime", &components.GameTime{
		CurrentDate:       initialDate,
		LastUpdated:       initialDate,
		IsPaused:          false,
		SpeedMultiplier:   1.0,
		RentCollectionDay: 1,
	})
	world.AddEntity(timeEntity)
}

func initializePlayer(world *ecs.World) {
	playerEntity := ecs.NewEntity(1)
	playerEntity.AddComponent("Player", &components.Player{
		ID:         1,
		Funds:      10000,
		Properties: []*components.Property{},
	})
	world.AddEntity(playerEntity)
	world.AddEntity(playerEntity)
}

func initializeProperties(world *ecs.World) {
	for _, property := range allProperties {
		entity := ecs.NewEntity(property.ID)
		entity.AddComponent("Property", property)
		world.AddEntity(entity)
	}
}

func initializeNeighborhoods(world *ecs.World) {
	neighborhoodMap := make(map[int]*components.Neighborhood)
	for _, neighborhood := range allNeighborhoods {
		neighborhoodMap[neighborhood.ID] = neighborhood
		neighborhoodEntity := ecs.NewEntity(neighborhood.ID)
		neighborhoodEntity.AddComponent("Neighborhood", neighborhood)
		world.AddEntity(neighborhoodEntity)
	}
}

func initializeSystems(world *ecs.World) {
	world.AddSystem(&systems.IncomeSystem{})
	world.AddSystem(&systems.NeighborhoodValueSystem{})
	world.AddSystem(&systems.PropertyManagementSystem{})
	world.AddSystem(&systems.TimeSystem{})
}
