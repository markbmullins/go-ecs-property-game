package systems_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/markbmullins/city-developer/pkg/components"
	"github.com/markbmullins/city-developer/pkg/ecs"
	"github.com/markbmullins/city-developer/pkg/systems"
	"github.com/markbmullins/city-developer/pkg/utils"
	"github.com/stretchr/testify/assert"
)

// applyUpgrade attempts to apply an upgrade to a property.
// It checks for prerequisites and updates the property's upgrade level if successful.
func applyUpgrade(world *ecs.World, propertyID int, upgradeID string, currentDate time.Time) bool {
	// Locate the property entity
	var property *components.Property
	for _, entity := range world.Entities {
		prop := entity.GetComponent("Property")
		if prop != nil && entity.ID == (ecs.EntityKey{EntityType: "Property", ID: propertyID}) {
			property = prop.(*components.Property)
			break
		}
	}

	// Find the upgrade by ID
	var targetUpgrade *components.Upgrade
	for i := range property.UpgradePaths {
		for _, upg := range property.UpgradePaths[i] {
			if upg.ID == upgradeID {
				targetUpgrade = &upg
				break
			}
		}
		if targetUpgrade != nil {
			break
		}
	}

	if targetUpgrade == nil {
		fmt.Printf("Upgrade with ID '%s' not found for property '%s'\n", upgradeID, property.Name)
		return false
	}

	// Check if prerequisites are met
	if targetUpgrade.Prerequisite != nil {
		prereqID := targetUpgrade.Prerequisite.ID
		prereqMet := false
		for _, upg := range property.Upgrades {
			if upg.ID == prereqID && upg.Level < targetUpgrade.Level {
				prereqMet = true
				break
			}
		}
		if !prereqMet {
			fmt.Printf("Prerequisite upgrade '%s' not met for upgrade '%s'\n", prereqID, upgradeID)
			return false
		}
	}

	// Apply the upgrade by creating a copy to avoid modifying the original UpgradePaths
	appliedUpgrade := *targetUpgrade // Create a copy
	appliedUpgrade.PurchaseDate = currentDate
	property.Upgrades = append(property.Upgrades, appliedUpgrade)

	fmt.Printf("Applied upgrade '%s' to property '%s'\n", upgradeID, property.Name)
	return true
}

// Helper function to create a test world with game time, player, and property
func createTestWorld(purchaseDate time.Time, baseRent float64) *ecs.World {
	world := ecs.NewWorld()

	// Add GameTime component
	gameTime := &components.GameTime{
		CurrentDate:     purchaseDate,
		IsPaused:        false,
		SpeedMultiplier: 1.0,
		LastUpdated:     purchaseDate,
	}

	timeEntity := ecs.NewGameTime()
	timeEntity.AddComponent("GameTime", gameTime)
	world.AddEntity(timeEntity)

	// Add Player component
	player := &components.Player{
		ID:    1,
		Funds: 0,
	}
	playerEntity := ecs.NewPlayer(1)
	playerEntity.AddComponent("Player", player)
	world.AddEntity(playerEntity)

	// Add Property component
	property := &components.Property{
		Name:         "Test Property",
		Owned:        true,
		BaseRent:     baseRent,
		PlayerID:     1,
		PurchaseDate: purchaseDate,
	}
	propertyEntity := ecs.NewProperty(2)
	propertyEntity.AddComponent("Property", property)
	world.AddEntity(propertyEntity)

	return world
}

// createTestWorldWithUpgrades sets up a test world with a property having specific upgrades and prerequisites.
func createTestWorldWithUpgrades(purchaseDate time.Time, baseRent float64) *ecs.World {
	world := ecs.NewWorld()

	// Add GameTime component
	gameTime := &components.GameTime{
		CurrentDate:     purchaseDate,
		IsPaused:        false,
		SpeedMultiplier: 1.0,
		LastUpdated:     purchaseDate,
	}
	timeEntity := ecs.NewGameTime()
	timeEntity.AddComponent("GameTime", gameTime)
	world.AddEntity(timeEntity)

	// Add Player component
	player := &components.Player{
		ID:    1,
		Funds: 0,
	}
	playerEntity := ecs.NewPlayer(1)
	playerEntity.AddComponent("Player", player)
	world.AddEntity(playerEntity)

	// Define Upgrades with Prerequisites
	renovatedInterior := components.Upgrade{
		ID:             "renovated_interior",
		Name:           "Renovated Interior",
		Level:          1,
		Cost:           10000,
		RentIncrease:   100.0,
		DaysToComplete: 7,
		Prerequisite:   nil,
	}

	smartHomeAutomation := components.Upgrade{
		ID:             "smart_home_automation",
		Name:           "Smart Home Automation",
		Level:          2,
		Cost:           20000,
		RentIncrease:   200.0,
		DaysToComplete: 14,
		Prerequisite:   &renovatedInterior,
	}

	premiumFixtures := components.Upgrade{
		ID:             "premium_fixtures",
		Name:           "Premium Fixtures",
		Level:          3,
		Cost:           30000,
		RentIncrease:   300.0,
		DaysToComplete: 21,
		Prerequisite:   &smartHomeAutomation,
	}

	// Create Property with Upgrade Paths
	property := &components.Property{
		Name:         "Test Property with Upgrades",
		Owned:        true,
		BaseRent:     baseRent,
		PlayerID:     1,
		PurchaseDate: purchaseDate,
		Upgrades:     []components.Upgrade{},
		UpgradePaths: map[string][]components.Upgrade{
			"Luxury": {renovatedInterior, smartHomeAutomation, premiumFixtures},
		},
	}

	// Add Property component
	propertyEntity := ecs.NewProperty(2)
	propertyEntity.AddComponent("Property", property)
	world.AddEntity(propertyEntity)

	return world
}

