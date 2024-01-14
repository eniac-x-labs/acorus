package logs

import (
	"fmt"
	"gorm.io/gorm/logger"
	"io"
	"log"
	"os"
	"time"

	"github.com/cornerstone-labs/acorus/common/global_const"
)

var myLogWriter dateFileWriter

var targetFile *os.File

var logTime string

type dateFileWriter struct {
	io.Writer
	LogLevel logger.LogLevel // 日志级别
}

type MyLogger struct {
	logger *log.Logger
}

func (m MyLogger) Printf(format string, v ...any) {

	// 检测到如果是gorm的sql语句，就截取sql语句
	if global_const.GormInfoFmt == format {
		sql := v[3].(string)
		if len(sql) > 1000 {
			sql = sql[:1000] + "..."
			v[3] = sql
		}
	}
	m.logger.Printf(" 走你"+format, v...)
}

func init() {
	RefreshLogFile()
	myLogWriter = dateFileWriter{
		LogLevel: logger.Info,
	}
	log.SetOutput(MyLogWriter())
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func StartLog() {
	fmt.Println("startLog")
}

func (l *dateFileWriter) Write(p []byte) (n int, err error) {
	if logTime != time.Now().Format(global_const.LogTimeFormat) {
		RefreshLogFile()
	}
	return targetFile.Write(p)
}

func RefreshLogFile() {
	logTime = time.Now().Format(global_const.LogTimeFormat)
	f, err := GetLogFile()
	if err != nil {
		fmt.Println(err)
		return
	}
	targetFile = f
	fmt.Println("日志文件切割成功~")
}

func MyLogWriter() io.Writer {
	return io.MultiWriter(os.Stdout, &myLogWriter)
}

func New() *MyLogger {
	myLogger := &MyLogger{log.New(MyLogWriter(), "\r\n", log.Ldate|log.Ltime|log.LstdFlags)}
	return myLogger
}
