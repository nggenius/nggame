package template

import (
	"github.com/nggenius/nggame/gameobject"
	"github.com/nggenius/nggame/gameobject/component"
)

type NpcObject struct {
	gameobject.BaseBehavior
	transform *component.Transform
	visible   *component.Visible
}

func NewNpcObject() *NpcObject {
	ro := new(NpcObject)
	ro.transform = component.NewTransform()
	ro.visible = component.NewVisible()
	return ro
}

func (r *NpcObject) GameObjectType() int {
	return gameobject.OBJECT_NPC
}

func (r *NpcObject) OnCreate() {
	r.AddComponent("transform", r.transform, true)
	r.AddComponent("visible", r.visible, true)
}

func (r *NpcObject) OnDestroy() {

}
