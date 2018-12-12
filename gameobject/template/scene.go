package template

import (
	"time"

	"github.com/nggenius/nggame/gameobject"
)

type SceneObject struct {
	gameobject.BaseObject
}

func (s *SceneObject) OnCreate() {
	s.BaseObject.OnCreate()
}

func (s *SceneObject) ObjectType() int {
	return gameobject.OBJECT_SCENE
}

func (s *SceneObject) OnDestroy() {
	s.BaseObject.OnDestroy()
}

func (s *SceneObject) Update(delta time.Duration) {
	s.BaseObject.Update(delta)
}
