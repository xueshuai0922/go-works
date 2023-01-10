package main

import "fmt"

// Person /**
/**
1.结构体
2.结构体方法 func(t T结构体 ) method(入参) 返回参数{}
3.方法值  结构体的变量.方法 ===》接收方法 t.method 返回的是方法
4.方法表达式 T.method 返回的是表达式
*/
/*
定义一个结构体,类似”对象“
*/
type Person struct {
	name string
}

// TMethod1 方法
//结构体的方法==》对象的方法，特点/**
/*
*

 */
func (person Person) TMethod1() {
	person.name = " struct name1"
	fmt.Println("this is a struct method1")

}
func (person *Person) TMethod2() {
	person.name = " struct name2"
	fmt.Println("this is a struct method2")
}

/*
*
普通参数方法
*/
func method1(person Person) {
	person.name = "new name1"
	fmt.Println("method1 ", person.name)
}

func method2(person *Person) {
	person.name = "new name2"
	fmt.Println("method2 ", person.name)

}

/*
*
方法值
*/
func methodValue(person Person) {
	tMethod2 := person.TMethod2
	tMethod2()
}

/**
方法表达式
*/

func methodTemplate() func(Person) {
	return Person.TMethod1
}
func main() {
	person := Person{
		name: "old name",
	}
	fmt.Println("method1 调用前 ", person.name)

	method1(person)
	fmt.Println("method1 调用后 ", person.name)

	fmt.Println("method2 调用前 ", person.name)
	method2(&person)
	fmt.Println("method2 调用后 ", person.name)

	fmt.Println("TMethod1 调用前 ", person.name)
	person.TMethod1()
	fmt.Println("TMethod1 调用后 ", person.name)

	fmt.Println("TMethod2 调用前 ", person.name)
	person.TMethod2()
	fmt.Println("TMethod2 调用后 ", person.name) // TMethod2 调用后   struct name2

	//2.方法值 测试
	methodValue(person)

	//3.todo 理解 方法表达式 测试
	template := methodTemplate()
	template(person)
}
