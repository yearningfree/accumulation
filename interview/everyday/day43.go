//1.下面的代码有什么问题？

import(
	"fmt"
	"log"
	"time"
)

func main()  {
	
}
//导入的包没有被使用

//2.下面代码输出什么？
func main()  {
	x := interface{}(nil)
	y := (*int)(nil)
	a := y == x
	b := y == nil
	_, c := x.(interface{})
	println(a, b, c)
}
A. true true false
B. false true true
C. true true true
D. false true false
//D

//3.下面代码有几处错误的地方？请说明原因。
func main()  {
	var s []int
	s = append(s, 1)

	var m map[string]int
	m["one"] = 1
}
//有 1 处错误，不能对 nil 的 map 直接赋值