package models

import (
	"github.com/nggenius/nggame/gameobject"
	"github.com/nggenius/nggame/gameobject/entity"
	"github.com/nggenius/nggame/gameobject/template"
	"github.com/nggenius/ngmodule/object"
)

const (
	GAME_PLAYER = "GamePlayer"
)

type GamePlayer struct {
	*template.RoleObject
	*entity.Player
}

func (p *GamePlayer) Ctor() {
	p.Player = entity.NewPlayer()
	p.RoleObject = template.NewRoleObject()
}

func (p *GamePlayer) Object() object.Object {
	return p.Player
}

func (p *GamePlayer) GameObject() gameobject.GameObject {
	return p.RoleObject
}

func (p *GamePlayer) EntityType() string {
	return GAME_PLAYER
}
