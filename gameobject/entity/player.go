// Code generated by data parser.
// DO NOT EDIT!
package entity

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/nggenius/ngmodule/object"

	"github.com/mysll/toolkit"
)

var _ = json.Marshal
var _ = toolkit.ParseNumber

// tuple LandPos 位置
type PlayerLandPos_t struct {
	X      float64 // X
	Y      float64 // Y
	Z      float64 // Z
	Orient float64 // Orient
}

// tuple LandPos construct
func NewPlayerLandPos() *PlayerLandPos_t {
	landpos := &PlayerLandPos_t{}
	return landpos
}

// tuple LandPos Set
func (landpos *PlayerLandPos_t) Set(x float64, y float64, z float64, orient float64) {

	landpos.X = x
	landpos.Y = y
	landpos.Z = z
	landpos.Orient = orient
}

// tuple LandPos Get
func (landpos *PlayerLandPos_t) Get() (x float64, y float64, z float64, orient float64) {
	return landpos.X, landpos.Y, landpos.Z, landpos.Orient
}

// tuple LandPos equal other
func (landpos *PlayerLandPos_t) Equal(other PlayerLandPos_t) bool {
	if (landpos.X == other.X) && (landpos.Y == other.Y) && (landpos.Z == other.Z) && (landpos.Orient == other.Orient) {
		return true
	}
	return false
}

// record Toolbox row define
type PlayerToolbox_c struct {
	Id     int64 //
	Amount int32 //
}

// record Toolbox 道具(表格测试)
type PlayerToolbox_r struct {
	root object.Object
	data [10]*PlayerToolbox_c
	Row  []*PlayerToolbox_c
}

// record  Toolbox  serial
type PlayerToolboxJson struct {
	ColName []string
	ColType []string
	Row     [][]interface{}
}

// record Toolbox construct
func NewPlayerToolbox(root object.Object) *PlayerToolbox_r {
	toolbox := &PlayerToolbox_r{root: root}
	toolbox.Row = toolbox.data[:0]
	return toolbox
}

// row count
func (r *PlayerToolbox_r) Rows() int {
	return len(r.Row)
}

// row cap
func (r *PlayerToolbox_r) Cap() int {
	return cap(r.data)
}

// get Id
func (r *PlayerToolbox_r) Id(rownum int) (int64, error) {
	if rownum < 0 || rownum >= len(r.Row) {
		return 0, fmt.Errorf("row num error")
	}
	return r.Row[rownum].Id, nil
}

// set Id
func (r *PlayerToolbox_r) SetId(rownum int, id int64) error {
	if r.root != nil && r.root.Dummy() && !r.root.Sync() {
		r.root.ChangeTable("Toolbox", rownum, 0, id)
		return nil
	}
	if rownum < 0 || rownum >= len(r.Row) {
		return fmt.Errorf("row num error")
	}
	if r.Row[rownum].Id != id {
		r.Row[rownum].Id = id
		if r.root != nil {
			r.root.ChangeTable("Toolbox", rownum, 0, id)
		}
	}
	return nil
}

// get Amount
func (r *PlayerToolbox_r) Amount(rownum int) (int32, error) {
	if rownum < 0 || rownum >= len(r.Row) {
		return 0, fmt.Errorf("row num error")
	}
	return r.Row[rownum].Amount, nil
}

// set Amount
func (r *PlayerToolbox_r) SetAmount(rownum int, amount int32) error {
	if r.root != nil && r.root.Dummy() && !r.root.Sync() {
		r.root.ChangeTable("Toolbox", rownum, 1, amount)
		return nil
	}
	if rownum < 0 || rownum >= len(r.Row) {
		return fmt.Errorf("row num error")
	}
	if r.Row[rownum].Amount != amount {
		r.Row[rownum].Amount = amount
		if r.root != nil {
			r.root.ChangeTable("Toolbox", rownum, 1, amount)
		}
	}
	return nil
}

