package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

//定义一个函数类型
type shellFunc func(w http.ResponseWriter, r *http.Request) error

//接收这个函数类型，并且返回一个可以被HandleFunc接受的func
func handleShellFunc(shell shellFunc) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := shell(w, r)
		if err != nil {
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			default:
				code = http.StatusInternalServerError
			}
			http.Error(w, http.StatusText(code), code)
		}
	}
}

func sayhelloName(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte("hello world rua!"))
	go func() {
		fmt.Println(222)
		time.Sleep(time.Second * 10)
		fmt.Println(111)
	}()

	go func() {
		fmt.Println(333)
		time.Sleep(time.Second * 10)
		fmt.Println(444)
	}()
	return nil
}

func main() {
	http.HandleFunc("/", handleShellFunc(sayhelloName)) //设置访问的路由
	err := http.ListenAndServe(":9090", nil)            //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
