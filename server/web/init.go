package web

import (
	"SLG_sg_server/db"
	"SLG_sg_server/server/web/controller"
	"SLG_sg_server/server/web/middleware"
	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	//测试数据库，并且初始化数据库
	db.TestDB()
	//还有别的初始化方法
	initRouter(router)
}

func initRouter(router *gin.Engine) {
	router.Use(middleware.Cors())
	router.Any("/account/register", controller.DefaultAccountController.Register)
}