// set row value
func (r *PlayerToolbox_r) SetRowValue(rownum int, id int64, amount int32) error {
	if r.root != nil && r.root.Dummy() && !r.root.Sync() {
		r.root.SetTableRowValue("Toolbox", rownum, id, amount)
		return nil
	}

	if rownum < 0 || rownum >= len(r.Row) {
		return fmt.Errorf("row num error")
	}
	/*
		if r.Row[rownum].Id != id {
			r.Row[rownum].Id = id
			if r.root != nil {
				r.root.ChangeTable("Toolbox", rownum, 0, id)
			}
		}
		if r.Row[rownum].Amount != amount {
			r.Row[rownum].Amount = amount
			if r.root != nil {
				r.root.ChangeTable("Toolbox", rownum, 1, amount)
			}
		}
	*/

	r.Row[rownum].Id = id
	r.Row[rownum].Amount = amount
	if r.root != nil {
		r.root.SetTableRowValue("Toolbox", rownum, id, amount)
	}
	return nil
}

// get row value
func (r *PlayerToolbox_r) RowValue(rownum int) (int64, int32, error) {
	var row PlayerToolbox_c
	if rownum < 0 || rownum >= len(r.Row) {
		return row.Id, row.Amount, fmt.Errorf("row num error")
	}

	row = *r.Row[rownum]
	return row.Id, row.Amount, nil
}

// add row
func (r *PlayerToolbox_r) AddRow(rownum int) (int, error) {
	if r.root != nil && r.root.Dummy() && !r.root.Sync() {
		r.root.AddTableRow("Toolbox", rownum)
		return -1, nil
	}
	if len(r.Row) > cap(r.data) { // full
		return -1, fmt.Errorf("record PlayerToolbox is full")
	}

	if rownum < -1 || rownum >= cap(r.data) { // out of range
		return -1, fmt.Errorf("record PlayerToolbox row %d out of range", rownum)
	}

	size := len(r.Row)
	row := &PlayerToolbox_c{}
	r.Row = r.data[:size+1]
	if rownum == -1 || rownum == size {
		r.Row[size] = row
		return size, nil
	}
	copy(r.Row[rownum+1:], r.Row[rownum:])
	r.Row[rownum] = row
	if r.root != nil {
		r.root.AddTableRow("Toolbox", rownum)
	}
	return rownum, nil
}

// add row value
func (r *PlayerToolbox_r) AddRowValue(rownum int, id int64, amount int32) (int, error) {
	if r.root != nil && r.root.Dummy() && !r.root.Sync() {
		r.root.AddTableRowValue("Toolbox", rownum, id, amount)
		return -1, nil
	}
	if len(r.Row) > cap(r.data) { // full
		return -1, fmt.Errorf("record PlayerToolbox is full")
	}

	if rownum < -1 || rownum >= cap(r.data) { // out of range
		return -1, fmt.Errorf("record PlayerToolbox row %d out of range", rownum)
	}

	size := len(r.Row)
	row := &PlayerToolbox_c{id, amount}
	r.Row = r.data[:size+1]
	if rownum == -1 || rownum == size {
		r.Row[size] = row
		if r.root != nil {
			r.root.AddTableRowValue("Toolbox", rownum, id, amount)
		}
		return size, nil
	}
	copy(r.Row[rownum+1:], r.Row[rownum:])
	r.Row[rownum] = row
	if r.root != nil {
		r.root.AddTableRowValue("Toolbox", rownum, id, amount)
	}
	return rownum, nil
}

// del row
func (r *PlayerToolbox_r) Del(rownum int) error {
	if r.root != nil && r.root.Dummy() && !r.root.Sync() {
		r.root.DelTableRow("Toolbox", rownum)
		return nil
	}
	if rownum < 0 || rownum >= len(r.Row) {
		return fmt.Errorf("row num error")
	}
	copy(r.Row[rownum:], r.Row[rownum+1:])
	r.Row = r.data[:len(r.Row)-1]
	if r.root != nil {
		r.root.DelTableRow("Toolbox", rownum)
	}
	return nil
}

// clear
func (r *PlayerToolbox_r) Clear() {
	if r.root != nil && r.root.Dummy() && !r.root.Sync() {
		r.root.ClearTable("Toolbox")
		return
	}
	r.Row = r.data[:0]
	if r.root != nil {
		r.root.ClearTable("Toolbox")
	}
}

// json encode interface
func (r *PlayerToolbox_r) Marshal() ([]byte, error) {
	return r.pack()
}

