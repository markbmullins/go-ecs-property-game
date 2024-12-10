package ecs

import (
	"reflect"
	"sync"
)

var (
	typeRegistry   = make(map[string]reflect.Type)
	typeRegistryMu sync.RWMutex
)

func registerType(t reflect.Type) {
	typeRegistryMu.Lock()
	defer typeRegistryMu.Unlock()
	typeRegistry[t.Name()] = t
}

func getTypeFromString(name string) (reflect.Type, bool) {
	typeRegistryMu.RLock()
	defer typeRegistryMu.RUnlock()
	t, ok := typeRegistry[name]
	return t, ok
}
