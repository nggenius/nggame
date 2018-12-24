package gameobject

import (
	"time"

	"github.com/nggenius/ngengine/core/service"

	"github.com/nggenius/ngmodule/object"
)

type Component interface {
	SetOwner(GameObject)
	Owner() GameObject
	Create()
	Start()
	Update(delta time.Duration)
	Destroy()
	Enable() bool
	SetEnable(e bool)
}

type GameComponent struct {
	owner  GameObject
	enable bool
}

// SetGameObject 设置当前附加的对象，由GameObject调用
func (g *GameComponent) SetOwner(obj GameObject) {
	g.owner = obj
}

// GameObject 获取当前附加的对象
func (g *GameComponent) Owner() GameObject {
	return g.owner
}

// Core
func (g *GameComponent) Core() service.CoreAPI {
	return g.owner.Spirit().Core()
}

// Spirit 获取数据对象
func (g *GameComponent) Spirit() object.Object {
	return g.owner.Spirit()
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
