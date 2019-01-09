package gameobject

import (
	"time"

	"github.com/nggenius/nggame/gameobject/entity"

	"github.com/nggenius/ngengine/core/service"
	"github.com/nggenius/ngengine/utils"
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
	BeforeSerialize()
	Serialize(ar *utils.StoreArchive) error
	Deserialize(ar *utils.LoadArchive) error
	Replication() bool
	SetReplication(v bool)
}

type GameComponent struct {
	gameobject  GameObject
	enable      bool
	replication bool
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

// Replicate 是否复制
func (g *GameComponent) Replication() bool {
	return g.replication
}

// SetReplicate set replicate
func (g *GameComponent) SetReplication(v bool) {
	g.replication = v
}

// Serialize 序列化
func (g *GameComponent) Serialize(ar *utils.StoreArchive) error {
	return nil
}

// Deserialize 反序列化
func (g *GameComponent) Deserialize(ar *utils.LoadArchive) error {
	return nil
}

// BeforeSerialize 准备序列化
func (g *GameComponent) BeforeSerialize() {

}
