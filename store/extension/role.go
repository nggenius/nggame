package extension

import (
	"fmt"

	"github.com/nggenius/nggame/gameobject/entity/inner"

	"github.com/nggenius/ngengine/core/rpc"
	"github.com/nggenius/ngengine/core/service"
	"github.com/nggenius/ngengine/protocol"
	"github.com/nggenius/ngengine/share"
	"github.com/nggenius/ngmodule/store"
)

type Role struct {
	store     *store.StoreModule
	ctx       service.CoreAPI
	role      string
	entity    string
	player    string
	playerbak string
	backupsql string
}

func NewRole(core service.CoreAPI, s *store.StoreModule, role, entity, player, playerbak string) *Role {
	r := &Role{}
	r.ctx = core
	r.store = s
	r.role = role
	r.entity = entity
	r.player = player
	r.playerbak = playerbak
	r.backupsql = fmt.Sprintf("insert into %s select * from %s where id=?", playerbak, player)
	return r
}

func (r *Role) RegisterCallback(svr rpc.Servicer) {
	svr.RegisterCallback("CreateRole", r.CreateRole)
	svr.RegisterCallback("DeleteRole", r.DeleteRole)
	svr.RegisterCallback("ChooseRole", r.ChooseRole)
	svr.RegisterCallback("SaveRole", r.SaveRole)
}

func (r *Role) CreateRole(sender, _ rpc.Mailbox, msg *protocol.Message) (errcode int32, reply *protocol.Message) {
	m := protocol.NewMessageReader(msg)
	role := new(inner.Role)
	player := r.store.CreateDBObj(r.entity)
	err := m.Get(role)
	if err != nil {
		return protocol.ReplyError(protocol.TINY, share.ERR_ARGS_ERROR, err.Error())
	}
	err = m.Get(player)
	if err != nil {
		return protocol.ReplyError(protocol.TINY, share.ERR_ARGS_ERROR, err.Error())
	}

	session := r.store.Sql().Session()
	defer session.Close()

	tmp := new(inner.Role)

	count, err := session.Where("`role_name`=? ", role.RoleName).Count(tmp)
	if err != nil {
		return protocol.ReplyError(protocol.TINY, store.ERR_STORE_SQL, err.Error())
	}
	if count != 0 {
		return protocol.ReplyError(protocol.TINY, store.ERR_STORE_ROLE_NAME, "")
	}

	count, err = session.Where("`index`=? and `account`=? and `deleted`=0", role.Index, role.Account).Count(tmp)
	if err != nil {
		return protocol.ReplyError(protocol.TINY, store.ERR_STORE_SQL, err.Error())
	}

	if count != 0 {
		return protocol.ReplyError(protocol.TINY, store.ERR_STORE_ROLE_INDEX, "index error")
	}

	session.Begin()
	_, err = session.Insert(role)
	if err != nil {
		return protocol.ReplyError(protocol.TINY, store.ERR_STORE_ERROR, err.Error())
	}
	_, err = session.Insert(player)
	if err != nil {
		return protocol.ReplyError(protocol.TINY, store.ERR_STORE_ERROR, err.Error())
	}

	session.Commit()
	return protocol.Reply(protocol.TINY)
}

func (r *Role) DeleteRole(sender, _ rpc.Mailbox, msg *protocol.Message) (errcode int32, reply *protocol.Message) {
	m := protocol.NewMessageReader(msg)
	roleid, err := m.ReadInt64()
	if err != nil {
		r.ctx.LogFatal("read roleid failed, ", err)
		return 0, nil
	}

	session := r.store.Sql().Session()
	defer session.Close()

	role := new(inner.Role)

	session.Begin()

	b, _ := session.Id(roleid).Get(role)
	if !b {
		return protocol.ReplyError(protocol.TINY, store.ERR_STORE_ROLE_NOT_FOUND, "player not found")
	}

	role.Delete()
	if _, err := session.Id(roleid).Cols("deleted", "delete_time").Update(role); err != nil {
		return protocol.ReplyError(protocol.TINY, store.ERR_STORE_ERROR, err.Error())
	}

	session.Commit()
	return protocol.Reply(protocol.TINY)
}

func (r *Role) ChooseRole(sender, _ rpc.Mailbox, msg *protocol.Message) (errcode int32, reply *protocol.Message) {
	var nest rpc.Mailbox
	var roleid int64
	if err := protocol.ParseArgs(msg, &nest, &roleid); err != nil {
		r.ctx.LogFatal("read args failed, ", err)
		return 0, nil
	}

	session := r.store.Sql().Session()
	defer session.Close()

	role := new(inner.Role)
	player := r.store.CreateDBObj(r.entity)

	session.Begin()

	b, _ := session.Id(roleid).Get(role)
	if !b {
		return protocol.ReplyError(protocol.TINY, store.ERR_STORE_ROLE_NOT_FOUND, "player not found")
	}

	if role.GetDeleted() {
		return protocol.ReplyError(protocol.TINY, store.ERR_STORE_ROLE_DELETED, "player status error")
	}

	if role.GetStatus() != 0 {
		return protocol.ReplyError(protocol.TINY, store.ERR_STORE_ROLE_STATUS_ERROR, "player status error")
	}

	role.Login(nest)
	if _, err := session.Id(roleid).Cols("last_log_time", "nest").Update(role); err != nil {
		return protocol.ReplyError(protocol.TINY, store.ERR_STORE_ERROR, err.Error())
	}

	b, _ = session.Id(roleid).Get(player)
	if !b {
		return protocol.ReplyError(protocol.TINY, store.ERR_STORE_ROLE_NOT_FOUND, "player not found")
	}
	session.Commit()
	return protocol.Reply(protocol.DEF, player)
}

func (r *Role) SaveRole(src rpc.Mailbox, _ rpc.Mailbox, msg *protocol.Message) (int32, *protocol.Message) {
	role := new(inner.Role)
	player := r.store.CreateDBObj(r.entity)

	var typ int8
	var roleid int64

	if err := protocol.ParseArgs(msg, &typ, &roleid, player); err != nil {
		r.ctx.LogFatal("read roleid failed, ", err)
		return 0, nil
	}

	session := r.store.Sql().Session()
	defer session.Close()

	session.Begin()

	role.Id = roleid
	role.Save(typ == store.STORE_SAVE_OFFLINE)
	var cols []string
	if typ == store.STORE_SAVE_OFFLINE {
		cols = []string{"save_time", "nest"}
	} else {
		cols = []string{"save_time"}
	}

	if _, err := session.Id(roleid).Cols(cols...).Update(role); err != nil {
		r.ctx.LogErr(err.Error())
		return protocol.ReplyError(protocol.TINY, store.ERR_STORE_SAVE_FAILED, err.Error())
	}

	if _, err := session.Id(roleid).Update(player); err != nil {
		r.ctx.LogErr(err.Error())
		return protocol.ReplyError(protocol.TINY, store.ERR_STORE_SAVE_FAILED, err.Error())
	}

	session.Commit()
	return protocol.Reply(protocol.TINY)
}
