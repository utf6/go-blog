package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

type App struct {
	JwtSecret 			string
	PageSize 			int
	RuntimeRootPath 	string

	PrefixUrl 			string
	ImageSavePath 		string
	ImageMaxSize 		int
	ImageAllowExts 		[]string

	LogSavePath			string
	LogSaveName			string
	LogFileExt			string
	TimeFormat 			string

	ExportSavePath 		string
	QrCodeSavePath		string
}
var AppSetting = &App{}

type Server struct {
	RunMode 		string
	HttpPort 		int
	ReadTimeout 	time.Duration
	WriteTimeOut 	time.Duration
}
var ServerSetting = &Server{}

type Database struct {
	Type 		string
	User 		string
	Password 	string
	Host 		string
	Name 		string
	TablePrefix string
}
var DatabaseSetting = &Database{}

type Redis struct {
	Host 		string
	Password 	string
	MaxIdle 	int
	MaxActive 	int
	IdleTimeout time.Duration
}
var RedisSetting = &Redis{}

var cfg *ini.File

func SetUp() {
	var err error

	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	mapTo("redis", RedisSetting)

	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeOut = ServerSetting.WriteTimeOut * time.Second
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second
}

func mapTo(section string, v interface{})  {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("cfg.mapto %s err: %v", section, err)
	}
}