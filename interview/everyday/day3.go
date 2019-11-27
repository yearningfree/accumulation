//1.下面两段代码输出什么。
package main

import "fmt"

//1.
func main() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}

//2.
func main() {
	s := make([]int, 0)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}

//输出
//1. [0 0 0 0 0 1 2 3]
//2. [1 2 3]

//2.下面这段代码有什么缺陷

func funcMui(x,y int)(sum int,error){
	return x+y,nil
}

//第二个返回值未命名

//正确写法
func funcMui(x,y int)(sum int,err error){
	sum = x+y
	err = nil
	return
}

//3.new() 与 make() 的区别

//两个函数都是用来分配内存，但适用类型不同
//new(T)会为T类型的值分配已置零的内存空间，并返回地址（指针）
//适用值类型

//make(T,args)返回初始化之后的T类型的值，这个值并不是T类型的零值，
//也不是指针*T，是经过初始化的T的引用。
//只适合slice、map、channel