// createTestWorldWithCircularUpgrades sets up a test world with circular upgrade dependencies.
func createTestWorldWithCircularUpgrades(purchaseDate time.Time, baseRent float64) *ecs.World {
	world := ecs.NewWorld()

	// Add GameTime component
	gameTime := &components.GameTime{
		CurrentDate:     purchaseDate,
		IsPaused:        false,
		SpeedMultiplier: 1.0,
		LastUpdated:     purchaseDate,
	}
	timeEntity := ecs.NewGameTime()
	timeEntity.AddComponent("GameTime", gameTime)
	world.AddEntity(timeEntity)

	// Add Player component
	player := &components.Player{
		ID:    1,
		Funds: 0,
	}
	playerEntity := ecs.NewPlayer(1)
	playerEntity.AddComponent("Player", player)
	world.AddEntity(playerEntity)

	// Define Upgrades with Circular Prerequisites
	upgradeA := components.Upgrade{
		ID:             "upgrade_a",
		Name:           "Upgrade A",
		Level:          1,
		Cost:           5000,
		RentIncrease:   100.0,
		DaysToComplete: 5,
		Prerequisite:   nil, // Will be set to Upgrade B
	}

	upgradeB := components.Upgrade{
		ID:             "upgrade_b",
		Name:           "Upgrade B",
		Level:          2,
		Cost:           7000,
		RentIncrease:   150.0,
		DaysToComplete: 10,
		Prerequisite:   &upgradeA, // Upgrade A requires Upgrade B, creating a cycle
	}

	// Creating the cycle
	upgradeA.Prerequisite = &upgradeB

	// Create Property with Upgrade Paths
	property := &components.Property{
		Name:         "Test Property with Circular Upgrades",
		Owned:        true,
		BaseRent:     baseRent,
		PlayerID:     1,
		PurchaseDate: purchaseDate,
		Upgrades:     []components.Upgrade{},
		UpgradePaths: map[string][]components.Upgrade{
			"CircularPath": {upgradeA, upgradeB},
		},
	}

	// Add Property component
	propertyEntity := ecs.NewProperty(2)
	propertyEntity.AddComponent("Property", property)
	world.AddEntity(propertyEntity)

	return world
}

// createTestWorldWithUpgradeChain sets up a test world with a chain of upgrade prerequisites.
func createTestWorldWithUpgradeChain(purchaseDate time.Time, baseRent float64) *ecs.World {
	world := ecs.NewWorld()

	// Add GameTime component
	gameTime := &components.GameTime{
		CurrentDate:     purchaseDate,
		IsPaused:        false,
		SpeedMultiplier: 1.0,
		LastUpdated:     purchaseDate,
	}
	timeEntity := ecs.NewGameTime()
	timeEntity.AddComponent("GameTime", gameTime)
	world.AddEntity(timeEntity)

	// Add Player component
	player := &components.Player{
		ID:    1,
		Funds: 0,
	}
	playerEntity := ecs.NewPlayer(1)
	playerEntity.AddComponent("Player", player)
	world.AddEntity(playerEntity)

	// Define Upgrades with Prerequisite Chain
	upgradedInterior := components.Upgrade{
		ID:             "upgrade_1",
		Name:           "Upgrade 1",
		Level:          1,
		Cost:           5000,
		RentIncrease:   100.0,
		DaysToComplete: 5,
		Prerequisite:   nil,
	}

	upgrade2 := components.Upgrade{
		ID:             "upgrade_2",
		Name:           "Upgrade 2",
		Level:          2,
		Cost:           7000,
		RentIncrease:   150.0,
		DaysToComplete: 10,
		Prerequisite:   &upgradedInterior,
	}

	upgrade3 := components.Upgrade{
		ID:             "upgrade_3",
		Name:           "Upgrade 3",
		Level:          3,
		Cost:           9000,
		RentIncrease:   200.0,
		DaysToComplete: 15,
		Prerequisite:   &upgrade2,
	}

	// Create Property with Upgrade Paths
	property := &components.Property{
		Name:         "Test Property with Upgrade Chain",
		Owned:        true,
		BaseRent:     baseRent,
		PlayerID:     1,
		PurchaseDate: purchaseDate,
		Upgrades:     []components.Upgrade{},
		UpgradePaths: map[string][]components.Upgrade{
			"ChainPath": {upgradedInterior, upgrade2, upgrade3},
		},
	}

	// Add Property component
	propertyEntity := ecs.NewProperty(2)
	propertyEntity.AddComponent("Property", property)
	world.AddEntity(propertyEntity)

	return world
}

