package ecs

type World struct {
	Entities map[int]*Entity
	Systems  []System
	Indexes  map[string]map[int]*Entity
}

func NewWorld() *World {
	return &World{
		Entities: make(map[int]*Entity),
		Indexes:  make(map[string]map[int]*Entity),
	}
}

func (w *World) AddEntity(entity *Entity) {
	w.Entities[entity.ID] = entity

	// Index components by type
	for compName := range entity.Components {
		if _, exists := w.Indexes[compName]; !exists {
			w.Indexes[compName] = make(map[int]*Entity)
		}
		w.Indexes[compName][entity.ID] = entity
	}
}

func (w *World) GetEntity(id int) *Entity {
    return w.Entities[id]
}

func (w *World) RemoveEntity(id int) {
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

func (w *World) LookupEntityByComponentAndID(componentName string, id int) *Entity {
	if compEntities, exists := w.Indexes[componentName]; exists {
		return compEntities[id]
	}
	return nil
}

func (w *World) AddSystem(system System) {
	w.Systems = append(w.Systems, system)
}

func (w *World) Update() {
	for _, system := range w.Systems {
		system.Update(w)
	}
}
