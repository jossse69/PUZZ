package ecs

type Entity uint32

type EntityManager struct {
	lastID   Entity
	entities map[Entity]bool
}

// NewEntityManager instantiates a new EntityManager.
func NewEntityManager() *EntityManager {
	return &EntityManager{
		entities: make(map[Entity]bool),
	}
}

// NewEntity creates a new entity.
func (manager *EntityManager) NewEntity() Entity {
	manager.lastID++
	manager.entities[manager.lastID] = true
	return manager.lastID
}

// DestroyEntity destroys an entity.
func (manager *EntityManager) DestroyEntity(entity Entity) {
	delete(manager.entities, entity)
}

// HasEntity returns true if the entity exists.
func (manager *EntityManager) HasEntity(entity Entity) bool {
	_, ok := manager.entities[entity]
	return ok
}

// GetAllEntities returns all entities.
func (manager *EntityManager) GetAllEntities() []Entity {
	entities := make([]Entity, 0, len(manager.entities))
	for entity := range manager.entities {
		entities = append(entities, entity)
	}
	return entities
}
