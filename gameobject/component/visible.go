package component

import "github.com/nggenius/nggame/gameobject"

type Visible struct {
	gameobject.GameComponent
}

func NewVisible() *Visible {
	v := new(Visible)
	return v
}

func (v *Visible) Create() {

}
