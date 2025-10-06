package main
/**
Go 错误处理的两种模式
	Error 模式：函数返回 error 类型，调用者检查错误
	Panic/Recover 模式：用于处理不可恢复的异常
自定义错误类型
	实现 Error() string 方法即可成为 error 类型
	可以包含更丰富的错误上下文信息
类型断言
	使用 switch v := err.(type) 进行类型判断
	可以根据不同错误类型进行不同处理

*/

import (
	"fmt"
	"os"
	"time"
)

// recover() 函数：
// 	只能在 defer 函数中使用
// 	用于捕获和处理 panic 异常
// 	如果没有异常，返回 nil
// panic 异常：
// 	运行时错误会自动触发 panic
// 	数组/切片越界访问是典型的 panic 场景

// 自定义异常
//自定义错误类型结构体
//包含路径、操作类型、创建时间和错误消息
type PathError struct {
	path       string
	op         string
	createTime string
	message    string
}

// 实现 error 接口，返回格式化的错误信息
func (p *PathError) Error() string {
	return fmt.Sprintf("path=%s \nop=%s \ncreateTime=%s \nmessage=%s",
		p.path, p.op, p.createTime, p.message)
}

//打开文件，成功则确保文件关闭后返回nil
//如果文件打开失败，返回自定义PathError
func Open(filename string) error {

	file, err := os.Open(filename)
	if err != nil {
		return &PathError{
			path:       filename,
			op:         "read",
			message:    err.Error(),
			createTime: fmt.Sprintf("%v", time.Now()),
		}
	}

	defer file.Close()
	return nil
}

func main() {

	//捕获系统抛出异常
	// defer func() {
	// 	//recover() 函数用于捕获 panic 异常
	// 	//如果捕获到异常（err != nil），则打印异常信息
	// 	if err := recover(); err != nil {
	// 		fmt.Println("捕获：", err)
	// 	}
	// }()

	//nums := []int{1, 2, 3}
	//fmt.Println(nums[4]) // 系统抛出异常
	// 捕获： runtime error: index out of range [4] with length 3

	//手动抛出异常并捕获
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println("捕获：", err)
	// 	}
	// }()
	// panic("出现异常！") // 手动抛出异常
	// 捕获： 出现异常！

	//自定义异常
	err := Open("test.txt")
	switch v := err.(type) {
	case *PathError:
		fmt.Println("get path error,", v)
	default:
	}
}
