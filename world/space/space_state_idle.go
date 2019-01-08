package space

import "github.com/nggenius/ngengine/common/fsm"

type Idle struct {
	fsm.Default
	owner *SpaceManage
	Idle  int32
}

func newIdle(o *SpaceManage) *Idle {
	s := new(Idle)
	s.owner = o
	return s
}

func (s *Idle) Init(r fsm.StateRegister) {
	r.AddHandle(EREGION_RESP, s.OnCheck)
}

func (s *Idle) OnTimer() string {
	s.Idle++
	return ""
}

func (s *Idle) OnCheck(event int, param interface{}) string {
	if s.owner.hasAllReady() {
		return SCREATE
	}

	return ""
}

func (s *Idle) Handle(event int, param interface{}) string {
	s.owner.ctx.Core.LogWarnf("idle state receive error event(%d)", event)
	return ""
}
