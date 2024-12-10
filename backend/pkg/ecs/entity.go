package ecs

import (
	"fmt"
	"reflect"
	"sync"
)

type ComponentName = string
type EntityName = string
type Entity struct {
	Key        EntityKey
	Components map[string]interface{}
	mu         sync.RWMutex
}

type EntityKey struct {
	EntityType EntityName
	ID         int
}

func NewEntityKey(entityType string, id int) string {
	return fmt.Sprintf("%s-%d", entityType, id)
}

func (k EntityKey) ToString() string {
	return fmt.Sprintf("%s-%d", k.EntityType, k.ID)
}

func (w *World) GetEntity(key EntityKey) *Entity {
	return w.Entities[key.ToString()]
}

func (w *World) GetEntityStr(key string) *Entity {
	return w.Entities[key]
}

func (w *World) GetEntityById(entityType string, id int) *Entity {
	return w.Entities[NewEntityKey(entityType, id)]
}

func NewEntity(entityType string, id int) *Entity {
	return &Entity{
		Key:        EntityKey{EntityType: entityType, ID: id},
		Components: make(map[string]interface{}),
	}
}

// AddComponent uses reflect.Type for type safety in the method signature.
// Internally, it converts the reflect.Type to a string key.
func AddComponent[T any](entity *Entity, component *T) {
	entity.mu.Lock()
	defer entity.mu.Unlock()

	compType := reflect.TypeOf((*T)(nil)).Elem()
	compKey := compType.Name()

	// Register the type so we can convert back to reflect.Type later if needed.
	registerType(compType)
	entity.Components[compKey] = component
}

// GetComponent still uses reflect.Type to find the component type,
// but this time we find the string key from the type registry.
func GetComponent[T any](entity *Entity) (*T, bool) {
	entity.mu.RLock()
	defer entity.mu.RUnlock()

	compType := reflect.TypeOf((*T)(nil)).Elem()
	compKey := compType.Name()

	component, exists := entity.Components[compKey]
	if !exists {
		return nil, false
	}
	return component.(*T), true
}
