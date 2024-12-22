package main

import "fmt"

func main() {
	//  Map
	//  Map 是 Go 语言中的键值对集合
	//  Map 是引用类型
	//  Map 是无序的
	//  Map 的键不能重复

	//  定义一个 Map
	var map1 = make(map[string]int)
	map1["a"] = 1
	map1["b"] = 2
	fmt.Println("map1 = ", map1) //  map1 =  map[a:1 b:2]

	//  获取元素
	fmt.Println("map1[\"a\"] = ", map1["a"]) //  map1["a"] =  1

	//  修改元素
	map1["a"] = 3
	fmt.Println("map1 = ", map1) //  map1 =  map[a:3 b:2]

	//  删除元素
	delete(map1, "a")
	fmt.Println("map1 = ", map1) //  map1 =  map[b:2]

	//  遍历 Map
	for k, v := range map1 {
		fmt.Println("map1[", k, "] = ", v)
	}

	//  空 Map
	var map2 = make(map[string]int)
	fmt.Println("map2 = ", map2) //  map2 =  map[]

	//  获取长度
	fmt.Println("len(map2) = ", len(map2)) //  len(map2) =  0
}
