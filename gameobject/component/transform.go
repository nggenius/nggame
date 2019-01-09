package component

import (
	"time"

	"github.com/nggenius/ngengine/utils"
	"github.com/nggenius/nggame/gameobject"
)

type Transform struct {
	gameobject.GameComponent
}

func NewTransform() *Transform {
	t := new(Transform)
	return t
}

func (t *Transform) Create() {
}

func (t *Transform) Destroy() {
}

func (t *Transform) Update(delta time.Duration) {

}

func (t *Transform) LookAtTarget(target Transform) {

}

func (t *Transform) LookAtPoint(x, y, z float32) {

}

func (t *Transform) RotateEulerAngles(x, y, z float32) {

}

func (t *Transform) RotateDirAngle(x, y, z float32, angle float32) {

}

func (t *Transform) RotateAngle(xAngle, yAngle, zAngle float32) {

}

func (t *Transform) Translate(x, y, z float32) {

}

// Serialize 序列化
func (t *Transform) Serialize(ar *utils.StoreArchive) error {
	return nil
}

// Deserialize 反序列化
func (t *Transform) Deserialize(ar *utils.LoadArchive) error {
	return nil
}
