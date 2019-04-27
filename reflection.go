package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Hello() {
	fmt.Println("Hello World.")
}

//结构体User的绑定方法
func (u User) CallReflect(name string, id, age int) {
	fmt.Printf("My name is %s, I am %d, My id is %d, ni hao %s", name, age, id, u.Name)
}

func main() {
	/*u := User{
		Id:   1,
		Name: "LPJ",
		Age:  2,
	}
	info(u)*/

	/*	m := 123
		vOfValue := reflect.ValueOf(&m)
		vOfValue.Elem().SetInt(999)
		fmt.Println(m)*/
	user := User{2, "lpj", 18}
	user.CallReflect("jjy", 10, 20)
	//获取user的反射值类型
	v := reflect.ValueOf(user)
	//获取方法
	mehoodU := v.MethodByName("CallReflect")
	//设置参数
	args := []reflect.Value{reflect.ValueOf("YJJ"), reflect.ValueOf(3), reflect.ValueOf(5)}
	mehoodU.Call(args)
}

//反射方法   传递一个空的接口
func info(o interface{}) {
	//使用反射获取这个空接口的类型集合
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name())

	//判断传入的类型是否正确
	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("传入了错误的类型")
		return
	}

	//使用反射获取这个空接口的值集合
	v := reflect.ValueOf(o)
	fmt.Println("Value:", v)

	fmt.Println("num:", t.NumField(), "Field:", t.Field(0))

	for i := 0; i < t.NumField(); i++ {
		fmt.Printf("%6s:%v => %v\n", t.Field(i).Name, t.Field(i).Type, v.Field(i).Interface())
	}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s:%v", m.Name, m.Type)
	}
}
