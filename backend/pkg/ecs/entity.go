package ecs

import (
	"errors"
	"fmt"
	"strings"
	"sync"
)

type Entity struct {
	ID         int
	Type       string
	Components map[string]interface{} // Type-safe storage for components
	mu         sync.RWMutex
}

// NewEntity creates a new entity with the specified type.
func NewEntity(entityType string) *Entity {
	return &Entity{
		ID:         -1, // ID assigned by the World
		Type:       entityType,
		Components: make(map[string]interface{}),
	}
}

// AddComponent adds a component to the entity.
func (e *Entity) AddComponent(component interface{}) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	typeName := typeNameOf(component)
	if _, exists := e.Components[typeName]; exists {
		return errors.New("component already exists")
	}

	e.Components[typeName] = component
	return nil
}

// GetComponent retrieves a component of the specified type.
func (e *Entity) GetComponent(componentType interface{}) (interface{}, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()

	typeName := typeNameOf(componentType)
	component, exists := e.Components[typeName]
	if !exists {
		return nil, errors.New("component not found")
	}
	return component, nil
}

// RemoveComponent removes a component of the specified type.
func (e *Entity) RemoveComponent(componentType interface{}) {
	e.mu.Lock()
	defer e.mu.Unlock()

	typeName := typeNameOf(componentType)
	delete(e.Components, typeName)
}

// Utility function to get the type name of a component.
func typeNameOf(component interface{}) string {
	fullTypeName := fmt.Sprintf("%T", component) // e.g., "*components.GameTime"
	if strings.HasPrefix(fullTypeName, "*components.") {
		return strings.TrimPrefix(fullTypeName, "*components.")
	}
	return fullTypeName
}
