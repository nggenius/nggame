package world

import (
	"github.com/nggenius/ngengine/core/service"
	"github.com/nggenius/nggame/world/space"
)

type World struct {
	service.BaseService
	worldspace *space.WorldSpaceModule
}

func (s *World) Prepare(core service.CoreAPI) error {
	s.CoreAPI = core
	s.worldspace = space.New()
	return nil
}

func (s *World) Init(opt *service.CoreOption) error {
	s.AddModule(s.worldspace)
	return nil
}

func (s *World) Start() error {
	s.BaseService.Start()
	return nil
}
