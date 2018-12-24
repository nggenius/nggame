package main

import (
	"flag"
	"net/http"

	"github.com/bitly/go-simplejson"

	"github.com/mysll/toolkit"

	_ "net/http/pprof"

	"github.com/nggenius/ngengine/core"
	"github.com/nggenius/nggame/login"
	"github.com/nggenius/nggame/nest"
	_ "github.com/nggenius/nggame/proto"
	"github.com/nggenius/nggame/region"
	"github.com/nggenius/nggame/store"
	"github.com/nggenius/nggame/world"
)

var (
	configPath = flag.String("s", "../../config/servers_test.cfg", "config path")
)

func getConfig(json *simplejson.Json, key string) string {
	j := json.Get(key)
	if j == nil {
		panic("key not found")
	}

	j = j.GetIndex(0)
	if j == nil {
		panic("key not found")
	}

	b, err := j.MarshalJSON()
	if err != nil {
		panic("key not found")
	}
	return string(b)
}

func main() {
	flag.Parse()
	f, err := toolkit.ReadFile(*configPath)
	if err != nil {
		panic(err)
	}
	cfg, err := simplejson.NewJson(f)
	if err != nil {
		panic(err)
	}
	// 捕获异常
	core.RegisterService("store", new(store.Store))
	core.RegisterService("login", new(login.Login))
	core.RegisterService("nest", new(nest.Nest))
	core.RegisterService("world", new(world.World))
	core.RegisterService("region", new(region.Region))
	_, err = core.CreateService("login", getConfig(cfg, "login"))
	if err != nil {
		panic(err)
	}

	_, err = core.CreateService("nest", getConfig(cfg, "nest"))
	if err != nil {
		panic(err)
	}

	_, err = core.CreateService("store", getConfig(cfg, "store"))
	if err != nil {
		panic(err)
	}

	_, err = core.CreateService("world", getConfig(cfg, "world"))
	if err != nil {
		panic(err)
	}

	_, err = core.CreateService("region", getConfig(cfg, "region"))
	if err != nil {
		panic(err)
	}

	core.RunAllService()

	go http.ListenAndServe(":9600", nil)
	core.Wait()
}
