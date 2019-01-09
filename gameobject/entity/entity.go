package entity

import (
	"github.com/nggenius/ngengine/core/rpc"
	"github.com/nggenius/ngengine/core/service"
	"github.com/nggenius/ngengine/utils"
)

const (
	EXPOSE_NONE  = 0
	EXPOSE_OWNER = 1
	EXPOSE_OTHER = 2
	EXPOSE_ALL   = EXPOSE_OWNER | EXPOSE_OTHER
)

// 数据对象
type DataObject interface {
	// set id
	SetId(val int64)
	// db id
	DBId() int64
	// Archive 存档对象
	Archive() interface{}
	// Type 类型(对应xml里面的type)
	Type() string
	// Entity 类型(对应xml里面的name)
	Entity() string
	// AttrType 获取某个属性的类型
	AttrType(name string) string
	// FindAttr 获取属性
	FindAttr(name string) interface{}
	// SetAttr 设置属性
	SetAttr(name string, value interface{}) error
	// Expose 导出状态
	Expose(name string) int
	// AllAttr 所有属性名
	AllAttr() []string
	// AttrIndex 属性的索引编号
	AttrIndex(name string) int
	// 设置目击者
	SetWitness(w Witness)
	Serialize(ar *utils.StoreArchive) error
	Deserialize(ar *utils.LoadArchive) error
}

type Witness interface {
	// Posses 附加到一个数据对象
	Posses(e *Entity)
	// Silence 沉默状态
	Silence() bool
	// SetSilence 设置沉默状态
	SetSilence(bool)
	// Dummy 是否是复制对象
	Dummy() bool
	// SetDummy 设置为复制对象
	SetDummy(c bool)
	// Sync 同步状态
	Sync() bool
	// SetSync 设置同步状态
	SetSync(bool)
	// Original 原始对象
	Original() *rpc.Mailbox
	// SetOriginal 设置原始对象
	SetOriginal(m *rpc.Mailbox)
	// UpdateAttr 属性变动回调
	UpdateAttr(name string, val interface{}, old interface{})
	// UpdateTuple tuple变动回调
	UpdateTuple(name string, val interface{}, old interface{})
	// AddTableRow 表格增加行回调
	AddTableRow(name string, row int)
	// AddTableRowValue 表格增加行并设置值回调
	AddTableRowValue(name string, row int, val ...interface{})
	// SetTableRowValue 设置表格行
	SetTableRowValue(name string, row int, val ...interface{})
	// DelTableRow 删除表格行
	DelTableRow(name string, row int)
	// ClearTable 清除表格
	ClearTable(name string)
	// ChangeTable 表格单元格变动
	ChangeTable(name string, row, col int, val interface{})
	// ExistDummy 是否存在某个副本对象
	ExistDummy(dummy rpc.Mailbox) bool
	// AddDummy 关联一个副本对象
	AddDummy(dummy rpc.Mailbox, state int)
	// RemoveDummy 移除一个副本对象
	RemoveDummy(dummy rpc.Mailbox)
	// ChangeDummyState 更新副本对象的状态
	ChangeDummyState(dummy rpc.Mailbox, state int) error
}

type Observer interface {
	// AddAttrObserver 增加一个属性观察者
	AddAttrObserver(name string, observer AttrObserver) error
	// RemoveAttrObserver 删除属性观察者
	RemoveAttrObserver(name string)
	// AddTableObserver 增加表格观察者
	AddTableObserver(name string, observer TableObserver) error
	// RemoveTableObserver 删除表格观察者
	RemoveTableObserver(name string)
}

type Entity struct {
	*CacheData            // 临时数据
	d          DataObject // 数据对象
	w          Witness    // witness
	core       service.CoreAPI
	objid      rpc.Mailbox
}

func NewEntity(typ string) *Entity {
	d := Create(typ)
	if d == nil {
		return nil
	}
	e := new(Entity)
	e.CacheData = NewData()
	e.d = d
	w := NewWitness()
	w.Posses(e)
	e.w = w
	return e
}

// Data 获取数据对象
func (e *Entity) Data() DataObject {
	return e.d
}

// Witness 获取目击者
func (e *Entity) Witness() Witness {
	return e.w
}

// SetCore 设置core
func (e *Entity) SetCore(c service.CoreAPI) {
	e.core = c
}

// Core 获取Core接口
func (e *Entity) Core() service.CoreAPI {
	return e.core
}

// ObjId 唯一ID
func (e *Entity) ObjId() rpc.Mailbox {
	return e.objid
}

// SetObjId 设置唯一ID
func (e *Entity) SetObjId(id rpc.Mailbox) {
	e.objid = id
}

// Serialize 序列化
func (e *Entity) Serialize(ar *utils.StoreArchive) error {
	err := e.d.Serialize(ar)
	if err != nil {
		return err
	}

	return nil
}

// Deserialize 反序列化
func (e *Entity) Deserialize(ar *utils.LoadArchive) error {
	err := e.d.Deserialize(ar)
	if err != nil {
		return err
	}

	return nil
}
