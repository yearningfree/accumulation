//1.通过指针变量 p 访问其成员变量 name，有哪几种方式？

A.p.name
B.(&p).name
C.(*p).name
D.p->name
//AC
//BD错，B为取地址，D无此方法

//2.下面这段代码能否通过编译？如果通过，输出什么？

package main

import "fmt"

type MyInt1 int
type MyInt2 = int

func main() {
	var i int = 0
	var i1 MyInt1 = i
	var i2 MyInt2 = i
	fmt.Println(i1, i2)
}
//不能，第21行无法将int类型的变量赋值给MyInt1类型的变量
//强制转换：var i1 MyInt1 = MyInt1(i)

//16行，基于int类型创建新类型MyInt1
//17行，创建int类型别名MyInt2
