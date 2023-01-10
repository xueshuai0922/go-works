package main

import "fmt"

/*
*
switch
*/
func main() {
	//1.switch 默认写法
	num := 1
	switch num {
	case 1:
		fmt.Println("第一季度")
		fallthrough //可以穿透到下一个case中
	case 2:
		fmt.Println("第二季度")
		fmt.Println("第二季度")
		fmt.Println("第二季度")
		break
		fmt.Println("第二季度")

	case 3:
		fmt.Println("第三季度")
	case 4:
		fmt.Println("第四季度")
	default:
		fmt.Println("季度有误")
	}
	//2.如果不写任何变量，默认是true
	switch {
	case true:
		fmt.Println("true") //输出true
	case false:
		fmt.Println("false")
	}
	//3.case 可以匹配多个值
	a := 3
	switch a {
	case 1, 3, 5, 7, 9:
		fmt.Println("这是奇数")
	case 2, 4, 6, 8, 10:
		fmt.Println("这是偶数")
	}
}
