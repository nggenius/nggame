package gameobject

import (
	"time"

	"github.com/nggenius/nggame/gameobject/entity"

	"github.com/nggenius/ngengine/core/service"
)

type ComponentInfo struct {
	started   bool
	comp      Component
	useUpdate bool
}

type Component interface {
	SetGameObject(GameObject)
	GameObject() GameObject
	Create()
	Start()
	Update(delta time.Duration)
	Destroy()
	Enable() bool
	SetEnable(e bool)
}

type GameComponent struct {
	gameobject GameObject
	enable     bool
}

// SetGameObject 设置当前附加的对象，由GameObject调用
func (g *GameComponent) SetGameObject(o GameObject) {
	g.gameobject = o
}

// GameObject 获取当前附加的对象
func (g *GameComponent) GameObject() GameObject {
	return g.gameobject
}

// Core
func (g *GameComponent) Core() service.CoreAPI {
	return g.gameobject.Core()
}

// Spirit 获取数据对象
func (g *GameComponent) Spirit() *entity.Entity {
	return g.gameobject.Spirit()
}

// Enable 获取当前组件的开启状态
func (g *GameComponent) Enable() bool {
	return g.enable
}

// SetEnable 设置当前组件的开启状态
func (g *GameComponent) SetEnable(e bool) {
	g.enable = e
}

// Create 组件被创建时调用
func (g *GameComponent) Create() {

}

// Start 组件开启时被调用
func (g *GameComponent) Start() {

}

// Update 组件的周期定时回调
func (g *GameComponent) Update(delta time.Duration) {

}

// Destroy 组件被销毁
func (g *GameComponent) Destroy() {

}
