package scene

import "github.com/nggenius/ngengine/common/fsm"

type Idle struct {
	fsm.Default
	owner *GameScene
	Idle  int32
}

func newIdle(o *GameScene) *Idle {
	s := new(Idle)
	s.owner = o
	return s
}

func (s *Idle) Init(r fsm.StateRegister) {
	//r.AddHandle(LOGIN, s.OnLogin)
}

func (s *Idle) OnTimer() string {
	s.Idle++
	return ""
}

func (s *Idle) OnHandle(event int, param interface{}) string {
	return ""
}
