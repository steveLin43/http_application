package global

import (
	"http_application/pkg/logger"
	"http_application/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettings
	AppSetting      *setting.AppSettings
	DatabaseSetting *setting.DatabaseSettings
	Logger          *logger.Logger
	JWTSetting      *setting.JWTSettings
)
