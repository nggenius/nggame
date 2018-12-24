package template

import (
	"github.com/nggenius/nggame/gameobject"
	"github.com/nggenius/nggame/gameobject/component"
)

type ItemObject struct {
	*gameobject.BaseObject
	visible *component.Visible
}

func NewItemObject() *ItemObject {
	ro := new(ItemObject)
	ro.BaseObject = new(gameobject.BaseObject)
	ro.visible = component.NewVisible()
	return ro
}

func (r *ItemObject) ObjectType() int {
	return gameobject.OBJECT_ITEM
}

func (r *ItemObject) OnCreate() {
	r.AddComponent("visible", r.visible, true)
}

func (r *ItemObject) OnDestroy() {

}
