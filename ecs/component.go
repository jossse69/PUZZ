package ecs

import "fmt"

type Component interface{}

type ComponentManager struct {
	components map[Entity]map[string]Component
}

func NewComponentManager() *ComponentManager {
	return &ComponentManager{
		components: make(map[Entity]map[string]Component),
	}
}

func (manager *ComponentManager) AddComponent(entity Entity, component Component) {
	componentName := getType(component)
	if manager.components[entity] == nil {
		manager.components[entity] = make(map[string]Component)
	}
	manager.components[entity][componentName] = component
}

func (manager *ComponentManager) GetComponent(entity Entity, componentName string) (Component, bool) {
	if components, ok := manager.components[entity]; ok {
		comp, ok := components[componentName]
		return comp, ok
	}
	return nil, false
}

// getType returns a unique string identifier for the component's type.
func getType(myvar interface{}) string {
	return fmt.Sprintf("%T", myvar)
}
