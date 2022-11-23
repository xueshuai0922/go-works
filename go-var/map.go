package main

import "fmt"

/*
*
map
*/
func main() {
	//1.map 第一种声明方式
	var content map[string]string
	fmt.Println(content) //map[]
	//1.1如果给空的map增加数据，报错
	addNilMapError(content)
	//1.2定义完后必须进行开创空间
	content = make(map[string]string)
	content["name"] = "xueshuai"
	content["name"] = "xueshuai"
	fmt.Println(content, len(content)) //map[name:xueshuai] 1
	content["age"] = "12"
	fmt.Println(content, len(content)) //map[age:12 name:xueshuai] 2

	//2.第二种声明方式，map中嵌套map
	mapmap := make(map[string]map[int]string)
	mapmap["ssr"] = make(map[int]string, 3)
	mapmap["ssr"][1] = "第一个"
	mapmap["ssr"][2] = "第2个"
	mapmap["ssd"] = make(map[int]string, 3)
	mapmap["ssd"][1] = "第一个"
	mapmap["ssd"][2] = "第2个"
	fmt.Println(mapmap) //map[ssd:map[1:第一个 2:第2个] ssr:map[1:第一个 2:第2个]]

}
func addNilMapError(content map[string]string) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("错误信息:%s\n", err) //assignment to entry in nil map
		}
	}()
	content["name"] = "xueshuai" //panic: assignment to entry in nil map 这里报错，只是给了变量，默认的空间为0，空的map无法进行存放key value
	fmt.Println(content)
}
