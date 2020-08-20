package main

import (
	"fmt"
	"gojpush/conf"
	"gojpush/dao"
	"gojpush/model"
	"gojpush/routers"

	"gopkg.in/ini.v1"
)

var (
	cfg = new(conf.AppConf)
)

func loadConfig() (err error) {
	err = ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		panic("map failed")
	}
	return
}
func main() {
	//加载配置
	err := loadConfig()
	if err != nil {
		fmt.Printf("Load ini failed,err:%v", err)
	}
	//初始化MySQL
	err = dao.InitMySQL(cfg.Mysql)
	if err != nil {
		fmt.Printf("Init mysql failed,err:%v", err)
	}
	dao.DB.AutoMigrate(&model.App{})

	r := routers.SetupRouters()
	r.Run(":8080")
}
