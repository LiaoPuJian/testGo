package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isMatch("mississippi", "mis*is*p*."))
}

func isMatch(s string, p string) bool {
	if p == ".*" || (s == "" && p == "") {
		return true
	}
	if len(s) >= 2 && len(p) == 1 {
		return false
	}

	var old int32

	for k, v := range s {
		var flag bool
		jump := true
		a := 1
		fmt.Println("第"+strconv.Itoa(k)+"次循环", string(s), string(p))
		for k1, v1 := range p {
			fmt.Println(string(v), string(v1), old)
			if v1 == '.' || v == v1 {
				flag = true
				old = v1
				fmt.Println("1将v1的值赋予old", v1)
				break
			}
			if v1 == '*' {
				fmt.Println("v1和old分别为:", string(v1), old, v)
				//判断前一个值是'.'或者是和v1相等的值
				if old == '.' || old == v {
					fmt.Println("匹配上了", string(v), old)
					flag = true
					//old = v1
					jump = false
					break
				} else {
					fmt.Println("没匹配上，继续匹配下一个", k1)
					a += k1 + 1
				}
			}
			old = v1
			fmt.Println("2将v1的值赋予old", v1)
		}

		if !flag {
			return false
		} else {
			if jump {
				//丢弃前面已经循环过的p的字符
				if a > len(p) {
					return false
				}
				p = string(p[a:])
				fmt.Println("丢弃字符串后的p:", p)
			}
		}
	}
	return true
}
