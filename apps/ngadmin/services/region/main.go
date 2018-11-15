package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"os"
	"runtime/debug"

	"github.com/mysll/toolkit"
	"github.com/nggenius/ngengine/core"
	"github.com/nggenius/nggame/region"
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
		panic("region parameter is empty")
	}

	core.RegisterService("region", new(region.Region))

	_, err := core.CreateService("region", *startPara)
	if err != nil {
		panic(err)
	}
	core.RunAllService()

	toolkit.WaitForQuit()
	core.CloseAllService()
	core.Wait()
	return
}