// createTestWorldWithUpgradeTree sets up a test world with a property having multiple upgrade paths.
func createTestWorldWithUpgradeTree(purchaseDate time.Time, baseRent float64) *ecs.World {
	world := ecs.NewWorld()

	// Add GameTime component
	gameTime := &components.GameTime{
		CurrentDate:     purchaseDate,
		IsPaused:        false,
		SpeedMultiplier: 1.0,
		LastUpdated:     purchaseDate,
	}
	timeEntity := ecs.NewGameTime()
	timeEntity.AddComponent("GameTime", gameTime)
	world.AddEntity(timeEntity)

	// Add Player component
	player := &components.Player{
		ID:    1,
		Funds: 0,
	}
	playerEntity := ecs.NewPlayer(1)
	playerEntity.AddComponent("Player", player)
	world.AddEntity(playerEntity)

	// Define Upgrades for "Luxury" path
	renovatedInterior := components.Upgrade{
		ID:             "renovated_interior",
		Name:           "Renovated Interior",
		Level:          1,
		Cost:           10000,
		RentIncrease:   100.0,
		DaysToComplete: 7,
		Prerequisite:   nil,
	}

	smartHomeAutomation := components.Upgrade{
		ID:             "smart_home_automation",
		Name:           "Smart Home Automation",
		Level:          2,
		Cost:           20000,
		RentIncrease:   200.0,
		DaysToComplete: 14,
		Prerequisite:   &renovatedInterior,
	}

	premiumFixtures := components.Upgrade{
		ID:             "premium_fixtures",
		Name:           "Premium Fixtures",
		Level:          3,
		Cost:           30000,
		RentIncrease:   300.0,
		DaysToComplete: 21,
		Prerequisite:   &smartHomeAutomation,
	}

	// Define Upgrades for "Efficiency" path
	solarPanels := components.Upgrade{
		ID:             "solar_panels",
		Name:           "Solar Panels",
		Level:          1,
		Cost:           8000,
		RentIncrease:   80.0,
		DaysToComplete: 10,
		Prerequisite:   nil,
	}

	energyEfficientWindows := components.Upgrade{
		ID:             "energy_efficient_windows",
		Name:           "Energy-efficient Windows",
		Level:          2,
		Cost:           12000,
		RentIncrease:   120.0,
		DaysToComplete: 15,
		Prerequisite:   &solarPanels,
	}

	highEfficiencyHVAC := components.Upgrade{
		ID:             "high_efficiency_hvac",
		Name:           "High-efficiency HVAC",
		Level:          3,
		Cost:           20000,
		RentIncrease:   200.0,
		DaysToComplete: 20,
		Prerequisite:   &energyEfficientWindows,
	}

	// Create Property with Upgrade Paths
	property := &components.Property{
		Name:         "Test Property with Multiple Paths",
		Owned:        true,
		BaseRent:     baseRent,
		PlayerID:     1,
		PurchaseDate: purchaseDate,
		Upgrades:     []components.Upgrade{},
		UpgradePaths: map[string][]components.Upgrade{
			"Luxury":     {renovatedInterior, smartHomeAutomation, premiumFixtures},
			"Efficiency": {solarPanels, energyEfficientWindows, highEfficiencyHVAC},
		},
	}

	// Add Property component
	propertyEntity := ecs.NewProperty(2)
	propertyEntity.AddComponent("Property", property)
	world.AddEntity(propertyEntity)

	return world
}

// TestProratedRent tests prorated rent calculation with various advancement speeds.
func TestProratedRent(t *testing.T) {
	purchaseDate := time.Date(2023, 1, 27, 0, 0, 0, 0, time.UTC)
	world := createTestWorld(purchaseDate, 1000.0)

	gameTime, err := utils.GetCurrentGameTime(world)
	assert.NoError(t, err)

	// Set LastUpdated to the purchase date
	gameTime.LastUpdated = purchaseDate

	// Set CurrentDate to February 1, 2023
	gameTime.CurrentDate = time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC)

	incomeSystem := systems.IncomeSystem{}
	incomeSystem.Update(world)

	player := world.GetPlayer(1).GetComponent("Player").(*components.Player)

	// Expected prorated rent calculation:
	// 1. Rent begins the day after the purchase date, i.e., January 28, 2023.
	// 2. Days owned in January: 4 (January 28–31 inclusive).
	// 3. Daily rent = Base rent / Days in January = 1000 / 31 ≈ 32.26.
	// 4. Prorated rent = Daily rent × Days owned = 32.26 × 4 = 129.04.
	// 5. Final rent is rounded down to the nearest multiple of 5: floor(129.04 / 5) * 5 = 125.00.
	// Therefore, the expected funds in the player's account after the update is $125.00.
	expectedProratedRent := 125.0

	assert.Equal(t, expectedProratedRent, player.Funds, "Prorated rent should be correct")
}

// TestFullMonthRent tests full rent collection for a single month at normal speed.
func TestPurchaseOnFirstDayOfMonth(t *testing.T) {
	purchaseDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	world := createTestWorld(purchaseDate, 1000.0)
	gameTime, err := utils.GetCurrentGameTime(world)
	assert.NoError(t, err)
	gameTime.CurrentDate = time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC)

	incomeSystem := systems.IncomeSystem{}
	incomeSystem.Update(world)

	player := world.GetPlayer(1).GetComponent("Player").(*components.Player)
	// Expected rent calculation:
	// 1. Rent starts the day after the purchase date. Since the property was purchased on January 1, rent begins on January 2.
	// 2. January has 31 days, so rent is calculated for 30 days (January 2–31 inclusive).
	// 3. Daily rent = Base rent / Days in January = 1000 / 31 ≈ 32.26.
	// 4. Total rent = Daily rent × 30 days = 32.26 × 30 ≈ 967.80.
	// 5. Rent is rounded down to the nearest multiple of 5: floor(967.80 / 5) * 5 = 965.00.
	// As a result, the expected rent collected is 965.00 instead of the full 1000.00.
	expectedFullRent := 965.0

	assert.Equal(t, expectedFullRent, player.Funds, "Full monthly rent should be collected")
}

func TestProratedAndFullMonthRent(t *testing.T) {
	// Set up the purchase date (January 1, 2023) and a property with a base monthly rent of $1000.00.
	purchaseDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	world := createTestWorld(purchaseDate, 1000.0)

	gameTime, err := utils.GetCurrentGameTime(world)
	assert.NoError(t, err)

	// Simulate collecting rent for January by advancing the game to February 1.
	gameTime.CurrentDate = time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC)
	incomeSystem := systems.IncomeSystem{}
	incomeSystem.Update(world)

	// Verify that January rent (prorated for 30 days) has been collected.
	player := world.GetPlayer(1).GetComponent("Player").(*components.Player)
	// 1. Rent starts the day after the purchase date. Since the property was purchased on January 1, rent begins on January 2.
	// 2. January has 31 days, so rent is calculated for 30 days (January 2–31 inclusive).
	// 3. Daily rent = Base rent / Days in January = 1000 / 31 ≈ 32.26.
	// 4. Total rent = Daily rent × 30 days = 32.26 × 30 ≈ 967.80.
	// 5. Rent is rounded down to the nearest multiple of 5: floor(967.80 / 5) * 5 = 965.00.
	expectedProratedRent := 965.0
	assert.Equal(t, expectedProratedRent, player.Funds, "Prorated January rent should be correct")

	// Advance to March 1, 2023, to collect rent for February (full month).
	gameTime.CurrentDate = time.Date(2023, 3, 1, 0, 0, 0, 0, time.UTC)
	incomeSystem.Update(world)

	// Verify that February's full rent has been collected.
	expectedTotalFunds := expectedProratedRent + 1000.0
	assert.Equal(t, expectedTotalFunds, player.Funds, "Full February rent should be collected")
}

