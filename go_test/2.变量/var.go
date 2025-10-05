package main

import "fmt"

//声明全局变量 方法一、二、三是可以的
var gA int = 100
var gB = 200

/*
* 用方法四不能用来声明全局变量
* :=只能够用在函数内部，不能在函数外使用
*/
//gC := 300

var ( //这种分解的写法,一般用于声明全局变量
        j int
        k bool
)

func main() {
	/* 四种变量的声明方式 */
	// 方法一：指定变量类型，声明后若不赋值，使用默认值0
	var a int
	fmt.Printf("a = %d\n", a)
	//判断变量类型使用%T
	fmt.Printf("type of a = %T\n", a)

	// 方法二：指定变量类型，声明后赋值
	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("type of b = %T\n", b)

	var bb string = "abcd"
	fmt.Printf("bb = %s, type of bb = %T\n", bb, bb)
	// 方法三：省略变量类型，根据值自行判断变量类型
	var c = 1000
	fmt.Println("c = ", c)
	fmt.Printf("type of c = %T\n", c)
	var cc = "abcd"
	fmt.Printf("cc = %s, type of cc = %T\n", cc, cc)

	// 方法四：(常用)省去var关键字，直接自动匹配,如果变量已经使用 var 声明过了，再使用 := 声明变量，就产生编译错误
	d := 3.14
	fmt.Println("d = ", d)
	fmt.Printf("type of d = %T\n", d)
	dd := "abcd"
	fmt.Printf("dd = %s,type of dd = %T\n",dd,dd)

	// =====
	fmt.Println("gA = ", gA, ", gB = ", gB)
	//fmt.Println("gC = ", gC)

	//声明多个变量
	var xx,yy int = 100,200
	fmt.Println("xx = ", xx, ", yy = ", yy)
	var kk,ll = 100,"abcd"
	fmt.Println("kk = ",kk,", ll = ",ll)
	
	//多行的变量声明
	var(
		vv  = 1000
		jj  = true
	)
	fmt.Println("vv = ", vv, ", jj = ", jj)
}
