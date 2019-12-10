package main

import (
	"fmt"
	"time"
)

//1.下面这段代码能否正常结束？

func main() {
	v := []int{1, 2, 3}
	for i := range v {
		v = append(v, i)
	}
}

//能，切片长度在循环前已经确定，循环内改变切片长度不影响循环次数

//2.下面这段代码输出什么？为什么？

func main() {

	var m = [...]int{1, 2, 3}

	for i, v := range m {
		go func() {
			fmt.Println(i, v)
		}()
	}

	time.Sleep(time.Second * 3)
}

//2 3
//2 3
//2 3
//一个闭包引用，i,v都是循环结束的值
//正确写法
//1.使用函数传递
for i, v := range m {
	go func(i,v int) {
		fmt.Println(i, v)
	}(i,v)
}
//2.使用临时变量保留当前值
for i, v := range m {
    i := i           // 这里的 := 会重新声明变量，而不是重用
    v := v
    go func() {
        fmt.Println(i, v)
    }()
}