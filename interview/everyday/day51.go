//1.下面的代码能否正确输出？
func main()  {
	var fn1 = func(){}
	var fn2 = func(){}

	if fn1 != fn2 {
		println("fn1 not equal fn2")
	}
}
//不能 编译错误，函数只能与nil比较

//2.下面代码输出什么？
type T struct{
	n int
}

func main()  {
	m := make(map[int]T)
	m[0].n = 1
	fmt.Println(m[0].n)
}
//compilation error，map中的struct是不可寻址的
//修复代码
func main()  {
	m := make(map[int]T)
	t := T{1}
	m[0]=t
	fmt.Println(m[0].n)
}