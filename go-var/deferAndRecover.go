package main

import "fmt"

/*
*
defer:延迟
defer函数作用：1.捕获异常 2.释放资源 3.输出日志
recover:Go语言提供了专用于“拦截”运行时panic的内建函数“recover”。它可以是当前的程序从运行时panic的状态中恢复并重新获得流程控制权。
*/

func deferDemo() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}

/*
*
模拟错误，并捕获错误打印日志
*/
func recoverDemo(i int) {
	var arr [10]int
	//这里有点类似Java中的匿名函数；错误捕获需要写在可能出现错误前边
	defer func() {
		err := recover()
		fmt.Println(err) //产生异常进行打印异常 ---runtime error: index out of range [10] with length 10
	}()
	arr[i] = 10               //下标越界了，会进行报错，defer延迟函数中对异常进行捕获，并不会影响下一步
	fmt.Println("1111111111") //这里输出并没有进行，正面recover只是对子函数中的方法进行拦截错误，然后对主函数中的流程不影响
}
func main() {
	recoverDemo(10)
	fmt.Println("报出越界错误后，仍然可以进行向下执行")
	deferDemo()
}

//结果
//runtime error: index out of range [10] with length 10
//报出越界错误后，仍然可以进行向下执行
//3
//2
//1
