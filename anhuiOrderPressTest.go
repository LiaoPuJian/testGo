package main

import (
	"fmt"
	"io/ioutil"
	"myUtil"
	"net/http"
	"runtime"
	"strings"
	"sync"
)

var (
	//生成订单的url
	url = "https://shouzhan1.51fubei.com/Home/AndroidElectricApi/actionElectricBillQuery"
	//声明一个等待组
	waitGroup sync.WaitGroup
	//设置单位时间并发数量
	num = 20
	//access_token
	access_token = "2019022118105067469428686"
	//户号
	bill_key = "7180005386"
)

/**
 * 用于安徽电网订单号并发的压力测试
 */
func main() {

	fmt.Println("开始压测")
	//设置cpu核数
	runtime.GOMAXPROCS(runtime.NumCPU())
	//设置计时器的数量
	waitGroup.Add(num)

	//参数
	params := make(map[string]interface{})

	params["access_token"] = access_token
	params["bill_key"] = bill_key

	//循环goroutine
	for i := 0; i < num; i++ {
		go doRequest(params)
	}
	//等待goroutine执行完毕
	waitGroup.Wait()
	fmt.Println("压测完毕")
}

/**
 * 安徽电网压测
 */
func doRequest(params map[string]interface{}) {
	content, err := myUtil.RequestConversion(params, "&")
	if err != nil {
		panic(err)
	}
	fmt.Println(content)
	//发送post请求
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(content))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	//读取响应
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	//等待组计时器减1
	waitGroup.Done()
}
