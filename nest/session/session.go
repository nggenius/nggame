package session

import (
	"time"

	"github.com/nggenius/ngengine/common/fsm"
	"github.com/nggenius/ngengine/core/rpc"
	"github.com/nggenius/nggame/gameobject"
	"github.com/nggenius/nggame/gameobject/entity/inner"
	"github.com/nggenius/nggame/proto/c2s"
	"github.com/nggenius/nggame/proto/s2c"
)

type SessionDB map[uint64]*Session

// session保存当前连接和角色的相关信息
type Session struct {
	*fsm.FSM
	ctx                *SessionModule
	id                 uint64
	Account            string
	Mailbox            *rpc.Mailbox
	delete             bool
	gameobject         gameobject.GameObject
	landscene          int64
	lx, ly, lz, orient float64
	region             rpc.Mailbox
	autoenter          bool
	offlineTimeout     time.Duration
	remainTime         time.Duration
}

func NewSession(id uint64, ctx *SessionModule) *Session {
	s := &Session{}
	s.ctx = ctx
	s.id = id
	s.FSM = initState(s)
	s.autoenter = true
	return s
}

func (s *Session) SetLandInfo(scene int64, x, y, z, o float64) {
	s.landscene = scene
	s.lx = x
	s.ly = y
	s.lz = z
	s.orient = o
}

func (s *Session) SetGameObject(g gameobject.GameObject) {
	if s.gameobject != nil {
		s.ctx.factory.Destroy(s.gameobject)
	}
	if s.Mailbox != nil {
		tr := gameobject.NewTransport(s.ctx.Core, *s.Mailbox)
		g.Behavior().SetTransport(tr)
	}
	s.gameobject = g
}

func (s *Session) GameObject() gameobject.GameObject {
	return s.gameobject
}

// 删除自己
func (s *Session) DestroySelf() {
	s.delete = true
	if s.gameobject != nil {
		s.ctx.factory.Destroy(s.gameobject)
		s.gameobject = nil
	}
	s.ctx.deleted.PushBack(s.id)
}

// 断开客户端的连接
func (s *Session) Break() {
	s.ctx.Core.Break(s.id)
}

// 验证token
func (s *Session) ValidToken(token string) bool {
	if s.ctx.cache.Valid(s.Account, token) {
		return true
	}
	return false
}

// 查询玩家信息
func (s *Session) QueryRoleInfo() bool {
	if err := s.ctx.account.requestRoleInfo(s); err == nil {
		return true
	}
	return false
}

// 发送角色信息
func (s *Session) SendRoleInfo(role []*inner.Role) {
	s.ctx.Core.LogDebug("role info", role)
	roles := &s2c.RoleInfo{}
	roles.Roles = make([]s2c.Role, 0, len(role))
	for k := range role {
		r := s2c.Role{}
		r.RoleId = role[k].Id
		r.Index = role[k].Index
		r.Name = role[k].RoleName
		roles.Roles = append(roles.Roles, r)
	}

	s.ctx.Core.Mailto(nil, s.Mailbox, "Account.Roles", roles)
}

// CreateRole 创建角色
func (s *Session) CreateRole(info c2s.CreateRole) error {
	return s.ctx.account.CreateRole(s, info)
}

// ChooseRole 选择角色
func (s *Session) ChooseRole(info c2s.ChooseRole) error {
	return s.ctx.account.ChooseRole(s, info)
}

// DeleteRole 删除角色
func (s *Session) DeleteRole(info c2s.DeleteRole) error {
	return s.ctx.account.DeleteRole(s, info)
}

// SaveRole 保存角色数据
func (s *Session) SaveRole(stype int) error {
	return s.ctx.account.SaveRole(s, stype)
}

func (s *Session) SyncData(data []byte) error {
	f := s.gameobject.Factory()
	return f.Sync(s.gameobject, data)
}

func (s *Session) FindRegion() error {
	return s.ctx.account.FindRegion(s, s.landscene, s.lx, s.ly, s.lz)
}

func (s *Session) EnterRegion(r rpc.Mailbox) error {
	s.region = r
	err := s.ctx.account.EnterRegion(s, r)
	if err != nil {
		s.ctx.Core.LogErr("enter region failed")
	}
	return err
}

func (s *Session) LevelRegion() error {
	return s.ctx.account.LeaveRegion(s)
}

func (s *Session) RegionRemove() error {
	return s.ctx.account.RemovePlayer(s)
}

func (s *Session) Error(errcode int32) {
	err := s2c.Error{}
	err.ErrCode = errcode
	s.ctx.Core.Mailto(nil, s.Mailbox, "system.Error", &err)
}
