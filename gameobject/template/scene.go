package template

import (
	"time"

	"github.com/nggenius/nggame/gameobject"
)

type SceneObject struct {
	gameobject.BaseBehavior
}

func NewSceneObject() *SceneObject {
	so := new(SceneObject)
	return so
}

func (s *SceneObject) OnCreate() {
}

func (s *SceneObject) GameObjectType() int {
	return gameobject.OBJECT_SCENE
}

func (s *SceneObject) OnDestroy() {

}

func (s *SceneObject) OnUpdate(delta time.Duration) {
}
