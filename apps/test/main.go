package main

import (
	"net/http"

	_ "net/http/pprof"

	"github.com/nggenius/ngengine/core"
	"github.com/nggenius/nggame/login"
	"github.com/nggenius/nggame/nest"
	_ "github.com/nggenius/nggame/proto"
	"github.com/nggenius/nggame/region"
	"github.com/nggenius/nggame/store"
	"github.com/nggenius/nggame/world"
)

var startnest = `{
	"ServId":5,
	"ServType": "nest",
	"AdminAddr":"127.0.0.1",
	"AdminPort":12500,
	"ServName": "nest_1",
	"ServAddr": "127.0.0.1",
	"ServPort": 0,
	"Expose": true,
	"OuterAddr":"192.168.21.76",
	"HostAddr": "0.0.0.0",
	"HostPort": 0,
	"LogFile":"log/nest.log",
	"Args": {
		"MainEntity":"entity.Player",
		"Role":"GamePlayer"
	}
}`

var startlogin = `{
	"ServId":4,
	"ServType": "login",
	"AdminAddr":"127.0.0.1",
	"AdminPort":12500,
	"ServName": "login_1",
	"ServAddr": "127.0.0.1",
	"ServPort": 0,
	"Expose": true,
	"OuterAddr":"192.168.21.76",
	"HostAddr": "0.0.0.0",
	"HostPort": 4000,
	"LogFile":"log/login.log",
	"Args": {}
}`

var startworld = `{
	"ServId":3,
	"ServType": "world",
	"AdminAddr":"127.0.0.1",
	"AdminPort":12500,
	"ServName": "world_1",
	"ServAddr": "127.0.0.1",
	"ServPort": 0,
	"Expose": false,
	"LogFile":"log/world.log",
	"ResRoot":"../../res/",
	"Args": {
		"Region":"region.json",
		"MinRegions":1
	}
}`

var startregion = `{
	"ServId":2,
	"ServType": "region",
	"AdminAddr":"127.0.0.1",
	"AdminPort":12500,
	"ServName": "region_1",
	"ServAddr": "127.0.0.1",
	"ServPort": 0,
	"Expose": false,
	"LogFile":"log/region.log",
	"ResRoot":"D:/home/work/github/ngengine/res/",
	"Args": {}
}`

var dbargs = `{
	"ServId":1,
	"ServType": "store",
	"AdminAddr":"127.0.0.1",
	"AdminPort":12500,
	"ServName": "db_1",
	"ServAddr": "127.0.0.1",
	"ServPort": 0,
	"Expose": false,
	"HostAddr": "",
	"HostPort": 0,
	"LogFile":"log/db.log",
	"Args": {
		"db":"mysql",
		"datasource":"root:123456@tcp(192.168.21.76:3306)/ngengine?charset=utf8",
		"showsql":false
	}
}`

func main() {
	// 捕获异常
	core.RegisterService("store", new(store.Store))
	core.RegisterService("login", new(login.Login))
	core.RegisterService("nest", new(nest.Nest))
	core.RegisterService("world", new(world.World))
	core.RegisterService("region", new(region.Region))
	_, err := core.CreateService("login", startlogin)
	if err != nil {
		panic(err)
	}

	_, err = core.CreateService("nest", startnest)
	if err != nil {
		panic(err)
	}

	_, err = core.CreateService("store", dbargs)
	if err != nil {
		panic(err)
	}

	_, err = core.CreateService("world", startworld)
	if err != nil {
		panic(err)
	}

	_, err = core.CreateService("region", startregion)
	if err != nil {
		panic(err)
	}

	core.RunAllService()

	go http.ListenAndServe(":9600", nil)
	core.Wait()
}
