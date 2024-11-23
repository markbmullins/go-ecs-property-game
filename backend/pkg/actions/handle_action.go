package actions

import (
	"encoding/json"
	"net/http"

	"github.com/markbmullins/city-developer/pkg/components"
	"github.com/markbmullins/city-developer/pkg/ecs"
	"github.com/markbmullins/city-developer/pkg/models"
	"github.com/markbmullins/city-developer/pkg/utils"
)

// ActionRequest defines the structure of an incoming action request.
type ActionRequest struct {
	Action  string      `json:"action"`
	Payload interface{} `json:"payload"`
}

// HandleAction processes an incoming action request.
func HandleAction(world *ecs.World, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.SendResponse(w, "error", "Invalid request method", nil, http.StatusMethodNotAllowed)
		return
	}

	var actionReq ActionRequest
	if err := json.NewDecoder(r.Body).Decode(&actionReq); err != nil {
		utils.SendResponse(w, "error", "Invalid request payload", nil, http.StatusBadRequest)
		return
	}

	switch actionReq.Action {
	case "buy_property":
		handleBuyProperty(world, actionReq.Payload, w)
	case "upgrade_property":
		handleUpgradeProperty(world, actionReq.Payload, w)
	case "sell_property":
		handleSellProperty(world, actionReq.Payload, w)
	case "control_time":
		handleControlTime(world, actionReq.Payload, w)
	default:
		utils.SendResponse(w, "error", "Unknown action", nil, http.StatusBadRequest)
	}
}

// handleControlTime handles the "control_time" action.
func handleControlTime(world *ecs.World, payload interface{}, w http.ResponseWriter) {
	data, ok := payload.(map[string]interface{})
	if !ok {
		utils.SendResponse(w, "error", "Invalid payload structure", nil, http.StatusBadRequest)
		return
	}

	action, ok := data["action"].(string)
	if !ok {
		utils.SendResponse(w, "error", "Missing or invalid action", nil, http.StatusBadRequest)
		return
	}

	speedMultiplier, _ := data["speed_multiplier"].(float64)

	gameTime, _ := utils.GetCurrentGameTime(world)

	if gameTime != nil {
		switch action {
		case "pause":
			gameTime.IsPaused = true
		case "start":
			gameTime.IsPaused = false
		case "set_speed":
			if speedMultiplier > 0 {
				gameTime.SpeedMultiplier = speedMultiplier
			}
		default:
			utils.SendResponse(w, "error", "Invalid control action", nil, http.StatusBadRequest)
			return
		}
		utils.SendResponse(w, "success", "Time control action performed successfully", gameTime, http.StatusOK)
		return
	} else {
		utils.SendResponse(w, "error", "Game time component not found", nil, http.StatusNotFound)

	}

}

// handleBuyProperty handles the "buy_property" action.
func handleBuyProperty(world *ecs.World, payload interface{}, w http.ResponseWriter) {
	data, ok := payload.(map[string]interface{})
	if !ok {
		utils.SendResponse(w, "error", "Invalid payload structure", nil, http.StatusBadRequest)
		return
	}

	propertyID := int(data["property_id"].(float64))
	playerID := int(data["player_id"].(float64))

	playerEntity, playerFound := world.Entities[playerID]
	propertyEntity, propertyFound := world.Entities[propertyID]
	gameTime, gameTimeFoundErr := utils.GetCurrentGameTime(world)

	if !playerFound || !propertyFound {
		utils.SendResponse(w, "error", "Player or Property not found", nil, http.StatusNotFound)
		return
	}

	if gameTimeFoundErr != nil {
		utils.SendResponse(w, "error", "Unable to fetch game time", nil, http.StatusNotFound)
		return
	}

	playerComp := playerEntity.GetComponent("PlayerComponent").(*components.PlayerComponent)
	propertyComp := propertyEntity.GetComponent("PropertyComponent").(*components.PropertyComponent)

	if playerComp.Player.Funds >= propertyComp.Property.Price {
		playerComp.Player.Funds -= propertyComp.Property.Price
		propertyComp.Property.Owned = true
		propertyComp.Property.PlayerID = playerID
		propertyComp.Property.ProrateRent = true
		propertyComp.Property.PurchaseDate = gameTime.CurrentDate

		utils.SendResponse(w, "success", "Property purchased successfully", world, http.StatusOK)
	} else {
		utils.SendResponse(w, "error", "Insufficient funds", nil, http.StatusForbidden)
	}
}

