package component

import (
	"time"

	"github.com/nggenius/ngengine/utils"
	"github.com/nggenius/nggame/gameobject"
)

type Visible struct {
	gameobject.GameComponent
}

func NewVisible() *Visible {
	v := new(Visible)
	return v
}

func (v *Visible) Create() {
}

func (v *Visible) Update(delta time.Duration) {

}

// Serialize 序列化
func (v *Visible) Serialize(ar *utils.StoreArchive) error {
	return nil
}

// Deserialize 反序列化
func (v *Visible) Deserialize(ar *utils.LoadArchive) error {
	return nil
}
