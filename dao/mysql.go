package dao

import (
	"fmt"
	"gojpush/conf"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	DB *gorm.DB
)

func InitMySQL(cfg conf.MySQLConf) (err error) {

	dst := fmt.Sprintf(`%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local`, cfg.Root, cfg.Psw, cfg.Ip, cfg.DataBase)
	fmt.Println(dst)
	DB, err = gorm.Open("mysql", dst)
	if err != nil {
		fmt.Printf("open mysql failed,err:%v", err)
	}

	return DB.DB().Ping()
}
