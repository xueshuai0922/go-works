package main

import "fmt"

// 多个变量定义的时候，变量之间必须进行换行
var (
	cp string
	np int
)

// mainvar:=23//全局变量不能简要定义变量
var mainvar = "mainvar"

func main() {
	变量()
	//多线程()

}

func 多线程() {
	for i := 0; i < 10000; i++ {
		go fmt.Println(i)
	}
}
func 变量() {
	//第一种：
	var name string = "xuehsuai"

	fmt.Println(name)
	//	多个变量设置
	var sex, age = "男", 12
	fmt.Println(sex)
	fmt.Println(age)
	//简化的赋值
	s := "xueshuai"
	fmt.Println(s)
	//【重点】 不要的返回用下划线标识，go的函数有多个返回参数
	_, money := 5, 100000000
	fmt.Println(money)

	var isFlag bool
	fmt.Println(isFlag) //默认false
	isFlag = true
	fmt.Println(isFlag) //true

}
