package ecs

type World struct {
    Entities map[int]*Entity
    Systems  []System
}

func NewWorld() *World {
    return &World{
        Entities: make(map[int]*Entity),
        Systems:  []System{},
    }
}

func (w *World) AddEntity(entity *Entity) {
    w.Entities[entity.ID] = entity
}

func (w *World) AddSystem(system System) {
    w.Systems = append(w.Systems, system)
}

func (w *World) Update() {
    for _, system := range w.Systems {
        system.Update(w)
    }
}
