package chat

import (
	"SLG_sg_server/net"
	"SLG_sg_server/server/chat/controller"
)

var Router = &net.Router{}

func Init() {
	initRouter()
}

func initRouter() {

	controller.ChatController.Router(Router)

}
