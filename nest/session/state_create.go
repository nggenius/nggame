package session

import (
	"github.com/nggenius/ngengine/common/fsm"
	"github.com/nggenius/nggame/define"
)

type CreateRole struct {
	fsm.Default
	owner *Session
	Idle  int32
}

func newCreateRole(o *Session) *CreateRole {
	s := new(CreateRole)
	s.owner = o
	return s
}

func (s *CreateRole) Init(r fsm.StateRegister) {
	r.AddHandle(EBREAK, s.OnBreak)
	r.AddHandle(ECREATED, s.OnCreated)
}

func (s *CreateRole) OnTimer() string {
	s.Idle++
	if s.Idle > 60 {
		s.owner.Error(define.ERR_CREATE_TIMEOUT)
		return SLOGGED
	}
	return ""
}

func (s *CreateRole) OnBreak(event int, param interface{}) string {
	s.owner.DestroySelf()
	return fsm.STOP
}

func (s *CreateRole) OnCreated(event int, param interface{}) string {
	errcode := param.(int32)
	if errcode != 0 {
		s.owner.Error(errcode)
		return SLOGGED
	}
	if !s.owner.QueryRoleInfo() {
		s.owner.Error(-1)
	}
	return SLOGGED
}

func (s *CreateRole) OnHandle(event int, param interface{}) string {
	s.owner.ctx.Core.LogWarnf("create role state receive error event(%d)", event)
	return ""
}
