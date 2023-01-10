package main

import "fmt"

/*
*
for 循环 go的写法;
break 退出循环
continue 退出本次循环
*/
func main() {
	//1，正常循环的写法
	for i := 1; i < 2; i++ {
		fmt.Println("hello word", i)
	}
	//2.嵌入循环，break continue只控制当前的循环体
	for i := 0; i < 10; i++ {
		if i == 5 {
			//break
			continue
		}
		fmt.Println("输出 ", i)
	}
	//3.可以给循环打上标签，并进行控制外层循环退出
out:
	for i := 0; i < 4; i++ {
		for j := 0; j < 6; j++ {
			if j == 3 {
				//break out
				continue out
			}
			fmt.Printf("i=%d,j=%d \n", i, j)
		}
	}
}
