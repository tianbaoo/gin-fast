package conf

import (
	"github.com/go-ini/ini"
	"time"
)

type SqlDataBase struct {
	Type     string
	Host     string
	Port     string
	User     string
	Password string
	DB       string
	Charset  string
	Prefix   string
}

type Jwt struct {
	SecretKey string
}

type Project struct {
	StaticUrlMapPath string
	TemplateGlob     string
	MediaFilePath	 string
}

type Server struct {
	Port string
	ReadTimeout time.Duration
	WriteTimeout time.Duration
}

type LoggerStruct struct {
	Mode string `json:"mode"`
	Port string `json:"port"`
	Level string `json:"level"`
	FilePath string `json:"filepath"`
	MaxSize int `json:"maxsize"`
	MaxAge int `json:"max_age"`
	MaxBackups int `json:"max_backups"`
}

type RedisStruct struct {
	Host string `json:"host"`
	Port string `json:"port"`
	User string `json:"user"`
	Password string `json:"password"`
	DB int `json:"db"`
	PoolSize int `json:"poolsize"`
}

var (
	DataBase     = &SqlDataBase{}
	JwtSecretKey = &Jwt{}
	ProjectCfg   = &Project{}
	HttpServer 	= &Server{}
	LoggerCfg   = &LoggerStruct{}
	RedisCfg   = &RedisStruct{}
)

func SetUp() {
	cfg, err := ini.Load("conf/conf.ini")
	if err != nil {
		panic(err)
	}
	if err := cfg.Section("mysql").MapTo(DataBase); err != nil {
		panic(err)
	}
	if err := cfg.Section("jwt").MapTo(JwtSecretKey); err != nil {
		panic(err)
	}
	if err := cfg.Section("project").MapTo(ProjectCfg); err != nil {
		panic(err)
	}
	if err := cfg.Section("server").MapTo(HttpServer); err != nil {
		panic(err)
	}
	if err := cfg.Section("log").MapTo(LoggerCfg); err != nil {
		panic(err)
	}
	if err := cfg.Section("redis").MapTo(RedisCfg); err != nil {
		panic(err)
	}
}
