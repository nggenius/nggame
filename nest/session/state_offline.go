package session

import (
	"time"

	"github.com/nggenius/ngengine/common/fsm"
	"github.com/nggenius/ngmodule/store"
)

type offline struct {
	fsm.Default
	owner      *Session
	Idle       int32
	remainTime time.Time
}

func (o *offline) Enter() {
	o.remainTime = time.Now().Add(o.owner.ctx.offlinetime)
}

func (o *offline) Handle(event int, param interface{}) string {
	switch event {
	case ETIMER:
		if time.Now().Sub(o.remainTime) > 0 {
			o.owner.SaveRole(store.STORE_SAVE_OFFLINE)
			return SLEAVING
		}
	case EREMAINTIME:
		rt := param.(int)
		o.remainTime = time.Now().Add(time.Duration(rt) * time.Second)
	default:
		o.owner.ctx.Core.LogWarnf("offline state receive error event(%d)", event)
	}
	return ""
}
