package main

import "fmt"

//1.下面代码段输出什么？

type Person struct {
	age int
}

func main() {
	person := &Person{28}

	// 1.
	defer fmt.Println(person.age)

	// 2.
	defer func(p *Person) {
		fmt.Println(p.age)
	}(person)

	// 3.
	defer func() {
		fmt.Println(person.age)
	}()

	person.age = 29
}

//29 29 28
//1.将28作为defer函数参数缓存入栈中
//2.传入缓存的是person地址，在最后person的值被修改，根据地址取得29
//3.闭包引用
