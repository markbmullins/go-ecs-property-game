package ecs

import (
	"errors"
	"fmt"
	"reflect"
	"sync"

	"github.com/markbmullins/city-developer/pkg/components"
)

type World struct {
	Entities          map[int]*Entity
	Systems           []System
	nextEntityID      int
	nextEntityIDMutex sync.Mutex
	//lookup table to quickly find which entities have a given component
	Indexes                  map[string]map[int]*Entity // componentName -> (entityId -> entityPointer)
	OwnedPropertiesIndex     map[int][]int              // ownerID -> propertyIDs
	GroupPropertiesIndex     map[int][]int              // groupID -> propertyIDs
	GroupUpgradedPercentages map[int]float64            // groupID -> upgradedPercentage
	GroupUpgradedCounts      map[int]int                // groupID -> number of properties with >=1 upgrade
	Players                  []*Entity
}

func NewWorld() *World {
	return &World{
		Entities:                 make(map[int]*Entity),
		Indexes:                  make(map[string]map[int]*Entity),
		OwnedPropertiesIndex:     make(map[int][]int),
		GroupPropertiesIndex:     make(map[int][]int),
		GroupUpgradedPercentages: make(map[int]float64),
		GroupUpgradedCounts:      make(map[int]int),
		nextEntityID:             1, // Start at 1 since 0 is reserved for GameTime
	}
}

func (w *World) AddComponentToIndex(entity *Entity, compType reflect.Type) {
	compName := compType.String()
	if w.Indexes[compName] == nil {
		w.Indexes[compName] = make(map[int]*Entity)
	}
	w.Indexes[compName][entity.ID] = entity
}

func (w *World) RemoveComponentFromIndex(entity *Entity, compType reflect.Type) {
	compName := compType.String()
	if index, exists := w.Indexes[compName]; exists {
		delete(index, entity.ID)
		if len(index) == 0 {
			delete(w.Indexes, compName) // Clean up empty index
		}
	}
}

func (w *World) AddSpecificEntity(id int, entity *Entity) {
	if _, exists := w.Entities[id]; exists {
		panic(fmt.Sprintf("Entity ID %d already exists!", id)) // Ensure no duplicate IDs
	}
	entity.ID = id
	w.Entities[id] = entity
}

func (w *World) GetEntity(id int) *Entity {
	return w.Entities[id]
}

func (w *World) AddEntity(entity *Entity) {
	w.nextEntityIDMutex.Lock()
	defer w.nextEntityIDMutex.Unlock()

	id := w.nextEntityID
	w.nextEntityID++
	entity.ID = id
	w.Entities[id] = entity
	for compKey := range entity.Components {
		if t, ok := getTypeFromString(compKey); ok {
			w.AddComponentToIndex(entity, t)
		}
	}

	if entity.Type == "Player" {
		w.Players = append(w.Players, entity)
	}

	// Update indexing for ownership and group
	if entity.Type == "Property" {
		if ownable, err := entity.GetOwnable(); err != nil && ownable.Owned {
			w.OwnedPropertiesIndex[ownable.OwnerID] = append(w.OwnedPropertiesIndex[ownable.OwnerID], entity.ID)
		}
		if groupable, err := entity.GetGroupable(); err != nil {
			w.GroupPropertiesIndex[groupable.GroupID] = append(w.GroupPropertiesIndex[groupable.GroupID], entity.ID)
		}
	}

}

func (w *World) RemoveEntity(id int) {
	entity, exists := w.Entities[id]
	if !exists {
		return
	}
	if entity.Type == "Player" {
		w.removePlayerFromIndex(entity)
	}
	for compKey := range entity.Components {
		// Convert back to reflect.Type before removal
		if t, ok := getTypeFromString(compKey); ok {
			w.RemoveComponentFromIndex(entity, t)
		}
	}
	delete(w.Entities, id)
}

func (w *World) QueryByComponent(componentName string) []*Entity {
	Entities := make([]*Entity, 0)
	for _, entity := range w.Indexes[componentName] {
		Entities = append(Entities, entity)
	}
	return Entities
}

func (w *World) AddSystem(system System) {
	w.Systems = append(w.Systems, system)
}

func (w *World) Update() {
	for _, system := range w.Systems {
		system.Update(w)
	}
}

