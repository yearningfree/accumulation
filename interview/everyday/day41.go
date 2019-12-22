//1.下面代码编译能通过吗？

func main()
{
	fmt.Println("hello world")
}
//不能，大括号不能单独放一行

//2.下面这段代码输出什么？

var x = []int{2:2,3,0:1}

func main()  {
	fmt.Println(x)
}
//[1 0 2 3]
//通过字面量初始化时可以指定索引，没有指定索引的元素会在前一个指定索引基础之上加一


//3.下面这段代码输出什么？

func incr(p *int) int {
	*p++
	return *p
}
func main()  {
	v := 1
	incr(&v)
	fmt.Println(v)
}
//2