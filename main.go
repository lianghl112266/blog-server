package main

import (
	"blogServer/core"
	"blogServer/global"
	"blogServer/routers"
	"fmt"
	"log"
)

func init() {
	var err error

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

}

func main() {
	r := routers.InitRouter()
	addr := fmt.Sprintf("%s:%d", global.Config.System.Host, global.Config.System.Port)
	global.Log.Infof("running on %s", addr)
	r.Run(addr)
}
