package systems

import (
	"fmt"
	"math"
	"time"

	"github.com/markbmullins/city-developer/pkg/components"
	"github.com/markbmullins/city-developer/pkg/ecs"
	"github.com/markbmullins/city-developer/pkg/models"
	"github.com/markbmullins/city-developer/pkg/utils"
)

// Purchase Month Rent Proration:

//     Purchase Day: No rent is collected.
//     Rent Collection Starts: From the day after the purchase date.

// Upgrade Application:

//     Upgrade Completion Day: No rent increase is applied on the day the upgrade is completed.
//     Rent Increase Starts: From the day after the upgrade completion date.

// General Rules:

//     Prorated Rent Calculation: Rent is calculated based on the number of days the property (and any upgrades) are owned excluding the purchase or upgrade completion day.
//     Full Rent Collection: Once a month has fully elapsed (excluding the purchase day), full rent (including any applicable upgrades) is collected.

// - in my game, i can gave different time advancement speeds. sometimes those speeds can go higher than 30 days
// - the first month a property is purchased, the rent should be prorated from the day of purchase. if you purchase a property on the first day you get the full rent minus 1 day. if you purchase a property on the last day of the month, prorated rent is 0 and you collect full rent the next month. the day of purchase DOES NOT GENERATE RENT.
// - upgrades should increase the rent the based on their RentIncrease
// - if you upgrade a property mid month, the upgrade boost should be prorated following the same rules as properties. It takes effect on the day purchaseDate + DaysToComplete (the day when the upgrade is finished)
// - careful consideration should be made for edge cases, first and last day in a month, leap years, etc.
// Similar to how, in the first month, the purchase month, rent should not be collected on the purchase date, it kicks in the next day. similarly, in the month an upgrade was completed, upgrades should not be counted on the day the upgrade was completed. rather, the next day. So if a property was purchased in Jan, upgraded in feb, base rent would be prorated in jan, the only the upgrade bonus would be prorated in feb, then in march the player would collect the full amount. if a property is purchased and upgraded in the same month, the base rent should be prorated from day after the purchase date and the upgrade rent would be prorated from the day after the upgrade is completed, the next month the player would collect the full rent and upgrade bonus.

/*
===========================================================
                     INCOME SYSTEM REQUIREMENTS
===========================================================

1. **Purchase Mechanics**
   - **No Rent on Purchase Day:** Rent is not collected on the day a property is purchased.
   - **Rent Collection Start:** Begins the day after the purchase date.
   - **First Month Proration:**
     - *First Day Purchase:* Full rent minus one day.
     - *Last Day Purchase:* Prorated rent is 0; full rent starts the next month.

2. **Upgrade Mechanics**
   - **No Rent Increase on Completion Day:** Rent increase from upgrades does not apply on the upgrade completion day.
   - **Rent Increase Start:** Begins the day after the upgrade completion date.
   - **Mid-Month Upgrades:** Rent increases are prorated based on active days, excluding the completion day.
   - **Multiple Upgrades:** Each upgrade's rent increase is prorated based on its respective completion date.

3. **Rent Collection Rules**
   - **Prorated Rent Calculation:** Based on days owned/upgraded, excluding purchase and upgrade completion days.
   - **Full Rent Collection:** After a full month has elapsed (excluding purchase day), full rent including upgrades is collected.
   - **Rent Composition:**
     - *Base Rent:* Defined per property.
     - *Upgrade Increases:* Added based on each upgrade's RentIncrease value.
     - *Total Rent:* Sum of Base Rent and all applicable Upgrade Increases.

4. **Time Advancement Considerations**
   - **Variable Speeds:** Supports multiple time advancement speeds, including cycles exceeding 30 days.
   - **Accurate Proration:** Rent calculations adjust based on the actual number of days elapsed, regardless of time speed.

5. **Edge Case Handling**
   - **Calendar Variations:** Correctly handles months with 28-31 days and leap years (February with 29 days).
   - **Same Month Transactions:**
     - *Purchase & Upgrade in Same Month:*
       - Base rent prorated from day after purchase.
       - Upgrade rent prorated from day after upgrade completion.
       - Full rent and upgrade bonus collected the following month.
   - **Boundary Days:** Ensures accurate calculations for purchases or upgrades on the first or last day of a month.

6. **Implementation Guidelines**
   - **Day Exclusions:** Exclude Purchase Day and Upgrade Completion Day from rent generation.
   - **Accurate Day Counting:** Implement precise mechanisms to count qualifying days for proration.
   - **Testing & Validation:**
     - Develop test cases for first/last day purchases, mid-month upgrades, leap years, and variable time speeds.
     - Verify all edge cases to ensure correct rent calculations.

**Key Rules Summary:**
- No rent is collected on Purchase Day or Upgrade Completion Day.
- Rent and upgrade bonuses commence the day after their respective initiation dates.
- Prorated rents are based on days owned or upgraded, excluding initiation days.
- Full rent is collected after a complete month, inclusive of all upgrades.
- The system accommodates variable time speeds and handles calendar anomalies accurately.

===========================================================
*/

