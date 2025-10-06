package main
import "fmt"


// ... 是 Go 的一种语法糖
// 用法 1：函数可以用来接受多个不确定数量的参数。
// 用法 2：slice 可以被打散进行传递

/*
* args ...string 表示函数可以接受任意数量的string类型参数
* 在函数内部，args 被当作 []string 切片来处理
* 使用 range 遍历所有传入的参数并打印
*/
func test(args ...string) {
	for _, v := range args {
		fmt.Println(v)
	}
}

func main() {
	var ss = []string{
		"abc",
		"efg",
		"hij",
		"123",
	}
	test(ss...)
}
