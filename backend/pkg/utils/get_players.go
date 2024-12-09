package utils

import (
	"github.com/markbmullins/city-developer/pkg/ecs"
)

func GetPlayers(world *ecs.World) []*ecs.Entity {
	var playerEntities []*ecs.Entity
	for _, entity := range world.Entities {
		if _, ok := entity.Components["Player"]; ok {
			playerEntities = append(playerEntities, entity)
		}
	}
	return playerEntities
}
