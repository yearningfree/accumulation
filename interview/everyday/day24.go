package main

import "fmt"

//1.下面这段代码输出什么？

func main() {
	m := map[int]string{0: "zero", 1: "one"}
	for k, v := range m {
		fmt.Println(k, v)
	}
}

//0 zero
//1 one
//或
//1 one
//0 zero
//map是无序的

//2.下面这段代码输出什么？

func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

//10 1 2 3
//20 0 2 2
//2 0 2 2
//1 1 3 4
