package gate

import (
	"SLG_sg_server/net"
	"SLG_sg_server/server/gate/controller"
)

var Router = &net.Router{}

func Init() {
	initRouter()
}

func initRouter() {
	controller.GateHandler.Router(Router)
}
