//1.关于 channel 下面描述正确的是？

A. close() 可以用于只接收通道；

B. 单向通道可以转换为双向通道；

C. 不能在单向通道上做逆向操作（例如：只发送通道用于接收）；
//C

//2.下面的代码有什么问题？
type T struct {
    n int
}

func getT() T {
    return T{}
}

func main() {
    getT().n = 1
}
//返回T{}无法寻址，编译错误
//修复
func main()  {
	t := getT()
	p := &t.n // <=> p = &(t.n)
	*p = 1
	fmt.Println(t.n)
}
