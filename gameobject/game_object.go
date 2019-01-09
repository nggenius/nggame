package gameobject

import (
	"github.com/nggenius/ngengine/utils"

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
	// Spirit 数据对象
	Spirit() *entity.Entity
	// Behavior 行为对象
	Behavior() Behavior
	// EntityType 数据对象类型
	EntityType() string
	// Serialize 序列化接口
	Serialize(ar *utils.StoreArchive) error
	// Deserialize 反序列化接口
	Deserialize(ar *utils.LoadArchive) error
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

func (g *BaseObject) Serialize(ar *utils.StoreArchive) error {
	g.object.Behavior().Base().BeforeSerialize()

	// 序列化自身数据
	err := g.object.Spirit().Serialize(ar)
	if err != nil {
		return err
	}

	// 序列化组件
	pos := ar.Len() // 占位符
	comps := g.object.Behavior().Base().component
	ar.Put(int16(0))
	var count int16
	for k, c := range comps {
		if c.comp.Replication() {
			ar.PutString(k)
			cbegin := ar.Len() // 组件开始位置，记录组件数据长度，以便组件不存在时可以跳过
			ar.Put(int16(0))
			if err := c.comp.Serialize(ar); err != nil {
				return err
			}
			l := int16(ar.Len() - cbegin - 2)
			if l < 0 {
				panic("size exceed")
			}

			ar.WriteAt(cbegin, int16(ar.Len()-cbegin)-2)
			count++
		}
	}
	err = ar.WriteAt(pos, count)
	if err != nil {
		return err
	}
	return nil
}

func (g *BaseObject) Deserialize(ar *utils.LoadArchive) error {
	err := g.object.Spirit().Deserialize(ar)
	if err != nil {
		return err
	}
	count, err := ar.GetUint16()
	if err != nil {
		return err
	}
	for i := 0; i < int(count); i++ {
		cname, err := ar.GetString()
		if err != nil {
			return err
		}
		len, err := ar.GetInt16()
		if err != nil {
			return err
		}
		comp := g.object.Behavior().GetComponent(cname)
		if comp == nil {
			ar.Seek(ar.Position()+int(len), 0)
			continue
		}
		err = comp.Deserialize(ar)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *BaseObject) SetCap(cap int) {
	g.object.Behavior().Base().SetCap(cap)
}
