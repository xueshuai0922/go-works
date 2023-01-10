package main

import "fmt"

/*
*
函数相关学习
1.go的函数可以返回多个值
2,main 函数只能在package main中
3.func 本身也是一个数据类型，理解成和其他基础数据类型一样的数据类型，也是可以赋值和作为参数的
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

	//4.函数
	fmt.Printf("函数的格式为 %T", func1) //函数的格式为 func(int, int) int
	func2 := func1                 //复制出来一个函数
	result := func2(1, 3)
	fmt.Println(result)

	//4.2 函数可以作为另一个函数的一个参数 真秒！！
	ss := funcInside(1, 2, func1)
	fmt.Println(ss)
	plus := funcInside(3, 5, func(a int, b int) (s int) {
		return a * b
	})
	fmt.Println("回调函数为参数，回调函数内a*b ", plus)

	//4.3 匿名函数
	func() {
		fmt.Println("这是匿名函数")
	}()

	func(a, b int) {
		fmt.Println(a + b)
	}(1, 3)

	res := func(a, b string) string {
		return a + b
	}("hello", " world")
	fmt.Println("匿名函数返回结果值", res)
	//5 可变参数
	variable(1, 2, 3, 4, 5)
	slices := []int{222, 12, 34, 44}
	variable(slices...)
	fmt.Println()
	//6.闭包
	f := close()
	i := f()
	i = f()
	i = f()
	fmt.Println("i闭包结果", i)

	f1 := close()
	j := f1()
	fmt.Println("j闭包结果", j)
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
func func1(a, b int) (s int) {
	fmt.Println(a, b)
	return a + b
}

/*
*
传入一个函数参数 和其他数据类型的参数，根据这个函数参数的不同，实现不同功能
这里f为回调函数
*/
func funcInside(a int, b int, f func(int, int) (s int)) int {
	return f(a, b)
}

func variable(nums ...int) {
	for _, v := range nums {
		fmt.Printf("%T,%d ", v, v)
	}
}

// 闭包结构：外层函数的局部变量和内层函数统一整体称为闭包；外层函数的局部变量
func close() func() int {
	i := 0
	return func() int {
		i++
		return i
	}

}
