package util

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

var (
	fileDir = "E:\\keke\\"
)

func SaveFile(fileSrc string) {
	res, err := http.Get(fileSrc)
	if err != nil {
		fmt.Println("A error occurred!")
		return
	}
	// defer后的为延时操作，通常用来释放相关变量
	defer res.Body.Close()
	// 获得get请求响应的reader对象
	reader := bufio.NewReaderSize(res.Body, 32*1024)

	fileName := path.Base(fileSrc)

	file, err := os.Create(fileDir + fileName)
	if err != nil {
		return
	}
	// 获得文件的writer对象
	writer := bufio.NewWriter(file)

	io.Copy(writer, reader)
}
