package gorm_log

import (
	"gorm.io/gorm/logger"
)

type CustomLogger struct {
	gprmLogger *logger.Writer
}

func (l *CustomLogger) Printf(format string, v ...interface{}) {
	// 在这里实现你的自定义逻辑
	// 这只是一个示例，将日志消息添加前缀并调用原始的 Printf 方法
	//prefix := "[MyLogger] "
	//l.Interface.
	//l.gprmLogger.(prefix+format, v...)
}
