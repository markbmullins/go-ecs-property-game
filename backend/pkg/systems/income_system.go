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
type IncomeSystem struct{}

// Update triggers the rent collection process, handling fast-forwarding of time.
func (s *IncomeSystem) Update(world *ecs.World) {
	gameTime, err := utils.GetCurrentGameTime(world)
	if err != nil {
		fmt.Println("Error: GameTime component not found")
		return
	}

	// fmt.Printf("GameTime - CurrentDate: %s, LastUpdated: %s, IsPaused: %v\n",
	// 	gameTime.CurrentDate.Format("2006-01-02"), gameTime.LastUpdated.Format("2006-01-02"), gameTime.IsPaused)

	if !gameTime.IsPaused {
		monthsPassed := calculateMonthsPassed(gameTime.LastUpdated, gameTime.CurrentDate)
		// fmt.Printf("Months passed: %d\n", monthsPassed)
		if monthsPassed > 0 {
			// Pass currentDate to collectRentForSkippedMonths
			collectRentForSkippedMonths(world, gameTime.LastUpdated, monthsPassed)
			gameTime.NewMonth = true
		}
		gameTime.LastUpdated = gameTime.CurrentDate
	}
}

// collectRentForSkippedMonths iterates through each skipped month and collects appropriate rent.
func collectRentForSkippedMonths(world *ecs.World, lastUpdated time.Time, monthsPassed int) {
	// fmt.Printf("Collecting rent for %d skipped months starting from %s\n", monthsPassed, lastUpdated.Format("2006-01-02"))

	location := lastUpdated.Location()

	// Calculate the first day of the next month after lastUpdated
	firstOfNextMonth := time.Date(lastUpdated.Year(), lastUpdated.Month(), 1, 0, 0, 0, 0, location).AddDate(0, 1, 0)

	// Step 1: Handle partial month if lastUpdated is not the last day of the month
	if lastUpdated.Day() != daysInMonth(lastUpdated) {
		// Partial month: from lastUpdated to end of month
		partialMonthStart := lastUpdated // Include the purchase day
		partialMonthEnd := time.Date(lastUpdated.Year(), lastUpdated.Month(), daysInMonth(lastUpdated), 0, 0, 0, 0, location)

		// fmt.Printf("Processing partial month: %s to %s\n", partialMonthStart.Format("2006-01-02"), partialMonthEnd.Format("2006-01-02"))

		for _, entity := range world.Entities {
			propComp := entity.GetComponent("PropertyComponent")
			if propComp != nil {
				property := propComp.(*components.PropertyComponent).Property
				if property.Owned {
					if property.PurchaseDate.Before(partialMonthStart) {
						// Full rent for the partial month
						collectFullRent(world, property, partialMonthStart, partialMonthEnd)
					} else if property.PurchaseDate.After(partialMonthEnd) {
						// No rent for this partial month
						continue
					} else {
						// Prorated rent for the partial month
						daysOwned := partialMonthEnd.Day() - property.PurchaseDate.Day()
						proratedRent := calculateProratedRent(property, daysOwned, float64(daysInMonth(partialMonthStart)), partialMonthStart, partialMonthEnd)
						distributeRentToOwner(world, property, proratedRent)
						// fmt.Printf("Collected prorated rent for property %s: $%.2f\n", property.Name, proratedRent)
					}
				}
			}
		}

		monthsPassed-- // Partial month already processed
	}

	// Step 2: Handle full months
	for i := 0; i < monthsPassed; i++ {
		// Calculate the specific month being processed
		monthStartDate := firstOfNextMonth.AddDate(0, i, 0)
		monthEndDate := monthStartDate.AddDate(0, 1, -1)

		// fmt.Printf("Processing full month: %s to %s\n", monthStartDate.Format("2006-01-02"), monthEndDate.Format("2006-01-02"))

		for _, entity := range world.Entities {
			propComp := entity.GetComponent("PropertyComponent")
			if propComp != nil {
				property := propComp.(*components.PropertyComponent).Property
				if property.Owned {
					if property.PurchaseDate.Before(monthStartDate) || property.PurchaseDate.Equal(monthStartDate) {
						// Full rent for the month
						collectFullRent(world, property, monthStartDate, monthEndDate)
					} else if property.PurchaseDate.After(monthEndDate) {
						// No rent for this month
						continue
					} else {
						// Prorated rent for the month
						// Add one to include putchase date
						daysOwned := daysInMonth(monthStartDate) - property.PurchaseDate.Day()
						proratedRent := calculateProratedRent(property, daysOwned, float64(daysInMonth(monthStartDate)), monthStartDate, monthEndDate)
						distributeRentToOwner(world, property, proratedRent)
						// fmt.Printf("Collected prorated rent for property %s: $%.2f\n", property.Name, proratedRent)
					}
				}
			}
		}
	}
}

