package login

import "github.com/nggenius/ngengine/common/fsm"

const (
	NONE         = iota
	TIMER        // 1秒钟的定时器
	BREAK        // 客户端断开连接
	LOGIN        // 客户端登录
	LOGIN_RESULT // 登录结果
	NEST_RESULT  // nest 登录结果
)

const (
	SIDLE    = "idle"
	SLOGGING = "logging"
	SLOGGED  = "logged"
)

func initState(s *Session) *fsm.FSM {
	fsm := fsm.NewFSM()
	fsm.Register(SIDLE, NewIdle(s))
	fsm.Register(SLOGGING, NewLogging(s))
	fsm.Register(SLOGGED, NewLogged(s))
	fsm.Start(SIDLE)
	return fsm
}
