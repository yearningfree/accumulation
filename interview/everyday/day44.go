//1.下面代码有什么问题？

func main() {
    m := make(map[string]int,2)
    cap(m) 
}
//cap函数不适用与map，可以使用len函数返回map元素个数
//使用 make 创建 map 变量时可以指定第二个参数，不过会被忽略

//2.下面的代码有什么问题？

func main() {  
    var x = nil 
    _ = x
}
//nil 用于表示 interface、函数、maps、slices 和 channels 的“零值”。如果不指定变量的类型，编译器猜不出变量的具体类型，导致编译错误

//3.下面代码能编译通过吗？

type info struct {
	result int
}

func work() (int,error) {
	return 13,nil
}

func main()  {
	var data info

	data.result, err := work()
	fmt.Printf("info: %+v\n",data)
}
//不能使用短变量声明设置结构体字段值
//正确写法
func main()  {
	var data info
	var err error
	data.result, err = work()
	if err !=nil{
		fmt.Println(err)
		return
	}
	fmt.Printf("info: %+v\n",data)
}