package inner

import (
	"time"
)

type Role struct {
	Id          int64
	Index       int8
	Account     string `xorm:"varchar(128) index"`
	RoleName    string `xorm:"varchar(128) unique"`
	CreateTime  time.Time
	LastLogTime time.Time
	LastAddress string `xorm:"varchar(32)"`
	Status      int8
	ServerId    int64
	Deleted     int8
	DeleteTime  time.Time
	SaveTime    time.Time
}

// set id
func (r *Role) SetId(val int64) {
	r.Id = val
}

// db id
func (r *Role) DBId() int64 {
	return r.Id
}

func (r *Role) SetStatus(s int8) {
	r.Status = s
}

func (r *Role) GetStatus() int8 {
	return r.Status
}

func (r *Role) UpdateLogTime() {
	r.LastLogTime = time.Now()
}

func (r *Role) GetDeleted() bool {
	return r.Deleted == 1
}

func (r *Role) Delete() {
	r.Deleted = 1
	r.DeleteTime = time.Now()
}

func (r *Role) Save() {
	r.SaveTime = time.Now()
}
