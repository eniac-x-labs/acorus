package logs

import (
	"log"
	"os"
	"time"

	"crypto/rand"
	"encoding/base64"

	"github.com/cornerstone-labs/acorus/common/global_const"
)

func GetLogFile() (*os.File, error) {
	dir := "./logs"
	b, mkdirErr := DirExists(dir)
	if !b && mkdirErr == nil {
		err := os.Mkdir(dir, os.ModePerm)
		if err != nil {
			log.Println("logs目录创建失败~")
			return nil, err
		} else {
			log.Println("logs目录创建成功~")
		}
	}
	//randomString, _ := generateRandomString(5)
	file := "./logs/" + time.Now().Format(global_const.LogTimeFormat) + "_info" + ".log"
	f, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		// 关闭
		defer f.Close()
		return nil, err
	}
	return f, err
}

// DirExists 判断目录是否存在
func DirExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func generateRandomString(length int) (string, error) {
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}
	randomString := base64.RawURLEncoding.EncodeToString(buffer)
	return randomString[:length], nil
}
