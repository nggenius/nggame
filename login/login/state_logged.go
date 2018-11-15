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

func (l *Logged) Handle(event int, param interface{}) string {
	switch event {
	case NEST_RESULT:
		args := param.([2]interface{})
		if l.owner.NestResult(args[0].(int32), args[1].(string)) {
			l.idle = 0 // 1 分钟后退出
			return ""
		}
		return SIDLE //重新登录
	case TIMER:
		l.idle++
		if l.idle > 60 {
			l.owner.Error(share.S2C_ERR_SERVICE_INVALID)
			l.owner.Break()
			return ""
		}
	case BREAK:
		l.owner.DestroySelf()
		return fsm.STOP
	}
	return ""
}
