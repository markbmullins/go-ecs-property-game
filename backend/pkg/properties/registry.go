// pkg/properties/registry.go

package properties

import (
	"sync"

	"github.com/markbmullins/city-developer/pkg/models"
)

type PropertyRegistry struct {
	properties map[int]*models.Property
	mutex      sync.RWMutex
}

func NewPropertyRegistry() *PropertyRegistry {
	return &PropertyRegistry{
		properties: make(map[int]*models.Property),
	}
}

func (pr *PropertyRegistry) Register(prop *models.Property) {
	pr.mutex.Lock()
	defer pr.mutex.Unlock()
	pr.properties[prop.ID] = prop
}

func (pr *PropertyRegistry) GetPropertyByID(id int) *models.Property {
	pr.mutex.RLock()
	defer pr.mutex.RUnlock()
	return pr.properties[id]
}

func (pr *PropertyRegistry) GetAllProperties() []*models.Property {
	pr.mutex.RLock()
	defer pr.mutex.RUnlock()
	props := make([]*models.Property, 0, len(pr.properties))
	for _, prop := range pr.properties {
		props = append(props, prop)
	}
	return props
}
