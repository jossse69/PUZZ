package ecs

type System interface {
	Update(dt float32)
	Draw() // for renderer systems
}

type SystemManager struct {
	systems []System
}

// NewSystemManager instantiates a new SystemManager.
func NewSystemManager() *SystemManager {
	return &SystemManager{}
}

// AddSystem adds a system to the SystemManager.
func (manager *SystemManager) AddSystem(system System) {
	manager.systems = append(manager.systems, system)
}

// UpdateSystems updates all systems in the SystemManager with the given delta time.
func (manager *SystemManager) UpdateSystems(dt float32) {
	for _, system := range manager.systems {
		system.Update(dt)
	}
}

// DrawSystems iterates through the systems in the SystemManager and calls the Draw method on each system.
func (manager *SystemManager) DrawSystems() {
	for _, system := range manager.systems {
		system.Draw()
	}
}

// GetAllSystems returns all systems in the SystemManager.
func (manager *SystemManager) GetAllSystems() []System {
	return manager.systems
}
