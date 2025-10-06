package main

import (
	"fmt"
)

func main() {
	//声明slice1是一个切片，并且初始化，默认值是1，2，3。 长度len是3
	//slices1 := []int{1,2,3}

	//声明slice1是一个切片，但是并没有给slice分配空间
	//var slice1 []int   // slice1 == nil

	//声明slice1是一个切片，同时给slice分配空间，3个空间，初始化值是0
	//slice1 = make([]int,3)  //[0 0 0]

	//声明slice1是一个切片，同时给slice分配空间，3个空间，初始化值是0, 通过:=推导出slice是一个切片
	slice1 := make([]int, 3) //[0 0 0]

	// len() 和 cap() 函数
	// len：长度，表示左指针⾄右指针之间的距离
	// cap：容量，表示指针至底层数组末尾的距离
	fmt.Printf("len:%d,cap:%d\n", len(slice1), cap(slice1))

	//判断一个slice是否为0
	if slice1 == nil {
		fmt.Println("slice1是一个空切片")
	} else {
		fmt.Println("slice1是一个非空切片")
	}
}