func TestFullMonthRent(t *testing.T) {
	purchaseDate := time.Date(2022, 12, 1, 0, 0, 0, 0, time.UTC)
	world := createTestWorld(purchaseDate, 1000.0)

	gameTime, err := utils.GetCurrentGameTime(world)
	assert.NoError(t, err)

	// Set LastUpdated to the start of February 2023.
	gameTime.LastUpdated = time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC)

	// Set CurrentDate to March 1, 2023, to collect the full rent for February 2023.
	gameTime.CurrentDate = time.Date(2023, 3, 1, 0, 0, 0, 0, time.UTC)

	incomeSystem := systems.IncomeSystem{}
	incomeSystem.Update(world)

	player := world.GetPlayer(1).GetComponent("Player").(*components.Player)

	// Full monthly rent is $1000.00.
	expectedFullRent := 1000.0

	assert.Equal(t, expectedFullRent, player.Funds, "Full February rent should be collected")
}

// TestMultipleMonthAdvancement tests rent collection when more than one month passes in a single update.
func TestMultipleMonthAdvancement(t *testing.T) {
	purchaseDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	world := createTestWorld(purchaseDate, 1000.0)
	gameTime, err := utils.GetCurrentGameTime(world)
	assert.NoError(t, err)
	gameTime.CurrentDate = time.Date(2023, 5, 1, 0, 0, 0, 0, time.UTC) // Advance by 4 months

	incomeSystem := systems.IncomeSystem{}
	incomeSystem.Update(world)

	player := world.GetPlayer(1).GetComponent("Player").(*components.Player)
	// Expected rent calculation:
	// 1. Rent starts the day after the purchase date. Since the property was purchased on January 1, rent begins on January 2.
	// 2. January has 31 days, so rent is calculated for 30 days (January 2–31 inclusive).
	//    - Daily rent = Base rent / Days in January = 1000 / 31 ≈ 32.26.
	//    - Prorated rent for January = Daily rent × 30 days = 32.26 × 30 ≈ 967.80.
	//    - Rent is rounded down to the nearest multiple of 5: floor(967.80 / 5) * 5 = 965.00.
	// 3. Rent for February, March, and April is collected as full months, each at $1000.00.
	//    - Total full-month rent = 1000 × 3 = 3000.00.
	// 4. Total rent collected over the 4-month period (January–April):
	//    - Total rent = Prorated January rent + Full-month rent for February–April.
	//    - Total rent = 965.00 + 3000.00 = 3965.00.
	//
	// Thus, the expected total rent collected is $3965.00.
	expectedTotalRent := 965.0 + 1000.0*3

	assert.Equal(t, expectedTotalRent, player.Funds, "Rent for 4 months should be collected")
}

// TestSpeedMultiplierEffect tests rent collection when the game speed multiplier is set to high values.
func TestSpeedMultiplierEffect(t *testing.T) {
	purchaseDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	world := createTestWorld(purchaseDate, 1000.0)
	gameTime, err := utils.GetCurrentGameTime(world)
	assert.NoError(t, err)
	gameTime.CurrentDate = time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC) // Advance by 12 months

	incomeSystem := systems.IncomeSystem{}
	incomeSystem.Update(world)

	player := world.GetPlayer(1).GetComponent("Player").(*components.Player)
	expectedTotalRent := 965.0 + 1000.0*11 // Correct total rent based on prorated first month

	assert.Equal(t, expectedTotalRent, player.Funds, "Rent for 12 months should be collected")
}

func TestLeapYearFebruary(t *testing.T) {
	purchaseDate := time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC) // 2024 is a leap year
	world := createTestWorld(purchaseDate, 1000.0)

	gameTime, err := utils.GetCurrentGameTime(world)
	assert.NoError(t, err)

	// Set LastUpdated to purchase date
	gameTime.LastUpdated = purchaseDate

	// Advance to March 1
	gameTime.CurrentDate = time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC)

	incomeSystem := systems.IncomeSystem{}
	incomeSystem.Update(world)

	player := world.GetPlayer(1).GetComponent("Player").(*components.Player)

	// Expected prorataed rent for February (29 days - purchase day)
	expectedRent := 965.0

	assert.Equal(t, expectedRent, player.Funds, "Full rent should be collected for leap year February")
}

func TestFuturePropertyPurchase(t *testing.T) {
	purchaseDate := time.Date(2023, 3, 1, 0, 0, 0, 0, time.UTC) // Purchased in the future
	world := createTestWorld(purchaseDate, 1000.0)

	gameTime, err := utils.GetCurrentGameTime(world)
	assert.NoError(t, err)

	// Set CurrentDate before the purchase date
	gameTime.CurrentDate = time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC)

	incomeSystem := systems.IncomeSystem{}
	incomeSystem.Update(world)

	player := world.GetPlayer(1).GetComponent("Player").(*components.Player)
	expectedRent := 0.0

	assert.Equal(t, expectedRent, player.Funds, "No rent should be collected for a future property purchase")
}

func TestRentNotCollectedWhenGamePaused(t *testing.T) {
	purchaseDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	world := createTestWorld(purchaseDate, 1000.0)

	gameTime, err := utils.GetCurrentGameTime(world)
	assert.NoError(t, err)

	// Set game as paused
	gameTime.IsPaused = true
	gameTime.CurrentDate = time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC)

	incomeSystem := systems.IncomeSystem{}
	incomeSystem.Update(world)

	player := world.GetPlayer(1).GetComponent("Player").(*components.Player)
	expectedRent := 0.0 // No rent when game is paused

	assert.Equal(t, expectedRent, player.Funds, "No rent should be collected when game is paused")
}

