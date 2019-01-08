package session

import (
	"github.com/nggenius/ngengine/common/fsm"
	"github.com/nggenius/nggame/define"
	"github.com/nggenius/nggame/gameobject"
)

type ChooseRole struct {
	fsm.Default
	owner *Session
	Idle  int32
}

func newChooseRole(o *Session) *ChooseRole {
	s := new(ChooseRole)
	s.owner = o
	return s
}

func (s *ChooseRole) Init(r fsm.StateRegister) {
	r.AddHandle(ECHOOSED, s.OnChoose)
	r.AddHandle(EBREAK, s.OnBreak)
}

func (s *ChooseRole) OnTimer() string {
	s.Idle++
	if s.Idle > 60 {
		s.owner.Error(define.ERR_CHOOSE_TIMEOUT)
		return SLOGGED
	}
	return ""
}

func (s *ChooseRole) OnChoose(event int, param interface{}) string {
	args := param.([2]interface{})
	errcode := args[0].(int32)
	if errcode != 0 {
		s.owner.Error(errcode)
		return SLOGGED
	}
	player := args[1].(gameobject.GameObject)
	if player == nil {
		s.owner.Error(define.ERR_CHOOSE_ROLE)
		return SLOGGED
	}

	ls, ok := player.Spirit().Data().(LandInfo)
	if !ok {
		s.owner.DestroySelf()
		s.owner.ctx.Core.LogErr("entity not define landpos ", s.owner.ctx.mainEntity)
		return ""
	}
	s.owner.ctx.Core.LogDebug("enter game")
	s.owner.SetGameObject(player)
	x, y, z, o := ls.LandPosXYZOrient()
	s.owner.SetLandInfo(ls.LandScene(), x, y, z, o)
	return SONLINE
}

func (s *ChooseRole) OnBreak(event int, param interface{}) string {
	s.owner.DestroySelf()
	return fsm.STOP
}

func (s *ChooseRole) OnHandle(event int, param interface{}) string {
	s.owner.ctx.Core.LogWarnf("choose role state receive error event(%d)", event)
	return ""
}
