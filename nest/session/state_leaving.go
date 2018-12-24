package session

import (
	"time"

	"github.com/nggenius/ngengine/common/fsm"
)

type leaving struct {
	fsm.Default
	owner       *Session
	Idle        int32
	saveTimeout time.Time
	errors      int
}

func (l *leaving) Enter() {
	l.saveTimeout = time.Now().Add(l.owner.ctx.saveTimeout)
}

func (l *leaving) Handle(event int, param interface{}) string {
	switch event {
	case ESTORED:
		if ok := param.(int32); ok == 0 {
			l.owner.Break()
			l.owner.DestroySelf()
		} else {
			l.owner.ctx.Core.LogWarn("save player failed")
		}
	case ETIMER:
		if time.Now().Sub(l.saveTimeout) > 0 {
			l.owner.SaveRole()
			l.saveTimeout = time.Now().Add(l.owner.ctx.saveTimeout)
			l.owner.ctx.Core.LogWarn("save role timeout")
		}
	default:
		l.owner.ctx.Core.LogWarnf("leaving state receive error event(%d)", event)
	}
	return ""
}