// collectFullRent collects the full monthly rent for a property, considering upgrades.
func collectFullRent(world *ecs.World, property *models.Property, monthStartDate, monthEndDate time.Time) {
	fullRent := calculateFullRent(property, monthStartDate, monthEndDate)
	distributeRentToOwner(world, property, fullRent)
	// fmt.Printf("Collected full monthly rent for property %s: $%.2f\n", property.Name, fullRent)
}

// calculateFullRent calculates full monthly rent, incorporating upgrades and their proration.
func calculateFullRent(property *models.Property, monthStartDate, monthEndDate time.Time) float64 {
	baseRent := property.BaseRent
	totalUpgradeIncrease := 0.0
	daysInMonth := float64(daysInMonth(monthStartDate))

	for i := 0; i < property.UpgradeLevel && i < len(property.Upgrades); i++ {
		upgrade := property.Upgrades[i]
		completionDate := upgrade.PurchaseDate.AddDate(0, 0, upgrade.DaysToComplete)

		if completionDate.Before(monthStartDate) {
			// Upgrade completed before the month starts; full RentIncrease applies
			totalUpgradeIncrease += upgrade.RentIncrease
			// fmt.Printf("Upgrade %d completed before month start: +$%.2f\n", i+1, upgrade.RentIncrease)
		} else if completionDate.Equal(monthStartDate) || completionDate.After(monthStartDate) && completionDate.Before(monthEndDate.AddDate(0, 0, 1)) {
			// Upgrade completed during the month; prorate RentIncrease
			daysAfterUpgrade := float64(monthEndDate.Sub(completionDate).Hours() / 24)
			if daysAfterUpgrade < 0 {
				daysAfterUpgrade = 0
			}
			proratedIncrease := (upgrade.RentIncrease * daysAfterUpgrade) / daysInMonth
			proratedIncrease = roundToNearest5(proratedIncrease)
			totalUpgradeIncrease += proratedIncrease
			// fmt.Printf("Upgrade %d completed on %s: Prorated +$%.2f\n", i+1, completionDate.Format("2006-01-02"), proratedIncrease)
		}
		// Upgrades completed after the month are not included
	}

	fullRent := baseRent + totalUpgradeIncrease
	// fmt.Printf("Total rent for property %s: $%.2f (Base: $%.2f + Upgrades: $%.2f)\n", property.Name, fullRent, baseRent, totalUpgradeIncrease)
	return fullRent
}

