package login

import (
	"github.com/nggenius/ngengine/common/fsm"
	"github.com/nggenius/ngengine/share"
	"github.com/nggenius/nggame/gameobject/entity/inner"
)

type Logging struct {
	fsm.Default
	owner *Session
	idle  int32
}

func NewLogging(s *Session) *Logging {
	state := new(Logging)
	state.owner = s
	return state
}

func (s *Logging) Init(r fsm.StateRegister) {
	r.AddHandle(LOGIN_RESULT, s.OnLogin)
	r.AddHandle(BREAK, s.OnBreak)
}

func (s *Logging) OnTimer() string {
	s.idle++
	if s.idle > 60 {
		s.owner.Error(share.S2C_ERR_SERVICE_INVALID)
		s.owner.Break()
		return ""
	}
	return ""
}

func (s *Logging) OnLogin(event int, param interface{}) string {
	args := param.([2]interface{})
	if s.owner.LoginResult(args[0].(int32), args[1].(*inner.Account)) {
		return SLOGGED
	}
	return ""
}

func (s *Logging) OnBreak(event int, param interface{}) string {
	s.owner.DestroySelf()
	return fsm.STOP
}

func (s *Logging) OnHandle(event int, param interface{}) string {
	s.owner.ctx.Core.LogWarnf("logging state receive error event(%d)", event)
	return ""
}
