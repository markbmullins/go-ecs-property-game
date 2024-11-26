package game

import (
	"log"
	"time"

	"github.com/markbmullins/city-developer/pkg/components"
	"github.com/markbmullins/city-developer/pkg/ecs"
	"github.com/markbmullins/city-developer/pkg/models"
	"github.com/markbmullins/city-developer/pkg/neighborhoods"
	"github.com/markbmullins/city-developer/pkg/properties"
	"github.com/markbmullins/city-developer/pkg/systems"
)

func InitializeGame() *ecs.World {
	world := ecs.NewWorld()

	initialDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	gameTimeModel := &models.GameTime{
		CurrentDate:     initialDate,
		LastUpdated:     initialDate,
		IsPaused:        false,
		SpeedMultiplier: 1.0,
	}

	// Add game time as a component to the world
	timeEntity := ecs.NewEntity(0)
	timeEntity.AddComponent("GameTime", &components.GameTime{
		Time: gameTimeModel,
	})
	world.AddEntity(timeEntity)

	// Create player entity
	player := &models.Player{ID: 1, Funds: 10000}
	playerEntity := ecs.NewEntity(1)
	playerEntity.AddComponent("PlayerComponent", &components.PlayerComponent{Player: player})
	world.AddEntity(playerEntity)

	// List of all neighborhoods
	allNeighborhoods := []*models.Neighborhood{
		neighborhoods.GetDowntownDistrictNeighborhood(),
		neighborhoods.GetHistoricHeightsNeighborhood(),
		neighborhoods.GetTechValleyNeighborhood(),
		neighborhoods.GetCedarGroveNeighborhood(),
		neighborhoods.GetWillowFlatsNeighborhood(),
	}

	// Initialize the property registry
	propertyRegistry := properties.NewPropertyRegistry()

	// Add properties to the ECS world
	for _, property := range propertyRegistry.GetAllProperties() {
		if err := addPropertyToWorld(world, property); err != nil {
			log.Printf("Error adding property ID %d: %v", property.ID, err)
		}
	}

	// Initialize and add systems
	neighborhoodSystem := &systems.NeighborhoodValueSystem{
		Neighborhoods: getNeighborhoodMap(allNeighborhoods),
	}

	world.AddSystem(&systems.IncomeSystem{})
	world.AddSystem(neighborhoodSystem)
	world.AddSystem(&systems.PropertyManagementSystem{})
	world.AddSystem(&systems.TimeSystem{})

	return world
}

// addPropertyToWorld adds a property entity to the ECS world
func addPropertyToWorld(world *ecs.World, prop models.Property) {
	// Create a new entity for the property using its unique ID
	propertyEntity := ecs.NewEntity(prop.ID)

	// Add PropertyComponent
	propertyComponent := &components.PropertyComponent{
		Property: &prop, // Assuming PropertyComponent holds a pointer to Property
	}
	propertyEntity.AddComponent("PropertyComponent", propertyComponent)

	// Add other necessary components here if needed

	// Add Entity to World
	world.AddEntity(propertyEntity)
}

// Placeholder for property retrieval by ID
func getPropertyByID(id int) *models.Property {
	// Implement this function to retrieve a property by its ID from all neighborhoods
	// Example implementation:
	allNeighborhoods := []*models.Neighborhood{
		neighborhoods.GetDowntownDistrictNeighborhood(),
		neighborhoods.GetHistoricHeightsNeighborhood(),
		neighborhoods.GetTechValleyNeighborhood(),
		neighborhoods.GetCedarGroveNeighborhood(),
		neighborhoods.GetWillowFlatsNeighborhood(),
	}

	for _, neighborhood := range allNeighborhoods {
		for _, propID := range neighborhood.PropertyIDs {
			if propID == id {
				// Retrieve the property from the respective slice
				// Implement the actual retrieval logic based on your property storage
			}
		}
	}

	return nil
}

// getNeighborhoodMap creates a map of Neighborhood ID to Neighborhood pointer
func getNeighborhoodMap(neighborhoods []*models.Neighborhood) map[int]*models.Neighborhood {
	neighborhoodMap := make(map[int]*models.Neighborhood)
	for _, neighborhood := range neighborhoods {
		neighborhoodMap[neighborhood.ID] = neighborhood
	}
	return neighborhoodMap
}