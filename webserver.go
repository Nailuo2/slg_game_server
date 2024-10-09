package main

import (
	"SLG_sg_server/config"
	"SLG_sg_server/server/web"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	//host := config.File.MustValue("web_server", "host", "127.0.0.1")
	port := config.File.MustValue("web_server", "port", "8088")

	router := gin.Default()
	//路由
	web.Init(router)
	s := &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	log.Println(err)
}
