package main

/*
import _ "fmt"
给 fmt 包一个匿名,⽆法使用该包的⽅法,但是会执行该包内部的 init() 方法

import aa "fmt"
给 fmt包起一个别名aa,可以用别名直接调用：aa.Println()

import . "fmt"
将 fmt 包中的全部方法，导入到当前包的作用域中,全部方法可以直接调用,无需fmt.API的形式
*/

import (
	"GolangStudy/5-init/lib1"
	// _ "GolangStudy/5-init/lib1"
	mylib2 "GolangStudy/5-init/lib2"
	//. "GolangStudy/5-init/lib2"
	"fmt"
)

//每个go程序都会一开始执行init()函数,可以用来做一些初始化操作

func init() {
	fmt.Println("init()函数执行了")
}
func main() {
	fmt.Println("hello world")
	lib1.Lib1Test()
	//lib2.Lib2Test()
	mylib2.Lib2Test()
	//Lib2Test()
}