func TestMultiplePropertiesWithDifferentPlayers(t *testing.T) {
	world := ecs.NewWorld()

	// Add GameTime component
	gameTime := &components.GameTime{
		CurrentDate:     time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC),
		IsPaused:        false,
		SpeedMultiplier: 1.0,
		LastUpdated:     time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	timeEntity := ecs.NewGameTime()
	timeEntity.AddComponent("GameTime", gameTime)
	world.AddEntity(timeEntity)

	// Add Player 1 and Property 1
	player1 := &components.Player{ID: 1, Funds: 0}
	playerEntity1 := ecs.NewPlayer(1)
	playerEntity1.AddComponent("Player", player1)
	world.AddEntity(playerEntity1)

	property1 := &components.Property{
		Name:         "Property1",
		Owned:        true,
		BaseRent:     1000.0,
		PlayerID:     player1.ID,
		PurchaseDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	propertyEntity1 := ecs.NewProperty(2)
	propertyEntity1.AddComponent("Property", property1)
	world.AddEntity(propertyEntity1)

	// Add Player 2 and Property 2
	player2 := &components.Player{ID: 2, Funds: 0}
	playerEntity2 := ecs.NewPlayer(player2.ID)
	playerEntity2.AddComponent("Player", player2)
	world.AddEntity(playerEntity2)

	property2 := &components.Property{
		Name:         "Property2",
		Owned:        true,
		BaseRent:     2000.0,
		PlayerID:     3, // Player ID corresponds to Player.ID
		PurchaseDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	propertyEntity2 := ecs.NewProperty(4)
	propertyEntity2.AddComponent("Property", property2)
	world.AddEntity(propertyEntity2)

	incomeSystem := systems.IncomeSystem{}
	incomeSystem.Update(world)

	// Expected rent calculation:
	// For Player 1's property (Property1):
	// 1. Property purchased on January 1, 2023, so rent starts January 2.
	// 2. January has 31 days, so rent is for 30 days (Jan 2–31).
	// 3. Daily base rent = 1000 / 31 ≈ 32.26.
	// 4. Total rent = 32.26 × 30 ≈ 967.80.
	// 5. Rounded down to nearest multiple of 5: floor(967.80/5)*5 = 965.00.

	// For Player 2's property (Property2):
	// 1. Property purchased on January 1, 2023, so rent starts January 2.
	// 2. January has 31 days, so rent is for 30 days (Jan 2–31).
	// 3. Daily base rent = 2000 / 31 ≈ 64.52.
	// 4. Total rent = 64.52 × 30 ≈ 1935.48.
	// 5. Rounded down: floor(1935.48/5)*5 = 1935.00.

	// Player 1 expected = 965.00
	// Player 2 expected = 1935.00

	// Assert funds for Player 1
	playerComp1 := world.GetPlayer(1).GetComponent("Player").(*components.Player)
	assert.Equal(t, 965.0, playerComp1.Funds, "Player 1's rent should be collected correctly")

	// Assert funds for Player 2
	playerComp2 := world.GetPlayer(2).GetComponent("Player").(*components.Player)
	assert.Equal(t, 1935.0, playerComp2.Funds, "Player 2's rent should be collected correctly")
}

func TestRentWithPropertyUpgrade(t *testing.T) {
	// Set Purchase Date to December 31, 2022 (before January 2023)
	purchaseDate := time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC)
	world := createTestWorld(purchaseDate, 1000.0)

	// Configure Property Upgrades
	propertyComp := world.GetProperty(2).GetComponent("Property").(*components.Property)
	propertyComp.Upgrades = []components.Upgrade{
		{RentIncrease: 200.0}, // Level 1
	}

	// Set CurrentDate to January 31, 2023
	gameTime, err := utils.GetCurrentGameTime(world)
	assert.NoError(t, err)
	gameTime.CurrentDate = time.Date(2023, 1, 31, 0, 0, 0, 0, time.UTC)

	// Update Income System
	incomeSystem := systems.IncomeSystem{}
	incomeSystem.Update(world)

	// Assert Player Funds
	player := world.GetPlayer(1).GetComponent("Player").(*components.Player)
	expectedRent := 1200.0 // 1000 base + 200 upgrade

	assert.Equal(t, expectedRent, player.Funds, "Rent should include property upgrade effect")
}

func TestProratedRentWithUpgrade(t *testing.T) {
	// Set Purchase Date to January 15, 2023 (mid-month)
	baseRent := 1000.0
	purchaseDate := time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC)
	world := createTestWorld(purchaseDate, baseRent)

	// So rent kicks in Jan 16 to Jan 31 inclusive! 16 days

	// Configure Property Upgrades
	propertyComp := world.GetProperty(2).GetComponent("Property").(*components.Property)
	propertyComp.Upgrades = []components.Upgrade{
		{
			RentIncrease:   200.0,
			PurchaseDate:   purchaseDate,
			DaysToComplete: 0, // Upgrade completed on the purchase day
		},
	}

	// Set CurrentDate to February 1, 2023
	gameTime, err := utils.GetCurrentGameTime(world)
	assert.NoError(t, err)
	gameTime.CurrentDate = time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC)

	// Update Income System
	incomeSystem := systems.IncomeSystem{}
	incomeSystem.Update(world)

	// Assert Player Funds
	// Calculate Expected Rent

	player := world.GetPlayer(1).GetComponent("Player").(*components.Player)
	// TODO: which rent is correct?
	// expectedProratedRent := 715.0
	// (baseRent * 16)/31 rounded down to nearest 5 + 200 * 16 / 31 rounded down to nearest 5
	expectedProratedRent := 615.0

	// verify if the IncomeSystem correctly calculates daysOwned.
	assert.Equal(t, expectedProratedRent, player.Funds, "Prorated rent should include property upgrade effect")
}

