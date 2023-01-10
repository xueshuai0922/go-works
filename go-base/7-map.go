package main

import (
	"fmt"
	"sort"
)

/*
*
map kev:value 无序
*/
func main() {
	//1.创建map 第一种声明方式 默认是只是创建了，没有初始化 nil
	var content map[string]string
	fmt.Println(content) //map[]
	//1.1如果给空的map增加数据，报错
	addNilMapError(content)
	//1.2 make这种方式是初始化了的，完后必须进行开创空间
	content = make(map[string]string)
	//1.3 修改更新map map[xxx]=value  如果xxx存在则更新，如果xxx不存在就增加
	content["name"] = "xueshuai"
	content["name"] = "xueshuai2"
	fmt.Println(content, len(content)) //map[name:xueshuai] 1
	content["age"] = "12"
	fmt.Println(content, len(content)) //map[age:12 name:xueshuai] 2

	//2.创建map第二种声明方式，map中嵌套map
	mapmap := make(map[string]map[int]string)
	mapmap["ssr"] = make(map[int]string, 3)
	mapmap["ssr"][1] = "第一个"
	mapmap["ssr"][2] = "第2个"
	mapmap["ssd"] = make(map[int]string, 3)
	mapmap["ssd"][1] = "第一个"
	mapmap["ssd"][2] = "第2个"
	fmt.Println(mapmap) //map[ssd:map[1:第一个 2:第2个] ssr:map[1:第一个 2:第2个]]

	//3.删除
	delete(content, "name")
	fmt.Println(content) //map[age:12]

	//4.获取map中的key的值
	value, ok := content["age"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("没有获取到值")
	}

	//5.map的遍历
	content["sex"] = "男"
	content["height"] = "188"
	for key, value := range content {
		fmt.Println(key, "---->", value)
	}
	keySlice := make([]string, 0)
	for keys := range content {
		keySlice = append(keySlice, keys)
	}
	//sort内置函数进行排序
	sort.Strings(keySlice)
	fmt.Println(keySlice)
	for i := 0; i < len(keySlice); i++ {
		fmt.Println(keySlice[i], "--->", content[keySlice[i]])
	}

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
