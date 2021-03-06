package scene

import (
	"github.com/nggenius/ngengine/core/rpc"
	"github.com/nggenius/ngengine/protocol"
	"github.com/nggenius/nggame/define"
)

type RegionCreate struct {
	ctx *SceneModule
}

func NewRegionCreate(ctx *SceneModule) *RegionCreate {
	s := new(RegionCreate)
	s.ctx = ctx
	return s
}

func (s *RegionCreate) RegisterCallback(srv rpc.Servicer) {
	srv.RegisterCallback("Query", s.Query)
	srv.RegisterCallback("Create", s.Create)
}

func (s *RegionCreate) Query(src rpc.Mailbox, dest rpc.Mailbox, msg *protocol.Message) (int32, *protocol.Message) {
	return protocol.Reply(protocol.TINY, s.ctx.Core.Mailbox().ServiceId())
}

func (s *RegionCreate) Create(src rpc.Mailbox, dest rpc.Mailbox, msg *protocol.Message) (int32, *protocol.Message) {
	var r define.Region
	if err := protocol.ParseArgs(msg, &r); err != nil {
		s.ctx.Core.LogErr("parse args error")
		return 0, nil
	}
	mb, err := s.ctx.scenes.CreateScene(r)
	if err != nil {
		return protocol.ReplyError(protocol.TINY, define.ERR_REGION_CREATE_FAILED, err.Error())
	}

	s.ctx.Core.LogInfo("create scene ", r)
	return protocol.Reply(protocol.TINY, mb)
}
