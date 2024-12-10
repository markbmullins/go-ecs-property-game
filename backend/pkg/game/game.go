package game

import (
	"slices"
	"time"

	"github.com/markbmullins/city-developer/pkg/ecs"
	"github.com/markbmullins/city-developer/pkg/entities"
	"github.com/markbmullins/city-developer/pkg/neighborhoods"
	"github.com/markbmullins/city-developer/pkg/systems"
)

func InitializeGame() *ecs.World {
	world := ecs.NewWorld()
	initialDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

	timeEntity := entities.CreateGameTime(initialDate, 1)
	world.AddSpecificEntity(0, timeEntity)

	playerEntity := entities.CreatePlayer("Mark", 100000000)
	world.AddEntity(playerEntity)

	initializeProperties(world)
	initializeSystems(world)

	return world
}

// var allNeighborhoods = []*components.Neighborhood{
// 	neighborhoods.GetDowntownDistrictNeighborhood(),
// 	neighborhoods.GetHistoricHeightsNeighborhood(),
// 	neighborhoods.GetTechValleyNeighborhood(),
// 	neighborhoods.GetCedarGroveNeighborhood(),
// 	neighborhoods.GetWillowFlatsNeighborhood(),
// }

// var allProperties = slices.Concat(
// 	neighborhoods.GetDowntownProperties(),
// 	neighborhoods.GetHistoricHeightsProperties(),
// 	neighborhoods.GetTechValleyProperties(),
// 	neighborhoods.GetCedarGroveProperties(),
// 	neighborhoods.GetWillowFlatsProperties(),
// )

func initializeProperties(world *ecs.World) {
	neighborhoods.InitializeCedarGroveUpgrades()
	var cedarGroveProperties = neighborhoods.GetCedarGroveProperties()

	var allProperties = slices.Concat(cedarGroveProperties)

	for _, property := range allProperties {
		world.AddEntity(property)
	}
}

func initializeSystems(world *ecs.World) {
	world.AddSystem(&systems.RentCollectionSystem{})
	world.AddSystem(&systems.PropertyManagementSystem{})
	world.AddSystem(&systems.TimeSystem{})
}
