package main

import (
	"SLG_sg_server/config"
	"SLG_sg_server/net"
	"SLG_sg_server/server/chat"
)

func main() {
	host := config.File.MustValue("chat_server", "host", "127.0.0.1")
	port := config.File.MustValue("chat_server", "port", "8002")
	s := net.NewServer(host + ":" + port)
	s.NeedSecret(false)
	chat.Init()
	s.Router(chat.Router)
	s.Start()
}