func TestMultiLevelUpgrades(t *testing.T) {
	purchaseDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	world := createTestWorld(purchaseDate, 1000.0)

	// Configure Property with multiple upgrade levels
	propertyComp := world.GetProperty(2).GetComponent("Property").(*components.Property)
	propertyComp.Upgrades = []components.Upgrade{
		{RentIncrease: 100.0, DaysToComplete: 0, PurchaseDate: purchaseDate},                  // Level 1
		{RentIncrease: 200.0, DaysToComplete: 0, PurchaseDate: purchaseDate.AddDate(0, 0, 1)}, // Level 2
		{RentIncrease: 300.0, DaysToComplete: 0, PurchaseDate: purchaseDate.AddDate(0, 0, 2)}, // Level 3
	}

	gameTime, err := utils.GetCurrentGameTime(world)
	assert.NoError(t, err)
	gameTime.CurrentDate = time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC)

	incomeSystem := systems.IncomeSystem{}
	incomeSystem.Update(world)

	player := world.GetPlayer(1).GetComponent("Player").(*components.Player)
	expectedRent := 1520.0

	assert.Equal(t, expectedRent, player.Funds, "Rent with multiple upgrades should be collected correctly")
}

func TestExcessiveUpgradeLevel(t *testing.T) {
	purchaseDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	world := createTestWorld(purchaseDate, 1000.0)

	// Configure Property with an out-of-bounds upgrade level
	property := world.GetProperty(2).GetComponent("Property").(*components.Property)
	property.Upgrades = []components.Upgrade{
		{RentIncrease: 100.0}, // Level 1
	}

	gameTime, err := utils.GetCurrentGameTime(world)
	assert.NoError(t, err)
	gameTime.CurrentDate = time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC)

	incomeSystem := systems.IncomeSystem{}
	incomeSystem.Update(world)

	player := world.GetPlayer(1).GetComponent("Player").(*components.Player)
	expectedRent := 1060.0 // Correct total rent based on prorated base rent + available upgrades

	assert.Equal(t, expectedRent, player.Funds, "Rent should include available upgrades even if UpgradeLevel exceeds available upgrades")
}

func TestFractionalSpeedMultiplier(t *testing.T) {
	purchaseDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	world := createTestWorld(purchaseDate, 1000.0)

	gameTime, err := utils.GetCurrentGameTime(world)
	assert.NoError(t, err)
	gameTime.SpeedMultiplier = 1.5 // Fast-forwards by 1.5x
	gameTime.CurrentDate = time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC)

	incomeSystem := systems.IncomeSystem{}
	incomeSystem.Update(world)

	player := world.GetPlayer(1).GetComponent("Player").(*components.Player)
	expectedRent := 965.0

	assert.Equal(t, expectedRent, player.Funds, "Fractional speed multiplier should not partially advance month")
}

func TestNewMonthFlag(t *testing.T) {
	purchaseDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	world := createTestWorld(purchaseDate, 1000.0)

	gameTime, err := utils.GetCurrentGameTime(world)
	assert.NoError(t, err)
	gameTime.CurrentDate = time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC)

	incomeSystem := systems.IncomeSystem{}
	incomeSystem.Update(world)

	// Confirm that the NewMonth flag is set
	assert.True(t, gameTime.NewMonth, "NewMonth flag should be set after month advancement")
}

func TestMissingProperty(t *testing.T) {
	world := ecs.NewWorld()
	world.AddEntity(ecs.NewGameTime()) // Only adding GameTime, no Property

	gameTime := &components.GameTime{
		CurrentDate: time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC),
		IsPaused:    false,
	}
	world.GetGameTime().AddComponent("GameTime", gameTime)

	incomeSystem := systems.IncomeSystem{}
	incomeSystem.Update(world)

	// Test should pass without error even though Property is missing
}

func TestLastUpdatedAfterCurrentDate(t *testing.T) {
	purchaseDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	world := createTestWorld(purchaseDate, 1000.0)

	gameTime, err := utils.GetCurrentGameTime(world)
	assert.NoError(t, err)
	gameTime.LastUpdated = time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC) // Set LastUpdated ahead of CurrentDate
	gameTime.CurrentDate = time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC)

	incomeSystem := systems.IncomeSystem{}
	incomeSystem.Update(world)

	player := world.GetPlayer(1).GetComponent("Player").(*components.Player)
	expectedRent := 0.0

	assert.Equal(t, expectedRent, player.Funds, "No rent should be collected if LastUpdated is after CurrentDate")
}

func TestPurchaseAfterLeapYearFebruary(t *testing.T) {
	purchaseDate := time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC) // Right after leap year February
	world := createTestWorld(purchaseDate, 1000.0)

	gameTime, err := utils.GetCurrentGameTime(world)
	assert.NoError(t, err)
	gameTime.CurrentDate = time.Date(2024, 4, 1, 0, 0, 0, 0, time.UTC)

	incomeSystem := systems.IncomeSystem{}
	incomeSystem.Update(world)

	player := world.GetPlayer(1).GetComponent("Player").(*components.Player)
	expectedRent := 965.0 // Correct total rent based on prorated first month

	assert.Equal(t, expectedRent, player.Funds, "Rent should be collected correctly after leap year February")
}

func TestBackdatedRentCollection(t *testing.T) {
	purchaseDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	world := createTestWorld(purchaseDate, 1000.0)

	gameTime, err := utils.GetCurrentGameTime(world)
	assert.NoError(t, err)
	gameTime.LastUpdated = time.Date(2023, 5, 1, 0, 0, 0, 0, time.UTC) // Far future date
	gameTime.CurrentDate = time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC) // Backdated

	incomeSystem := systems.IncomeSystem{}
	incomeSystem.Update(world)

	player := world.GetPlayer(1).GetComponent("Player").(*components.Player)
	expectedRent := 0.0 // No rent due to backdated dates

	assert.Equal(t, expectedRent, player.Funds, "No rent should be collected when CurrentDate is backdated")
}

func TestPropertyPurchaseDuringFebruaryNonLeapYear(t *testing.T) {
	purchaseDate := time.Date(2023, 2, 15, 0, 0, 0, 0, time.UTC) // February of non-leap year
	world := createTestWorld(purchaseDate, 1000.0)

	gameTime, err := utils.GetCurrentGameTime(world)
	assert.NoError(t, err)
	gameTime.CurrentDate = time.Date(2023, 3, 1, 0, 0, 0, 0, time.UTC)

	incomeSystem := systems.IncomeSystem{}
	incomeSystem.Update(world)

	player := world.GetPlayer(1).GetComponent("Player").(*components.Player)
	expectedProratedRent := 460.0

	assert.Equal(t, expectedProratedRent, player.Funds, "Prorated rent should be correct for February in non-leap year")
}