// json decode interface
func (r *PlayerToolbox_r) Unmarshal(data []byte) error {
	return r.unpack(data)
}

// xorm encode interface
func (r *PlayerToolbox_r) ToDB() ([]byte, error) {
	return r.pack()
}

// xorm decode interface
func (r *PlayerToolbox_r) FromDB(data []byte) error {
	return r.unpack(data)
}

// gob encode interface
func (r *PlayerToolbox_r) GobEncode() ([]byte, error) {
	return r.pack()
}

// gob decode interface
func (r *PlayerToolbox_r) GobDecode(data []byte) error {
	return r.unpack(data)
}

// record Toolbox pack
func (r *PlayerToolbox_r) pack() ([]byte, error) {
	j := &PlayerToolboxJson{}
	j.ColName = make([]string, 2)
	j.ColType = make([]string, 2)

	j.ColName[0] = "Id"
	j.ColType[0] = "int64"
	j.ColName[1] = "Amount"
	j.ColType[1] = "int32"

	j.Row = make([][]interface{}, len(r.Row))
	for k, row := range r.Row {
		if row == nil {
			panic("row is nil")
		}
		j.Row[k] = make([]interface{}, 0, 2)
		j.Row[k] = append(j.Row[k], row.Id, row.Amount)
	}

	return json.Marshal(j)
}

// record Toolbox unpack
func (r *PlayerToolbox_r) unpack(data []byte) error {
	r.Row = r.data[:0]
	j := &PlayerToolboxJson{}
	err := json.Unmarshal(data, j)
	if err != nil {
		return err
	}

	for _, row := range j.Row {
		if len(r.Row) > cap(r.data) {
			break
		}
		toolboxrow := &PlayerToolbox_c{}
		for k, col := range row {
			switch j.ColName[k] {
			case "Id":
				if j.ColType[k] == "int64" {
					toolkit.ParseNumber(col, &toolboxrow.Id)
				}
			case "Amount":
				if j.ColType[k] == "int32" {
					toolkit.ParseNumber(col, &toolboxrow.Amount)
				}
			}
		}
		r.Row = r.data[:len(r.Row)+1]
		r.Row[len(r.Row)-1] = toolboxrow
	}
	return nil
}

// tuple Pos 位置
type PlayerPos_t struct {
	X float32 //
	Y float32 //
	Z float32 //
}

// tuple Pos construct
func NewPlayerPos() *PlayerPos_t {
	pos := &PlayerPos_t{}
	return pos
}

// tuple Pos Set
func (pos *PlayerPos_t) Set(x float32, y float32, z float32) {

	pos.X = x
	pos.Y = y
	pos.Z = z
}

// tuple Pos Get
func (pos *PlayerPos_t) Get() (x float32, y float32, z float32) {
	return pos.X, pos.Y, pos.Z
}

// tuple Pos equal other
func (pos *PlayerPos_t) Equal(other PlayerPos_t) bool {
	if (pos.X == other.X) && (pos.Y == other.Y) && (pos.Z == other.Z) {
		return true
	}
	return false
}

// Player archive
type PlayerArchive struct {
	root object.Object `xorm:"-"`
	flag int           `xorm:"-"`

	Id        int64
	Name      string           `xorm:"varchar(128)"` // 玩家名
	LandScene int64            // 场景编号
	LandPos   *PlayerLandPos_t `xorm:"json"` // 位置
	Toolbox   *PlayerToolbox_r `xorm:"json"` // 道具(表格测试)
	Pos       *PlayerPos_t     `xorm:"json"` // 位置
	Orient    float32          // 朝向(弧度)

}

// Player archive construct
func NewPlayerArchive(root object.Object) *PlayerArchive {
	archive := &PlayerArchive{root: root}

	archive.LandPos = NewPlayerLandPos()
	archive.Toolbox = NewPlayerToolbox(root)
	archive.Pos = NewPlayerPos()

	return archive
}

// archive table name
func (a PlayerArchive) TableName() string {
	return "player"
}

// set id
func (a *PlayerArchive) SetId(val int64) {
	a.Id = val
}

// db id
func (a *PlayerArchive) DBId() int64 {
	return a.Id
}

