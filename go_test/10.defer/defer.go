package main

import "fmt"

/**
* defer语句被用于预定对一个函数的调用
* 可以把这类被defer语句调用的函数称为延迟函数。
* defer作用
	* 释放占用的资源
	* 捕捉处理异常
	* 输出日志
*/

//关于 defer 和 return 谁先谁后
//return之后的语句先执⾏，defer后的语句后执⾏
func deferFunc() int {
	fmt.Println("defer func called...")
	return 0
}

func returnFunc() int {
	fmt.Println("return func called...")
	return 0
}

func returnAndDefer() int {
	defer deferFunc()
	return returnFunc()
}


func main() {
	//写入defer关键字
	//defer 声明的语句会在当前函数执行完之后调用
	//如果一个函数中有多个defer语句，它们会以LIFO（后进先出）的顺序执行
	defer fmt.Println("main end1")
	defer fmt.Println("main end2")

	fmt.Println("main::hello go 1")
	fmt.Println("main::hello go 2")
	returnAndDefer()
}