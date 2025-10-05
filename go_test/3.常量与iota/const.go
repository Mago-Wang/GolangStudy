package main

import (
	"fmt"
	"unsafe"
)

// const 来定义枚举类型
const (
	//可以在const() 添加一个关键字itoa,每行的itoa都会累加1，第一行的itoa的默认值是0
	BEIJING = 10*iota  //10*0
	SHANGHAI  // 10*1
	SHENZHEN  // 10*2
)

const (
	a,b = iota+1,iota+2  // iota = 0,a = 1,b =2
	c,d // iota = 1,c = 2,d = 3
	e,f // iota = 2,e = 3,f = 4

	g,h = iota*2,iota*3 // iota = 3,g = 6,h = 9
	i,k //iota = 4,i =8,k = 12
)

/*
常量可以用len(), cap(), unsafe.Sizeof()常量计算表达式的值。
常量表达式中，函数必须是内置函数，否则编译不过
cap() 可以测量切片最长可以达到多少
*/
const (
	aa = "abc"
	bb = len(aa)
	cc = unsafe.Sizeof(aa)
	/*
	*在 Go 语言中，字符串的内部结构是这样的：
		一个指向字符串数据的指针（8字节，在 64 位系统上）
		一个表示字符串长度的整型值（8字节，在 64 位系统上）
	* unsafe.Sizeof(aa) 返回的是字符串类型本身的大小（即字符串结构体的大小），而不是字符串内容的长度
	* 字符串类型的内存大小：8字节指针 + 8字节长度 = 16字节
	*/
)

func main() {

	//常量(只读属性)
	//显示类型定义
	const length int = 10;
	//隐式类型定义
	const width = 5;
    //length = 100; 常量是不允许修改的
	fmt.Println("length = ", length)
	fmt.Println("width = ", width)

	fmt.Println("BEIJIGN = ", BEIJING)
	fmt.Println("SHANGHAI = ", SHANGHAI)
	fmt.Println("SHENZHEN = ", SHENZHEN)

	fmt.Println("a = ", a, "b = ", b)
	fmt.Println("c = ", c, "d = ", d)
	fmt.Println("e = ", e, "f = ", f)

	fmt.Println("g = ", g, "h = ", h)
	fmt.Println("i = ", i, "k = ", k)

	// iota 只能够配合const()一起使用，iota只有在const进行累加效果
	//var a int = iota

	println(aa,bb,cc)

}