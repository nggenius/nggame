package space

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/nggenius/ngengine/common/event"
	"github.com/nggenius/ngengine/common/fsm"
	"github.com/nggenius/ngengine/core/rpc"
	"github.com/nggenius/ngengine/protocol"
	"github.com/nggenius/ngengine/share"
	"github.com/nggenius/ngengine/utils"
	"github.com/nggenius/nggame/define"

	"github.com/mysll/toolkit"
)

const (
	REGION_CREATING = iota + 1
	REGION_RUNNING
	REGION_CLOSING
	REGION_DELETING
	REGION_FAILED
)

type RegionInfo struct {
	define.Region
	Where   rpc.Mailbox
	Dest    rpc.Mailbox
	Players int
	Status  int
}

const (
	RS_NONE = iota
	RS_QUERY
	RS_RUNNING
	RS_OFFLINE
)

type RegionState struct {
	mailbox rpc.Mailbox
	regions []int
	players int
	state   int
}

func NewRegionState(mb rpc.Mailbox) *RegionState {
	s := new(RegionState)
	s.mailbox = mb
	s.regions = make([]int, 0, 10)
	return s
}

// Load 负载量，每运行一个场景（即使没有玩家)折算成10个玩家+玩家总数
func (r RegionState) Load() int {
	return len(r.regions)*10 + r.players
}

// HasRegion 是否存在某个区域
func (r RegionState) HasRegion(id int) bool {
	for k := range r.regions {
		if r.regions[k] == id {
			return true
		}
	}
	return false
}

// AddRegion 增加一个区域
func (r *RegionState) AddRegion(id int) {
	if r.HasRegion(id) {
		return
	}

	r.regions = append(r.regions, id)
}

// RemoveRegion 移除一个区域
func (r *RegionState) RemoveRegion(id int) {
	for k := range r.regions {
		if r.regions[k] == id {
			copy(r.regions[k:], r.regions[k+1:])
			r.regions = r.regions[:len(r.regions)-1]
		}
	}
}

type SpaceManage struct {
	ctx         *WorldSpaceModule
	MinRegions  int
	regiondef   map[int]define.Region
	regionmap   map[int]*RegionInfo
	regionstate []*RegionState
	firstLoad   sync.Once
	fsm         *fsm.FSM
}

func NewSpaceManage(ctx *WorldSpaceModule) *SpaceManage {
	s := new(SpaceManage)
	s.ctx = ctx
	s.regionmap = make(map[int]*RegionInfo)
	s.regiondef = make(map[int]define.Region)
	s.regionstate = make([]*RegionState, 0, 10)
	s.fsm = initState(s)
	return s
}

// RegionState 获取regionstate
func (s *SpaceManage) RegionState(id share.ServiceId) *RegionState {
	for k := range s.regionstate {
		if s.regionstate[k].mailbox.ServiceId() == id {
			return s.regionstate[k]
		}
	}

	return nil
}

// AddRegionState 增加一个regionstate
func (s *SpaceManage) AddRegionState(rs *RegionState) {
	s.regionstate = append(s.regionstate, rs)
}

func (s *SpaceManage) onServiceReady(e string, args ...interface{}) {
	info := args[0].(event.EventArgs)
	srv := s.ctx.Core.LookupService(info["id"].(share.ServiceId))
	if srv.Type == "region" {
		rs := s.RegionState(srv.Id)
		if rs == nil {
			rs = NewRegionState(*srv.Mailbox())
			s.AddRegionState(rs)
		}
		rs.state = RS_QUERY
		s.ctx.Core.MailtoAndCallback(nil, srv.Mailbox(), "Region.Query", s.onRegionStateQuery, srv.Id)
	}
}

// RefreshRegionState 刷新region state
func (s *SpaceManage) RefreshRegionState() {
	srvs := s.ctx.Core.LookupAllServiceByType("region")
	for _, srv := range srvs {
		rs := s.RegionState(srv.Id)
		if rs == nil {
			rs = NewRegionState(*srv.Mailbox())
			s.AddRegionState(rs)
		}

		rs.state = RS_QUERY

		s.ctx.Core.MailtoAndCallback(nil, srv.Mailbox(), "Region.Query", s.onRegionStateQuery, srv.Id)
	}
}

