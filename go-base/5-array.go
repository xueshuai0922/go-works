package main

import "fmt"

/*
*

		数组 ：  var 数组名称 [数组长度] 数组的类型；数组也是值类型
	    值类型：开辟新的内存空间，将值复制过来，是值传递（备份）不是内存地址传递
	            int,float,string,数组
	    引用类型：slice ,map
*/
func main() {
	//1.第一种定义方式
	var arr [5]int
	arr[0] = 1
	fmt.Println(arr)
	//2.快捷的方式
	ss := [...]int{1, 2, 3, 4, 5}
	fmt.Println(ss)
	fmt.Println(len(ss))
	//3循环数组
	sum := 0
	//for index,value := range 数组
	for _, i := range arr {
		sum += i
	}
	fmt.Println("和为", sum)
	//4.测试数组值传递
	acopy := arr
	acopy[0] = 100
	fmt.Println(arr)          //[1 0 0 0 0]
	fmt.Println(acopy)        //[100 0 0 0 0]
	fmt.Println(arr == acopy) //false
	acopy[0] = 1
	fmt.Println(arr == acopy) //true
}
