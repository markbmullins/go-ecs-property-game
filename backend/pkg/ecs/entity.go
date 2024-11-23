package ecs

type Entity struct {
    ID         int
    Components map[string]Component
}

func NewEntity(id int) *Entity {
    return &Entity{
        ID:         id,
        Components: make(map[string]Component),
    }
}

func (e *Entity) AddComponent(name string, component Component) {
    e.Components[name] = component
}

func (e *Entity) GetComponent(name string) Component {
    return e.Components[name]
}
