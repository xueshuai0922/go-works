package main

import "fmt"

/*
  - go中的
    1.slice就是Java中的list 可变数组=动态数组,底层指向就是一个数组
    2.实例化过程 ： make([]type,length,capacity ) type指的是数据类型，length是动态数组的长度（实际占用），capacity是数组的容量（可以占用长度）
    make指的是开创空间
*/
func main() {
	//1.切片的第一种方式
	list := make([]int, 2, 5)
	fmt.Println(cap(list), len(list), list) //5 2 [0 0]
	printSlice(list)                        //len=2 ,cap =5, slice=[0 0]

	var nullArr []int
	printSlice(nullArr) //len=0 ,cap =0, slice=[]
	if nullArr == nil {
		fmt.Println("切片是空的")
	}

	//2.切片的第二种方式，直接赋值
	numbers := []int{1, 2, 3, 4, 5, 6, 7}
	printSlice(numbers)      //addres=0xc00001a200, len=7 ,cap =7, slice=[1 2 3 4 5 6 7]
	printSlice(numbers[1:4]) //addres=0xc00001a208, len=3 ,cap =6, slice=[2 3 4]

	printSlice(numbers[:4]) //[1 2 3 4]
	printSlice(numbers[3:]) //[4 5 6 7]
	numOne := numbers[:3]
	printSlice(numOne) //len=3 ,cap =7, slice=[1 2 3]
	ints := numbers[:6]
	printSlice(ints) //len=6 ,cap =7, slice=[1 2 3 4 5 6]

	//3.切片的append 增加切片的内容
	numbers = append(numbers, 8)
	printSlice(numbers)              //len=8 ,cap =14, slice=[1 2 3 4 5 6 7 8] 这里我们可以注意到如果append增加的内容超过了容量，切片容量扩容到2倍,指向的
	numbers = append(numbers, 9, 10) //可以一次性增加多个
	printSlice(numbers)              //len=8 ,cap =14, slice=[1 2 3 4 5 6 7 8 9 10]

	//4.切片的赋值 copy
	var copynums []int
	copynums = numbers
	//copy numbers内容到copynums
	//copy(copynums, numbers)
	printSlice(copynums) //len=0 ,cap =0, slice=[] 因为这里的容量为0，所以复制内容无法存放
	copynums = make([]int, 8, 16)
	copy(copynums, numbers)
	printSlice(copynums) //这里的容量为16，复制内容可以存放 len=8 ,cap =16, slice=[1 2 3 4 5 6 7 8]

}
func printSlice(list []int) {
	fmt.Printf("addres=%p, len=%d ,cap =%d, slice=%v\n", list, len(list), cap(list), list)
}
