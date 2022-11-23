package main

import "fmt"

/*
*
函数相关学习
1.go的函数可以返回多个值
2,main 函数只能在package main中
*/

func init() {
	fmt.Println("init method")
}
func main() {
	a, b := swap("1", "2")
	println(a, b)
}
func swap(x, y string) (string, string) {
	return y, x
}
