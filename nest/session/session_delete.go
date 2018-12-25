package session

import (
	"github.com/nggenius/ngengine/common/fsm"
	"github.com/nggenius/ngengine/share"
)

type Deleting struct {
	fsm.Default
	owner *Session
	Idle  int32
}

func newDeleting(o *Session) *Deleting {
	s := new(Deleting)
	s.owner = o
	return s
}

func (s *Deleting) Init(r fsm.StateRegister) {
	r.AddHandle(EDELETED, s.OnDeleting)
	r.AddHandle(EBREAK, s.OnBreak)
}

func (s *Deleting) OnTimer() string {
	s.Idle++
	if s.Idle > 60 {
		s.owner.Error(share.ERR_TIME_OUT)
		return SLOGGED
	}
	return ""
}

func (s *Deleting) OnDeleting(event int, param interface{}) string {
	errcode := param.(int32)
	if errcode != 0 {
		s.owner.Error(errcode)
		return SLOGGED
	}
	s.owner.QueryRoleInfo()
	return SLOGGED
}

func (s *Deleting) OnBreak(event int, param interface{}) string {
	s.owner.DestroySelf()
	return fsm.STOP
}

func (s *Deleting) OnHandle(event int, param interface{}) string {
	s.owner.ctx.Core.LogWarnf("deleting state receive error event(%d)", event)
	return ""
}
