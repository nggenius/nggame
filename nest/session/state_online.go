package session

import (
	"github.com/nggenius/ngengine/common/fsm"
	"github.com/nggenius/ngengine/core/rpc"
	"github.com/nggenius/nggame/define"
)

type Online struct {
	fsm.Default
	owner  *Session
	Idle   int32
	online bool
}

func newOnline(o *Session) *Online {
	s := new(Online)
	s.owner = o
	return s
}

func (s *Online) Init(r fsm.StateRegister) {
	r.AddHandle(EFREGION, s.OnFindRegion)
	r.AddHandle(EONLINE, s.Online)
	r.AddHandle(EBREAK, s.OnBreak)
}

func (s *Online) Enter() {
	if s.owner.enterregion {
		s.owner.FindRegion()
		s.online = false
		return
	}
	s.owner.Dispatch(EONLINE, nil)
}

func (s *Online) OnTimer() string {
	s.Idle++
	if !s.online {
		if s.Idle > 60 {
			s.owner.Error(define.ERR_ENTER_REGION_FAILED)
			return SLOGGED
		}
	}
	return ""
}

func (s *Online) OnFindRegion(event int, param interface{}) string {
	args := param.([2]interface{})
	errcode := args[0].(int32)
	if errcode != 0 {
		s.owner.Error(errcode)
		return SLOGGED
	}

	r := args[1].(rpc.Mailbox)
	if r == rpc.NullMailbox {
		s.owner.Error(define.ERR_REGION_NOT_FOUND)
		return SLOGGED
	}

	if err := s.owner.EnterRegion(r); err != nil {
		s.owner.Error(define.ERR_ENTER_REGION_FAILED)
		return SLOGGED
	}
	return ""
}

func (s *Online) Online(event int, param interface{}) string {
	s.online = true
	s.Idle = 0
	s.owner.ctx.Core.LogDebug("player online")
	return ""
}

func (s *Online) OnBreak(event int, param interface{}) string {
	s.owner.ctx.Core.LogInfo("client break")
	return SOFFLINE
}

func (s *Online) OnHandle(event int, param interface{}) string {
	s.owner.ctx.Core.LogWarnf("online state receive error event(%d)", event)
	return ""
}

func (s *Online) Exit() {
	if !s.online {
		if s.owner.gameobject != nil {
			s.owner.ctx.factory.Destroy(s.owner.gameobject)
			s.owner.gameobject = nil
		}
	}
}
