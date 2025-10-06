package main

import "fmt"

//slice、map、channel都是引用类型，声明后还需要初始化分配内存，即make
//map中的键对值是无序的
func main() {

	//===> 第一种声明方式
	// 声明myMap1是一种map类型 key是string，value是string
	var myMap1 map[string]string
	fmt.Println(myMap1 == nil) // true

	// 使用map前，需要先用make给map分配数据空间
	myMap1 = make(map[string]string, 10)
	myMap1["one"] = "java"
	myMap1["two"] = "python"
	myMap1["three"] = "c++"
	fmt.Println(myMap1) //map[one:java three:c++ two:python]

	//===> 第二种声明方式
	myMap2 := make(map[int]string)
	myMap2[1] = "java"
	myMap2[2] = "python"
	myMap2[3] = "c++"
	fmt.Println(myMap2) //map[1:java 2:python 3:c++]

	//===> 第三种声明方式
	myMap3 := map[string]string{
		"one":   "java",
		"two":   "python",
		"three": "c++",
		//注意最后一个也要加入，
	}
	fmt.Println(myMap3) //map[one:java three:c++ two:python]
}
