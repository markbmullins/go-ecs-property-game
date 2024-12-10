package ecs

func (w *World) GetPlayer(id int) *Entity {
	key := NewEntityKey("Player", id)
	return w.Entities[key]
}

func (w *World) GetProperty(id int) *Entity {
	key := NewEntityKey("Property", id)
	return w.Entities[key]
}

func (w *World) GetAllProperties() []*Entity {
	entities := make([]*Entity, 0)
	for _, entity := range w.Entities {
		if entity.Key.EntityType == "Property" {
			entities = append(entities, entity)
		}
	}
	return entities
}

func (w *World) GetAllPropertiesMap() map[string]*Entity {
	entities := map[string]*Entity{}
	for _, entity := range w.Entities {
		if entity.Key.EntityType == "Property" {
			entities[entity.Key.ToString()] = entity
		}
	}
	return entities
}
