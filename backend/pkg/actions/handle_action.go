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

	gameTime, _ := world.GetCurrentGameTime()

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
	gameTime, _ := world.GetCurrentGameTime()

	if !playerFound || !propertyFound {
		utils.SendResponse(w, http.StatusBadRequest, "Player or Property not found", nil)
		return
	}
	funds, _ := ecs.GetComponent[components.Funds](playerEntity)
	purchaseable, _ := ecs.GetComponent[components.Purchaseable](propertyEntity)
	ownable, _ := ecs.GetComponent[components.Ownable](propertyEntity)

	log.Printf("Player funds: %f, Property price: %f\n", funds.Amount, purchaseable.Cost)
	if funds.Amount >= purchaseable.Cost {
		funds.Amount -= purchaseable.Cost
		ownable.Owned = true
		ownable.OwnerID = playerID
		purchaseable.PurchaseDate = gameTime.CurrentDate

		// Append the property to the player's list of properties
		world.BuyProperty(propertyID, playerID)
		utils.SendResponse(w, http.StatusOK, "Property purchased successfully", world)
	} else {
		utils.SendResponse(w, http.StatusBadRequest, "Insufficient funds", nil)
	}
}

func handleUpgradeProperty(world *ecs.World, data UpgradePropertyPayload, w http.ResponseWriter) {
	propertyID := data.PropertyID
	upgradePathName := data.PathName

	// Retrieve the property entity
	propertyEntity := world.GetProperty(propertyID)
	propertyFound := propertyEntity != nil
	if !propertyFound {
		utils.SendResponse(w, http.StatusNotFound, "Property not found", nil)
		return
	}

	var ownable, _ = ecs.GetComponent[components.Ownable](propertyEntity)
	if !ownable.Owned {
		utils.SendResponse(w, http.StatusBadRequest, "Property is not owned", nil)
		return
	}

	upgradable, _ := ecs.GetComponent[components.Upgradable](propertyEntity)
	if upgradable == nil {
		utils.SendResponse(w, http.StatusBadRequest, "Property is not upgradable", nil)
		return
	}

	upgradePath, exists := upgradable.PossibleUpgrades[upgradePathName]
	if !exists || len(upgradePath) <= len(upgradable.AppliedUpgrades) {
		utils.SendResponse(w, http.StatusBadRequest, "Invalid upgrade path or max level reached", nil)
		return
	}

	currentLevel := len(upgradePath)

	// Check if the current level is below the maximum for the upgrade path
	if currentLevel >= len(upgradePath)-1 {
		utils.SendResponse(w, http.StatusBadRequest, "Max upgrade level reached in this path", nil)
		return
	}

	// Retrieve the next upgrade details
	nextUpgrade := upgradePath[currentLevel+1]

	playerEntity := world.GetPlayer(ownable.OwnerID)
	playerFunds, _ := ecs.GetComponent[components.Funds](playerEntity)

	// Deduct the upgrade cost
	playerFunds.Amount -= nextUpgrade.Cost

	// Get current game time
	gameTime, _ := world.GetCurrentGameTime()

	// Set the PurchaseDate to current game time
	purchaseDate := gameTime.CurrentDate

	// Create a new Upgrade instance with PurchaseDate
	newUpgrade := components.Upgrade{
		Name:           nextUpgrade.Name,
		Level:          currentLevel + 1,
		Cost:           nextUpgrade.Cost,
		RentIncrease:   nextUpgrade.RentIncrease,
		DaysToComplete: nextUpgrade.DaysToComplete,
		PurchaseDate:   purchaseDate,
		Prerequisite:   getPrerequisiteUpgrade(propertyEntity, upgradePathName),
	}

	// Append the new upgrade to the Upgrades slice
	upgradable.AppliedUpgrades = append(upgradable.AppliedUpgrades, &newUpgrade)

	// Optionally, handle concurrency or lock the property during upgrade
	// For example, prevent further upgrades until this one completes

	// Send success response
	responseData := map[string]interface{}{
		"property_id":      propertyID,
		"upgrade_level":    upgradable.MaxUpgradeLevel(),
		"purchase_date":    purchaseDate.Format("2006-01-02"),
		"rent_increase":    nextUpgrade.RentIncrease,
		"days_to_complete": nextUpgrade.DaysToComplete,
	}
	utils.SendResponse(w, http.StatusOK, "Property upgraded successfully", responseData)
}

func getPrerequisiteUpgrade(property *ecs.Entity, pathName string) *components.Upgrade {
	var upgradable, _ = ecs.GetComponent[components.Upgradable](property)
	var currentLevel = upgradable.CurrentUpgradeLevel(pathName)
	if currentLevel == 0 {
		return nil // No prerequisite for first upgrade
	}
	if currentLevel > len(upgradable.PossibleUpgrades[pathName])-1 {
		return nil // Inconsistent data
	}
	prereq := upgradable.PossibleUpgrades[pathName][currentLevel-1]
	return prereq
}

func handleSellProperty(world *ecs.World, data SellPropertyPayload, w http.ResponseWriter) {
	propertyID := data.PropertyID
	propertyEntity := world.GetProperty(propertyID)

	propertyFound := propertyEntity != nil
	if !propertyFound {
		utils.SendResponse(w, http.StatusBadRequest, "Property not found", nil)
		return
	}

	var ownable, owned = ecs.GetComponent[components.Ownable](propertyEntity)
	if !owned {
		utils.SendResponse(w, http.StatusBadRequest, "Property is not owned", nil)
		return
	}
	ownerEntity := world.GetPlayer(ownable.OwnerID)

	if ownerEntity == nil {
		utils.SendResponse(w, http.StatusBadRequest, "Owner not found", nil)
		return
	}

	var purchaseable, _ = ecs.GetComponent[components.Purchaseable](propertyEntity)
	salePrice := purchaseable.Cost * 0.8
	var fundsComponent, _ = ecs.GetComponent[components.Funds](ownerEntity)
	fundsComponent.Amount += salePrice

	// Remove the property from the player's Properties array
	world.SellProperty(propertyID)
	ownable.Owned = false
	ownable.OwnerID = 0
	utils.SendResponse(w, http.StatusOK, "Property sold successfully", world)
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
