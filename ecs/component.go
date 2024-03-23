package ecs

import "fmt"

type Component interface{}

type ComponentManager struct {
	components map[Entity]map[string]Component
}

// NewComponentManager creates a new instance of ComponentManager.
func NewComponentManager() *ComponentManager {
	return &ComponentManager{
		components: make(map[Entity]map[string]Component),
	}
}

// AddComponent adds a component a entity
func (manager *ComponentManager) AddComponent(entity Entity, component Component) {
	componentName := getType(component)
	if manager.components[entity] == nil {
		manager.components[entity] = make(map[string]Component)
	}
	manager.components[entity][componentName] = component
}

// GetComponent returns the component if it exists.
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

// HasComponent returns true if the entity has the component.
func (manager *ComponentManager) HasComponent(entity Entity, componentName string) bool {
	_, ok := manager.components[entity][componentName]
	return ok
}

// HasComponents returns true if the entity has all the components in the list.
func (manager *ComponentManager) HasComponents(entity Entity, componentNames ...string) bool {
	for _, componentName := range componentNames {
		if !manager.HasComponent(entity, componentName) {
			return false
		}
	}
	return true
}
