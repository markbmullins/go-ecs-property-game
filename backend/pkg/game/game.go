package game

import (
	"time"

	"github.com/markbmullins/city-developer/pkg/components"
	"github.com/markbmullins/city-developer/pkg/ecs"
	"github.com/markbmullins/city-developer/pkg/models"
	"github.com/markbmullins/city-developer/pkg/systems"
)

func InitializeGame() *ecs.World {
	world := ecs.NewWorld()

	gameTimeModel := &models.GameTime{
		CurrentDate:     time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
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

	// Create neighborhood and properties
	neighborhood := &models.Neighborhood{
		ID:                 1,
		Name:               "Downtown",
		PropertyIDs:        []int{2, 3},
		RentBoostThreshold: 50.0, // Configurable threshold
		RentBoostAmount:    10.0, // Configurable boost percentage
	}

	luxuryUpgrades := []models.Upgrade{
		{ID: "luxury_1", Name: "Renovated Interior", Cost: 10000, RentIncrease: 500, DaysToComplete: 7},
		{ID: "luxury_2", Name: "Smart Home Automation", Cost: 20000, RentIncrease: 1000, DaysToComplete: 14},
		{ID: "luxury_3", Name: "Premium Fixtures", Cost: 30000, RentIncrease: 1500, DaysToComplete: 21},
	}

	efficiencyUpgrades := []models.Upgrade{
		{ID: "efficiency_1", Name: "Solar Panels", Cost: 8000, RentIncrease: 300, DaysToComplete: 10},
		{ID: "efficiency_2", Name: "Energy-efficient Windows", Cost: 12000, RentIncrease: 500, DaysToComplete: 15},
		{ID: "efficiency_3", Name: "High-efficiency HVAC", Cost: 20000, RentIncrease: 800, DaysToComplete: 20},
	}

	capacityUpgrades := []models.Upgrade{
		{ID: "capacity_1", Name: "Expand Seating Area", Cost: 15000, RentIncrease: 700, DaysToComplete: 12},
		{ID: "capacity_2", Name: "Add Outdoor Seating", Cost: 25000, RentIncrease: 1200, DaysToComplete: 18},
	}

	technologyUpgrades := []models.Upgrade{
		{ID: "tech_1", Name: "Install POS System", Cost: 5000, RentIncrease: 200, DaysToComplete: 5},
		{ID: "tech_2", Name: "Automated Inventory Management", Cost: 10000, RentIncrease: 400, DaysToComplete: 10},
		{ID: "tech_3", Name: "Customer Loyalty App", Cost: 15000, RentIncrease: 600, DaysToComplete: 14},
	}

	// Set upgrade prerequisites
	luxuryUpgrades[1].Prerequisite = &luxuryUpgrades[0] // "Smart Home Automation" requires "Renovated Interior"
	luxuryUpgrades[2].Prerequisite = &luxuryUpgrades[1] // "Premium Fixtures" requires "Smart Home Automation"

	efficiencyUpgrades[1].Prerequisite = &efficiencyUpgrades[0] // "Energy-efficient Windows" requires "Solar Panels"
	efficiencyUpgrades[2].Prerequisite = &efficiencyUpgrades[1] // "High-efficiency HVAC" requires "Energy-efficient Windows"

	capacityUpgrades[1].Prerequisite = &capacityUpgrades[0] // "Add Outdoor Seating" requires "Expand Seating Area"

	technologyUpgrades[1].Prerequisite = &technologyUpgrades[0] // "Automated Inventory Management" requires "Install POS System"
	technologyUpgrades[2].Prerequisite = &technologyUpgrades[1] // "Customer Loyalty App" requires "Automated Inventory Management"

	// Create property entities
	properties := []*models.Property{
		{
			Name:         "Residential 1",
			Type:         models.Residential,
			Subtype:      models.SingleFamily,
			BaseRent:     1000,
			Owned:        false,
			UpgradeLevel: 0,
			UpgradePaths: map[string][]models.Upgrade{
				"Luxury":     luxuryUpgrades,
				"Efficiency": efficiencyUpgrades,
			},
			Price:          10000,
			NeighborhoodID: neighborhood.ID,
		},
		{
			Name:         "Downtown Restaurant",
			Type:         models.Commercial,
			Subtype:      models.Restaurant,
			BaseRent:     5000,
			Owned:        false,
			UpgradeLevel: 0,
			UpgradePaths: map[string][]models.Upgrade{
				"Capacity":   capacityUpgrades,
				"Technology": technologyUpgrades,
			},
			Price:          50000,
			NeighborhoodID: neighborhood.ID,
		},
	}

	for i, prop := range properties {
		propertyEntity := ecs.NewEntity(i + 2)
		propertyEntity.AddComponent("PropertyComponent", &components.PropertyComponent{Property: prop})
		world.AddEntity(propertyEntity)
		neighborhood.PropertyIDs = append(neighborhood.PropertyIDs, i+2)
	}

	neighborhoodSystem := &systems.NeighborhoodValueSystem{
		Neighborhoods: map[int]*models.Neighborhood{
			neighborhood.ID: neighborhood,
		},
	}

	// Add systems
	world.AddSystem(&systems.IncomeSystem{})
	world.AddSystem(neighborhoodSystem)
	world.AddSystem(&systems.PropertyManagementSystem{})
	world.AddSystem(&systems.TimeSystem{})

	return world
}
