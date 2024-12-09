package ecs

import "fmt"

type ComponentName = string
type EntityName = string
type Entity struct {
	ID         EntityKey
	Components map[ComponentName]interface{}
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

func NewEntity(entityType string, id int) *Entity {
	return &Entity{
		ID:         EntityKey{EntityType: entityType, ID: id},
		Components: make(map[string]interface{}),
	}
}

func (e *Entity) AddComponent(name string, component Component) {
	e.Components[name] = component
}

func (e *Entity) GetComponent(name string) Component {
	return e.Components[name]
}
