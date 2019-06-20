package util

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"time"
)

var (
	fileDir = "E:\\keke\\"
)

func SaveFile(fileSrc string) {

	fileName := path.Base(fileSrc)

	if has, _ := PathExists(fileDir + fileName); has {
		client := http.Client{}
		client.Timeout = time.Second * 150
		res, err := client.Get(fileSrc)

		//res, err := http.Get(fileSrc)
		if err != nil {
			fmt.Println("保存图片时报错了！", err)
			return
		}
		// defer后的为延时操作，通常用来释放相关变量
		defer res.Body.Close()
		// 获得get请求响应的reader对象
		reader := bufio.NewReaderSize(res.Body, 32*1024)

		file, err := os.Create(fileDir + fileName)
		if err != nil {
			return
		}
		// 获得文件的writer对象
		writer := bufio.NewWriter(file)

		io.Copy(writer, reader)
	}

}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
