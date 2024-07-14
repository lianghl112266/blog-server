package main

import (
	"blogServer/core"
	"blogServer/global"
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
	global.Log.Info("Log init success.")

	if err = core.InitDB(); err != nil {
		global.Log.Error(err)
		panic(err)
	} else {
		global.Log.Info("db init success.")
	}

}

func main() {

}
