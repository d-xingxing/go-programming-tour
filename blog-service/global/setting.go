package global

import (
	"github.com/d-xingxing/go-programming-tour/blog-service/pkg/logger"
	"github.com/d-xingxing/go-programming-tour/blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
	JWTSetting      *setting.JWTSettingS
	EmailSetting    *setting.EmailSettingS
)
