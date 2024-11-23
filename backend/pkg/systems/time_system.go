package systems

import (
	"fmt"
	"time"

	"github.com/markbmullins/city-developer/pkg/ecs"
	"github.com/markbmullins/city-developer/pkg/utils"
)

type TimeSystem struct{}

func (s *TimeSystem) Update(world *ecs.World) {
	gameTime, _ := utils.GetCurrentGameTime(world)

	if !gameTime.IsPaused {
		// Store the original date before advancing
		originalDate := gameTime.CurrentDate

		// Advance game time by one day, adjusted by the speed multiplier
		newDate := originalDate.Add(time.Duration(24 * float64(time.Hour) * gameTime.SpeedMultiplier))
		gameTime.CurrentDate = newDate

		fmt.Printf("Time set to %s\n", newDate.Format("January-02-2006"))

		// Check if we crossed into a new month
		// Since the number of days in the month may not be divisible by the speed multipler
		// we could skip the 1st during the update, so compare months directly
		if originalDate.Month() != newDate.Month() {
			fmt.Println("First of the month: Collecting rent.")
			gameTime.NewMonth = true // Signal for monthly rent collection
		} else {
			gameTime.NewMonth = false
		}
	}
}
