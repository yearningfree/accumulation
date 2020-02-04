// 1.下面的代码有什么问题？
type X struct{}

func (x *X) test (){
	println(x)
}

func func main() {
	var a *X
	a.test()
	X{}.test()
}
//X{} 是不可寻址的，不能直接调用方法
//修复
var x = X{}
x.test()

// 2.下面代码有什么不规范的地方吗？
func main()  {
	x := map[string]string{"one":"a","two":"","three":"c"}
	if v := x["two"]; v == "" {
		fmt.Println("no entry")
	}
}
//检查 map 是否含有某一元素，直接判断元素的值并不是一种合适的方式。最可靠的操作是使用访问 map 时返回的第二个值。
//修复
if _,ok := x["two"];!ok {
	fmt.Println("no entry")
}