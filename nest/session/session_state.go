package session

import (
	"github.com/nggenius/ngengine/common/fsm"
)

const (
	NONE          = iota
	EBREAK        // 客户端断开连接
	ELOGIN        // 客户端登录
	EROLEINFO     // 角色列表
	ECREATE       // 创建角色
	ECREATED      // 创建完成
	ECHOOSE       // 选择角色
	ECHOOSED      // 选择角色成功
	EDELETE       // 删除角色
	EDELETED      // 删除成功
	ESTORED       // 存档完成
	EONLINE       // 进入场景
	EFREGION      // 查找场景
	ESWREGION     // 切换场景
	EREMAINTIME   // 更新离线存活时间
	EREGIONREMOVE // 场景已经删除玩家
)

const (
	SIDLE    = "idle"
	SLOGGED  = "logged"
	SCREATE  = "create"
	SCHOOSE  = "choose"
	SDELETE  = "delete"
	SONLINE  = "online"
	SOFFLINE = "offline"
	SLEAVING = "leaving"
)

func initState(s *Session) *fsm.FSM {
	fsm := fsm.NewFSM()
	fsm.Register(SIDLE, newIdle(s))
	fsm.Register(SLOGGED, newLogged(s))
	fsm.Register(SCREATE, newCreateRole(s))
	fsm.Register(SCHOOSE, newChooseRole(s))
	fsm.Register(SDELETE, newDeleting(s))
	fsm.Register(SONLINE, newOnline(s))
	fsm.Register(SOFFLINE, newOffline(s))
	fsm.Register(SLEAVING, newLeaving(s))
	fsm.Start(SIDLE)
	return fsm
}
