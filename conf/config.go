package conf

type AppConf struct {
	Jpush JpushConf `ini:"jpush"`
	Mysql MySQLConf `ini:"mysql"`
}
type JpushConf struct {
	AppKey       string
	MasterSecret string
}
type MySQLConf struct {
	DataBase string `ini:"DataBase"`
	Ip       string `ini:"IP"`
	Root     string `ini:"Root"`
	Psw      string `ini:"Password"`
}
type AppInfo struct {
	AppKey string
	ResID  string
}
