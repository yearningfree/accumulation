//1.下面这段代码输出什么？
type T struct{
	ls []int
}

func foo(t T){
	t.ls[0] = 100
}

func main()  {
	var t = T{
		ls : []int{1,2,3}
	}
	foo(t)
	fmt.Println(t.ls[0])
}
//B

//2.下面代码输出什么？
func main()  {
	isMatch := func (i int) bool {
		switch(i){
		case 1:
		case 2:
			return true
		}
		return false
	}
	fmt.Println(isMatch(1))
	fmt.Println(isMatch(2))
}
//false true,Go 语言的 switch 语句虽然没有"break"，但如果 case 完成程序会默认 break