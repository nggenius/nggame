package gameobject

import (
	"bytes"
	"encoding/gob"

	"github.com/nggenius/ngengine/core/rpc"
	"github.com/nggenius/ngengine/core/service"
	"github.com/nggenius/nggame/gameobject/entity"
	"github.com/nggenius/ngmodule/object"
)

const (
	OBJECT_NONE = iota
	OBJECT_SCENE
	OBJECT_PLAYER
	OBJECT_ITEM
	OBJECT_NPC
	OBJECT_ASSIST
	OBJECT_MAX
)

// GameObjectEqual 判断两个对象是否相等
func GameObjectEqual(l GameObject, r GameObject) bool {
	if l.Spirit() == nil || r.Spirit() == nil {
		panic("object is nil")
	}

	return l.Spirit().ObjId() == r.Spirit().ObjId()
}

type GameObject interface {
	Spirit() *entity.Entity
	Behavior() Behavior
	EntityType() string
	Serialize() ([]byte, error)
	Deserialize([]byte) error
	// ObjectType 获取对象类型
	ObjectType() string
	// ObjId 唯一ID
	ObjId() rpc.Mailbox
	// SetObjId 设置唯一ID
	SetObjId(id rpc.Mailbox)
	// Factory 所属的工厂
	Factory() *object.Factory
	// SetFactory 设置工厂，由工厂主动调用
	SetFactory(f *object.Factory)
	// Core 获取Core接口
	Core() service.CoreAPI
	// SetCap 设置对象容量
	SetCap(int)
}

type BaseObject struct {
	service.CoreAPI
	object  GameObject
	factory *object.Factory
}

func (g *BaseObject) Init(o interface{}) {
	g.object = o.(GameObject)
	g.object.Behavior().Base().Init(g.object)
}

func (g *BaseObject) Index() int {
	return g.object.Behavior().Base().Index()
}

func (g *BaseObject) SetIndex(i int) {
	g.object.Behavior().Base().SetIndex(i)
}

func (g *BaseObject) SetObjId(mb rpc.Mailbox) {
	g.object.Spirit().SetObjId(mb)
}

func (g *BaseObject) ObjId() rpc.Mailbox {
	return g.object.Spirit().ObjId()
}

func (g *BaseObject) Factory() *object.Factory {
	return g.factory
}

func (g *BaseObject) SetFactory(f *object.Factory) {
	g.factory = f
}

func (g *BaseObject) Core() service.CoreAPI {
	return g.CoreAPI
}

func (g *BaseObject) SetCore(c service.CoreAPI) {
	g.CoreAPI = c
	g.object.Spirit().SetCore(c)
}

func (g *BaseObject) Prepare() {
	g.object.Behavior().Base().Prepare()
}

func (g *BaseObject) Create() {
	g.object.Behavior().Base().Create()
}

func (g *BaseObject) Destroy() {
	g.object.Behavior().Base().Destroy()
}

func (g *BaseObject) Delete() {
	g.object.Behavior().Base().Delete()
}

func (g *BaseObject) Alive() bool {
	return g.object.Behavior().Base().Alive()
}

func (g *BaseObject) SetDelegate(d object.Delegate) {
	g.object.Behavior().Base().SetDelegate(d)
}

func (g *BaseObject) Serialize() ([]byte, error) {
	d := g.object.Spirit().Data()
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(d)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (g *BaseObject) Deserialize(b []byte) error {
	d := g.object.Spirit().Data()
	buf := bytes.NewBuffer(b)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(d)
	if err != nil {
		return err
	}
	return nil
}

func (g *BaseObject) SetCap(cap int) {
	g.object.Behavior().Base().SetCap(cap)
}
