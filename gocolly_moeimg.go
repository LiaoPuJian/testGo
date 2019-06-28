package main

import (
	"github.com/keima/spider"
	"sync"
)

var (
	wg = sync.WaitGroup{}
)

func main() {
	wg.Add(2)
	//爬取moeimg的内容
	go func() {
		spider.MoeimgRun(5)
		wg.Done()
	}()
	//爬取yande的内容
	go func() {
		spider.YandeRun(5)
		wg.Done()
	}()
	wg.Wait()
}