func TestUpgradePrerequisiteEnforcement(t *testing.T) {
	// Initialize test world with upgrades
	purchaseDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	world := createTestWorldWithUpgrades(purchaseDate, 1000.0)

	// Attempt to apply "smart_home_automation" without applying "renovated_interior"
	success := applyUpgrade(world, 2, "smart_home_automation", purchaseDate)
	assert.False(t, success, "Upgrade should not be applied without prerequisites")

	// Verify that UpgradeLevel remains 0
	property := world.GetProperty(2).GetComponent("Property").(*components.Property)
	assert.Equal(t, 0, len(property.Upgrades), "Upgrade level should remain 0 when prerequisites are not met")
}

func TestUpgradeApplicationOrder(t *testing.T) {
	// Initialize test world with upgrades
	purchaseDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	world := createTestWorldWithUpgrades(purchaseDate, 1000.0)

	// Apply "renovated_interior" (Level 1)
	success := applyUpgrade(world, 2, "renovated_interior", purchaseDate)
	assert.True(t, success, "Upgrade 'renovated_interior' should be applied successfully")

	// Attempt to apply "premium_fixtures" (Level 3) without applying "smart_home_automation" (Level 2)
	success = applyUpgrade(world, 2, "premium_fixtures", purchaseDate.AddDate(0, 0, 10))
	assert.False(t, success, "Upgrade 'premium_fixtures' should not be applied without 'smart_home_automation'")

	// Apply "smart_home_automation" (Level 2)
	success = applyUpgrade(world, 2, "smart_home_automation", purchaseDate.AddDate(0, 0, 5))
	assert.True(t, success, "Upgrade 'smart_home_automation' should be applied successfully")

	// Now apply "premium_fixtures" (Level 3)
	success = applyUpgrade(world, 2, "premium_fixtures", purchaseDate.AddDate(0, 0, 15))
	assert.True(t, success, "Upgrade 'premium_fixtures' should be applied successfully after prerequisites")

	// Verify that UpgradeLevel is 3
	property := world.GetProperty(2).GetComponent("Property").(*components.Property)
	assert.Equal(t, 3, len(property.Upgrades), "Upgrade level should be 3 after applying three upgrades")
}

func TestUpgradePathIndependence(t *testing.T) {
	// Initialize test world with multiple upgrade paths
	purchaseDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	world := createTestWorldWithUpgradeTree(purchaseDate, 1000.0) // Assuming a helper that sets up multiple paths

	// Attempt to apply "energy_efficient_windows" before applying "solar_panels"
	success := applyUpgrade(world, 2, "energy_efficient_windows", purchaseDate.AddDate(0, 0, 10))
	assert.False(t, success, "Upgrade 'energy_efficient_windows' should not be applied without 'solar_panels'")

	// Apply "solar_panels" from "Efficiency" path
	success = applyUpgrade(world, 2, "solar_panels", purchaseDate)
	assert.True(t, success, "Upgrade 'solar_panels' should be applied successfully")

	// Now apply "energy_efficient_windows" after applying "solar_panels"
	success = applyUpgrade(world, 2, "energy_efficient_windows", purchaseDate.AddDate(0, 0, 10))
	assert.True(t, success, "Upgrade 'energy_efficient_windows' should be applied successfully after 'solar_panels'")

	// Verify that upgrades from "Efficiency" path do not affect "Luxury" path
	// Attempt to apply "renovated_interior" from "Luxury" path
	success = applyUpgrade(world, 2, "renovated_interior", purchaseDate.AddDate(0, 0, 15))
	assert.True(t, success, "Upgrade 'renovated_interior' should be applied successfully without affecting 'Efficiency' path")

	// Verify that UpgradeLevel reflects all applied upgrades
	property := world.GetProperty(2).GetComponent("Property").(*components.Property)
	assert.Equal(t, 3, len(property.Upgrades), "Upgrade level should be 3 after applying three independent upgrades")
}

func TestUpgradeCompletionTiming(t *testing.T) {
	// Initialize test world with upgrades
	purchaseDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	world := createTestWorldWithUpgrades(purchaseDate, 1000.0)

	// Apply "renovated_interior" on January 10, completes on January 17
	success := applyUpgrade(world, 2, "renovated_interior", time.Date(2023, 1, 10, 0, 0, 0, 0, time.UTC))
	assert.True(t, success, "Upgrade 'renovated_interior' should be applied successfully")

	// Advance to February 1, 2023
	gameTime, err := utils.GetCurrentGameTime(world)
	assert.NoError(t, err)
	gameTime.LastUpdated = purchaseDate
	gameTime.CurrentDate = time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC)

	// Update Income System
	incomeSystem := systems.IncomeSystem{}
	incomeSystem.Update(world)

	player := world.GetPlayer(1).GetComponent("Player").(*components.Player)
	expectedProratedRent := 1010.0

	assert.Equal(t, expectedProratedRent, player.Funds, "Prorated rent should include upgrade effect after completion")
}

func TestCircularUpgradePrerequisites(t *testing.T) {
	// Initialize test world with circular upgrade dependencies
	purchaseDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	world := createTestWorldWithCircularUpgrades(purchaseDate, 1000.0)

	// Attempt to apply "upgrade_a" which requires "upgrade_b"
	success := applyUpgrade(world, 2, "upgrade_a", purchaseDate)
	assert.False(t, success, "Upgrade 'upgrade_a' should not be applied without 'upgrade_b'")

	// Attempt to apply "upgrade_b" which requires "upgrade_a"
	success = applyUpgrade(world, 2, "upgrade_b", purchaseDate.AddDate(0, 0, 5))
	assert.False(t, success, "Upgrade 'upgrade_b' should not be applied without 'upgrade_a'")

	// Verify that no upgrades have been applied
	property := world.GetProperty(2).GetComponent("Property").(*components.Property)
	assert.Equal(t, 0, len(property.Upgrades), "No upgrades should be applied due to circular prerequisites")
}

