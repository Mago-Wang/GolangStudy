package main
/**
 * 切片的三要素：指针、长度、容量
 * 动态扩容机制：切片如何在需要时自动增长
 *性能考虑：预设合适的容量可以避免频繁的内存分配
 */
import "fmt"

func main() {
	//参数含义：make([]int, length, capacity)
	var numbers = make([]int, 3, 5)    //第一种初始化方式
    //%v：打印切片的值
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)
    

	//场景1：容量未满时追加元素
	//向numbers切片追加一个元素1, numbers len = 4， [0,0,0,1], cap = 5
	numbers = append(numbers, 1)
    fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)
	//向numbers切片追加一个元素2, numbers len = 5， [0,0,0,1,2], cap = 5
	numbers = append(numbers, 2)
    fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	//场景2：容量已满时追加元素
	//向一个容量cap已经满的slice 追加元素，触发扩容机制
	//如果长度增加后超过容量，则将容量增加 2 倍
	numbers = append(numbers, 3)
    fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	fmt.Println("-=-------")
	var numbers2 = make([]int, 3)  //第二种初始化方式，只指定长度，容量默认等于长度
	// 3 3 [0 0 0]
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers2), cap(numbers2), numbers2)
	numbers2 = append(numbers2, 1)
	// 4 6 [0 0 0 1]
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers2), cap(numbers2), numbers2)
}
