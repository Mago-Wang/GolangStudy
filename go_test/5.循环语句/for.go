package main

import (
	"fmt"
)

/**
 * for 循环三种形式
 * for init;condition;post {}  类似for
 * init 一般为赋值表达式，给控制变量赋初值
 * condition 一般为关系或逻辑表达式，控制循环的结束条件
 * post 一般为赋值表达式，控制变量的更新
 * for condition {}  类似while
 * for {}
 */

func main() {
	numbers := [6]int{1, 2, 3, 5}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	i := 0
	for i < len(numbers) {
		fmt.Println(numbers[i])
		i++
	}

	for i, x := range numbers {
		fmt.Printf("index = %d,value = %d\n", i, x)
	}

	//无限循环
	// for {
	// 	fmt.Println("endless...")
	// }

	//for循环的range格式可以对数组、切片、map、字符串等进行迭代循环
	numbers1 := []int{1, 2, 3, 4, 5, 6}

	//忽略value，只取index，支持string/array/slice/map
	for i := range numbers1 {
		fmt.Println(i) // 0 1 2 3 4 5
	}

	//忽略index
	for _, value := range numbers1 {
		fmt.Println(value)
	}
	// 忽略全部返回值，仅迭代
	// for range numbers {
	// }

	//map无序的键值对
	m := map[string]int{"a": 1, "b": 2}
	for k, v := range m {
		fmt.Println(k, v)
	}

	//注意:range会复制对象
	a := [3]int{0, 1, 2}
	fmt.Printf("a的类型是%T\n", a)
	for i, v := range a { // i,v都是从复制品中取出的
		if i == 0 {
			a[1], a[2] = 999, 999
			fmt.Println(a)
		}
		a[i] = v + 100 // v的值仍然来自于复制品，所以不受影响
	}
	fmt.Println(a)
}
