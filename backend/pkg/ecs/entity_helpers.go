package ecs

import "github.com/markbmullins/city-developer/pkg/components"

func (e *Entity) GetFunds() (*components.Funds, error) {
	component, err := e.GetComponent(&components.Funds{})
	if err != nil {
		return nil, err
	}
	return component.(*components.Funds), nil
}

func (e *Entity) GetOwnable() (*components.Ownable, error) {
	component, err := e.GetComponent(&components.Ownable{})
	if err != nil {
		return nil, err
	}
	return component.(*components.Ownable), nil
}

func (e *Entity) GetGameTime() (*components.GameTime, error) {
	component, err := e.GetComponent(&components.GameTime{})
	if err != nil {
		return nil, err
	}
	return component.(*components.GameTime), nil
}

func (e *Entity) GetGroupable() (*components.Groupable, error) {
	component, err := e.GetComponent(&components.Groupable{})
	if err != nil {
		return nil, err
	}
	return component.(*components.Groupable), nil
}

func (e *Entity) GetPurchaseable() (*components.Purchaseable, error) {
	component, err := e.GetComponent(&components.Purchaseable{})
	if err != nil {
		return nil, err
	}
	return component.(*components.Purchaseable), nil
}

func (e *Entity) GetUpgradable() (*components.Upgradable, error) {
	component, err := e.GetComponent(&components.Upgradable{})
	if err != nil {
		return nil, err
	}
	return component.(*components.Upgradable), nil
}

func (e *Entity) GetRentable() (*components.Rentable, error) {
	component, err := e.GetComponent(&components.Rentable{})
	if err != nil {
		return nil, err
	}
	return component.(*components.Rentable), nil
}

func (e *Entity) GetRentBoostable() (*components.RentBoostable, error) {
	component, err := e.GetComponent(&components.RentBoostable{})
	if err != nil {
		return nil, err
	}
	return component.(*components.RentBoostable), nil
}

func (e *Entity) AddFunds(funds *components.Funds) error {
	return e.AddComponent(funds)
}

func (e *Entity) AddUpgradable(upgradable *components.Upgradable) error {
	return e.AddComponent(upgradable)
}

func (w *World) GetAllProperties() []*Entity {
	entities := make([]*Entity, 0)
	for _, entity := range w.Entities {
		if entity.Type == "Property" {
			entities = append(entities, entity)
		}
	}
	return entities
}

func (w *World) GetAllPropertiesMap() map[int]*Entity {
	entities := map[int]*Entity{}
	for _, entity := range w.Entities {
		if entity.Type == "Property" {
			entities[entity.ID] = entity
		}
	}
	return entities
}
