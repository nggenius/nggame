package models

import (
	"github.com/nggenius/nggame/gameobject/entity"
	"github.com/nggenius/nggame/gameobject/template"
)

const (
	GAME_PLAYER = "GamePlayer"
)

type GamePlayer struct {
	template.RoleObject
	*entity.Player
}

func (p *GamePlayer) Ctor() {
	p.Player = entity.NewPlayer()
	p.SetSpirit(p.Player)
}

func (p *GamePlayer) EntityType() string {
	return GAME_PLAYER
}
