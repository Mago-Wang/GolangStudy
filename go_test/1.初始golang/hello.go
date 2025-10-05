package main //程序的包名，表示一个可独立执行的程序，每个Go 程序必须有一个main包

//import "fmt"  实现格式化IO的函数

//导入多个包方式一
/* import "fmt"
import "time" */

//导入多个包方式二
import (
	"fmt"
	"time"
)

// main函数
func main() { // 函数的{必须和函数名同行，否则编译报错
	/* 简单的程序 万能的hello world */
	fmt.Println("hello go")
	time.Sleep(1 * time.Second) //程序暂停一秒
}
