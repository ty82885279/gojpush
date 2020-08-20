package model

import (
	"fmt"
	"gojpush/conf"
	"gojpush/dao"
)

type App struct {
	ID           int64  `gorm:"AUTO_INCREMENT;primary_key;column:id"`
	AppID        string `gorm:"column:appid"`
	AppKey       string `gorm:"column:appkey"`
	MasterSecret string `gorm:"column:mastersecret"`
	Address      string `gorm:"column:address"`
}

func SendPush(Appinfo *conf.AppInfo) (appObj *App) {
	appObj = new(App)

	var count int
	dao.DB.Model(&App{}).Where(" appkey = ? ", Appinfo.AppKey).Count(&count)
	if count != 0 {
		err := dao.DB.Where(" appkey = ? ", Appinfo.AppKey).Find(appObj).Error
		if err != nil {
			fmt.Printf("select app failed,err:%v\n", err)
		}
	}
	return
}
