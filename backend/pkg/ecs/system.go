package ecs

type System interface {
    Update(world *World)
}
