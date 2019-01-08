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
	leaving    bool
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
	// 在场景中
	if !s.owner.region.IsNil() {
		s.owner.LevelRegion()
		s.leaving = true
		s.Idle = 0
	}
}

func (s *Offline) OnTimer() string {
	s.Idle++
	if s.leaving {
		if s.Idle > 60 { // 退出超时
			if !s.owner.region.IsNil() {
				s.owner.LevelRegion()
				s.Idle = 0
			}
		}
		return ""
	}

	if time.Now().Sub(s.remainTime) > 0 {
		s.owner.SaveRole(store.STORE_SAVE_OFFLINE)
		s.owner.RegionRemove()
		return SLEAVING
	}
	return ""
}

func (s *Offline) OnRemainTime(event int, param interface{}) string {
	args := param.([2]interface{})
	t1 := args[0].(time.Duration)
	data := args[1].([]byte)
	s.remainTime = time.Now().Add(t1)
	if err := s.owner.SyncData(data); err != nil {
		panic(err)
	}
	s.leaving = false
	return ""
}

func (s *Offline) OnHandle(event int, param interface{}) string {
	s.owner.ctx.Core.LogWarnf("offline state receive error event(%d)", event)
	return ""
}