// Player archive
type PlayerArchiveBak struct {
	Id        int64
	Name      string           `xorm:"varchar(128)"` // 玩家名
	LandScene int64            // 场景编号
	LandPos   *PlayerLandPos_t `xorm:"json"` // 位置
	Toolbox   *PlayerToolbox_r `xorm:"json"` // 道具(表格测试)
	Pos       *PlayerPos_t     `xorm:"json"` // 位置
	Orient    float32          // 朝向(弧度)
}

// archive table name
func (a PlayerArchiveBak) TableName() string {
	return "player_bak"
}

// set id
func (a *PlayerArchiveBak) SetId(val int64) {
	a.Id = val
}

// db id
func (a *PlayerArchiveBak) DBId() int64 {
	return a.Id
}

// Player attr
type PlayerAttr struct {
	root object.Object

	GroupId     int32 // 分组
	Invisible   byte  // 是否不可见(1不可见)
	VisualRange int32 // 可视范围
}

// Player attr construct
func NewPlayerAttr(root object.Object) *PlayerAttr {
	attr := &PlayerAttr{root: root}

	return attr
}

// Player
type Player struct {
	object.ObjectWitness
	archive *PlayerArchive // archive
	attr    *PlayerAttr    // attr
}

// Player construct
func NewPlayer() *Player {
	o := &Player{}
	o.archive = NewPlayerArchive(o)
	o.attr = NewPlayerAttr(o)
	o.Witness(o)
	return o
}

// Player store
func (o *Player) Store() {
}

// Player type
func (o *Player) Type() string {
	return "player"
}

// Player entity name
func (o *Player) Entity() string {
	return "Player"
}

// Player load
func (o *Player) Load() {
}

// set id
func (o *Player) SetId(val int64) {
	o.archive.SetId(val)
}

// db id
func (o *Player) DBId() int64 {
	return o.archive.DBId()
}

// get archive
func (o *Player) Archive() interface{} {
	return o.archive
}

// get attr
func (o *Player) Attr() interface{} {
	return o.attr
}

// set Name 玩家名
func (o *Player) SetName(name string) {
	if o.Dummy() && !o.Sync() {
		o.UpdateAttr("Name", name, nil)
		return
	}
	if o.archive.Name == name {
		return
	}
	old := o.archive.Name
	o.archive.Name = name
	o.UpdateAttr("Name", name, old)
}

// get Name 玩家名
func (o *Player) Name() string {
	return o.archive.Name
}

// set LandScene 场景编号
func (o *Player) SetLandScene(landscene int64) {
	if o.Dummy() && !o.Sync() {
		o.UpdateAttr("LandScene", landscene, nil)
		return
	}
	if o.archive.LandScene == landscene {
		return
	}
	old := o.archive.LandScene
	o.archive.LandScene = landscene
	o.UpdateAttr("LandScene", landscene, old)
}

// get LandScene 场景编号
func (o *Player) LandScene() int64 {
	return o.archive.LandScene
}

// set LandPos 位置
func (o *Player) SetLandPos(landpos PlayerLandPos_t) {
	if o.Dummy() && !o.Sync() {
		o.UpdateTuple("LandPos", landpos, nil)
		return
	}
	old := *o.archive.LandPos
	*o.archive.LandPos = landpos
	o.UpdateTuple("LandPos", landpos, old)
}

// set LandPos detail
func (o *Player) SetLandPosXYZOrient(x float64, y float64, z float64, orient float64) {
	if o.Dummy() && !o.Sync() {
		val := PlayerLandPos_t{x, y, z, orient}
		o.UpdateTuple("LandPos", val, nil)
		return
	}
	old := *o.archive.LandPos
	o.archive.LandPos.Set(x, y, z, orient)
	o.UpdateTuple("LandPos", *o.archive.LandPos, old)
}

// get LandPos 位置
func (o *Player) LandPos() PlayerLandPos_t {
	return *o.archive.LandPos
}

// get LandPos detail
func (o *Player) LandPosXYZOrient() (x float64, y float64, z float64, orient float64) {
	return o.archive.LandPos.Get()
}

