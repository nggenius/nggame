package template

import "github.com/nggenius/nggame/gameobject"

type AssistObject struct {
	*gameobject.BaseObject
}

func NewAssistObject() *AssistObject {
	ro := new(AssistObject)
	ro.BaseObject = new(gameobject.BaseObject)
	return ro
}

func (r *AssistObject) ObjectType() int {
	return gameobject.OBJECT_ASSIST
}

func (r *AssistObject) OnCreate() {
}

func (r *AssistObject) OnDestroy() {

}
