package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

//这个是一个爬虫程序，用于爬取一拳超人的漫画。

func main() {

	//要爬取的页面
	url := "https://manhua.fzdm.com/132/151/index_0.html"

	//获取一个收集器
	c := colly.NewCollector()

	//设置代理和请求头
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36"

	//请求之前的设置
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Upgrade-Insecure-Requests", "1")
	})

	//收到相应之后的处理
	c.OnResponse(func(resp *colly.Response) {
		/*response := string(resp.Body)
		fmt.Println(response)*/
	})

	//爬取到html后
	c.OnHTML("#mhimg0 > a > img", func(e *colly.HTMLElement) {
		//todo 将爬取到的图片下载，并且放入到桌面的文件夹中
		//link := e.ChildAttr("#")
		fmt.Println(e)
	})

	//错误时的报错信息
	c.OnError(func(resp *colly.Response, errHttp error) {
		fmt.Println(resp, errHttp)
	})

	err := c.Visit(url)

	if err != nil {
		fmt.Println("报错啦！", err)
	}
}
