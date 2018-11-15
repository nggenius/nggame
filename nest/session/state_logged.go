package session

import (
	"github.com/nggenius/ngengine/common/fsm"
	"github.com/nggenius/ngengine/share"
	"github.com/nggenius/nggame/gameobject/entity/inner"
	"github.com/nggenius/nggame/proto/c2s"
)

type logged struct {
	fsm.Default
	owner *Session
	Idle  int32
}

func (s *logged) Handle(event int, param interface{}) string {
	switch event {
	case ETIMER:
		s.Idle++
		if s.Idle > 60 {
			s.owner.Break()
			return ""
		}
	case EBREAK:
		s.owner.DestroySelf()
		return fsm.STOP
	case EROLEINFO:
		args := param.([2]interface{})
		errcode := args[0].(int32)
		roles := args[1].([]*inner.Role)
		if errcode != 0 {
			s.owner.Error(errcode)
			return ""
		}

		s.owner.SendRoleInfo(roles)
		s.Idle = 0
	case EDELETE:
		args := param.(c2s.DeleteRole)
		if err := s.owner.DeleteRole(args); err != nil {
			s.owner.Error(share.ERR_SYSTEM_ERROR)
			return ""
		}
		return SDELETE
	case ECREATE:
		args := param.(c2s.CreateRole)
		if err := s.owner.CreateRole(args); err != nil {
			s.owner.Error(share.ERR_SYSTEM_ERROR)
			return ""
		}
		return SCREATE
	case ECHOOSE:
		args := param.(c2s.ChooseRole)
		if err := s.owner.ChooseRole(args); err != nil {
			s.owner.Error(share.ERR_SYSTEM_ERROR)
			return ""
		}
		return SCHOOSE
	default:
		s.owner.ctx.Core.LogWarnf("logged state receive error event(%d)", event)
	}
	return ""
}
