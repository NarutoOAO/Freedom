package main

import (
	"9900project/conf"
	util "9900project/pkg/utils"
	"9900project/repository/db/dao"
	"9900project/router"
	"9900project/service"
	"fmt"
	"log"
	"net/http"
)

func main() {
	conf.Init()
	dao.InitMySQL()
	//cache.InitCache()
	util.InitLog()
	r := router.NewRouter()
	// define a go routine to start websocket
	go func() {
		service.ConnectWebSocket()
		log.Println("Server started on http://localhost:8080")
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			fmt.Println("connection failed")
			log.Fatal("Server error: ", err)
		}
	}()
	_ = r.Run(conf.HttpPort)
}
