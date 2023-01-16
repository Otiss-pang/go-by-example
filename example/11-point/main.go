package main

import "fmt"

// 默认go方法传参为传值，而不是传引用
func add2(n int) {
	n += 2
}

// 使用*type声明一个指针变量，使用*对变量解引用，使用&获取一个变量的指针（引用）
func add2ptr(n *int) {
	fmt.Println(n)
	*n += 2
}

func main() {
	n := 5
	add2(n)
	fmt.Println(n) // 5
	add2ptr(&n)
	fmt.Println(n) // 7
}