// set Toolbox 道具(表格测试)
func (o *Player) SetToolbox(toolbox *PlayerToolbox_r) {
	panic("Toolbox can't set")
}

// get Toolbox 道具(表格测试)
func (o *Player) Toolbox() *PlayerToolbox_r {
	return o.archive.Toolbox
}

// set GroupId 分组
func (o *Player) SetGroupId(groupid int32) {
	if o.Dummy() && !o.Sync() {
		o.UpdateAttr("GroupId", groupid, nil)
		return
	}
	if o.attr.GroupId == groupid {
		return
	}
	old := o.attr.GroupId
	o.attr.GroupId = groupid
	o.UpdateAttr("GroupId", groupid, old)
}

// get GroupId 分组
func (o *Player) GroupId() int32 {
	return o.attr.GroupId
}

// set Invisible 是否不可见(1不可见)
func (o *Player) SetInvisible(invisible byte) {
	if o.Dummy() && !o.Sync() {
		o.UpdateAttr("Invisible", invisible, nil)
		return
	}
	if o.attr.Invisible == invisible {
		return
	}
	old := o.attr.Invisible
	o.attr.Invisible = invisible
	o.UpdateAttr("Invisible", invisible, old)
}

// get Invisible 是否不可见(1不可见)
func (o *Player) Invisible() byte {
	return o.attr.Invisible
}

// set VisualRange 可视范围
func (o *Player) SetVisualRange(visualrange int32) {
	if o.Dummy() && !o.Sync() {
		o.UpdateAttr("VisualRange", visualrange, nil)
		return
	}
	if o.attr.VisualRange == visualrange {
		return
	}
	old := o.attr.VisualRange
	o.attr.VisualRange = visualrange
	o.UpdateAttr("VisualRange", visualrange, old)
}

// get VisualRange 可视范围
func (o *Player) VisualRange() int32 {
	return o.attr.VisualRange
}

// set Pos 位置
func (o *Player) SetPos(pos PlayerPos_t) {
	if o.Dummy() && !o.Sync() {
		o.UpdateTuple("Pos", pos, nil)
		return
	}
	old := *o.archive.Pos
	*o.archive.Pos = pos
	o.UpdateTuple("Pos", pos, old)
}

// set Pos detail
func (o *Player) SetPosXYZ(x float32, y float32, z float32) {
	if o.Dummy() && !o.Sync() {
		val := PlayerPos_t{x, y, z}
		o.UpdateTuple("Pos", val, nil)
		return
	}
	old := *o.archive.Pos
	o.archive.Pos.Set(x, y, z)
	o.UpdateTuple("Pos", *o.archive.Pos, old)
}

// get Pos 位置
func (o *Player) Pos() PlayerPos_t {
	return *o.archive.Pos
}

// get Pos detail
func (o *Player) PosXYZ() (x float32, y float32, z float32) {
	return o.archive.Pos.Get()
}

// set Orient 朝向(弧度)
func (o *Player) SetOrient(orient float32) {
	if o.Dummy() && !o.Sync() {
		o.UpdateAttr("Orient", orient, nil)
		return
	}
	if o.archive.Orient == orient {
		return
	}
	old := o.archive.Orient
	o.archive.Orient = orient
	o.UpdateAttr("Orient", orient, old)
}

// get Orient 朝向(弧度)
func (o *Player) Orient() float32 {
	return o.archive.Orient
}

// attr type
func (o *Player) AttrType(name string) string {
	switch name {
	case "Name":
		return "string"
	case "LandScene":
		return "int64"
	case "LandPos":
		return "tuple"
	case "Toolbox":
		return "table"
	case "GroupId":
		return "int32"
	case "Invisible":
		return "byte"
	case "VisualRange":
		return "int32"
	case "Pos":
		return "tuple"
	case "Orient":
		return "float32"
	default:
		return "unknown"
	}
}

// attr expose info
func (o *Player) Expose(name string) int {
	switch name {
	case "Name":
		return object.EXPOSE_OWNER
	case "LandScene":
		return object.EXPOSE_NONE
	case "LandPos":
		return object.EXPOSE_NONE
	case "Toolbox":
		return object.EXPOSE_NONE
	case "GroupId":
		return object.EXPOSE_NONE
	case "Invisible":
		return object.EXPOSE_NONE
	case "VisualRange":
		return object.EXPOSE_NONE
	case "Pos":
		return object.EXPOSE_ALL
	case "Orient":
		return object.EXPOSE_ALL
	default:
		panic("unknown")
	}
}