// IncomeSystem handles rent collection, including prorated rent for the first month and upgrades.
// type IncomeSystem struct{}

// Update triggers the rent collection process, handling fast-forwarding of time.
type IncomeSystem struct{}

// Update triggers the rent collection process, handling fast-forwarding of time.
func (s *IncomeSystem) Update(world *ecs.World) {
	gameTime, err := utils.GetCurrentGameTime(world)
	if err != nil {
		fmt.Println("Error: GameTime component not found")
		return
	}

	if !gameTime.IsPaused {
		monthsPassed := calculateMonthsPassed(gameTime.LastUpdated, gameTime.CurrentDate)
		if monthsPassed > 0 {
			processRent(world, gameTime.LastUpdated, monthsPassed)
			gameTime.NewMonth = true
		}
		gameTime.LastUpdated = gameTime.CurrentDate
	}
}

// processRent collects rent for all elapsed months, handling partial and full months.
func processRent(world *ecs.World, lastUpdated time.Time, monthsPassed int) {
	// location := lastUpdated.Location()
	firstOfNextMonth := nextMonthStart(lastUpdated)

	// Handle partial month
	if isPartialMonth(lastUpdated) {
		processPartialMonth(world, lastUpdated, monthEnd(lastUpdated))
		monthsPassed--
	}

	// Handle full months
	for i := 0; i < monthsPassed; i++ {
		startOfMonth := firstOfNextMonth.AddDate(0, i, 0)
		endOfMonth := monthEnd(startOfMonth)
		processFullMonth(world, startOfMonth, endOfMonth)
	}
}

// processPartialMonth collects prorated rent for a partial month.
func processPartialMonth(world *ecs.World, startDate, endDate time.Time) {
	for _, entity := range world.Entities {
		processEntityRent(world, entity, startDate, endDate, true)
	}
}

// processFullMonth collects full rent for a given month.
func processFullMonth(world *ecs.World, startOfMonth, endOfMonth time.Time) {
	for _, entity := range world.Entities {
		processEntityRent(world, entity, startOfMonth, endOfMonth, false)
	}
}

// processEntityRent calculates and distributes rent for a single entity.
func processEntityRent(world *ecs.World, entity *ecs.Entity, startDate, endDate time.Time, isPartial bool) {
	propComp := entity.GetComponent("PropertyComponent")
	if propComp == nil {
		return
	}

	property := propComp.(*components.PropertyComponent).Property
	if !property.Owned {
		return
	}

	if isPartial {
		if property.PurchaseDate.After(endDate) {
			return
		}
		collectProratedRent(world, property, startDate, endDate)
	} else if property.PurchaseDate.Before(startDate) || property.PurchaseDate.Equal(startDate) {
		collectFullRent(world, property, startDate, endDate)
	} else {
		collectProratedRent(world, property, startDate, endDate)
	}
}

// collectProratedRent calculates prorated rent for a property.
func collectProratedRent(world *ecs.World, property *models.Property, startDate, endDate time.Time) {
	daysOwned := calculateDaysOwned(startDate, endDate, property.PurchaseDate)
	proratedRent := calculateProratedRent(property, float64(daysOwned), float64(daysInMonth(startDate)))
	distributeRentToOwner(world, property, proratedRent)
}

