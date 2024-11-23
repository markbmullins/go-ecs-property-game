package utils

import (
	"errors"

	"github.com/markbmullins/city-developer/pkg/components"
	"github.com/markbmullins/city-developer/pkg/ecs"
	"github.com/markbmullins/city-developer/pkg/models"
)

// GetCurrentGameTime retrieves the current game time from the world.
func GetCurrentGameTime(world *ecs.World) (*models.GameTime, error) {
	for _, entity := range world.Entities {
		timeComp := entity.GetComponent("GameTime")
		if timeComp != nil {
			gameTimeComp, ok := timeComp.(*components.GameTime)
			if ok && gameTimeComp.Time != nil {
				return gameTimeComp.Time, nil
			}
		}
	}
	return nil, errors.New("GameTime component not found in the world")
}
