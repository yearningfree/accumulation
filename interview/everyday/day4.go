package main

import "fmt"

//1.下面这段代码能否通过编译，不能的话原因是什么；如果能，输出什么。

func main() {
	list := new([]int)
	list = append(list, 1)
	fmt.Println(list)
}

//不能，new(T)返回的是指针list，指针不能适用append

//2.下面这段代码能否通过编译，如果可以，输出什么？

func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1, s2)
	fmt.Println(s1)
}
//不能，如果append的第二个参数是slice，需写成append(s1, s2...)

//3.下面这段代码能否通过编译，如果可以，输出什么？

var(
	size := 1024
	max_size = size * 2
)

func main(){
	fmt.Println(size, max_size)
}
//不能，max_size语法错误，正确写法max_size := size * 2