func (s *SpaceManage) onRegionStateQuery(p interface{}, rpcerr *rpc.Error, ar *utils.LoadArchive) {
	if rpcerr != nil && protocol.CheckRpcError(rpcerr) {
		s.ctx.Core.LogErr("rpc error:", rpcerr.Error())
		return
	}

	id := p.(share.ServiceId)

	rs := s.RegionState(id)
	if rs == nil {
		s.ctx.Core.LogWarn("region state not found")
		return
	}

	//TODO: 这里需要同步原来服务器的信息，主要是world异常关闭后进行重建
	rs.state = RS_RUNNING
	s.fsm.Dispatch(EREGION_RESP, nil)
	//s.CreateRegion(1)
}

// hasAllReady 检查所有的状态是否准备好
func (s *SpaceManage) hasAllReady() bool {
	for _, rs := range s.regionstate {
		if rs.state != RS_RUNNING {
			return false
		}
	}

	if len(s.regionstate) < s.MinRegions {
		return false
	}

	return true
}

// createAllRegion 创建所有的场景
func (s *SpaceManage) createAllRegion() error {
	for k := range s.regiondef {
		if s.FindRegionById(k) == nil {
			err := s.CreateRegion(k)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// FindRegionById 通过ID查找场景
func (s *SpaceManage) FindRegionById(id int) *RegionInfo {
	if r, has := s.regionmap[id]; has {
		return r
	}

	return nil
}

// findLowerLoad 获取一个负载较小的区域服务
func (s *SpaceManage) findLowerLoad() *RegionState {
	if len(s.regionstate) == 0 {
		return nil
	}
	low := s.regionstate[0].Load()
	rs := s.regionstate[0]
	for _, r := range s.regionstate {
		if r.Load() < low {
			rs = r
			low = r.Load()
		}
	}

	return rs
}

// CreateRegion 创建区域
func (s *SpaceManage) CreateRegion(id int) error {
	if _, has := s.regionmap[id]; has {
		return fmt.Errorf("region already created")
	}

	def, has := s.regiondef[id]
	if !has {
		return fmt.Errorf("region def not find")
	}

	rs := s.findLowerLoad()
	if rs == nil {
		return fmt.Errorf("region not found")
	}

	var r RegionInfo
	r.Id = id
	r.Region = def
	r.Status = REGION_CREATING
	r.Where = rs.mailbox

	s.regionmap[id] = &r
	rs.AddRegion(id)

	return s.ctx.Core.MailtoAndCallback(nil, &rs.mailbox, "Region.Create", s.onCreateRegion, id, r.Region)
}

func (s *SpaceManage) onCreateRegion(p interface{}, rpcerr *rpc.Error, ar *utils.LoadArchive) {
	id := p.(int)
	ri := s.FindRegionById(id)
	if ri == nil {
		s.ctx.Core.LogErr("region not found")
		return
	}

	if rpcerr != nil {
		ri.Status = REGION_FAILED
		s.ctx.Core.LogErr("region create failed", rpcerr.Error())
		return
	}

	var mb rpc.Mailbox
	err := ar.Get(&mb)
	if err != nil {
		s.ctx.Core.LogErr("get mailbox error")
		return
	}
	ri.Dest = mb
	ri.Status = REGION_RUNNING

	//s.ctx.Core.Mailto(nil, &mb, "GameScene.Test", "test")
	s.ctx.Core.LogInfo("region created,", ri)
	s.fsm.Dispatch(EREGION_CREATED, id)
}

// LoadResource 加载资源
func (s *SpaceManage) LoadResource(f string) bool {
	data, err := toolkit.ReadFile(f)
	if err != nil {
		return false
	}

	regions := make(map[string][]define.Region)
	err = json.Unmarshal(data, &regions)
	if err != nil {
		panic(err)
	}

	if r, ok := regions["Regions"]; ok {
		for k := range r {
			s.regiondef[r[k].Id] = r[k]
		}
		return true
	}

	return false
}