// handleUpgradeProperty handles the "upgrade_property" action.// handleUpgradeProperty handles the "upgrade_property" action.
func handleUpgradeProperty(world *ecs.World, payload interface{}, w http.ResponseWriter) {
	data, ok := payload.(map[string]interface{})
	if !ok {
		utils.SendResponse(w, "error", "Invalid payload structure", nil, http.StatusBadRequest)
		return
	}

	// Extract property ID and upgrade path name from payload
	propertyIDFloat, exists := data["property_id"]
	if !exists {
		utils.SendResponse(w, "error", "property_id is required", nil, http.StatusBadRequest)
		return
	}
	propertyID := int(propertyIDFloat.(float64))

	pathName, pathExists := data["path_name"].(string)
	if !pathExists {
		utils.SendResponse(w, "error", "path_name is required", nil, http.StatusBadRequest)
		return
	}

	// Retrieve the property entity
	propertyEntity, propertyFound := world.Entities[propertyID]
	if !propertyFound {
		utils.SendResponse(w, "error", "Property not found", nil, http.StatusNotFound)
		return
	}

	// Get the PropertyComponent
	propComp, compExists := propertyEntity.GetComponent("PropertyComponent").(*components.PropertyComponent)
	if !compExists || propComp == nil {
		utils.SendResponse(w, "error", "PropertyComponent missing or invalid", nil, http.StatusInternalServerError)
		return
	}

	// Validate the upgrade path
	upgradePath, pathValid := propComp.Property.UpgradePaths[pathName]
	if !pathValid {
		utils.SendResponse(w, "error", "Invalid upgrade path", nil, http.StatusBadRequest)
		return
	}

	currentLevel := propComp.Property.UpgradeLevel

	// Check if the current level is below the maximum for the upgrade path
	if currentLevel >= len(upgradePath)-1 {
		utils.SendResponse(w, "error", "Max upgrade level reached in this path", nil, http.StatusForbidden)
		return
	}

	// Retrieve the next upgrade details
	nextUpgrade := upgradePath[currentLevel+1]

	// Retrieve the owner entity
	ownerEntity, ownerFound := world.Entities[propComp.Property.PlayerID]
	if !ownerFound {
		utils.SendResponse(w, "error", "Owner not found", nil, http.StatusNotFound)
		return
	}

	// Get the PlayerComponent
	ownerComp, ownerCompExists := ownerEntity.GetComponent("PlayerComponent").(*components.PlayerComponent)
	if !ownerCompExists || ownerComp == nil {
		utils.SendResponse(w, "error", "PlayerComponent missing or invalid", nil, http.StatusInternalServerError)
		return
	}

	// Check if the owner has sufficient funds
	if ownerComp.Player.Funds < nextUpgrade.Cost {
		utils.SendResponse(w, "error", "Insufficient funds for upgrade", nil, http.StatusForbidden)
		return
	}

	// Deduct the upgrade cost
	ownerComp.Player.Funds -= nextUpgrade.Cost

	// Get current game time
	gameTime, err := utils.GetCurrentGameTime(world)
	if err != nil {
		utils.SendResponse(w, "error", "Failed to retrieve game time", nil, http.StatusInternalServerError)
		return
	}

	// Set the PurchaseDate to current game time
	purchaseDate := gameTime.CurrentDate

	// Create a new Upgrade instance with PurchaseDate
	newUpgrade := models.Upgrade{
		Name:           nextUpgrade.Name,
		ID:             nextUpgrade.ID,
		Level:          currentLevel + 1,
		Cost:           nextUpgrade.Cost,
		RentIncrease:   nextUpgrade.RentIncrease,
		DaysToComplete: nextUpgrade.DaysToComplete,
		PurchaseDate:   purchaseDate,
		Prerequisite:   getPrerequisiteUpgrade(propComp.Property, currentLevel),
	}

	// Append the new upgrade to the Upgrades slice
	propComp.Property.Upgrades = append(propComp.Property.Upgrades, newUpgrade)

	// Increment the UpgradeLevel
	propComp.Property.UpgradeLevel++

	// Optionally, handle concurrency or lock the property during upgrade
	// For example, prevent further upgrades until this one completes
	// This depends on your game design requirements

	// Send success response
	responseData := map[string]interface{}{
		"property_id":      propertyID,
		"upgrade_level":    propComp.Property.UpgradeLevel,
		"purchase_date":    purchaseDate.Format("2006-01-02"),
		"rent_increase":    nextUpgrade.RentIncrease,
		"days_to_complete": nextUpgrade.DaysToComplete,
	}
	utils.SendResponse(w, "success", "Property upgraded successfully", responseData, http.StatusOK)
}

// getPrerequisiteUpgrade retrieves the prerequisite upgrade based on the current level.
func getPrerequisiteUpgrade(property *models.Property, currentLevel int) *models.Upgrade {
	if currentLevel == 0 {
		return nil // No prerequisite for first upgrade
	}
	if currentLevel > len(property.Upgrades)-1 {
		return nil // Inconsistent data
	}
	prereq := property.Upgrades[currentLevel-1]
	return &prereq
}

// handleSellProperty handles the "sell_property" action.
func handleSellProperty(world *ecs.World, payload interface{}, w http.ResponseWriter) {
	data, ok := payload.(map[string]interface{})
	if !ok {
		utils.SendResponse(w, "error", "Invalid payload structure", nil, http.StatusBadRequest)
		return
	}

	propertyID := int(data["property_id"].(float64))
	propertyEntity, propertyFound := world.Entities[propertyID]

	if !propertyFound {
		utils.SendResponse(w, "error", "Property not found", nil, http.StatusNotFound)
		return
	}

	propertyComp := propertyEntity.GetComponent("PropertyComponent").(*components.PropertyComponent)
	ownerEntity, ownerFound := world.Entities[propertyComp.Property.PlayerID]

	if ownerFound && propertyComp.Property.Owned {
		ownerComp := ownerEntity.GetComponent("PlayerComponent").(*components.PlayerComponent)
		salePrice := propertyComp.Property.Price * 0.8
		ownerComp.Player.Funds += salePrice
		propertyComp.Property.Owned = false
		propertyComp.Property.PlayerID = 0

		utils.SendResponse(w, "success", "Property sold successfully", world, http.StatusOK)
	} else {
		utils.SendResponse(w, "error", "Property is not owned or owner not found", nil, http.StatusForbidden)
	}
}
