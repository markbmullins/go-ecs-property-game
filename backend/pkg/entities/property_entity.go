package entities

import (
	"time"

	"github.com/markbmullins/city-developer/pkg/components"
	"github.com/markbmullins/city-developer/pkg/ecs"
)

/** Creates a residential or commercial property entity in the game.
 * A property entity has the following components:  
 * Nameable: The name of the property.
 * Addressable: The address of the property.
 * Describable: The description of the property.
 * Classifiable: The type and subtype of the property.
 * Rentable: The base rent and rent boost of the property.
 * Purchasable: The cost to purchase the property.
 * Ownable: The owner of the property.
 * Upgradable: The possible upgrades and applied upgrades of the property.
 * Groupable: The group ID of the property.
 */
func CreateProperty(
	id int,
	name string,
	address string,
	description string,
	propertyType components.PropertyType,
	subtype components.PropertySubtype,
	baseRent float64,
	price float64,
	groupID int,
) *ecs.Entity {
	property := ecs.NewEntity("Property", id)

	ecs.AddComponent(property, &components.Nameable{Name: name})
	ecs.AddComponent(property, &components.Addressable{Address: address})
	ecs.AddComponent(property, &components.Describable{Description: description})
	ecs.AddComponent(property, &components.Classifiable{Type: propertyType, Subtype: subtype})
	ecs.AddComponent(property, &components.Rentable{BaseRent: baseRent, RentBoost: 0, LastRentCollectionDate: time.Time{}})
	ecs.AddComponent(property, &components.Purchaseable{Cost: price, PurchaseDate: time.Time{}})
	ecs.AddComponent(property, &components.Ownable{OwnerID: 0, Owned: false})
	ecs.AddComponent(property, &components.Upgradable{PossibleUpgrades: map[string][]*components.Upgrade{}, AppliedUpgrades: []*components.Upgrade{}})
	ecs.AddComponent(property, &components.Groupable{GroupID: groupID})

	return property
}

func AddUpgradesToProperty(property *ecs.Entity, upgradePaths map[string][]*components.Upgrade) {
	upgradable, exists := ecs.GetComponent[components.Upgradable](property)
	if !exists {
		// If the component doesn't exist, create it
		upgradable = &components.Upgradable{
			PossibleUpgrades: upgradePaths,
		}
	} else {
		// If it exists, modify its PossibleUpgrades field
		upgradable.PossibleUpgrades = upgradePaths
	}
	// Add or replace the component in the entity
	ecs.AddComponent(property, upgradable)
}

func CreateUpgrade(
	name string,
	level int,
	cost float64,
	rentIncrease float64,
	daysToComplete int,
	prerequisite *components.Upgrade,
) *components.Upgrade {
	return &components.Upgrade{
		Name:           name,
		Level:          level,
		Cost:           cost,
		RentIncrease:   rentIncrease,
		DaysToComplete: daysToComplete,
		PurchaseDate:   time.Time{}, // Default to zero time
		Prerequisite:   prerequisite,
		Applied:        false,       // Default to not applied
	}
}

func CreateUpgradableComponent(upgradePaths map[string][]*components.Upgrade) *components.Upgradable {
	return &components.Upgradable{
		PossibleUpgrades: upgradePaths,
		AppliedUpgrades:  []*components.Upgrade{},
	}
}
