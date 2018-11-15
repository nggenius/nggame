package models

import "github.com/nggenius/ngmodule/object"

type reg interface {
	Register(oc object.ObjectCreate)
}

func Register(r reg) {
	r.Register(new(GamePlayer))
}
