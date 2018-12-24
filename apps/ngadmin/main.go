package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"os"
	"runtime/debug"

	"github.com/mysll/toolkit"
	"github.com/nggenius/ngengine/ngadmin"
)

var (
	configPath = flag.String("p", "../config/", "config path")
	appdef     = flag.String("a", "/app.cfg", "app config file")
	srvdef     = flag.String("s", "/servers.cfg", "services config file")
)

func main() {
	defer func() {
		if x := recover(); x != nil {
			d := fmt.Sprintf("panic(%v)\n%s", x, debug.Stack())
			ioutil.WriteFile("dump.log", []byte(d), 0666)

			os.Exit(0)
		}
	}()

	flag.Parse()

	var config ngadmin.Options

	err := config.Load(fmt.Sprintf("%s/%s", *configPath, *appdef), fmt.Sprintf("%s/%s", *configPath, *srvdef))
	if err != nil {
		panic(err)
	}

	ngadmin := ngadmin.New(&config)
	ngadmin.Main()
	toolkit.WaitForQuit()

	ngadmin.Exit()
}
