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

	// Add game time as a component to the world
	timeEntity := ecs.NewEntity(0)
	timeEntity.AddComponent("GameTime", &components.GameTime{
		CurrentDate:       initialDate,
		LastUpdated:       initialDate,
		IsPaused:          false,
		SpeedMultiplier:   1.0,
		RentCollectionDay: 1,
	})
	world.AddEntity(timeEntity)

	// Create player entity
	playerEntity := ecs.NewEntity(1)
	playerEntity.AddComponent("Player", &components.Player{ID: 1, Funds: 10000})
	world.AddEntity(playerEntity)

	// List of all neighborhoods
	allNeighborhoods := []*components.Neighborhood{
		neighborhoods.GetDowntownDistrictNeighborhood(),
		neighborhoods.GetHistoricHeightsNeighborhood(),
		neighborhoods.GetTechValleyNeighborhood(),
		neighborhoods.GetCedarGroveNeighborhood(),
		neighborhoods.GetWillowFlatsNeighborhood(),
	}

	allProperties := slices.Concat(neighborhoods.GetDowntownProperties(),
		neighborhoods.GetHistoricHeightsProperties(),
		neighborhoods.GetTechValleyProperties(),
		neighborhoods.GetCedarGroveProperties(),
		neighborhoods.GetWillowFlatsProperties())

	// Add properties to the ECS world
	for _, property := range allProperties {
		addPropertyToWorld(world, &property)
	}

	// Initialize and add systems
	neighborhoodValueSystem := &systems.NeighborhoodValueSystem{
		Neighborhoods: getNeighborhoodMap(allNeighborhoods),
	}

	world.AddSystem(&systems.IncomeSystem{})
	world.AddSystem(neighborhoodValueSystem)
	world.AddSystem(&systems.PropertyManagementSystem{})
	world.AddSystem(&systems.TimeSystem{})

	return world
}

func addPropertyToWorld(world *ecs.World, prop *components.Property) {
	entity := ecs.NewEntity(prop.ID)
	entity.AddComponent("Property", prop)
	world.AddEntity(entity)
}

// getNeighborhoodMap creates a map of Neighborhood ID to Neighborhood pointer
func getNeighborhoodMap(neighborhoods []*components.Neighborhood) map[int]*components.Neighborhood {
	neighborhoodMap := make(map[int]*components.Neighborhood)
	for _, neighborhood := range neighborhoods {
		neighborhoodMap[neighborhood.ID] = neighborhood
	}
	return neighborhoodMap
}
