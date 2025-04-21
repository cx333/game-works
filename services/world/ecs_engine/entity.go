package ecs_engine

import "time"

/**
 * @Author: wgl
 * @Description:
 * @File: entity
 * @Version: 1.0.0
 * @Date: 2025/4/21 22:21
 */

// EntityId 唯一实体ID类型
type EntityId int64

// Component 所有组件需实现的接口
type Component interface{}

// System 系统接口，每帧更新
type System interface {
	Update(world *World, deltaTime time.Duration)
}

// Entity 实体
type Entity struct {
	//mu        sync.RWMutex
	id         EntityId
	components map[string]*Component
}

// AddComponent 添加组件
func (e *Entity) AddComponent(name string, comp *Component) {
	e.components[name] = comp
}

// GetComponent 获取组件
func (e *Entity) GetComponent(name string) (Component, bool) {
	comp, ok := e.components[name]
	return comp, ok
}
