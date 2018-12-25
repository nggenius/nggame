package session

import (
	"github.com/nggenius/ngengine/common/fsm"
	"github.com/nggenius/ngengine/share"
)

type Idle struct {
	fsm.Default
	owner *Session
	Idle  int32
}

func newIdle(o *Session) *Idle {
	s := new(Idle)
	s.owner = o
	return s
}

func (s *Idle) Init(r fsm.StateRegister) {
	r.AddHandle(ELOGIN, s.OnLogin)
	r.AddHandle(EBREAK, s.OnBreak)
}

func (s *Idle) OnTimer() string {
	s.Idle++
	if s.Idle > 60 {
		s.owner.Break()
		return ""
	}
	return ""
}

func (s *Idle) OnLogin(event int, param interface{}) string {
	token := param.(string)
	if s.owner.ValidToken(token) {
		// TODO: 这里要进行排队检查
		if !s.owner.QueryRoleInfo() {
			s.owner.Error(share.ERR_SYSTEM_ERROR)
			return ""
		}
		return SLOGGED
	}
	// 验证失败直接踢下线
	s.owner.Break()
	return ""
}

func (s *Idle) OnBreak(event int, param interface{}) string {
	s.owner.DestroySelf()
	return fsm.STOP
}

func (s *Idle) OnHandle(event int, param interface{}) string {
	s.owner.ctx.Core.LogWarnf("idle state receive error event(%d)", event)
	return ""
}
