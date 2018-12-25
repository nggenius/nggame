package login

import (
	"github.com/nggenius/ngengine/common/fsm"
	"github.com/nggenius/ngengine/share"
)

type Logged struct {
	fsm.Default
	owner *Session
	idle  int32
}

func NewLogged(s *Session) *Logged {
	state := new(Logged)
	state.owner = s
	return state
}

func (s *Logged) Init(r fsm.StateRegister) {
	r.AddHandle(NEST_RESULT, s.OnNest)
	r.AddHandle(BREAK, s.OnBreak)
}

func (s *Logged) OnTimer() string {
	s.idle++
	if s.idle > 60 {
		s.owner.Error(share.S2C_ERR_SERVICE_INVALID)
		s.owner.Break()
		return ""
	}
	return ""
}

func (s *Logged) OnNest(event int, param interface{}) string {
	args := param.([2]interface{})
	if s.owner.NestResult(args[0].(int32), args[1].(string)) {
		s.idle = 0 // 1 分钟后退出
		return ""
	}
	return SIDLE //重新登录
}

func (s *Logged) OnBreak(event int, param interface{}) string {
	s.owner.DestroySelf()
	return fsm.STOP
}

func (s *Logged) OnHandle(event int, param interface{}) string {
	s.owner.ctx.Core.LogWarnf("logged state receive error event(%d)", event)
	return ""
}
