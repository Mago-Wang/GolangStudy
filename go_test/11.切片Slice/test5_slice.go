package main

import "fmt"

//slice截取是浅拷贝，若想深拷贝需要使用copy
//可以通过设置下限以及上限设置截取切片 [lower-bound: upper-bound]
func main() {
	s := []int{1, 2, 3} //len = 3, cap = 3, [1,2,3]
    
	/*打印原始切片*/
	fmt.Println(s)

	//[0, 2)
	/* 打印子切片从索引0(包含)到索引2(不包含) */
	//浅拷贝，新切片 s1 和原切片 s 共享同一个底层数组
	//修改其中一个切片的元素会影响到另一个切片
	s1 := s[0:2] // [1, 2]
	fmt.Println(s1)
	
	/* 默认下限为 0 */
	fmt.Println(s[:1])  // [1]

	/* 默认上限为 len(s) */
	fmt.Println(s[0:]) // [1, 2, 3]
	
	s1[0] = 100

	fmt.Println(s)
	fmt.Println(s1)

	//copy可以将底层数组的slice一起进行拷贝
	//copy是深拷贝
	// 新切片 s2 有自己独立的底层数组
	// 修改其中一个切片不会影响另一个切片
	s2 := make([]int, 3) //s2 = [0,0,0]

	//将s中的值 依次拷贝到s2中
	copy(s2, s)
	fmt.Println(s2)

}