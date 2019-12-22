//1.请指出下面代码的错误？
package main

func main()  {
	var one int
	two := 2
	var three int
	three = 3

	func (unused string)  {
		fmt.Println("Unused arg. No compile error")
	}("what?")
}
//变量声明未使用

//2.下面代码输出什么？

type ConfigOne struct{
	Daemon string
}

func (c *ConfigOne) String() string {
	return fmt.Sprintf("print: %v", c)
}

func main()  {
	c := &ConfigOne{}
	c.String()
}
//如果类型实现string方法，当格式化输出时会自动使用String方法，但是本题的String方法在调用格式化输出，形成递归调用，导致栈溢出

//3.下面代码输出什么？

func main()  {
	var a = []int{1,2,3,4,5}
	var r = make([]int, 0)

	for i, v := range a{
		if i == 0 {
			a = append(a, 6, 7)
		}

		r = append(r, v)
	}
	fmt.Println(r)
}
//[1 2 3 4 5]
//a进入range时使用的是a的副本，所以尽管a已经添加到7个元素，r仍只能获取a副本的底层元素