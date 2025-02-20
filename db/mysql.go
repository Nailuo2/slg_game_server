package db

import (
	"SLG_sg_server/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	"log"
)

var Engine *xorm.Engine

func TestDB() {

	mysqlConfig, err := config.File.GetSection("mysql")
	if err != nil {
		log.Println("数据库配置缺失", err)
		panic(any(err))
	}
	dbConn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		mysqlConfig["user"],
		mysqlConfig["password"],
		mysqlConfig["host"],
		mysqlConfig["port"],
		mysqlConfig["dbname"],
	)
	Engine, err = xorm.NewEngine("mysql", dbConn)
	if err != nil {
		log.Println("数据库连接失败", err)
		panic(any(err))
	}
	err = Engine.Ping()
	if err != nil {
		log.Println("数据库ping不通", err)
		panic(any(err))
	}
	maxIdle := config.File.MustInt("mysql", "max_idle", 2)
	maxConn := config.File.MustInt("mysql", "max_conn", 2)
	Engine.SetMaxIdleConns(maxIdle)
	Engine.SetMaxOpenConns(maxConn)
	Engine.ShowSQL(true)
	log.Println("数据库初始化完成...")
}
