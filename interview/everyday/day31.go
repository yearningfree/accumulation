package main

import "fmt"

//1.下面这段代码输出什么？

func change(s ...int) {
	s = append(s, 3)
}

func main() {
	slice := make([]int, 5, 5)
	slice[0] = 1
	slice[1] = 2
	change(slice...)
	fmt.Println(slice)
	change(slice[0:2]...)
	fmt.Println(slice)
}

//[1 2 0 0 0]
//[1 2 3 0 0]

//2.下面这段代码输出什么？

func main() {
	var a = []int{1, 2, 3, 4, 5}
	var r [5]int

	for i, v := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
}

//[1 12 13 4 5]
//[1 12 13 4 5]
//此处a是切片
