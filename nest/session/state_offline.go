package session

import (
	"time"

	"github.com/nggenius/ngengine/common/fsm"
	"github.com/nggenius/ngmodule/store"
)

type Offline struct {
	fsm.Default
	owner      *Session
	Idle       int32
	remainTime time.Time
}

func newOffline(o *Session) *Offline {
	s := new(Offline)
	s.owner = o
	return s
}

func (s *Offline) Init(r fsm.StateRegister) {
	r.AddHandle(EREMAINTIME, s.OnRemainTime)
}

func (s *Offline) Enter() {
	s.remainTime = time.Now().Add(s.owner.ctx.offlinetime)
}

func (s *Offline) OnTimer() string {
	if time.Now().Sub(s.remainTime) > 0 {
		s.owner.SaveRole(store.STORE_SAVE_OFFLINE)
		return SLEAVING
	}
	return ""
}

func (s *Offline) OnRemainTime(event int, param interface{}) string {
	rt := param.(int)
	s.remainTime = time.Now().Add(time.Duration(rt) * time.Second)
	return ""
}

func (s *Offline) OnHandle(event int, param interface{}) string {
	s.owner.ctx.Core.LogWarnf("offline state receive error event(%d)", event)
	return ""
}
