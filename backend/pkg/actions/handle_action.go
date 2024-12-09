package actions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
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
		utils.SendResponse(w, http.StatusMethodNotAllowed, "Invalid request method", nil)
		return
	}

	var actionReq ActionRequest
	if err := json.NewDecoder(r.Body).Decode(&actionReq); err != nil {
		utils.SendResponse(w, http.StatusBadRequest, "Invalid request payload", nil)
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
		utils.SendResponse(w, http.StatusBadRequest, "Unknown action", nil)
	}
}

func handleControlTime(world *ecs.World, payload interface{}, w http.ResponseWriter) {
	data, ok := payload.(map[string]interface{})
	if !ok {
		utils.SendResponse(w, http.StatusBadRequest, "Invalid payload structure", nil)
		return
	}

	action, ok := data["action"].(string)
	if !ok {
		utils.SendResponse(w, http.StatusBadRequest, "Missing or invalid action", nil)
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
			utils.SendResponse(w, http.StatusBadRequest, "Invalid control action", nil)
			return
		}
		utils.SendResponse(w, http.StatusOK, "Time control action performed successfully", gameTime)
		return
	} else {
		utils.SendResponse(w, http.StatusNotFound, "Game time component not found", nil)

	}

}

func handleBuyProperty(world *ecs.World, data BuyPropertyPayload, w http.ResponseWriter) {
	log.Printf("handleBuyProperty called with data: %+v\n", data)
	propertyID := data.PropertyID
	playerID := data.PlayerID

	playerEntity := world.GetPlayer(playerID)
	propertyEntity := world.GetProperty(propertyID)

	playerFound := playerEntity != nil
	propertyFound := propertyEntity != nil
	gameTime := world.GetGameTime().GetComponent("GameTime").(*components.GameTime)

	if !playerFound || !propertyFound {
		utils.SendResponse(w, http.StatusBadRequest, "Player or Property not found", nil)
		return
	}

	player := playerEntity.GetComponent("Player").(*components.Player)
	property := propertyEntity.GetComponent("Property").(*components.Property)

	log.Printf("Player funds: %f, Property price: %f\n", player.Funds, property.Price)
	if player.Funds >= property.Price {
		player.Funds -= property.Price
		property.Owned = true
		property.PlayerID = playerID
		property.PurchaseDate = gameTime.CurrentDate

		// Append the property to the player's list of properties
		player.Properties = append(player.Properties, property)

		log.Printf("player after purchase: %+v\n property after purchase: %+v", player, property)
		log.Printf("player from world: %+v\n property from world: %+v", world.GetPlayer(playerID), world.GetProperty(propertyID))
		utils.SendResponse(w, http.StatusOK, "Property purchased successfully", world)
	} else {
		utils.SendResponse(w, http.StatusBadRequest, "Insufficient funds", nil)
	}
}

func handleUpgradeProperty(world *ecs.World, data UpgradePropertyPayload, w http.ResponseWriter) {
	propertyID := data.PropertyID
	pathName := data.PathName

	// Retrieve the property entity
	propertyEntity := world.GetProperty(propertyID)
	propertyFound := propertyEntity != nil
	if !propertyFound {
		utils.SendResponse(w, http.StatusNotFound, "Property not found", nil)
		return
	}

	// Get the Property
	property, ok := propertyEntity.GetComponent("Property").(*components.Property)
	if !ok || property == nil {
		utils.SendResponse(w, http.StatusInternalServerError, "Property missing or invalid", nil)
		return
	}

	// Validate the upgrade path
	upgradePath, pathValid := property.UpgradePaths[pathName]
	if !pathValid {
		utils.SendResponse(w, http.StatusBadRequest, "Invalid upgrade path", nil)
		return
	}

	currentLevel := len(property.Upgrades)

	// Check if the current level is below the maximum for the upgrade path
	if currentLevel >= len(upgradePath)-1 {
		utils.SendResponse(w, http.StatusBadRequest, "Max upgrade level reached in this path", nil)
		return
	}

	// Retrieve the next upgrade details
	nextUpgrade := upgradePath[currentLevel+1]

	// Retrieve the owner entity
	ownerEntity := world.GetPlayer(property.PlayerID)
	ownerFound := ownerEntity != nil
	if !ownerFound {
		utils.SendResponse(w, http.StatusNotFound, "Owner not found", nil)
		return
	}

	// Get the Player
	player, playerExists := ownerEntity.GetComponent("Player").(*components.Player)
	if !playerExists {
		utils.SendResponse(w, http.StatusInternalServerError, "Player missing or invalid", nil)
		return
	}

	// Check if the owner has sufficient funds
	if player.Funds < nextUpgrade.Cost {
		utils.SendResponse(w, http.StatusInternalServerError, "Insufficient funds for upgrade", nil)
		return
	}

	// Deduct the upgrade cost
	player.Funds -= nextUpgrade.Cost

	// Get current game time
	gameTime, err := utils.GetCurrentGameTime(world)
	if err != nil {
		utils.SendResponse(w, http.StatusInternalServerError, "Failed to retrieve game time", nil)
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
	utils.SendResponse(w, http.StatusOK, "Property upgraded successfully", responseData)
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
		utils.SendResponse(w, http.StatusBadRequest, "Property not found", nil)
		return
	}

	property := propertyEntity.GetComponent("Property").(*components.Property)
	ownerEntity := world.GetPlayer(property.PlayerID)

	ownerFound := ownerEntity != nil
	if ownerFound && property.Owned {
		player := ownerEntity.GetComponent("Player").(*components.Player)
		salePrice := property.Price * 0.8
		player.Funds += salePrice
		
		// Remove the property from the player's Properties array
		var newProperties []*components.Property
		for _, p := range player.Properties {
			if p.ID != propertyID {
				newProperties = append(newProperties, p)
			}
		}
		player.Properties = newProperties

		property.Owned = false
		property.PlayerID = 0

		utils.SendResponse(w, http.StatusOK, "Property sold successfully", world)
	} else {
		utils.SendResponse(w, http.StatusBadRequest, "Property is not owned or owner not found", nil)
	}
}

func decodePayload(input interface{}, target interface{}, w http.ResponseWriter) bool {
	// Convert the interface{} to JSON bytes
	jsonData, err := json.Marshal(input)
	if err != nil {
		utils.SendResponse(w, http.StatusBadRequest, "Failed to process payload", nil)
		return false
	}

	// Decode the JSON bytes into the target struct
	if err := json.NewDecoder(bytes.NewReader(jsonData)).Decode(target); err != nil {
		utils.SendResponse(w, http.StatusBadRequest, fmt.Sprintf("Invalid payload structure: %v", err), nil)
		return false
	}
	return true
}
