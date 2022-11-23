package main

import "fmt"

/*
*
常量：不会被修改的量

	只能是boolean ，数值、字符串不能是切片、map
*/
func main() {
	//显示展示
	const name string = "xueshuai"
	//隐示展示
	const age = 12

	const flag = true

	const (
		computerName = "unicom-xs"
		cpus         = 12
	)
	fmt.Println(name, age, flag, computerName, len(computerName), cpus)
	//【重点】------------iota 自增量------------------
	const (
		one = iota
		two
		three
	)
	fmt.Println(one, two, three)

	type Allergen int

	const (
		IgEggs         Allergen = 1 << iota // 1 << 0 which is 00000001
		IgChocolate                         // 1 << 1 which is 00000010
		IgNuts                              // 1 << 2 which is 00000100
		IgStrawberries                      // 1 << 3 which is 00001000
		IgShellfish                         // 1 << 4 which is 00010000
	)
	fmt.Println(IgEggs, IgChocolate, IgNuts, IgStrawberries, IgShellfish)

	fmt.Println(IgEggs | IgChocolate | IgNuts) //相加？

	//【注意】 iota中自增加，换行了不是立即就才会增加
	const (
		s, ss = iota + 1, iota + 2
		sss, ssss
	)
	fmt.Println(s, ss, sss, ssss) //1 2 2 3

}
