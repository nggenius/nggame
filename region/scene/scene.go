package scene

import (
	"container/list"

	"github.com/nggenius/ngengine/common/fsm"
	"github.com/nggenius/ngengine/core/rpc"
	"github.com/nggenius/nggame/define"
	"github.com/nggenius/nggame/gameobject"
	"github.com/nggenius/nggame/gameobject/entity"
	"github.com/nggenius/nggame/gameobject/template"
	"github.com/nggenius/ngmodule/object"
)

const GAME_SCENE = "GameScene"

type GameScene struct {
	gameobject.BaseObject
	ctx     *SceneModule
	spirit  *entity.Entity
	scene   *template.SceneObject
	factory *object.Factory
	region  define.Region
	fsm     *fsm.FSM
	players *list.List
}

func (s *GameScene) Ctor() {
	s.spirit = entity.NewEntity(entity.SCENE)
	s.scene = template.NewSceneObject()
	s.fsm = initState(s)
	s.players = list.New()
}

func (s *GameScene) Spirit() *entity.Entity {
	return s.spirit
}

func (s *GameScene) Behavior() gameobject.Behavior {
	return s.scene
}

func (s *GameScene) ObjectType() string {
	return GAME_SCENE
}

func (s *GameScene) EntityType() string {
	return entity.SCENE
}

func (s *GameScene) LoadRes(res string) bool {
	return true
}

func (s *GameScene) addPlayer(player gameobject.GameObject) {
	s.players.PushBack(player)
}

func (s *GameScene) findPlayerByOrigin(src rpc.Mailbox) gameobject.GameObject {
	for e := s.players.Front(); e != nil; e = e.Next() {
		gameobj := e.Value.(gameobject.GameObject)
		origin := gameobj.Spirit().Witness().Original()
		if e.Value != nil && *origin == src {
			return gameobj
		}
	}
	return nil
}

func (s *GameScene) removePlayer(id rpc.Mailbox) {
	for e := s.players.Front(); e != nil; e = e.Next() {
		if e.Value != nil && e.Value.(gameobject.GameObject).Spirit().ObjId() == id {
			s.players.Remove(e)
			break
		}
	}
}

func (s *GameScene) removePlayerByOrigin(src rpc.Mailbox) {
	for e := s.players.Front(); e != nil; e = e.Next() {
		origin := e.Value.(gameobject.GameObject).Spirit().Witness().Original()
		if e.Value != nil && *origin == src {
			s.players.Remove(e)
			break
		}
	}
}
