package lib1

import "fmt"


//当前lib1包提供的API
//首字母必须大写，才能导入，表示对外开放
func Lib1Test() {
	fmt.Println("Lib1Test()")
}

func init() {
	fmt.Println("lib1.init()....")
}