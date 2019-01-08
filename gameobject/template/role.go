package template

import (
	"time"

	"github.com/nggenius/nggame/gameobject"
	"github.com/nggenius/nggame/gameobject/component"
)

type RoleObject struct {
	gameobject.BaseBehavior
	transform *component.Transform
	visible   *component.Visible
}

func NewRoleObject() *RoleObject {
	ro := new(RoleObject)
	ro.transform = component.NewTransform()
	ro.visible = component.NewVisible()
	return ro
}

func (r *RoleObject) GameObjectType() int {
	return gameobject.OBJECT_PLAYER
}

func (r *RoleObject) OnCreate() {
	r.AddComponent("transform", r.transform, true)
	r.AddComponent("visible", r.visible, true)
}

func (r *RoleObject) OnDestroy() {

}

func (r *RoleObject) OnUpdate(delta time.Duration) {

}
