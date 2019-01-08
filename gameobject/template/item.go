package template

import (
	"github.com/nggenius/nggame/gameobject"
	"github.com/nggenius/nggame/gameobject/component"
)

type ItemObject struct {
	gameobject.BaseBehavior
	visible *component.Visible
}

func NewItemObject() *ItemObject {
	ro := new(ItemObject)
	ro.visible = component.NewVisible()
	return ro
}

func (r *ItemObject) GameObjectType() int {
	return gameobject.OBJECT_ITEM
}

func (r *ItemObject) OnCreate() {
	r.AddComponent("visible", r.visible, true)
}

func (r *ItemObject) OnDestroy() {

}
