package template

import "github.com/nggenius/nggame/gameobject"

type AssistObject struct {
	gameobject.BaseBehavior
}

func NewAssistObject() *AssistObject {
	ro := new(AssistObject)
	return ro
}

func (r *AssistObject) GameObjectType() int {
	return gameobject.OBJECT_ASSIST
}

func (r *AssistObject) OnCreate() {
}

func (r *AssistObject) OnDestroy() {

}