// calculateProratedRent calculates prorated rent based on days owned in the specified month, including upgrades.
func calculateProratedRent(property *models.Property, daysOwned int, daysInMonth float64, monthStartDate, monthEndDate time.Time) float64 {
	// Prorate base rent
	proratedBaseRent := (property.BaseRent * float64(daysOwned)) / daysInMonth

	// Sum prorated upgrades
	proratedUpgradeIncrease := 0.0

	for i := 0; i < property.UpgradeLevel && i < len(property.Upgrades); i++ {
		upgrade := property.Upgrades[i]
		completionDate := upgrade.PurchaseDate.AddDate(0, 0, upgrade.DaysToComplete)

		if completionDate.Before(monthStartDate) {
			// Upgrade completed before the month starts; full RentIncrease applies
			proratedUpgradeIncrease += upgrade.RentIncrease
			// fmt.Printf("Upgrade %d completed before month start: +$%.2f\n", i+1, upgrade.RentIncrease)
		} else if completionDate.Equal(monthStartDate) || completionDate.After(monthStartDate) && completionDate.Before(monthEndDate.AddDate(0, 0, 1)) {
			// Upgrade completed during the month; prorate RentIncrease
			// Calculate days after upgrade completion (inclusive of the day after completion)
			daysAfterUpgrade := int(monthEndDate.Sub(completionDate).Hours() / 24) // Floor to int
			if daysAfterUpgrade < 0 {
				daysAfterUpgrade = 0
			}
			proratedIncrease := (upgrade.RentIncrease * float64(daysAfterUpgrade)) / daysInMonth
			proratedIncrease = roundToNearest5(proratedIncrease)
			proratedUpgradeIncrease += proratedIncrease
			// fmt.Printf("Upgrade %d completed on %s: Prorated +$%.2f\n", i+1, completionDate.Format("2006-01-02"), proratedIncrease)
		}
		// Upgrades completed after the month are not included
	}

	// Calculate total prorated rent
	totalProratedRent := proratedBaseRent + proratedUpgradeIncrease
	totalProratedRent = roundToNearest5(totalProratedRent)

	logProratedRentCalculation(property, daysInMonth, daysOwned, proratedBaseRent, proratedUpgradeIncrease, totalProratedRent)
	return totalProratedRent
}

// logProratedRentCalculation logs details of the prorated rent calculation for debugging.
func logProratedRentCalculation(property *models.Property, daysInMonth float64, daysOwned int, proratedBaseRent float64, proratedUpgradeIncrease float64, totalProratedRent float64) {
	fmt.Printf("Calculating prorated rent for property: %s\n", property.Name)
	fmt.Printf("Days in month: %.0f\n", daysInMonth)
	fmt.Printf("Days owned in month: %d\n", daysOwned)
	fmt.Printf("Prorated Base Rent: $%.2f\n", proratedBaseRent)
	fmt.Printf("Prorated Upgrade Increases: $%.2f\n", proratedUpgradeIncrease)
	fmt.Printf("Total Prorated Rent: $%.2f\n", totalProratedRent)
}

// distributeRentToOwner calculates and assigns rent to the owner of each property.
func distributeRentToOwner(world *ecs.World, property *models.Property, rent float64) {
	var ownerComp *components.PlayerComponent

	// Iterate through all entities to find the player with matching Player.ID
	for _, entity := range world.Entities {
		comp := entity.GetComponent("PlayerComponent")
		if comp != nil {
			playerComp := comp.(*components.PlayerComponent)
			if playerComp.Player.ID == property.PlayerID {
				ownerComp = playerComp
				break
			}
		}
	}

	if ownerComp == nil {
		fmt.Printf("Owner entity for player ID %d not found!\n", property.PlayerID)
		return
	}

	ownerComp.Player.Funds += rent
	// fmt.Printf("Distributed $%.2f to player %d. New funds: $%.2f\n", rent, ownerComp.Player.ID, ownerComp.Player.Funds)
}

// roundToNearest5 rounds a number down to the nearest multiple of 5.
func roundToNearest5(num float64) float64 {
	rounded := math.Floor(num/5.0) * 5.0
	return rounded
}

// calculateMonthsPassed determines how many full months have passed between two dates.
func calculateMonthsPassed(lastUpdated, currentDate time.Time) int {
	// fmt.Printf("Calculating months passed from %s to %s\n", lastUpdated.Format("2006-01-02"), currentDate.Format("2006-01-02"))

	if currentDate.Before(lastUpdated) {
		fmt.Println("Current date is before last updated date; months passed: 0")
		return 0
	}

	monthsPassed := 0
	firstOfNextMonth := time.Date(lastUpdated.Year(), lastUpdated.Month(), 1, 0, 0, 0, 0, lastUpdated.Location()).AddDate(0, 1, 0)

	for !firstOfNextMonth.After(currentDate) {
		monthsPassed++
		firstOfNextMonth = firstOfNextMonth.AddDate(0, 1, 0)
		// Add limit to prevent excessive iteration
		if monthsPassed > 1200 { // Example cap at 100 years
			fmt.Println("Exceeded max months count, breaking out of loop")
			break
		}
	}
	return monthsPassed
}

// Helper function to get number of days in a month
func daysInMonth(date time.Time) int {
	return date.AddDate(0, 1, -date.Day()).Day()
}
