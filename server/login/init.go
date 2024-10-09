package login

import (
	"SLG_sg_server/db"
	"SLG_sg_server/net"
	"SLG_sg_server/server/login/controller"
)

var Router = net.NewRouter()

func Init() {
	//测试数据库，并且初始化
	db.TestDB()
	initRouter()
}

func initRouter() {
	controller.DefaultAccount.Router(Router)
}
