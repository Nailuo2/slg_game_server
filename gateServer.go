package main

import (
	"SLG_sg_server/config"
	"SLG_sg_server/net"
	"SLG_sg_server/server/gate"
	"log"
)

func main() {
	//host := config.File.MustValue("gate_server", "host", "127.0.0.1")
	port := config.File.MustValue("gate_server", "port", "8004")
	s := net.NewServer(":" + port)
	s.NeedSecret(true)
	gate.Init()
	s.Router(gate.Router)
	s.Start()
	log.Println("网关服务启动成功")
}
