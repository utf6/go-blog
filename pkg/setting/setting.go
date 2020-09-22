package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

type App struct {
	JwtSecret string
	PageSize int
	RuntimeRootPath string

	ImagePrefixUrl string
	ImageSavePath string
	ImageMaxSize int
	ImageAllowExts []string

	LogSavePath string
	LogSaveName string
	LogFileExt string
	TimeFormat string
}
var AppSetting = &App{}

type Server struct {
	RunMode string
	HttpPort int
	ReadTimeout time.Duration
	WriteTimeOut time.Duration
}
var ServerSetting = &Server{}

type Database struct {
	Type string
	User string
	Password string
	Host string
	Name string
	TablePrefix string
}
var DatabaseSetting = &Database{}

func SetUp()  {
	Cfg, err := ini.Load("conf/app.in")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.in': %v", err)
	}

	err = Cfg.Section("app").MapTo(AppSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo AppSetting error: %v", err)
	}

	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024

	err = Cfg.Section("server").MapTo(AppSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo ServerSetting error: %v", err)
	}

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeOut = ServerSetting.WriteTimeOut * time.Second

	err = Cfg.Section("database").MapTo(AppSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo DatabaseSetting error: %v", err)
	}
}
