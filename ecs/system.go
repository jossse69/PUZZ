package ecs

type System interface {
	Update(dt float32)
	Draw() // for renderer systems
}

type SystemManager struct {
	systems []System
}

func NewSystemManager() *SystemManager {
	return &SystemManager{}
}

func (manager *SystemManager) AddSystem(system System) {
	manager.systems = append(manager.systems, system)
}

func (manager *SystemManager) UpdateSystems(dt float32) {
	for _, system := range manager.systems {
		system.Update(dt)
	}
}

func (manager *SystemManager) DrawSystems() {
	for _, system := range manager.systems {
		system.Draw()
	}
}
