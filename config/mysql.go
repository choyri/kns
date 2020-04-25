package config

import (
	"fmt"
	"github.com/choyri/kns/support"
)

type MySQL struct {
	Host      string
	Port      uint64
	Username  string
	Password  string
	MaxIdle   int
	MaxOpen   int
	EnableLog bool
	Database  string
}

func GetMySQL() MySQL {
	return MySQL{
		Host:      support.GetStringEnv("DB_HOST", "localhost"),
		Port:      support.GetUintEnv("DB_PORT", 3306),
		Username:  support.GetStringEnv("DB_USERNAME", "root"),
		Password:  support.GetStringEnv("DB_PASSWORD", "root"),
		MaxIdle:   int(support.GetUintEnv("DB_MAX_IDLE", 20)),
		MaxOpen:   int(support.GetUintEnv("DB_MAX_OPEN", 50)),
		EnableLog: support.GetBoolEnv("DB_ENABLE_LOG"),
		Database:  support.GetStringEnv("DB_DATABASE", "kns"),
	}
}

func (cfg *MySQL) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?collation=utf8mb4_unicode_ci&loc=Local&parseTime=true",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database)
}
