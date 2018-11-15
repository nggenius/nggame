package login

import (
	"github.com/nggenius/ngengine/common/fsm"
	"github.com/nggenius/ngengine/core/rpc"
	"github.com/nggenius/ngengine/share"
	"github.com/nggenius/nggame/gameobject/entity/inner"
	"github.com/nggenius/nggame/proto/c2s"
	"github.com/nggenius/nggame/proto/s2c"
)

type Session struct {
	*fsm.FSM
	ctx     *LoginModule
	id      uint64
	Account string
	Mailbox *rpc.Mailbox
	nest    share.ServiceId
	delete  bool
}

func NewSession(id uint64, ctx *LoginModule) *Session {
	c := &Session{}
	c.ctx = ctx
	c.id = id
	c.FSM = initState(c)
	return c
}

func (c *Session) DestroySelf() {
	c.delete = true
	c.ctx.deleted.PushBack(c.id)
}

func (c *Session) SetAccount(acc string) {
	c.Account = acc
}

func (c *Session) SetMailbox(mb *rpc.Mailbox) {
	c.Mailbox = mb
}

func (c *Session) Login(login *c2s.Login) {
	c.ctx.account.sendLogin(c, login)
}

func (c *Session) LoginResult(errcode int32, accinfo *inner.Account) bool {
	if errcode != 0 {
		c.Error(errcode)
		return false
	}

	if accinfo.Id != 0 {
		srv := c.ctx.account.findNest(c)
		if srv != nil {
			c.nest = srv.Id
			return true
		}
		return false
	}
	c.Error(share.S2C_ERR_NAME_PASS)
	return false
}

func (c *Session) NestResult(errcode int32, token string) bool {
	if errcode != 0 {
		c.Error(errcode)
		return false
	}

	srv := c.ctx.Core.LookupService(c.nest)
	if srv == nil {
		c.Error(share.S2C_ERR_SERVICE_INVALID)
		return false
	}

	nest := &s2c.NestInfo{}
	nest.Addr = srv.OuterAddr
	nest.Port = int32(srv.OuterPort)
	nest.Token = token

	if err := c.ctx.Core.Mailto(nil, c.Mailbox, "Login.Nest", nest); err != nil {
		return false
	}

	return true
}

func (c *Session) Break() {
	c.ctx.Core.Break(c.id)
}

func (c *Session) Error(err int32) {
	result := s2c.Error{}
	result.ErrCode = err
	c.ctx.Core.Mailto(nil, c.Mailbox, "system.Error", result)
}
