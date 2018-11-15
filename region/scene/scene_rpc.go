package scene

import (
	"github.com/nggenius/ngengine/core/rpc"
	"github.com/nggenius/ngengine/protocol"
	"github.com/nggenius/ngengine/share"
	"github.com/nggenius/nggame/define"
	"github.com/nggenius/nggame/gameobject"
)

func (s *GameScene) RegisterCallback(svr rpc.Servicer) {
	svr.RegisterCallback("AddPlayer", s.AddPlayer)
}

func (s *GameScene) AddPlayer(src rpc.Mailbox, dest rpc.Mailbox, msg *protocol.Message) (int32, *protocol.Message) {
	s.Core().LogDebug("add player")
	var data []byte
	err := protocol.ParseArgs(msg, &data)
	if err != nil {
		return protocol.ReplyError(protocol.TINY, share.ERR_ARGS_ERROR, err.Error())
	}

	obj, err := s.factory.Decode(data)
	if err != nil {
		return protocol.ReplyError(protocol.TINY, define.ERR_ENTER_REGION_FAILED, err.Error())
	}

	s.addPlayer(obj.(gameobject.GameObject))

	s.Core().LogDebug("add player succeed")
	return 0, nil
}
