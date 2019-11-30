//1.关于init函数，下面说法正确的是()

A. 一个包中，可以包含多个 init 函数；
B. 程序编译时，先执行依赖包的 init 函数，再执行 main 包内的 init 函数；
C. main 包中，不能有 init 函数；
D. init 函数可以被其他函数调用；
//AB

//2.下面这段代码输出什么以及原因？

func hello() []string {
	return nil
}

func main() {
	h := hello
	if h == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
}
//not nil，给h赋值的是函数hello不是函数返回值

//3.下面这段代码能否编译通过？如果可以，输出什么？

func GetValue() int {
	return 1
}

func main() {
	i := GetValue()
	switch i.(type) {
	case int:
		println("int")
	case string:
		println("string")
	case interface{}:
		println("interface")
	default:
		println("unknown")
	}
}
//编译失败，当i为接口类型时才可使用i.(type)语法