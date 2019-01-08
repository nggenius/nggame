package gameobject

import (
	"fmt"
	"time"

	"github.com/nggenius/ngengine/core/rpc"
	"github.com/nggenius/nggame/gameobject/entity"
	"github.com/nggenius/ngmodule/object"
)

type Behavior interface {
	// Base get BaseBehavior
	Base() *BaseBehavior
	// GameObject 获取GameObject
	GameObject() GameObject
	// 获取游戏对象类型
	GameObjectType() int
	// Spirit 获取数据对象
	Spirit() *entity.Entity
	// SetTransport 设置连接
	SetTransport(t *Transport)
	// Transport 获取连接
	Transport() *Transport
	// AddComponent 增加组件
	AddComponent(name string, com Component, update bool) error
	// RemoveComponent 移除组件
	RemoveComponent(name string)
	// GetComponent 获取组件
	GetComponent(name string) Component
	// Parent 父对象
	Parent() GameObject
	// SetParent 设置父对象
	SetParent(p GameObject)
	// ParentIndex 获取在父对象中的位置
	ParentIndex() int
	// SetParentIndex 设置在父对象中的位置
	SetParentIndex(pos int)
	// Cap 获取容量，返回-1表示不限容量
	Cap() int
	// SetCap 设置容量，并初始化容器
	SetCap(cap int) error
	// 回调
	OnCreate()
	OnDestroy()
	OnUpdate(delta time.Duration)
	OnDelete()
}

type BaseBehavior struct {
	c          *Container
	delete     bool
	index      int // 在factory中的索引
	client     rpc.Mailbox
	delegate   object.Delegate
	component  map[string]ComponentInfo
	transport  *Transport
	gameobject GameObject
	parent     GameObject
	pi         int  // 在父对象容器中的位置
	update     bool // 是否每一帧进行调用
}

// Base get BaseBehavior
func (b *BaseBehavior) Base() *BaseBehavior {
	return b
}

// Init 初始化
func (b *BaseBehavior) Init(e GameObject) {
	b.gameobject = e
}

// Parent 父对象
func (b *BaseBehavior) Parent() GameObject {
	return b.parent
}

// SetParent 设置父对象
func (b *BaseBehavior) SetParent(p GameObject) {
	b.parent = p
}

// ParentIndex 获取在父对象中的位置
func (b *BaseBehavior) ParentIndex() int {
	return b.pi
}

// SetParentIndex 设置在父对象中的位置
func (b *BaseBehavior) SetParentIndex(pi int) {
	b.pi = pi
}

// GameObject GameObject
func (b *BaseBehavior) GameObject() GameObject {
	return b.gameobject
}

// SetTransport 设置连接
func (b *BaseBehavior) SetTransport(t *Transport) {
	b.transport = t
}

// Transport 获取连接
func (b *BaseBehavior) Transport() *Transport {
	return b.transport
}

// Prepare 预处理
func (b *BaseBehavior) Prepare() {
	b.component = make(map[string]ComponentInfo)
}

// ObjectType 获取对象类型
func (b *BaseBehavior) ObjectType() int {
	return OBJECT_NONE
}

// Create 构造函数
func (b *BaseBehavior) Create() {
	if b.delegate != nil && b.gameobject.Spirit() != nil {
		b.delegate.Invoke(E_ON_CREATE, b.gameobject.Spirit().ObjId(), rpc.NullMailbox)
	}

	b.gameobject.Behavior().OnCreate()
}

// Spirit 获取数据对象
func (b *BaseBehavior) Spirit() *entity.Entity {
	return b.gameobject.Spirit()
}

// Destroy 准备销毁
func (b *BaseBehavior) Destroy() {
	if b.delegate != nil && b.gameobject.Spirit() != nil {
		b.delegate.Invoke(E_ON_DESTROY, b.gameobject.Spirit().ObjId(), rpc.NullMailbox)
	}
	b.delete = true

	for name, comp := range b.component {
		comp.comp.Destroy() // 销毁组件
		delete(b.component, name)
	}

	b.gameobject.Behavior().OnDestroy()
}

// Delete 正式开始删除
func (b *BaseBehavior) Delete() {
	b.gameobject.Behavior().OnDelete()
}

// Alive 是否还活着
func (b *BaseBehavior) Alive() bool {
	return !b.delete
}

// SetIndex 设置索引，由factory调用，不要手工调用
func (b *BaseBehavior) SetIndex(index int) {
	b.index = index
}

// Index 获取索引
func (b *BaseBehavior) Index() int {
	return b.index
}

// SetDelegate 设置事件代理
func (b *BaseBehavior) SetDelegate(d object.Delegate) {
	b.delegate = d
}

// Client 客户端地址
func (b *BaseBehavior) Client() rpc.Mailbox {
	return b.client
}

// SetClient 设置客户端地址
func (b *BaseBehavior) SetClient(mb rpc.Mailbox) {
	b.client = mb
}

// Update
func (b *BaseBehavior) Update(delta time.Duration) {
	for _, comp := range b.component {
		if !comp.comp.Enable() {
			continue
		}
		if !comp.started {
			comp.comp.Start()
			comp.started = true
		}
		if comp.useUpdate {
			comp.comp.Update(delta)
		}
	}

	b.gameobject.Behavior().OnUpdate(delta)
}

// GetComponent 获取组件
func (b *BaseBehavior) GetComponent(name string) Component {
	if comp, has := b.component[name]; has {
		return comp.comp
	}
	return nil
}

// AddComponent 增加组件
func (b *BaseBehavior) AddComponent(name string, com Component, update bool) error {
	if _, has := b.component[name]; has {
		return fmt.Errorf("component has register twice, %s ", name)
	}

	b.component[name] = ComponentInfo{
		started:   false,
		comp:      com,
		useUpdate: update,
	}

	com.SetGameObject(b.gameobject)
	com.SetEnable(true)
	// 调用初始化函数
	com.Create()
	return nil
}

// RemoveComponent 移除组件
func (b *BaseBehavior) RemoveComponent(name string) {
	if comp, has := b.component[name]; has {
		comp.comp.Destroy() // 销毁组件
		delete(b.component, name)
	}
}

func (b *BaseBehavior) OnCreate() {

}

func (b *BaseBehavior) OnDestroy() {

}

func (b *BaseBehavior) OnDelete() {

}

func (b *BaseBehavior) OnUpdate(delta time.Duration) {

}
