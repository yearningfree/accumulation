// 1.下面的代码有什么问题？
func (n N) value {
	n++
	fmt.Printf("v:%p,%v\n",&n,n)
}

func (n *N) pointer(){
	*n++
	fmt.Printf("v:%p,%v\n",n,*n)
}

func main()  {
	var a N = 25
	p := &a
	p1 := &p
	
	p1.value()
	p1.pointer()
}
//编译错误，不能使用多级指针调用方法

//2.下面的代码输出什么？
type N int
 
 func (n N) test(){
     fmt.Println(n)
 }
 
 func main()  {
    var n N = 10
    fmt.Println(n)

    n++
    f1 := N.test
    f1(n)

    n++
    f2 := (*N).test
    f2(&n)
}
//10  10 12