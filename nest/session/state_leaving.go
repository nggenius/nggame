package session

import (
	"time"

	"github.com/nggenius/ngengine/common/fsm"
	"github.com/nggenius/ngmodule/store"
)

type Leaving struct {
	fsm.Default
	owner       *Session
	Idle        int32
	saveTimeout time.Time
	errors      int
}

func newLeaving(o *Session) *Leaving {
	s := new(Leaving)
	s.owner = o
	return s
}

func (s *Leaving) Init(r fsm.StateRegister) {
	r.AddHandle(ESTORED, s.OnStored)
}

func (s *Leaving) Enter() {
	s.saveTimeout = time.Now().Add(s.owner.ctx.saveTimeout)
}

func (s *Leaving) OnTimer() string {
	if time.Now().Sub(s.saveTimeout) > 0 {
		s.owner.SaveRole(store.STORE_SAVE_OFFLINE)
		s.saveTimeout = time.Now().Add(s.owner.ctx.saveTimeout)
		s.owner.ctx.Core.LogWarn("save role timeout")
	}
	return ""
}

func (s *Leaving) OnStored(event int, param interface{}) string {
	args := param.([2]interface{})
	if ok := args[0].(int32); ok == 0 {
		s.owner.Break()
		s.owner.DestroySelf()
	} else {
		s.owner.ctx.Core.LogWarn("save player failed")
	}
	return ""
}

func (s *Leaving) OnHandle(event int, param interface{}) string {
	s.owner.ctx.Core.LogWarnf("leaving state receive error event(%d)", event)
	return ""
}
