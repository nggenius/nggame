package login

import (
	"github.com/nggenius/ngengine/common/fsm"
	"github.com/nggenius/nggame/proto/c2s"
)

type Idle struct {
	fsm.Default
	owner *Session
	idle  int32
}

func (s *Idle) Handle(event int, param interface{}) string {
	switch event {
	case LOGIN:
		s.owner.Login(param.(*c2s.Login))
		return "Logging"
	case TIMER:
		s.idle++
		if s.idle > 60 {
			s.owner.Break()
			return ""
		}
	case BREAK:
		s.owner.DestroySelf()
		return fsm.STOP
	}
	return ""
}
