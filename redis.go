package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	//跟redis建立连接
	conn, err := redis.Dial("tcp", "192.168.1.135:6379")
	if err != nil {
		fmt.Println("连接redis失败,", err)
		return
	}
	defer conn.Close()

	auth := authPassword(conn)
	if !auth {
		fmt.Println("redis密码验证失败")
		return
	}

	operaKeyValue(conn)
}

/**
验证redis密码
*/
func authPassword(conn redis.Conn) bool {
	_, err := conn.Do("auth", "redis123456")
	if err != nil {
		fmt.Println("auth err", err)
		return false
	}
	return true
}

/**
操作键值对
*/
func operaKeyValue(conn redis.Conn) {
	//设置键值对
	_, err := conn.Do("Set", "name", "Keima!")
	if err != nil {
		fmt.Println("Set err,", err)
		return
	}
	//通过键获取值
	value, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("Get err,", err)
		return
	}
	fmt.Println(value)
}

//操作hash
func operaHash(conn redis.Conn) {

}
