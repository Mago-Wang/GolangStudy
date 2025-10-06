package main

import "fmt"

/**
 * 调用函数，可以通过两种方式来传递参数：
 * 值传递,在调用过程中不影响实际参数
 * 引用传递，在调用过程中将实际参数的地址传递到函数中
 */

//值传递
//函数收到的是a的副本
//在函数内修改p =10只影响副本，不影响原变量a
func changevalue1(p int) {
	p = 10
}

//引用传递(地址和内存)
// 指针传递共享同一内存地址
func changevalue2(p *int) {
	*p = 10
}

//在函数中交换两数的值
func swap1(a, b int) {
	var temp int
	temp = a
	a = b
	b = temp
}

func swap2(a *int, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
}

func main() {
	var a int = 1
	changevalue1(a)
	fmt.Println(a)
	changevalue2(&a)
	fmt.Println(a)
	b, c := 1, 2
	fmt.Println("交换前:", b, c)
	swap2(&b, &c) //传地址
	fmt.Println("交换后:", b, c)

	var p *int
	p = &a
	fmt.Println(&a) //打印地址
	fmt.Println(p)  //打印地址

	var pp **int //二级指针

	pp = &p
	fmt.Println(&p)
	fmt.Println(pp)
}
