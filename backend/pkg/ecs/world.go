package ecs

type World struct {
	Entities map[string]*Entity
	Systems  []System
	Indexes  map[string]map[string]*Entity
}

func NewWorld() *World {
	return &World{
		Entities: make(map[string]*Entity),
		Indexes:  make(map[string]map[string]*Entity),
	}
}

func (w *World) AddEntity(entity *Entity) {
	w.Entities[entity.ID.ToString()] = entity
}

func (w *World) GetEntityByType(entityType string, id int) *Entity {
	return w.Entities[NewEntityKey(entityType, id)]
}

func (w *World) RemoveEntity(key EntityKey) {
	id := key.ToString()
	entity, exists := w.Entities[id]
	if !exists {
		return
	}
	for compName := range entity.Components {
		delete(w.Indexes[compName], id)
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
