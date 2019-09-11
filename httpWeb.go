package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type shellFunc func(w http.ResponseWriter, r *http.Request) error

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
	path := r.URL.Path[len("/"):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	w.Write(all)
	return nil
}

func main() {
	http.HandleFunc("/", handleShellFunc(sayhelloName)) //设置访问的路由
	err := http.ListenAndServe(":9090", nil)            //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
