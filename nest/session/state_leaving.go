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
	saved       bool
	removed     bool
}

func newLeaving(o *Session) *Leaving {
	s := new(Leaving)
	s.owner = o
	return s
}

func (s *Leaving) Init(r fsm.StateRegister) {
	r.AddHandle(ESTORED, s.OnStored)
	r.AddHandle(EREGIONREMOVE, s.OnRegionRemove)
}

func (s *Leaving) Enter() {
	s.saveTimeout = time.Now().Add(s.owner.ctx.saveTimeout)
}

func (s *Leaving) OnTimer() string {
	if !s.saved && time.Now().Sub(s.saveTimeout) > 0 {
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
		s.saved = true
		if s.removed {
			s.owner.DestroySelf()
			return fsm.STOP
		}
	} else {
		s.owner.ctx.Core.LogWarn("save player failed")
	}
	return ""
}

func (s *Leaving) OnRegionRemove(event int, param interface{}) string {
	s.removed = true
	if s.saved {
		s.owner.DestroySelf()
		return fsm.STOP
	}
	return ""
}

func (s *Leaving) OnHandle(event int, param interface{}) string {
	s.owner.ctx.Core.LogWarnf("leaving state receive error event(%d)", event)
	return ""
}
