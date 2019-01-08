package scene

import (
	"github.com/nggenius/ngengine/core/rpc"
	"github.com/nggenius/ngengine/protocol"
	"github.com/nggenius/ngengine/share"
	"github.com/nggenius/nggame/define"
	"github.com/nggenius/nggame/gameobject"
)

func (s *GameScene) RegisterCallback(svr rpc.Servicer) {
	svr.RegisterCallback("EnterRegion", s.EnterRegion)
	svr.RegisterCallback("LeaveRegion", s.LeaveRegion)
	svr.RegisterCallback("RemovePlayer", s.RemovePlayer)
}

func (s *GameScene) EnterRegion(src rpc.Mailbox, dest rpc.Mailbox, msg *protocol.Message) (int32, *protocol.Message) {
	s.LogDebug("enter region")
	var data []byte
	err := protocol.ParseArgs(msg, &data)
	if err != nil {
		return protocol.ReplyError(protocol.TINY, share.ERR_ARGS_ERROR, err.Error())
	}

	obj, err := s.factory.Decode(data)
	if err != nil {
		return protocol.ReplyError(protocol.TINY, define.ERR_ENTER_REGION_FAILED, err.Error())
	}

	gameobject := obj.(gameobject.GameObject)
	gameobject.Spirit().Witness().SetOriginal(&src)
	s.addPlayer(gameobject)

	s.spirit.Core().LogDebug("add player succeed")
	return 0, nil
}

func (s *GameScene) LeaveRegion(src rpc.Mailbox, dest rpc.Mailbox, msg *protocol.Message) (int32, *protocol.Message) {
	s.LogDebug("leave region")
	pl := s.findPlayerByOrigin(src)
	if pl == nil {
		return protocol.ReplyError(protocol.TINY, define.ERR_REGION_OBJECT_NOT_FOUND, "player not found")
	}

	data, err := s.factory.Encode(pl)
	if err != nil {
		panic(err)
	}

	//s.removePlayerByOrigin(src)

	return protocol.Reply(protocol.DEF, s.ctx.keeptime, data)
}

func (s *GameScene) RemovePlayer(src rpc.Mailbox, dest rpc.Mailbox, msg *protocol.Message) (int32, *protocol.Message) {
	s.LogDebug("leave region")
	pl := s.findPlayerByOrigin(src)
	if pl == nil {
		return 0, nil
	}
	s.factory.Destroy(pl)
	s.removePlayerByOrigin(src)
	return 0, nil
}
