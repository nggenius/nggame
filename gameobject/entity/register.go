package entity

import (
	"github.com/nggenius/nggame/gameobject/entity/inner"
)

const (
	ACCOUNT = "inner.Account"
	ROLE    = "inner.Role"
)

type createhelper interface {
	new() DataObject
	makeArchive() interface{}
	makeArchiveSlice() interface{}
	TableName() string
}

var objreg = make(map[string]createhelper)

type Register interface {
	Register(name string, obj interface{}, objslice interface{}) error
}

func RegisterToDB(r Register) {
	r.Register(ACCOUNT, &inner.Account{}, []*inner.Account{})
	r.Register(ROLE, &inner.Role{}, []*inner.Role{})
	for k, v := range objreg {
		if v.TableName() != "" {
			r.Register(k, v.makeArchive(), v.makeArchiveSlice())
		}
	}
}

func registObject(typ string, c createhelper) {
	if _, has := objreg[typ]; has {
		panic("register object twice")
	}

	objreg[typ] = c
}

func Create(typ string) DataObject {
	if c, ok := objreg[typ]; ok {
		return c.new()
	}
	return nil
}
