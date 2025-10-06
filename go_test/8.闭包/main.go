package main
import "fmt"

//闭包机制


//返回类型：func() int - 返回一个无参数、返回值为int的函数
func a() func() int {
	i := 0   //声明一个整型变量
	//b是一个匿名函数(闭包)
	// 访问并修改外层函数的变量 i
	// 每次调用时让 i 自增1
	// 打印当前的 i 值
	// 返回 i 的值
	b := func() int {
		i++
		fmt.Println(i)
		return i
	}
	return b
}

func main() {
	c := a()   //调用a()，获取返回的闭包函数
	c() // 第1次调用闭包，输出: 1
	c() // 第2次调用闭包，输出: 2
	c() // 第3次调用闭包，输出: 3

	a() //  只调用a()但不保存返回值，无输出
}
