package main

import "fmt"

/*
基于数组进行切片

	1.数组发生了变化，那么基于数组的变化的切片也会发生变化
	2，切片的深拷贝
*/
func main() {
	a := [4]int{1, 2, 4, 5}
	fmt.Println(a)
	fmt.Printf("address=%p,type=%T\n", &a, a)
	s := a[1:3]
	a[2] = 111111
	fmt.Println(s)
	fmt.Printf("address=%p,type=%T\n", s, s)
	fmt.Printf("address=%p,type=%T\n", &s, &s)

	//切片的copy是深拷贝
	s2 := make([]int, 10, 10)
	resultSlice := make([]int, 20, 30)
	ints := append(s2, 333, 2222, 444)
	copy(resultSlice, ints)
	fmt.Println(resultSlice)
}
