package utils

import (
	"errors"

	"github.com/markbmullins/city-developer/pkg/components"
	"github.com/markbmullins/city-developer/pkg/ecs"
)

// GetCurrentGameTime retrieves the current game time from the world.
func GetCurrentGameTime(world *ecs.World) (*components.GameTime, error) {
	for _, entity := range world.Entities {
		timeComp := entity.GetComponent("GameTime")
		if timeComp != nil {
			gameTimeComp, ok := timeComp.(*components.GameTime)
			if ok && gameTimeComp != nil {
				return gameTimeComp, nil
			}
		}
	}
	return nil, errors.New("GameTime component not found in the world")
}
