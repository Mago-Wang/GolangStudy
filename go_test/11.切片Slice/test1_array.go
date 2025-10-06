/**
* Go语言切片(动态数组)是对数组的抽象
* Golang 默认都是采用值传递，有些值天生就是指针：slice、map、channel。
* 注意：定长数组是值传递，slice 是指针传递
 */

package main

import "fmt"

func printArray(myArray [4]int) {
	//值拷贝
	for index, value := range myArray {
		fmt.Println("index = ", index, ", value = ", value)
	}
	myArray[0] = 111
}

func main() {
	//固定数组的声明方式
	//Array := [...]int{1, 2, 3, 4} 是自动计算数组长度，但并不是引用传递
	var array1 [10]int
	array2 := [10]int{1, 2, 3, 4}
	array3 := [4]int{11, 22, 33, 44}

	for i := 0; i < len(array1); i++ {
		fmt.Println(array1[i])
	}
	for index, value := range array2 {
		fmt.Println("index = ", index, ", value = ", value)
	}

	//查看数组的数据类型
	fmt.Printf("array1 types = %T\n", array1)
	fmt.Printf("array2 types = %T\n", array2)
	fmt.Printf("array3 types = %T\n", array3)
	printArray(array3)
	fmt.Println(" ------ ")
	for index, value := range array3 {
		fmt.Println("index = ", index, ", value = ", value)
	}
}
