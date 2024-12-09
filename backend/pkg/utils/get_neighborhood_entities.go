package utils

import (
	"github.com/markbmullins/city-developer/pkg/components"
	"github.com/markbmullins/city-developer/pkg/ecs"
)

func GetNeighborhoodEntities(world *ecs.World) []*components.Neighborhood {
	neighborhoodEntities := make([]*components.Neighborhood, 0)
	for _, entity := range world.Entities {
		if neighborhoodComp, ok := entity.GetComponent("Neighborhood").(*components.Neighborhood); ok {
			neighborhoodEntities = append(neighborhoodEntities, neighborhoodComp)
		}
	}
	return neighborhoodEntities
}
