package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"os"
	"runtime/debug"

	"github.com/mysll/toolkit"
	"github.com/nggenius/ngengine/core"
	"github.com/nggenius/nggame/store"
)

var startPara = flag.String("p", "", "startPara")

func main() {
	defer func() {
		if x := recover(); x != nil {
			d := fmt.Sprintf("panic(%v)\n%s", x, debug.Stack())
			ioutil.WriteFile("dump.log", []byte(d), 0666)

			os.Exit(0)
		}
	}()

	flag.Parse()

	if *startPara == "" {
		flag.PrintDefaults()
		panic("store parameter is empty")
	}

	core.RegisterService("store", new(store.Store))

	_, err := core.CreateService("store", *startPara)
	if err != nil {
		panic(err)
	}
	core.RunAllService()

	toolkit.WaitForQuit()
	core.CloseAllService()
	core.Wait()
	return
}
