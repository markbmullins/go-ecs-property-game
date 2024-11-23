package systems

import (
	"fmt"
	"time"

	"github.com/markbmullins/city-developer/pkg/components"
	"github.com/markbmullins/city-developer/pkg/ecs"
	"github.com/markbmullins/city-developer/pkg/models"
)

type UpgradeSystem struct{}

func (s *UpgradeSystem) Update(world *ecs.World) {
	// Loop over each property entity and check for upgrades
	for _, entity := range world.Entities {
		propComp := entity.GetComponent("PropertyComponent")
		if propComp != nil {
			property := propComp.(*components.PropertyComponent).Property
			if property.Owned && len(property.Upgrades) > 0 {
				handleUpgrades(world, property)
			}
		}
	}
}

// handleUpgrades applies rent increases based on upgrades and their completion dates.
func handleUpgrades(world *ecs.World, property *models.Property) {
	for _, upgrade := range property.Upgrades {
		completionDate := upgrade.PurchaseDate.AddDate(0, 0, upgrade.DaysToComplete)
		if completionDate.Before(time.Now()) && !upgrade.Applied {
			applyUpgradeToProperty(world, property, &upgrade)
		}
	}
}

// applyUpgradeToProperty updates the rent based on the upgrade and marks it as applied.
func applyUpgradeToProperty(world *ecs.World, property *models.Property, upgrade *models.Upgrade) {
	property.RentBoost += upgrade.RentIncrease
	upgrade.Applied = true
	fmt.Printf("Applied upgrade '%s' to property '%s' with RentIncrease: %.2f\n", upgrade.Name, property.Name, upgrade.RentIncrease)
}
