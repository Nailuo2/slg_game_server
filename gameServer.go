package main

import (
	"SLG_sg_server/config"
	"SLG_sg_server/net"
	"SLG_sg_server/server/game"
	"log"
)

func main() {
	host := config.File.MustValue("game_server", "host", "127.0.0.1")
	port := config.File.MustValue("game_server", "port", "8001")
	s := net.NewServer(host + ":" + port)
	s.NeedSecret(false)
	game.Init()
	s.Router(game.Router)
	s.Start()
	log.Println("游戏服务启动成功")
}