func TestUpgradePrerequisiteChain(t *testing.T) {
	// Initialize test world with an upgrade chain
	purchaseDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	world := createTestWorldWithUpgradeChain(purchaseDate, 1000.0)

	// Attempt to apply "upgrade_3" without applying "upgrade_1" and "upgrade_2"
	success := applyUpgrade(world, 2, "upgrade_3", purchaseDate.AddDate(0, 0, 10))
	assert.False(t, success, "Upgrade 'upgrade_3' should not be applied without 'upgrade_1' and 'upgrade_2'")

	// Apply "upgrade_1"
	success = applyUpgrade(world, 2, "upgrade_1", purchaseDate.AddDate(0, 0, 5))
	assert.True(t, success, "Upgrade 'upgrade_1' should be applied successfully")

	// Attempt to apply "upgrade_3" without applying "upgrade_2"
	success = applyUpgrade(world, 2, "upgrade_3", purchaseDate.AddDate(0, 0, 15))
	assert.False(t, success, "Upgrade 'upgrade_3' should not be applied without 'upgrade_2'")

	// Apply "upgrade_2"
	success = applyUpgrade(world, 2, "upgrade_2", purchaseDate.AddDate(0, 0, 10))
	assert.True(t, success, "Upgrade 'upgrade_2' should be applied successfully")

	// Now apply "upgrade_3"
	success = applyUpgrade(world, 2, "upgrade_3", purchaseDate.AddDate(0, 0, 15))
	assert.True(t, success, "Upgrade 'upgrade_3' should be applied successfully after prerequisites")

	// Verify that UpgradeLevel is 3
	property := world.GetProperty(2).GetComponent("Property").(*components.Property)
	assert.Equal(t, 3, len(property.Upgrades), "Upgrade level should be 3 after applying three upgrades")

	// Advance to February 1, 2023 and collect rent
	gameTime, err := utils.GetCurrentGameTime(world)
	assert.NoError(t, err)
	gameTime.LastUpdated = purchaseDate
	gameTime.CurrentDate = time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC)

	incomeSystem := systems.IncomeSystem{}
	incomeSystem.Update(world)

	player := world.GetPlayer(1).GetComponent("Player").(*components.Player)
	expectedRent := 1080.0

	assert.Equal(t, expectedRent, player.Funds, "Rent should include all applied upgrades")
}
func TestUpgradeTreeComprehensive(t *testing.T) {
	// Initialize test world with multiple upgrade paths
	purchaseDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	world := createTestWorldWithUpgradeTree(purchaseDate, 1000.0)

	// Apply "renovated_interior" from "Luxury" path on Jan3
	success := applyUpgrade(world, 2, "renovated_interior", purchaseDate.AddDate(0, 0, 2))
	assert.True(t, success, "Upgrade 'renovated_interior' should be applied successfully")

	// Attempt to apply "premium_fixtures" without applying "smart_home_automation" on Jan5
	success = applyUpgrade(world, 2, "premium_fixtures", purchaseDate.AddDate(0, 0, 4))
	assert.False(t, success, "Upgrade 'premium_fixtures' should not be applied without 'smart_home_automation'")

	// Apply "smart_home_automation" from "Luxury" path on Jan5
	success = applyUpgrade(world, 2, "smart_home_automation", purchaseDate.AddDate(0, 0, 5))
	assert.True(t, success, "Upgrade 'smart_home_automation' should be applied successfully")

	// Now apply "premium_fixtures" from "Luxury" path on Jan7
	success = applyUpgrade(world, 2, "premium_fixtures", purchaseDate.AddDate(0, 0, 7))
	assert.True(t, success, "Upgrade 'premium_fixtures' should be applied successfully after prerequisites")

	// Apply "solar_panels" from "Efficiency" path on Jan3
	success = applyUpgrade(world, 2, "solar_panels", purchaseDate.AddDate(0, 0, 3))
	assert.True(t, success, "Upgrade 'solar_panels' should be applied successfully")

	// Attempt to apply "high_efficiency_hvac" without applying "energy_efficient_windows" on Jan6
	success = applyUpgrade(world, 2, "high_efficiency_hvac", purchaseDate.AddDate(0, 0, 6))
	assert.False(t, success, "Upgrade 'high_efficiency_hvac' should not be applied without 'energy_efficient_windows'")

	// Apply "energy_efficient_windows" from "Efficiency" path on Jan8
	success = applyUpgrade(world, 2, "energy_efficient_windows", purchaseDate.AddDate(0, 0, 8))
	assert.True(t, success, "Upgrade 'energy_efficient_windows' should be applied successfully")

	// Now apply "high_efficiency_hvac" from "Efficiency" path on Jan10
	success = applyUpgrade(world, 2, "high_efficiency_hvac", purchaseDate.AddDate(0, 0, 10))
	assert.True(t, success, "Upgrade 'high_efficiency_hvac' should be applied successfully after prerequisites")

	// Verify that UpgradeLevel is 6
	property := world.GetProperty(2).GetComponent("Property").(*components.Property)
	assert.Equal(t, 6, len(property.Upgrades), "Upgrade level should be 6 after applying six upgrades")

	// Advance to February 1, 2023 and collect rent
	gameTime, err := utils.GetCurrentGameTime(world)
	assert.NoError(t, err)
	gameTime.LastUpdated = purchaseDate
	gameTime.CurrentDate = time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC)

	incomeSystem := systems.IncomeSystem{}
	incomeSystem.Update(world)
	player := world.GetPlayer(1).GetComponent("Player").(*components.Player)

	expectedRent := 1195.0
	assert.Equal(t, expectedRent, player.Funds, "Rent should include all applied upgrades from multiple paths")
}
