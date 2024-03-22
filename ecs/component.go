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

// GetEntitiesWithComponents returns a list of entities with all the specified components.
func (manager *ComponentManager) GetEntitiesWithComponents(componentNames ...string) []Entity {
	entities := make([]Entity, 0)
	for entity, components := range manager.components {
		match := true
		for _, componentName := range componentNames {
			if _, ok := components[componentName]; !ok {
				match = false
				break
			}
		}
		if match {
			entities = append(entities, entity)
		}
	}
	return entities
}
