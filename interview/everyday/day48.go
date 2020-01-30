//1.下面代码有什么问题？
type foo struct{
	bar int
}

func main()  {
	var f foo
	f.bar, tmp := 1, 2
}
//:=操作符不能用于结构体字段赋值


//2.下面的代码输出什么？
func main()  {
	fmt.Println(~2)
}
//编译错误，go的取反运算符是^