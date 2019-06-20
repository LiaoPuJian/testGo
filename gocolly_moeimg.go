package main

import (
	"bufio"
	"fmt"
	"github.com/gocolly/colly"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
	"sync"
)

var (
	imgFile = "E:\\keke\\"
	wg      = sync.WaitGroup{}
)

func main() {

	wg.Add(1000000)

	//要爬取的页面
	url := "http://moeimg.net/"

	//获取一个收集器
	c := colly.NewCollector()

	//设置代理和请求头
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36"

	//请求之前的设置
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Upgrade-Insecure-Requests", "1")
		r.Headers.Set("Accept-Encoding", "gzip, deflate")
		r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3")
		r.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Connection", "moeimg.net")
		r.Headers.Set("Host", "keep-alive")
		r.Headers.Set("If-None-Match", "3bed-58bbc967940f8")
		r.Headers.Set("Cookie", "__cfduid=df97a01336e673c1cdc31c97929055bdf1561017278; _ga=GA1.2.1164906477.1561017281; _gid=GA1.2.11570314.1561017281")
	})

	//收到相应之后的处理
	c.OnResponse(func(resp *colly.Response) {
		/*response := string(resp.Body)
		fmt.Println(response)*/
	})

	//爬取到html后
	c.OnHTML("a[href],img", func(e *colly.HTMLElement) {
		if e.Name == "a" {
			link := e.Attr("href")
			if -1 == strings.Index(link, "http") {
				link = url + link
			}
			fmt.Println("获取连接：", link)
			go c.Visit(link)
		} else {
			imgSrc := e.Attr("src")
			if -1 == strings.Index(imgSrc, "http") {
				imgSrc = url + imgSrc
			}
			fmt.Println("获取图片url:", imgSrc)
			go saveFile(imgSrc)
		}
		wg.Done()
	})

	//错误时的报错信息
	c.OnError(func(resp *colly.Response, errHttp error) {
		fmt.Println(errHttp)
	})

	err := c.Visit(url)

	if err != nil {
		fmt.Println("报错啦！", err)
	}

	wg.Wait()
}

func saveFile(imgSrc string) {
	res, err := http.Get(imgSrc)
	if err != nil {
		fmt.Println("A error occurred!")
		return
	}
	// defer后的为延时操作，通常用来释放相关变量
	defer res.Body.Close()
	// 获得get请求响应的reader对象
	reader := bufio.NewReaderSize(res.Body, 32*1024)

	fileName := path.Base(imgSrc)

	file, err := os.Create(imgFile + fileName)
	if err != nil {
		return
	}
	// 获得文件的writer对象
	writer := bufio.NewWriter(file)

	io.Copy(writer, reader)
}
