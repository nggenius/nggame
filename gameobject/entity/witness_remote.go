package entity

import (
	"fmt"

	"github.com/nggenius/ngengine/core/rpc"
)

// ExistDummy 是否存在某个副本
func (o *EntityWitness) ExistDummy(dummy rpc.Mailbox) bool {
	_, ok := o.dummys[dummy]
	return ok
}

// 关联一个副本对象
func (o *EntityWitness) AddDummy(dummy rpc.Mailbox, state int) {
	if _, ok := o.dummys[dummy]; ok {
		return
	}

	o.dummys[dummy] = state
}

// 移除一个副本对象
func (o *EntityWitness) RemoveDummy(dummy rpc.Mailbox) {
	delete(o.dummys, dummy)
}

// 更新副本对象的状态
func (o *EntityWitness) ChangeDummyState(dummy rpc.Mailbox, state int) error {
	if _, ok := o.dummys[dummy]; ok {
		o.dummys[dummy] = state
		return nil
	}

	return fmt.Errorf("dummy not found, %v", dummy)
}

// 是否是复制对象
func (o *EntityWitness) Dummy() bool {
	return o.dummy
}

// 设置为复制对象
func (o *EntityWitness) SetDummy(c bool) {
	o.dummy = c
}

// 同步状态
func (o *EntityWitness) Sync() bool {
	return o.sync
}

// 设置同步状态
func (o *EntityWitness) SetSync(s bool) {
	o.sync = s
}

// 原始对象
func (o *EntityWitness) Original() *rpc.Mailbox {
	return o.original
}

// 设置原始对象
func (o *EntityWitness) SetOriginal(m *rpc.Mailbox) {
	o.original = m
}

// RemoteUpdateAttr 远程更新属性
func (o *EntityWitness) RemoteUpdateAttr(name string, val interface{}) {
	if o.original == nil {
		o.e.Core().LogErr("original is nil")
		return
	}

	o.e.Core().Mailto(&o.e.objid, o.original, "object.UpdateAttr", name, val)
}

// RemoteUpdateTuple 远程更新tuple
func (o *EntityWitness) RemoteUpdateTuple(name string, val interface{}) {
	if o.original == nil {
		o.e.Core().LogErr("original is nil")
		return
	}
	o.e.Core().Mailto(&o.e.objid, o.original, "object.UpdateTuple", name, val)
}

func (o *EntityWitness) RemoteAddTableRow(name string, row int) {
	if o.original == nil {
		o.e.Core().LogErr("original is nil")
		return
	}
	o.e.Core().Mailto(&o.e.objid, o.original, "object.AddTableRow", name, row)
}

func (o *EntityWitness) RemoteAddTableRowValue(name string, row int, val ...interface{}) {
	if o.original == nil {
		o.e.Core().LogErr("original is nil")
		return
	}
	o.e.Core().Mailto(&o.e.objid, o.original, "object.AddTableRowValue", name, row, val)
}

func (o *EntityWitness) RemoteSetTableRowValue(name string, row int, val ...interface{}) {
	if o.original == nil {
		o.e.Core().LogErr("original is nil")
		return
	}
	o.e.Core().Mailto(&o.e.objid, o.original, "object.SetTableRowValue", name, row, val)
}

func (o *EntityWitness) RemoteDelTableRow(name string, row int) {
	if o.original == nil {
		o.e.Core().LogErr("original is nil")
		return
	}
	o.e.Core().Mailto(&o.e.objid, o.original, "object.DelTableRow", name, row)
}

func (o *EntityWitness) RemoteClearTable(name string) {
	if o.original == nil {
		o.e.Core().LogErr("original is nil")
		return
	}
	o.e.Core().Mailto(&o.e.objid, o.original, "object.ClearTable", name)
}

func (o *EntityWitness) RemoteChangeTable(name string, row, col int, val interface{}) {
	if o.original == nil {
		o.e.Core().LogErr("original is nil")
		return
	}

	o.e.Core().Mailto(&o.e.objid, o.original, "object.ChangeTable", name, row, col, val)
}
