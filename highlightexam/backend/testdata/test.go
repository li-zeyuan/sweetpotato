package testdata

import (
	"context"

	"github.com/li-zeyuan/common/mylogger"
	"github.com/li-zeyuan/common/mysqlstore"
	"github.com/li-zeyuan/sweetpotato/highlightexam/backend/config"
)

func InitServer() (context.Context, error) {
	err := config.LoadConfigFile("/Users/zeyuan.li/Desktop/workspace/code/src/github.com/li-zeyuan/sun/deployment/templates/highlightexam-backend-dev/config.yaml")
	if err != nil {
		return nil, err
	}

	mylogger.Init(&config.AppCfg.Logging)

	return context.Background(), mysqlstore.New(&config.AppCfg.Mysql)
}
