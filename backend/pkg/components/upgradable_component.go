package components

import "time"

type Upgrade struct {
	Name           string
	Level          int // TODO: Do I need this field?
	Cost           float64
	RentIncrease   float64 // The amount of rent increase the upgrade provides to the property
	DaysToComplete int
	PurchaseDate   time.Time
	Prerequisite   *Upgrade
	Applied        bool // Tracks if upgrade has been applied to property
}

type Upgradable struct {
	// A map of upgrade paths or just a list of all possible upgrades
	PossibleUpgrades map[string][]*Upgrade
	AppliedUpgrades  []*Upgrade
}

func (upgradable *Upgradable) CurrentUpgradeLevel(pathName string) int {
	pathUpgrades, exists := upgradable.PossibleUpgrades[pathName]
	if !exists {
		// Path does not exist, level is 0
		return 0
	}

	// Create a quick lookup for applied upgrades to improve efficiency
	appliedSet := make(map[*Upgrade]bool, len(upgradable.AppliedUpgrades))
	for _, applied := range upgradable.AppliedUpgrades {
		appliedSet[applied] = true
	}

	level := 0
	// Count how many from this path are in the applied set
	for _, upgrade := range pathUpgrades {
		if appliedSet[upgrade] {
			level++
		}
	}

	return level
}

func (upgradable *Upgradable) MaxUpgradeLevel() int {
	maxLevel := 0
	for pathName := range upgradable.PossibleUpgrades {
		level := upgradable.CurrentUpgradeLevel(pathName)
		if level > maxLevel {
			maxLevel = level
		}
	}
	return maxLevel
}
