package systems

import (
	"fmt"

	"github.com/markbmullins/city-developer/pkg/components"
	"github.com/markbmullins/city-developer/pkg/ecs"
	"github.com/markbmullins/city-developer/pkg/models"
)

type NeighborhoodValueSystem struct {
	Neighborhoods map[int]*models.Neighborhood
}

// Update calculates the neighborhood value and applies rent boosts if applicable.
func (s *NeighborhoodValueSystem) Update(world *ecs.World) {
	for _, neighborhood := range s.Neighborhoods {
		totalValue := 0.0
		upgradedCount := 0
		totalProperties := len(neighborhood.PropertyIDs)

		// Calculate the total neighborhood value
		for _, propID := range neighborhood.PropertyIDs {
			propertyEntity, found := world.Entities[propID]
			if !found {
				continue
			}

			propComp := propertyEntity.GetComponent("PropertyComponent").(*components.PropertyComponent)
			property := propComp.Property
			totalValue += property.Price

			if len(property.Upgrades) > 0 {
				upgradedCount++
			}
		}

		neighborhood.AveragePropertyValue = totalValue / float64(totalProperties)
		upgradedPercentage := float64(upgradedCount) / float64(totalProperties) * 100

		if upgradedPercentage > neighborhood.RentBoostThreshold {
			for _, propID := range neighborhood.PropertyIDs {
				propertyEntity, found := world.Entities[propID]
				if !found {
					continue
				}

				propComp := propertyEntity.GetComponent("PropertyComponent").(*components.PropertyComponent)
				property := propComp.Property
				property.RentBoost = (neighborhood.RentBoostPercent / 100) * property.BaseRent
				fmt.Printf("Applied rent boost of %.2f%% to property %s in neighborhood %s\n",
					neighborhood.RentBoostPercent, property.Name, neighborhood.Name)
			}
		}
	}
}
