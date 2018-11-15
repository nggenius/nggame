package login

import (
	"github.com/nggenius/ngengine/common/event"
	"github.com/nggenius/ngengine/core/service"
	"github.com/nggenius/ngengine/share"
	"github.com/nggenius/nggame/gameobject/entity"
	"github.com/nggenius/nggame/login/login"
	_ "github.com/nggenius/nggame/proto"
	"github.com/nggenius/ngmodule/store"
	"github.com/nggenius/ngmodule/timer"
)

// service
type Login struct {
	service.BaseService
	login *login.LoginModule
	store *store.StoreModule
	timer *timer.TimerModule
}

func (l *Login) Prepare(core service.CoreAPI) error {
	l.CoreAPI = core
	l.login = login.New()
	l.store = store.New()
	l.timer = timer.New()
	return nil
}

func (l *Login) Init(opt *service.CoreOption) error {
	l.CoreAPI.AddModule(l.store)
	l.CoreAPI.AddModule(l.login)
	l.CoreAPI.AddModule(l.timer)
	l.store.SetMode(store.STORE_CLIENT)
	entity.RegisterToDB(l.store)
	return nil
}

func (l *Login) Start() error {
	l.BaseService.Start()
	return nil
}

func (l *Login) OnEvent(e string, args event.EventArgs) {
	switch e {
	case share.EVENT_USER_CONNECT:
		l.CoreAPI.LogDebug("new user")
	case share.EVENT_USER_LOST:
		l.CoreAPI.LogDebug("lost user")
	}
}
