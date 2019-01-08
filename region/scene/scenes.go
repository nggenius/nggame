package scene

import (
	"time"

	"github.com/nggenius/ngengine/core/rpc"
	"github.com/nggenius/nggame/define"
	"github.com/nggenius/ngmodule/object"
)

type Scenes struct {
	ctx    *SceneModule
	scenes map[int]*GameScene
}

func NewScenes(ctx *SceneModule) *Scenes {
	s := new(Scenes)
	s.ctx = ctx
	s.scenes = make(map[int]*GameScene)
	return s
}

func (s *Scenes) CreateScene(r define.Region) (rpc.Mailbox, error) {
	fid := object.OBJECT_TYPE_SCENE_OFFSET + r.Id
	err := s.ctx.object.AddFactory(fid)
	if err != nil {
		return rpc.NullMailbox, err
	}

	f := s.ctx.object.Factory(fid)
	scene, err := f.Create(GAME_SCENE)
	if err != nil {
		return rpc.NullMailbox, err
	}

	gamescene := scene.(*GameScene)
	gamescene.factory = f
	gamescene.region = r
	gamescene.ctx = s.ctx
	s.scenes[r.Id] = gamescene

	return gamescene.spirit.ObjId(), nil
}

func (s *Scenes) UpdateAllScene(t time.Duration) {
	for k := range s.scenes {
		s.scenes[k].scene.Update(t)
	}
}
