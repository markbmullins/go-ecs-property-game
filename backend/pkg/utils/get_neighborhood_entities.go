package utils

import (
	"github.com/markbmullins/city-developer/pkg/components"
	"github.com/markbmullins/city-developer/pkg/ecs"
)

func GetNeighorhoodEntities(world *ecs.World) []*components.Neighborhood {
	neighborhoodEntities := make([]*components.Neighborhood, 0)
	for _, entity := range world.Entities {
		neighborhoodComp := entity.GetComponent("Neighborhood")
		if neighborhoodComp != nil {
			neighborhoodEntities = append(neighborhoodEntities, neighborhoodComp.(*components.Neighborhood))
		}
	}
	return neighborhoodEntities
}
