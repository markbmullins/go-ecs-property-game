package actions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/markbmullins/city-developer/pkg/components"
	"github.com/markbmullins/city-developer/pkg/ecs"
	"github.com/markbmullins/city-developer/pkg/utils"
)

type ControlTimePayload struct {
	Action          string  `json:"action"`
	SpeedMultiplier float64 `json:"speed_multiplier,omitempty"`
}

type BuyPropertyPayload struct {
	PropertyID int `json:"property_id"`
	PlayerID   int `json:"player_id"`
}

type UpgradePropertyPayload struct {
	PropertyID int    `json:"property_id"`
	PathName   string `json:"path_name"`
}

type SellPropertyPayload struct {
	PropertyID int `json:"property_id"`
}

type ActionRequest struct {
	Action  string      `json:"action"`
	Payload interface{} `json:"payload"`
}

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
		var payload BuyPropertyPayload
		if !decodePayload(actionReq.Payload, &payload, w) {
			return
		}
		handleBuyProperty(world, payload, w)
	case "upgrade_property":
		var payload UpgradePropertyPayload
		if !decodePayload(actionReq.Payload, &payload, w) {
			return
		}
		handleUpgradeProperty(world, payload, w)
	case "sell_property":
		var payload SellPropertyPayload
		if !decodePayload(actionReq.Payload, &payload, w) {
			return
		}
		handleSellProperty(world, payload, w)
	case "control_time":
		var payload ControlTimePayload
		if !decodePayload(actionReq.Payload, &payload, w) {
			return
		}
		handleControlTime(world, payload, w)
	default:
		utils.SendResponse(w, "error", "Unknown action", nil, http.StatusBadRequest)
	}
}

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

func handleBuyProperty(world *ecs.World, data BuyPropertyPayload, w http.ResponseWriter) {

	propertyID := data.PropertyID
	playerID := data.PlayerID

	playerEntity := world.GetPlayer(playerID)
	propertyEntity := world.GetProperty(propertyID)

	playerFound := playerEntity != nil
	propertyFound := propertyEntity != nil
	gameTime := world.GetGameTime().GetComponent("GameTime").(*components.GameTime)

	if !playerFound || !propertyFound {
		utils.SendResponse(w, "error", "Player or Property not found", nil, http.StatusNotFound)
		return
	}

	player := playerEntity.GetComponent("Player").(*components.Player)
	property := propertyEntity.GetComponent("Property").(*components.Property)

	if player.Funds >= property.Price {
		player.Funds -= property.Price
		property.Owned = true
		property.PlayerID = playerID
		property.PurchaseDate = gameTime.CurrentDate

		utils.SendResponse(w, "success", "Property purchased successfully", world, http.StatusOK)
	} else {
		utils.SendResponse(w, "error", "Insufficient funds", nil, http.StatusForbidden)
	}
}

