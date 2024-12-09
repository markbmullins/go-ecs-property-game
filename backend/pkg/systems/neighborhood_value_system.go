package systems

import (
	"fmt"

	"github.com/markbmullins/city-developer/pkg/components"
	"github.com/markbmullins/city-developer/pkg/ecs"
	"github.com/markbmullins/city-developer/pkg/utils"
)

type NeighborhoodValueSystem struct{}

// Update calculates the neighborhood value and applies rent boosts if applicable.
func (s *NeighborhoodValueSystem) Update(world *ecs.World) {
	neighborhoods := utils.GetNeighorhoodEntities(world)
	for _, neighborhood := range neighborhoods {
		totalValue := 0.0
		upgradedCount := 0
		totalProperties := len(neighborhood.PropertyIDs)

		// Calculate the total neighborhood value
		for _, propID := range neighborhood.PropertyIDs {
			propertyEntity := world.GetProperty(propID)
			if propertyEntity == nil {
				continue
			}

			property := propertyEntity.GetComponent("Property").(*components.Property)
			totalValue += property.Price

			if len(property.Upgrades) > 0 {
				upgradedCount++
			}
		}

		neighborhood.AveragePropertyValue = totalValue / float64(totalProperties)
		upgradedPercentage := float64(upgradedCount) / float64(totalProperties) * 100

		if upgradedPercentage > neighborhood.RentBoostThreshold {
			for _, propID := range neighborhood.PropertyIDs {
				propertyEntity := world.GetProperty(propID)
				if propertyEntity != nil {
					continue
				}

				property := propertyEntity.GetComponent("Property").(*components.Property)
				property.RentBoost = (neighborhood.RentBoostPercent / 100) * property.BaseRent
				fmt.Printf("Applied rent boost of %.2f%% to property %s in neighborhood %s\n",
					neighborhood.RentBoostPercent, property.Name, neighborhood.Name)
			}
		}
	}
}
