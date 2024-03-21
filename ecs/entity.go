package ecs

type Entity uint32

type EntityManager struct {
	lastID   Entity
	entities map[Entity]bool
}

func NewEntityManager() *EntityManager {
	return &EntityManager{
		entities: make(map[Entity]bool),
	}
}

func (manager *EntityManager) NewEntity() Entity {
	manager.lastID++
	manager.entities[manager.lastID] = true
	return manager.lastID
}
