package main

import "fmt"

/*
*
函数相关学习
1.go的函数可以返回多个值
2,main 函数只能在package main中
*/

// 先执行包中的init方法
func init() {
	fmt.Println("init method")
}
func main() {
	//1.go函数可以返回多个值
	x, y := "1", "2"
	fmt.Printf("交换前 x 的值为 : %s\n", x)
	fmt.Printf("交换前 y 的值为 : %s\n", y)
	a, b := swap(x, y)
	println(a, b) //2 1
	//2.go是值传递,简单讲就是将变量中的值复制出来进行传递
	swapValue(x, y) //1 2，swapValue中x,y只是对x,y复制过去进行
	fmt.Printf("交换后 x 的值为 : %s\n", x)
	fmt.Printf("交换后 y 的值为 : %s\n", y)
	//3.go可以使用指针---引用传递
	fmt.Printf("x的引用地址%d,y的引用地址%d\n\n", &x, &y)
	swapRef(&x, &y)
	fmt.Printf("x的引用地址%d,y的引用地址%d\n\n", &x, &y)
	fmt.Printf("引用交换后 x 的值为 : %s\n", x)
	fmt.Printf("引用交换后 y 的值为 : %s\n", y)
}
func swap(x, y string) (string, string) {
	return y, x
}
func swapValue(x, y string) {
	var temp string
	temp = x
	x = y
	y = temp
	fmt.Println(x, y) //2 1
}
func swapRef(x *string, y *string) {
	var temp string
	temp = *x           //保存x地址上的值
	*x = *y             //将y地址上的值赋值给x
	*y = temp           //
	fmt.Println(*x, *y) //2 1
}
