package main

import (
	"blogServer/core"
	"blogServer/flag"
	"blogServer/global"
	"blogServer/routers"
	"fmt"
	"log"
	"os"
)

func init() {
	var err error

	option := flag.Parse()

	if err = core.InitConf(); err != nil {
		log.Panicln(err)
		panic(err)
	} else {
		log.Println("config yamlFile load init success.")
	}

	core.InitLogger()
	global.Log.Infof("Log init success.")

	if err = core.InitDB(); err != nil {
		global.Log.Error(err)
		panic(err)
	} else {
		global.Log.Infof("db init success.")
	}

	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		os.Exit(1)
	}
}

func main() {
	r := routers.InitRouter()
	addr := fmt.Sprintf("%s:%d", global.Config.System.Host, global.Config.System.Port)
	global.Log.Infof("running on %s", addr)
	if err := r.Run(addr); err != nil {
		global.Log.Errorf(err.Error())
	}
}
