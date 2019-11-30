//1.关于channel，下面语法正确的是()

A. var ch chan int
B. ch := make(chan int)
C. <- ch
D. ch <-
//ABC
//D向channel写数据未带上值

//2.下面这段代码输出什么？

A.0
B.1
C.Compilation error  

type person struct{
	name string
}

func main(){
	var m map[person]int
	p := person{"mike"}
	fmt.Println(m[p])
}
//A，p未加入m中，此时打印的是零值

//3.下面这段代码输出什么？

A.18
B.5
C.Compilation error  

func hello(num ...int){
	num[0] = 18
}

func main(){
	i := []int{5,6,7}
	hello(i...)
	fmt.Println(i[0])
}
//A