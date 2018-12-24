package store

import (
	"github.com/nggenius/ngengine/core/service"
	"github.com/nggenius/nggame/gameobject/entity"
	"github.com/nggenius/nggame/store/extension"
	"github.com/nggenius/ngmodule/store"
)

type Store struct {
	service.BaseService
	store *store.StoreModule
	role  *extension.Role
}

func (d *Store) Prepare(core service.CoreAPI) error {
	d.CoreAPI = core
	d.store = store.New()

	return nil
}

func (d *Store) Init(opt *service.CoreOption) error {
	d.role = extension.NewRole(
		d.CoreAPI,
		d.store,
		opt.Args.MustString("Role", "inner.Role"),
		opt.Args.MustString("MainEntity", "entity.Player"),
		opt.Args.MustString("PlayerTable", "player"),
		opt.Args.MustString("PlayerBackup", "player_bak"),
	)
	d.CoreAPI.AddModule(d.store)
	d.store.SetMode(store.STORE_SERVER)
	d.store.Extend("role", d.role)
	entity.RegisterToDB(d.store)
	return nil
}

func (d *Store) Start() error {
	d.CoreAPI.Watch("all")
	return nil
}