// get all attr name
func (o *Player) AllAttr() []string {
	return []string{"Name", "LandScene", "LandPos", "Toolbox", "GroupId", "Invisible", "VisualRange", "Pos", "Orient"}
}

// get attr index by name
func (o *Player) AttrIndex(name string) int {
	switch name {
	case "Name":
		return 0
	case "LandScene":
		return 1
	case "LandPos":
		return 2
	case "Toolbox":
		return 3
	case "GroupId":
		return 4
	case "Invisible":
		return 5
	case "VisualRange":
		return 6
	case "Pos":
		return 7
	case "Orient":
		return 8
	default:
		return -1
	}
}

// get attr value
func (o *Player) FindAttr(name string) interface{} {
	switch name {
	case "Name":
		return o.archive.Name
	case "LandScene":
		return o.archive.LandScene
	case "LandPos":
		return *o.archive.LandPos
	case "Toolbox":
		return o.archive.Toolbox
	case "GroupId":
		return o.attr.GroupId
	case "Invisible":
		return o.attr.Invisible
	case "VisualRange":
		return o.attr.VisualRange
	case "Pos":
		return *o.archive.Pos
	case "Orient":
		return o.archive.Orient
	default:
		return nil
	}
}

// set attr value
func (o *Player) SetAttr(name string, value interface{}) error {
	switch name {
	case "Name":
		if v, ok := value.(string); ok {
			o.SetName(v)
			return nil
		}
		return fmt.Errorf("attr Name type not match")
	case "LandScene":
		if v, ok := value.(int64); ok {
			o.SetLandScene(v)
			return nil
		}
		return fmt.Errorf("attr LandScene type not match")
	case "LandPos":
		if v, ok := value.(PlayerLandPos_t); ok {
			o.SetLandPos(v)
			return nil
		}
		return fmt.Errorf("attr LandPos type not match")
	case "Toolbox":
		if v, ok := value.(*PlayerToolbox_r); ok {
			o.SetToolbox(v)
			return nil
		}
		return fmt.Errorf("attr Toolbox type not match")
	case "GroupId":
		if v, ok := value.(int32); ok {
			o.SetGroupId(v)
			return nil
		}
		return fmt.Errorf("attr GroupId type not match")
	case "Invisible":
		if v, ok := value.(byte); ok {
			o.SetInvisible(v)
			return nil
		}
		return fmt.Errorf("attr Invisible type not match")
	case "VisualRange":
		if v, ok := value.(int32); ok {
			o.SetVisualRange(v)
			return nil
		}
		return fmt.Errorf("attr VisualRange type not match")
	case "Pos":
		if v, ok := value.(PlayerPos_t); ok {
			o.SetPos(v)
			return nil
		}
		return fmt.Errorf("attr Pos type not match")
	case "Orient":
		if v, ok := value.(float32); ok {
			o.SetOrient(v)
			return nil
		}
		return fmt.Errorf("attr Orient type not match")
	default:
		return fmt.Errorf("attr %s not found", name)
	}
}

// gob interface
func (o *Player) GobEncode() ([]byte, error) {
	w := new(bytes.Buffer)
	encoder := gob.NewEncoder(w)
	var err error

	err = encoder.Encode(o.archive)
	if err != nil {
		return nil, err
	}
	err = encoder.Encode(o.attr)
	if err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}

func (o *Player) GobDecode(buf []byte) error {
	r := bytes.NewBuffer(buf)
	decoder := gob.NewDecoder(r)
	var err error

	err = decoder.Decode(o.archive)
	if err != nil {
		return err
	}
	err = decoder.Decode(o.attr)
	if err != nil {
		return err
	}
	return nil
}

// gob register
func init() {
	gob.Register(&Player{})
	gob.Register(&PlayerArchive{})
	gob.Register([]*Player{})
	gob.Register([]*PlayerArchive{})
	registObject("entity.Player", func() object.Object { return NewPlayer() })
}