// collectFullRent calculates and distributes full rent for a property.
func collectFullRent(world *ecs.World, property *models.Property, monthStart, monthEnd time.Time) {
	fullRent := calculateFullRent(property, monthStart, monthEnd)
	distributeRentToOwner(world, property, fullRent)
}

// calculateProratedRent computes prorated rent for a property, including upgrades.
func calculateProratedRent(property *models.Property, daysOwned, daysInMonth float64) float64 {
	baseProrated := (property.BaseRent * daysOwned) / daysInMonth
	upgradeProrated := calculateUpgradeProration(property, daysOwned, daysInMonth)
	return roundToNearest5(baseProrated + upgradeProrated)
}

// calculateFullRent computes full rent for a property, including upgrades.
func calculateFullRent(property *models.Property, monthStart, monthEnd time.Time) float64 {
	baseRent := property.BaseRent
	upgradeIncrease := calculateUpgradeRent(property, monthStart, monthEnd)
	return baseRent + upgradeIncrease
}

// calculateUpgradeRent computes the rent increase from upgrades for a given month.
func calculateUpgradeRent(property *models.Property, monthStart, monthEnd time.Time) float64 {
	total := 0.0
	for _, upgrade := range property.Upgrades {
		completionDate := upgrade.PurchaseDate.AddDate(0, 0, upgrade.DaysToComplete)
		if completionDate.Before(monthStart) {
			total += upgrade.RentIncrease
		} else if completionDate.After(monthStart) && completionDate.Before(monthEnd.AddDate(0, 0, 1)) {
			daysActive := float64(monthEnd.Sub(completionDate).Hours() / 24)
			total += (upgrade.RentIncrease * daysActive) / float64(daysInMonth(monthStart))
		}
	}
	return total
}

// calculateUpgradeProration computes prorated upgrade rent for a given month.
func calculateUpgradeProration(property *models.Property, daysOwned, daysInMonth float64) float64 {
	total := 0.0
	for _, upgrade := range property.Upgrades {
		completionDate := upgrade.PurchaseDate.AddDate(0, 0, upgrade.DaysToComplete)
		if completionDate.Before(property.PurchaseDate) {
			total += upgrade.RentIncrease
		} else {
			daysActive := daysOwned - float64(completionDate.Day())
			if daysActive > 0 {
				total += (upgrade.RentIncrease * daysActive) / daysInMonth
			}
		}
	}
	return total
}

func distributeRentToOwner(world *ecs.World, property *models.Property, rent float64) {
	for _, entity := range world.Entities {
		playerComp := entity.GetComponent("PlayerComponent")
		if playerComp != nil && playerComp.(*components.PlayerComponent).Player.ID == property.PlayerID {
			playerComp.(*components.PlayerComponent).Player.Funds += rent
			return
		}
	}
}

func calculateMonthsPassed(lastUpdated, currentDate time.Time) int {
	if currentDate.Before(lastUpdated) {
		return 0
	}

	monthsPassed := 0
	nextMonth := nextMonthStart(lastUpdated)
	for !nextMonth.After(currentDate) {
		monthsPassed++
		nextMonth = nextMonth.AddDate(0, 1, 0)
	}
	return monthsPassed
}

func calculateDaysOwned(startDate, endDate, purchaseDate time.Time) int {
	if purchaseDate.Before(startDate) {
		return int(endDate.Sub(startDate).Hours()/24) + 1
	}
	return int(endDate.Sub(purchaseDate).Hours() / 24)
}

func roundToNearest5(value float64) float64 {
	return math.Floor(value/5) * 5
}

func daysInMonth(date time.Time) int {
	return date.AddDate(0, 1, -date.Day()).Day()
}

func nextMonthStart(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location()).AddDate(0, 1, 0)
}

func monthEnd(date time.Time) time.Time {
	return date.AddDate(0, 1, -date.Day())
}

func isPartialMonth(date time.Time) bool {
	return date.Day() != daysInMonth(date)
}
