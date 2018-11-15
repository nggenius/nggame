package login

import (
	"github.com/nggenius/ngengine/common/fsm"
	"github.com/nggenius/ngengine/share"
	"github.com/nggenius/nggame/gameobject/entity/inner"
)

type Logging struct {
	fsm.Default
	owner *Session
	idle  int32
}

func (l *Logging) Handle(event int, param interface{}) string {
	switch event {
	case LOGIN_RESULT:
		args := param.([2]interface{})
		if l.owner.LoginResult(args[0].(int32), args[1].(*inner.Account)) {
			return "Logged"
		}
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