func handleUpgradeProperty(world *ecs.World, data UpgradePropertyPayload, w http.ResponseWriter) {
	propertyID := data.PropertyID
	pathName := data.PathName

	// Retrieve the property entity
	propertyEntity := world.GetProperty(propertyID)
	propertyFound := propertyEntity != nil
	if !propertyFound {
		utils.SendResponse(w, "error", "Property not found", nil, http.StatusNotFound)
		return
	}

	// Get the Property
	property, propertyExists := propertyEntity.GetComponent("Property").(*components.Property)
	if !propertyExists || property == nil {
		utils.SendResponse(w, "error", "Property missing or invalid", nil, http.StatusInternalServerError)
		return
	}

	// Validate the upgrade path
	upgradePath, pathValid := property.UpgradePaths[pathName]
	if !pathValid {
		utils.SendResponse(w, "error", "Invalid upgrade path", nil, http.StatusBadRequest)
		return
	}

	currentLevel := len(property.Upgrades)

	// Check if the current level is below the maximum for the upgrade path
	if currentLevel >= len(upgradePath)-1 {
		utils.SendResponse(w, "error", "Max upgrade level reached in this path", nil, http.StatusForbidden)
		return
	}

	// Retrieve the next upgrade details
	nextUpgrade := upgradePath[currentLevel+1]

	// Retrieve the owner entity
	ownerEntity := world.GetPlayer(property.PlayerID)
	ownerFound := ownerEntity != nil
	if !ownerFound {
		utils.SendResponse(w, "error", "Owner not found", nil, http.StatusNotFound)
		return
	}

	// Get the Player
	player, playerExists := ownerEntity.GetComponent("Player").(*components.Player)
	if !playerExists || player == nil {
		utils.SendResponse(w, "error", "Player missing or invalid", nil, http.StatusInternalServerError)
		return
	}

	// Check if the owner has sufficient funds
	if player.Funds < nextUpgrade.Cost {
		utils.SendResponse(w, "error", "Insufficient funds for upgrade", nil, http.StatusForbidden)
		return
	}

	// Deduct the upgrade cost
	player.Funds -= nextUpgrade.Cost

	// Get current game time
	gameTime, err := utils.GetCurrentGameTime(world)
	if err != nil {
		utils.SendResponse(w, "error", "Failed to retrieve game time", nil, http.StatusInternalServerError)
		return
	}

	// Set the PurchaseDate to current game time
	purchaseDate := gameTime.CurrentDate

	// Create a new Upgrade instance with PurchaseDate
	newUpgrade := components.Upgrade{
		Name:           nextUpgrade.Name,
		ID:             nextUpgrade.ID,
		Level:          currentLevel + 1,
		Cost:           nextUpgrade.Cost,
		RentIncrease:   nextUpgrade.RentIncrease,
		DaysToComplete: nextUpgrade.DaysToComplete,
		PurchaseDate:   purchaseDate,
		Prerequisite:   getPrerequisiteUpgrade(property, currentLevel),
	}

	// Append the new upgrade to the Upgrades slice
	property.Upgrades = append(property.Upgrades, newUpgrade)

	// Optionally, handle concurrency or lock the property during upgrade
	// For example, prevent further upgrades until this one completes
	// This depends on your game design requirements

	// Send success response
	responseData := map[string]interface{}{
		"property_id":      propertyID,
		"upgrade_level":    len(property.Upgrades),
		"purchase_date":    purchaseDate.Format("2006-01-02"),
		"rent_increase":    nextUpgrade.RentIncrease,
		"days_to_complete": nextUpgrade.DaysToComplete,
	}
	utils.SendResponse(w, "success", "Property upgraded successfully", responseData, http.StatusOK)
}

func getPrerequisiteUpgrade(property *components.Property, currentLevel int) *components.Upgrade {
	if currentLevel == 0 {
		return nil // No prerequisite for first upgrade
	}
	if currentLevel > len(property.Upgrades)-1 {
		return nil // Inconsistent data
	}
	prereq := property.Upgrades[currentLevel-1]
	return &prereq
}

func handleSellProperty(world *ecs.World, data SellPropertyPayload, w http.ResponseWriter) {
	propertyID := data.PropertyID
	propertyEntity := world.GetProperty(propertyID)

	propertyFound := propertyEntity != nil
	if !propertyFound {
		utils.SendResponse(w, "error", "Property not found", nil, http.StatusNotFound)
		return
	}

	property := propertyEntity.GetComponent("Property").(*components.Property)
	ownerEntity := world.GetPlayer(property.PlayerID)

	ownerFound := ownerEntity != nil
	if ownerFound && property.Owned {
		player := ownerEntity.GetComponent("Player").(*components.Player)
		salePrice := property.Price * 0.8
		player.Funds += salePrice
		property.Owned = false
		property.PlayerID = 0

		utils.SendResponse(w, "success", "Property sold successfully", world, http.StatusOK)
	} else {
		utils.SendResponse(w, "error", "Property is not owned or owner not found", nil, http.StatusForbidden)
	}
}

func decodePayload(input interface{}, target interface{}, w http.ResponseWriter) bool {
	// Convert the interface{} to JSON bytes
	jsonData, err := json.Marshal(input)
	if err != nil {
		utils.SendResponse(w, "error", "Failed to process payload", nil, http.StatusBadRequest)
		return false
	}

	// Decode the JSON bytes into the target struct
	if err := json.NewDecoder(bytes.NewReader(jsonData)).Decode(target); err != nil {
		utils.SendResponse(w, "error", fmt.Sprintf("Invalid payload structure: %v", err), nil, http.StatusBadRequest)
		return false
	}
	return true
}
