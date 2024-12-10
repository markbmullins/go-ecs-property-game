package entities

import (
	"time"

	"github.com/markbmullins/city-developer/pkg/components"
	"github.com/markbmullins/city-developer/pkg/ecs"
)

/** Creates a game time entity in the game.
 * A game time entity has the following components:
 * GameTime: Tracks the current game time, pause state, and speed multiplier.
 */
func CreateGameTime(
	currentDate time.Time,
	rentCollectionDay int,
) *ecs.Entity {
	gameTime := ecs.NewEntity("GameTime")

	gameTime.AddComponent(&components.GameTime{
		CurrentDate:       currentDate,
		IsPaused:          false,
		SpeedMultiplier:   1.0,
		NewMonth:          false,
		LastUpdated:       currentDate,
		RentCollectionDay: rentCollectionDay,
	})

	return gameTime
}
