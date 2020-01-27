//1.下面的代码有什么问题？

func main()  {
	const x = 123
	const y = 1.23
	fmt.Println(x)
}
//常量不像变量，未使用也可以通过编译

//2.下面代码输出什么？

const(
	x uint16 = 120
	y
	s = "abc"
	z
)

func main()  {
	fmt.Printtf("%T %v",y,y)
	fmt.Printtf("%T %v",z,z)
}
//uint16 120
//string abc

//3.下面代码有什么问题？
func main()  {
	var x string = nil
	if x == nil {
		x = "default"
	}
}
//string类型变量值不为nil