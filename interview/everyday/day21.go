//1.下面的两个切片声明中有什么区别？哪个更可取？

A. var a []int
B. a := []int{}
//A是nil切片
//B是空切片，长度、容量都为0
//第一种不会分配内存，优先选择

//2. A、B、C、D 哪些选项有语法错误？

type S struct {
}

func f(x interface{}) {
}

func g(x *interface{}) {
}

func main() {
    s := S{}
    p := &s
    f(s) //A
    g(s) //B
    f(p) //C
    g(p) //D
}
//B应写成g(&s)
//D
//interface参数可以接收任何类型参数，包括指针，不要写*interface{}

//3.下面 A、B 两处应该填入什么代码，才能确保顺利打印出结果？

type S struct {
    m string
}

func f() *S {
    return __  //A
}

func main() {
    p := __    //B
    fmt.Println(p.m) //print "foo"
}
//&S{"foo"}
//*f()或f()