package ecs_engine

import (
	"github.com/cx333/game-works/services/room"
	"sync"
	"time"
)

/**
 * @Author: wgl
 * @Description:
 * @File: world
 * @Version: 1.0.0
 * @Date: 2025/4/21 22:21
 */

// World ECS世界，统一管理所有实体和系统
type World struct {
	mu       sync.Mutex
	entities map[EntityId]*Entity
	systems  []System
	nextId   EntityId
}

// NewWorld 创建新世界
func NewWorld() *World {
	return &World{
		entities: make(map[EntityId]*Entity),
	}
}

// CreateEntity 新建实体
func (w *World) CreateEntity() *Entity {
	id := w.nextId
	w.nextId++
	e := &Entity{
		id:         id,
		components: make(map[string]*Component),
	}
	w.entities[id] = e
	return e
}

// AddSystem 添加系统
func (w *World) AddSystem(s System) {
	w.systems = append(w.systems, s)
}

func (w *World) Update(world *World, deltaTime time.Duration) {
	for _, e := range world.entities {
		posComp, hasPos := e.GetComponent("position")
		velComp, hasVel := e.GetComponent("velocity")

		if hasPos && hasVel {
			pos := posComp.(*room.Position)
			vel := velComp.(*room.Velocity)

			pos.X += vel.VX * deltaTime.Seconds()
			pos.Y += vel.VY * deltaTime.Seconds()
		}
	}
}
