package template

import (
	"time"

	"github.com/nggenius/nggame/gameobject"
	"github.com/nggenius/nggame/gameobject/component"
)

type RoleObject struct {
	gameobject.BaseObject
	transform *component.Transform
	visible   *component.Visible
}

func (r *RoleObject) ObjectType() int {
	return gameobject.OBJECT_PLAYER
}

func (r *RoleObject) OnCreate() {
	r.transform = component.NewTransform()
	r.visible = component.NewVisible()
	r.AddComponent("transform", r.transform, true)
	r.AddComponent("visible", r.visible, true)

	r.BaseObject.OnCreate()
}

func (r *RoleObject) OnDestroy() {
	r.BaseObject.OnDestroy()
}

func (r *RoleObject) Update(delta time.Duration) {
	r.BaseObject.Update(delta)
}
