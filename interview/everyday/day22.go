//1.下面的代码有几处语法问题，各是什么？

package main

import (
	"fmt"
)

func main() {
	var x string = nil
	if x == nil {
		x = "default"
	}
	fmt.Println(x)
}

//第8行，只有引用类型才能赋值nil
//第9行，字符串不能跟nil比较

//2.return 之后的 defer 语句会执行吗，下面这段代码输出什么？

var a bool = true

func main() {
	defer func() {
		fmt.Println("1")
	}()
	if a == true {
		fmt.Println("2")
		return
	}
	defer func() {
		fmt.Println("3")
	}()
}

//输出：2
//     1
//retuen之后的defer不执行
