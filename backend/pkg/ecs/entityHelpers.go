package ecs

func (w *World) GetPlayer(id int) *Entity {
	key := NewEntityKey("Player", id)
	return w.Entities[key]
}

func (w *World) GetProperty(id int) *Entity {
	key := NewEntityKey("Property", id)
	return w.Entities[key]
}

func (w *World) GetNeighborhood(id int) *Entity {
	key := NewEntityKey("Neighborhood", id)
	return w.Entities[key]
}

func (w *World) GetGameTime() *Entity {
	key := NewEntityKey("GameTime", 0)
	return w.Entities[key]
}

func NewPlayer(id int) *Entity {
	return &Entity{
		ID:         EntityKey{EntityType: "Player", ID: id},
		Components: make(map[string]interface{}),
	}
}

func NewGameTime() *Entity {
	return &Entity{
		ID:         EntityKey{EntityType: "GameTime", ID: 0},
		Components: make(map[string]interface{}),
	}
}

func NewNeighborhood(id int) *Entity {
	return &Entity{
		ID:         EntityKey{EntityType: "Neighborhood", ID: id},
		Components: make(map[string]interface{}),
	}
}

func NewProperty(id int) *Entity {
	return &Entity{
		ID:         EntityKey{EntityType: "Property", ID: id},
		Components: make(map[string]interface{}),
	}
}
