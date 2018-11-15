package scene

import (
	"container/list"

	"github.com/nggenius/ngengine/common/fsm"
	"github.com/nggenius/ngengine/core/rpc"
	"github.com/nggenius/nggame/define"
	"github.com/nggenius/nggame/gameobject"
	"github.com/nggenius/nggame/gameobject/entity"
	"github.com/nggenius/ngmodule/object"
)

const GAME_SCENE = "GameScene"

type GameScene struct {
	*entity.Scene
	gameobject.SceneObject
	factory *object.Factory
	region  define.Region
	fsm     *fsm.FSM
	players *list.List
}

func (s *GameScene) Ctor() {
	s.Scene = entity.NewScene()
	s.fsm = initState(s)
	s.players = list.New()
}

func (s *GameScene) EntityType() string {
	return GAME_SCENE
}

func (s *GameScene) LoadRes(res string) bool {
	return true
}

func (s *GameScene) addPlayer(player gameobject.GameObject) {
	s.players.PushBack(player)
}

func (s *GameScene) removePlayer(id rpc.Mailbox) {
	for e := s.players.Front(); e != nil; e = e.Next() {
		if e.Value != nil && e.Value.(gameobject.GameObject).Spirit().ObjId() == id {
			s.players.Remove(e)
			break
		}
	}
}
