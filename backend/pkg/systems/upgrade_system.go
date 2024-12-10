package systems

import (
	"github.com/markbmullins/city-developer/pkg/components"
	"github.com/markbmullins/city-developer/pkg/ecs"
)

type UpgradeSystem struct{}

func (s *UpgradeSystem) Update(world *ecs.World) {
	var gameTime, _ = world.GetCurrentGameTime()
	var upgradableEntities = world.QueryByComponent("Upgradeable")

	// Checks if upgrades have completed and applies them
	for _, property := range upgradableEntities {
		var ownableComponent, _ = ecs.GetComponent[components.Ownable](property)
		var upgradableComponent, _ = ecs.GetComponent[components.Upgradable](property)
		var upgrades = upgradableComponent.AppliedUpgrades
		if ownableComponent.Owned && len(upgrades) > 0 {
			for _, upgrade := range upgrades {
				completionDate := upgrade.PurchaseDate.AddDate(0, 0, upgrade.DaysToComplete)
				if completionDate.Before(gameTime.CurrentDate) && !upgrade.Applied {
					world.ApplyUpgradeToProperty(property, upgrade)
				}
			}
		}
	}
}
