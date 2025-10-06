package main

import "fmt"

func printArray(myArray []int) {
	fmt.Println("myArray = ", myArray)
	myArray[0] = 10   //动态数组是引用传递
}

func main() {
	//声明一个未指定大小的数组来定义切片
	myArray := []int{1,2,3,4} //动态数组，切片slice
	fmt.Printf("myArray type is %T\n", myArray)
    printArray(myArray)
	fmt.Println(" ==== ")
	fmt.Println(myArray)  //[10 2 3 4]
}