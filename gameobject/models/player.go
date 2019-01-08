package models

import (
	"github.com/nggenius/nggame/gameobject"
	"github.com/nggenius/nggame/gameobject/entity"
	"github.com/nggenius/nggame/gameobject/template"
)

const (
	GAME_PLAYER = "GamePlayer"
)

type GamePlayer struct {
	gameobject.BaseObject
	role   *template.RoleObject
	spirit *entity.Entity
}

func (p *GamePlayer) Ctor() {
	p.role = template.NewRoleObject()
	p.spirit = entity.NewEntity(entity.PLAYER)
}

func (p *GamePlayer) Spirit() *entity.Entity {
	return p.spirit
}

func (p *GamePlayer) Behavior() gameobject.Behavior {
	return p.role
}

func (p *GamePlayer) ObjectType() string {
	return GAME_PLAYER
}

func (p *GamePlayer) EntityType() string {
	return entity.PLAYER
}
