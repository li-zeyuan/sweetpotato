package dao

import (
	"github.com/li-zeyuan/common/mylogger"
	"github.com/li-zeyuan/common/mysqlstore"
	"github.com/li-zeyuan/sweetpotato/highlightexam/backend/config"
)

func InitDao() error {
	cfg := &config.Config{
		Mysql: mysqlstore.Config{
			DSN:     "root:root@tcp(localhost:3306)/government_exam?charset=utf8mb4&parseTime=True&loc=UTC",
			MaxConn: 1,
			MaxOpen: 1,
			Timeout: 10,
		},
		Logging: mylogger.LoggerCfg{},
	}
	config.AppCfg = cfg
	mylogger.Init(nil)
	New(cfg)

	return mysqlstore.New(&cfg.Mysql)
}