func (w *World) ChangePropertyOwnership(propertyID int, oldOwnerID, newOwnerID int) {
	// Remove from old owner
	if oldProps, ok := w.OwnedPropertiesIndex[oldOwnerID]; ok {
		w.OwnedPropertiesIndex[oldOwnerID] = removeIntFromSlice(oldProps, propertyID)
	}
	// Add to new owner
	w.OwnedPropertiesIndex[newOwnerID] = append(w.OwnedPropertiesIndex[newOwnerID], propertyID)
}
func (w *World) BuyProperty(propertyID int, ownerID int) {
	w.OwnedPropertiesIndex[ownerID] = append(w.OwnedPropertiesIndex[ownerID], propertyID)
}

func (w *World) SellProperty(propertyID int) {
	propertyEntity := w.Entities[propertyID]
	ownableComponent, _ := propertyEntity.GetOwnable()
	ownerId := ownableComponent.OwnerID

	// Remove from old owner
	if oldProps, ok := w.OwnedPropertiesIndex[ownerId]; ok {
		w.OwnedPropertiesIndex[ownerId] = removeIntFromSlice(oldProps, propertyID)
	}
}

// Utility func
func removeIntFromSlice(s []int, val int) []int {
	for i, v := range s {
		if v == val {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func (w *World) ChangePropertyGroup(propertyID int, oldGroupID, newGroupID int) {
	if oldGroupProps, ok := w.GroupPropertiesIndex[oldGroupID]; ok {
		w.GroupPropertiesIndex[oldGroupID] = removeIntFromSlice(oldGroupProps, propertyID)
	}
	w.GroupPropertiesIndex[newGroupID] = append(w.GroupPropertiesIndex[newGroupID], propertyID)
}

func (w *World) GetOwnedEntities(ownerID int) []*Entity {
	var results []*Entity
	propertyIDs, ok := w.OwnedPropertiesIndex[ownerID]
	if !ok {
		return results
	}
	for _, pid := range propertyIDs {
		propertyEntity := w.Entities[pid]
		if propertyEntity != nil {
			results = append(results, propertyEntity)
		}
	}
	return results
}

func (w *World) GetCurrentGameTime() (*components.GameTime, error) {
	timeComp := w.Entities[0]
	if timeComp != nil {
		gameTimeComp, err := timeComp.GetGameTime()
		if err == nil {
			return gameTimeComp, nil
		}
	}
	return nil, errors.New("GameTime component not found in the world")
}

func (w *World) ApplyUpgradeToProperty(property *Entity, upgrade *components.Upgrade) error {
	upgradable, err := property.GetUpgradable()
	if err != nil {
		return errors.New("property not upgradable")
	}

	// Check if property currently has zero upgrades applied
	hadNoUpgrades := (len(upgradable.AppliedUpgrades) == 0)

	// Apply the new upgrade
	upgrade.Applied = true
	upgradable.AppliedUpgrades = append(upgradable.AppliedUpgrades, upgrade)
	property.AddComponent(upgradable) // update property component

	// If this is the first applied upgrade, we need to update the group's count
	if hadNoUpgrades {
		groupable, _ := property.GetGroupable()
		groupID := groupable.GroupID
		w.GroupUpgradedCounts[groupID]++ // increment count of upgraded props in this group

		// Recalculate the upgraded percentage for this group
		w.recalculateGroupUpgradedPercentage(groupID)
	}

	return nil
}

func (w *World) recalculateGroupUpgradedPercentage(groupID int) {
	properties := w.GroupPropertiesIndex[groupID]
	totalProperties := len(properties)
	if totalProperties == 0 {
		w.GroupUpgradedPercentages[groupID] = 0
		return
	}

	upgradedCount := w.GroupUpgradedCounts[groupID]
	percentage := float64(upgradedCount) / float64(totalProperties) * 100.0
	w.GroupUpgradedPercentages[groupID] = percentage
}

func (w *World) removePlayerFromIndex(entity *Entity) {
	for i, p := range w.Players {
		if p == entity {
			w.Players = append(w.Players[:i], w.Players[i+1:]...)
			break
		}
	}
}

func (w *World) GetEntitiesInGroup(groupID int) []*Entity {
	var groupPropertyIds = w.GroupPropertiesIndex[groupID]

	var entities []*Entity
	for _, id := range groupPropertyIds {
		entities = append(entities, w.Entities[id])
	}

	return entities
}
