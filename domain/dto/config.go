package dto

import "github.com/linchengzhi/go-clean-backend/Infra/logger"

type Config struct {
	AppName   string         `yaml:"app_name"`
	Env       string         `yaml:"env"`
	DebugPort string         `yaml:"debug_port"`
	HTTP      HTTP           `yaml:"http"`
	Log       *logger.Config `yaml:"log"`
	Mysql     Mysql          `yaml:"mysql"`
	Redis     Redis          `yaml:"redis"`
}

type HTTP struct {
	Port string `yaml:"port"`
}

type Mysql struct {
	Host         string `json:"host" yaml:"host"`                     // 服务器地址
	Port         string `json:"port" yaml:"port"`                     // 端口
	Dbname       string `json:"dbname" yaml:"dbname"`                 // 数据库名
	Username     string `json:"username" yaml:"username"`             // 数据库用户名
	Password     string `json:"password" yaml:"password"`             // 数据库密码
	Config       string `json:"config" yaml:"config"`                 // 高级配置
	MaxIdleConns int    `json:"max_idle_conns" yaml:"max_idle_conns"` // 空闲中的最大连接数
	MaxOpenConns int    `json:"max_open_conns" yaml:"max_open_conns"` // 打开到数据库的最大连接数
	MaxLifeTime  int    `json:"max_life_time" yaml:"max_life_time"`
	LogLevel     string `json:"logLevel" yaml:"log_level"`
}

type Redis struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}
