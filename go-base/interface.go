package main

import "fmt"

/*
*

	interface 接口类似Java中的object一样，所有的类似都实现了该接口，作为参数传入后，要进行断言
*/

func main() {
	var a interface{} = "nihao"
	funcName(a)

}

func funcName(a interface{}) {

	s, ok := a.(string)
	if !ok {
		fmt.Println("this is not string type")
	} else {
		fmt.Printf("this is a %T ,value %s", s, s)
	}
}
func wrongMethod(a interface{}) {
	// fmt.Println(string(a)) //Cannot convert an expression of the type 'interface{}' to the type 'string'

}
