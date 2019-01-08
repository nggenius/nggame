package entity

// Witness的作用是监听数据对象的变动，作为一个目击者。由数据对象直接持有。
// 目击者作为一个事件集散地，对所有的数据变动的事件进行调度。由第三方注册。
// 目击者不关注变动的细节，只进行转发，由第三方进行细节的处理。
import (
	"fmt"

	"github.com/nggenius/ngengine/core/rpc"
)

const (
	TABLE_INIT = iota + 1
	TABLE_ADD_ROW
	TABLE_REMOVE_ROW
	TABLE_CLEAR_ROW
	TABLE_GRID_CHANGE
	TABLE_SET_ROW
)

const (
	DUMMY_STATE_NONE = iota
	DUMMY_STATE_REPLICATE
	DUMMY_STATE_READY
	DUMMY_STATE_REMOVE
)

type AttrObserver interface {
	Init(e *Entity)
	UpdateAttr(name string, val interface{}, old interface{})
	UpdateTuple(name string, val interface{}, old interface{})
}

type TableObserver interface {
	Init(e *Entity)
	UpdateTable(name string, op_type, row, col int)
}

type EntityWitness struct {
	e             *Entity
	original      *rpc.Mailbox
	dummys        map[rpc.Mailbox]int
	dummy         bool // 是否是副本
	sync          bool // 同步状态
	silence       bool // 沉默状态
	attrobserver  map[string]AttrObserver
	tableobserver map[string]TableObserver
}

func NewWitness() Witness {
	w := new(EntityWitness)
	w.attrobserver = make(map[string]AttrObserver)
	w.tableobserver = make(map[string]TableObserver)
	w.dummys = make(map[rpc.Mailbox]int)
	return w
}

// Posses 附加到一个数据对象
func (o *EntityWitness) Posses(e *Entity) {
	o.e = e
	o.e.d.SetWitness(o)
}

// Silence 沉默状态
func (o *EntityWitness) Silence() bool {
	return o.silence
}

// SetSilence 设置沉默状态
func (o *EntityWitness) SetSilence(s bool) {
	o.silence = s
}

// AddAttrObserver 增加属性观察者,这里的name是观察者的标识符，不是属性名称
func (o *EntityWitness) AddAttrObserver(name string, observer AttrObserver) error {
	if _, dup := o.attrobserver[name]; dup {
		return fmt.Errorf("add attr observer twice %s", name)
	}

	o.attrobserver[name] = observer
	observer.Init(o.e)
	return nil
}

// RemoveAttrObserver 删除属性观察者
func (o *EntityWitness) RemoveAttrObserver(name string) {
	delete(o.attrobserver, name)
}

// AddTableObserver 增加表格观察者,这里的name是观察者的标识符，不是表格名称
func (o *EntityWitness) AddTableObserver(name string, observer TableObserver) error {
	if _, dup := o.tableobserver[name]; dup {
		return fmt.Errorf("add table observer twice %s", name)
	}

	o.tableobserver[name] = observer
	observer.Init(o.e)
	return nil
}

// RemoveTableObserver 删除表格观察者
func (o *EntityWitness) RemoveTableObserver(name string) {
	delete(o.tableobserver, name)
}

// UpdateAttr 对象属性变动(由object调用)
func (o *EntityWitness) UpdateAttr(name string, val interface{}, old interface{}) {
	if o.dummy && !o.sync { // 需要操作远程对象
		o.RemoteUpdateAttr(name, val)
		return
	}
	if o.silence {
		return
	}
	for _, observer := range o.attrobserver {
		observer.UpdateAttr(name, val, old)
	}
}

// UpdateTuple 对象tupele属性变动(由object调用)
func (o *EntityWitness) UpdateTuple(name string, val interface{}, old interface{}) {
	if o.dummy && !o.sync { // 需要操作远程对象
		o.RemoteUpdateTuple(name, val)
		return
	}
	if o.silence {
		return
	}
	for _, observer := range o.attrobserver {
		observer.UpdateTuple(name, val, old)
	}
}

// AddTableRow 对象表格增加一行(由object调用)
func (o *EntityWitness) AddTableRow(name string, row int) {
	if o.dummy && !o.sync { // 需要操作远程对象
		o.RemoteAddTableRow(name, row)
		return
	}
	if o.silence {
		return
	}
	for _, observer := range o.tableobserver {
		observer.UpdateTable(name, TABLE_ADD_ROW, row, 0)
	}
}

// AddTableRowValue 对象表格增加一行，并设置值(由object调用)
func (o *EntityWitness) AddTableRowValue(name string, row int, val ...interface{}) {
	if o.dummy && !o.sync { // 需要操作远程对象
		o.RemoteAddTableRowValue(name, row, val...)
		return
	}
	if o.silence {
		return
	}
	for _, observer := range o.tableobserver {
		observer.UpdateTable(name, TABLE_ADD_ROW, row, 0)
	}
}

// SetTableRowValue 设置表格一行的值(由object调用)
func (o *EntityWitness) SetTableRowValue(name string, row int, val ...interface{}) {
	if o.dummy && !o.sync { // 需要操作远程对象
		o.RemoteSetTableRowValue(name, row, val...)
		return
	}
	if o.silence {
		return
	}
	for _, observer := range o.tableobserver {
		observer.UpdateTable(name, TABLE_SET_ROW, row, 0)
	}
}

// DelTableRow 对象表格删除一行(由object调用)
func (o *EntityWitness) DelTableRow(name string, row int) {
	if o.dummy && !o.sync { // 需要操作远程对象
		o.RemoteDelTableRow(name, row)
		return
	}
	if o.silence {
		return
	}
	for _, observer := range o.tableobserver {
		observer.UpdateTable(name, TABLE_REMOVE_ROW, row, 0)
	}
}

// ClearTable 对象表格清除所有行(由object调用)
func (o *EntityWitness) ClearTable(name string) {
	if o.dummy && !o.sync { // 需要操作远程对象
		o.RemoteClearTable(name)
		return
	}
	if o.silence {
		return
	}
	for _, observer := range o.tableobserver {
		observer.UpdateTable(name, TABLE_CLEAR_ROW, 0, 0)
	}
}

// ChangeTable 对象表格单元格更新(由object调用)
func (o *EntityWitness) ChangeTable(name string, row, col int, val interface{}) {
	if o.dummy && !o.sync { // 需要操作远程对象
		o.RemoteChangeTable(name, row, col, val)
		return
	}
	if o.silence {
		return
	}
	for _, observer := range o.tableobserver {
		observer.UpdateTable(name, TABLE_GRID_CHANGE, row, col)
	}
}
