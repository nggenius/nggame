package template

import (
	"time"

	"github.com/nggenius/nggame/gameobject"
)

type SceneObject struct {
	*gameobject.BaseObject
}

func NewSceneObject() *SceneObject {
	so := new(SceneObject)
	so.BaseObject = new(gameobject.BaseObject)
	return so
}

func (s *SceneObject) OnCreate() {
}

func (s *SceneObject) ObjectType() int {
	return gameobject.OBJECT_SCENE
}

func (s *SceneObject) OnDestroy() {

}

func (s *SceneObject) OnUpdate(delta time.Duration) {
}
