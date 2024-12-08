package systems

import (
	"fmt"
	"math"
	"time"

	"github.com/markbmullins/city-developer/pkg/components"
	"github.com/markbmullins/city-developer/pkg/ecs"
	"github.com/markbmullins/city-developer/pkg/utils"
)

/*
===========================================================
                     INCOME SYSTEM
===========================================================

1. **Purchase Mechanics**
   - **No Rent on Purchase Day:** Rent is not collected on the day a property is purchased.
   - **Rent Collection Start:** Begins the day after the purchase date.
   - **First Month Proration:**
     - *First Day Purchase:* Full rent minus one day.
		 - *Mid Month Purchase:* Rent is collected from the day after purchase until the end of the month
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

**Key Rules Summary:**
- No rent is collected on Purchase Day or Upgrade Completion Day.
- Rent and upgrade bonuses commence the day after their respective initiation dates.
- Prorated rents are based on days owned or upgraded, excluding initiation days.
- Full rent is collected after a complete month, inclusive of all upgrades.
- The system accommodates variable time speeds and handles calendar anomalies accurately.

===========================================================
*/
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

func processRent(world *ecs.World, lastUpdated time.Time, monthsPassed int) {
	firstOfNextMonth := nextMonthStart(lastUpdated)

	// Handle partial month (if lastUpdated isn't at the end of its month)
	if isPartialMonth(lastUpdated) {
		processMonth(world, lastUpdated, monthEnd(lastUpdated))
		monthsPassed--
	}

	// Handle full months
	for i := 0; i < monthsPassed; i++ {
		startOfMonth := firstOfNextMonth.AddDate(0, i, 0)
		endOfMonth := monthEnd(startOfMonth)
		processMonth(world, startOfMonth, endOfMonth)
	}
}

func processMonth(world *ecs.World, startDate, endDate time.Time) {
	for _, entity := range world.Entities {
		processEntityRent(world, entity, startDate, endDate)
	}
}

func processEntityRent(world *ecs.World, entity *ecs.Entity, startDate, endDate time.Time) {
	property := entity.GetComponent("Property").(*components.Property)

	if !property.Owned {
		return
	}

	rent := calculateMonthlyRent(property, startDate, endDate)
	if rent > 0 {
		distributeRentToOwner(world, property, rent)
	}
}

// calculateMonthlyRent calculates the rent owed for a given property within the given month.
// It adheres to the following rules:
// - No rent on the purchase day; rent begins the day after purchase if within the month.
// - Each upgrade also begins contributing rent the day after it completes, if within the month.
// - Both base rent and upgrades are prorated based on the number of days active in the month.
// - After determining total active days for the property and any upgrades, it rounds the total rent down to the nearest multiple of 5.
func calculateMonthlyRent(property *components.Property, monthStart, monthEnd time.Time) float64 {
	daysInCurrentMonth := float64(daysInMonth(monthStart))
	if daysInCurrentMonth == 0 {
		// Safety check: Should never happen since daysInMonth should always return > 0
		return 0
	}

	// Determine when the property first becomes eligible to collect rent this month.
	// Rent starts the day after the purchase date, if that day falls within this month.
	propertyRentStartDate := maxTime(property.PurchaseDate.AddDate(0, 0, 1), monthStart)
	if propertyRentStartDate.After(monthEnd) {
		// The property wasn't active this month at all (e.g., purchased too late in the month).
		return 0
	}

	// Calculate the number of days the property is active in this month.
	propertyRentDays := countDaysInRange(propertyRentStartDate, monthEnd)

	// Calculate daily base rent by dividing the base rent by the total days in the month.
	baseDailyRent := (property.BaseRent + property.RentBoost)/ daysInCurrentMonth

	totalBaseRent := (baseDailyRent * float64(propertyRentDays))

	// Calculate the total rent from all upgrades active during this month.
	// Each upgrade also begins contributing rent the day after its completion date.
	totalUpgradeRent := 0.0
	for _, upgrade := range property.Upgrades {
		upgradeCompletionDate := upgrade.PurchaseDate.AddDate(0, 0, upgrade.DaysToComplete)
		upgradeRentStart := upgradeCompletionDate.AddDate(0, 0, 1)

		// Determine when the upgrade becomes active in this month.
		upgradeActiveStart := maxTime(upgradeRentStart, monthStart)
		if upgradeActiveStart.After(monthEnd) {
			// This upgrade wasn't active during this month.
			continue
		}

		// An upgrade only contributes rent if the property itself is active.
		// Find the intersection of the property's active period and the upgrade's active period.
		upgradeIntersectionStart := maxTime(propertyRentStartDate, upgradeActiveStart)
		if upgradeIntersectionStart.After(monthEnd) {
			// No intersection means no contribution from this upgrade.
			continue
		}

		// Count how many days this upgrade was both active and within the property's active period this month.
		intersectionDays := countDaysInRange(upgradeIntersectionStart, monthEnd)
		if intersectionDays > 0 {
			upgradeDailyRent := upgrade.RentIncrease / daysInCurrentMonth
			totalUpgradeRent += (upgradeDailyRent * float64(intersectionDays))
		}
	}

	// Total rent is the sum of the prorated base rent and the prorated upgrades rent.
	totalRent := totalBaseRent + totalUpgradeRent

	// Round down to the nearest multiple of 5 per the given rounding rule.
	return roundToNearest5(totalRent)
}


// distribute rent to correct player
func distributeRentToOwner(world *ecs.World, property *components.Property, rent float64) {
	for _, entity := range world.Entities {
		player := entity.GetComponent("Player").(*components.Player)
		if player != nil && player.ID == property.PlayerID {
			player.Funds += rent
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

func roundToNearest5(value float64) float64 {
	return math.Floor(value/5) * 5
}

// countDaysInRange counts INCLUSIVE days from startDate to endDate.
// For example:
// If startDate = 2024-12-01 and endDate = 2024-12-03,
// the days are counted as:
// 2024-12-01 (Day 1),
// 2024-12-02 (Day 2),
// 2024-12-03 (Day 3).
// Total = 3 days (inclusive).
func countDaysInRange(startDate, endDate time.Time) int {
	if startDate.After(endDate) {
		return 0
	}
	return int(endDate.Sub(startDate).Hours()/24) + 1
}

// maxTime returns the max of two times.
func maxTime(a, b time.Time) time.Time {
	if a.After(b) {
		return a
	}
	return b
}
