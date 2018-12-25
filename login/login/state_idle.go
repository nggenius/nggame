package login

import (
	"github.com/nggenius/ngengine/common/fsm"
	"github.com/nggenius/nggame/proto/c2s"
)

type Idle struct {
	fsm.Default
	owner *Session
	idle  int32
}

func NewIdle(s *Session) *Idle {
	state := new(Idle)
	state.owner = s
	return state
}

func (s *Idle) Init(r fsm.StateRegister) {
	r.AddHandle(LOGIN, s.OnLogin)
	r.AddHandle(BREAK, s.OnBreak)
}

func (s *Idle) OnTimer() string {
	s.idle++
	if s.idle > 60 {
		s.owner.Break()
	}
	return ""
}

func (s *Idle) OnLogin(event int, param interface{}) string {
	s.owner.Login(param.(*c2s.Login))
	return SLOGGING
}

func (s *Idle) OnBreak(event int, param interface{}) string {
	s.owner.DestroySelf()
	return fsm.STOP
}

func (s *Idle) OnHandle(event int, param interface{}) string {
	s.owner.ctx.Core.LogWarnf("idle state receive error event(%d)", event)
	return ""
}
