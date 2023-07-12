package main

import (
	"9900project/conf"
	util "9900project/pkg/utils"
	"9900project/repository/db/dao"
	"9900project/router"
)

func main() {
	conf.Init()
	dao.InitMySQL()
	//cache.InitCache()
	util.InitLog()
	r := router.NewRouter()
	_ = r.Run(conf.HttpPort)
}
