package space

import "github.com/nggenius/ngengine/common/fsm"

type CreateRegion struct {
	fsm.Default
	owner *SpaceManage
	Idle  int32
}

func newCreateRegion(o *SpaceManage) *CreateRegion {
	s := new(CreateRegion)
	s.owner = o
	return s
}

func (s *CreateRegion) Init(r fsm.StateRegister) {
}

func (s *CreateRegion) Enter() {
	if err := s.owner.createAllRegion(); err != nil {
		s.owner.ctx.Core.LogErr(err)
	}
}

func (s *CreateRegion) OnTimer() string {
	s.Idle++
	return ""
}

func (s *CreateRegion) Handle(event int, param interface{}) string {
	s.owner.ctx.Core.LogWarnf("create state receive error event(%d)", event)
	return ""
}
