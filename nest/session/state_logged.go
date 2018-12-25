package session

import (
	"github.com/nggenius/ngengine/common/fsm"
	"github.com/nggenius/ngengine/share"
	"github.com/nggenius/nggame/gameobject/entity/inner"
	"github.com/nggenius/nggame/proto/c2s"
)

type Logged struct {
	fsm.Default
	owner *Session
	Idle  int32
}

func newLogged(o *Session) *Logged {
	s := new(Logged)
	s.owner = o
	return s
}

func (s *Logged) Init(r fsm.StateRegister) {
	r.AddHandle(EBREAK, s.OnBreak)
	r.AddHandle(EROLEINFO, s.OnRoleInfo)
	r.AddHandle(EDELETE, s.OnDelete)
	r.AddHandle(ECREATE, s.OnCreate)
	r.AddHandle(ECHOOSE, s.OnChoose)
}

func (s *Logged) OnTimer() string {
	s.Idle++
	if s.Idle > 60 {
		s.owner.Break()
		return ""
	}
	return ""
}

func (s *Logged) OnBreak(event int, param interface{}) string {
	s.owner.DestroySelf()
	return fsm.STOP
}

func (s *Logged) OnRoleInfo(event int, param interface{}) string {
	args := param.([2]interface{})
	errcode := args[0].(int32)
	roles := args[1].([]*inner.Role)
	if errcode != 0 {
		s.owner.Error(errcode)
		return ""
	}

	s.owner.SendRoleInfo(roles)
	s.Idle = 0
	return ""
}

func (s *Logged) OnDelete(event int, param interface{}) string {
	args := param.(c2s.DeleteRole)
	if err := s.owner.DeleteRole(args); err != nil {
		s.owner.Error(share.ERR_SYSTEM_ERROR)
		return ""
	}
	return SDELETE
}

func (s *Logged) OnCreate(event int, param interface{}) string {
	args := param.(c2s.CreateRole)
	if err := s.owner.CreateRole(args); err != nil {
		s.owner.Error(share.ERR_SYSTEM_ERROR)
		return ""
	}
	return SCREATE
}

func (s *Logged) OnChoose(event int, param interface{}) string {
	args := param.(c2s.ChooseRole)
	if err := s.owner.ChooseRole(args); err != nil {
		s.owner.Error(share.ERR_SYSTEM_ERROR)
		return ""
	}
	return SCHOOSE
}

func (s *Logged) OnHandle(event int, param interface{}) string {
	s.owner.ctx.Core.LogWarnf("logged state receive error event(%d)", event)
	return ""
}
