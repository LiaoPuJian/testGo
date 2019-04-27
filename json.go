package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/pquerna/ffjson/ffjson"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	jsonA, err := json.Marshal(a)
	if err != nil {
		panic("报错了！")
	}

	fmt.Println(string(jsonA))

	//对jsonA解码
	var a1 interface{}
	json.Unmarshal([]byte(jsonA), &a1)
	fmt.Println(a1)

	ffJosn, err := ffjson.Marshal(a)
	if err != nil {
		panic("ff报错了！")
	}
	fmt.Println(string(ffJosn))

	md5A := md5.New()

	md5A.Write([]byte("pipikai shi ge zhizhang"))
	str := md5A.Sum([]byte(""))
	fmt.Printf("%x\n\n", str)

}
