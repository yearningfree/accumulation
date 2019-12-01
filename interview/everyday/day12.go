//1.下面属于关键字的是（）

A.func
B.struct
C.class
D.defer
//ABD

//2.下面这段代码输出什么？

func main() {  
    i := -5
    j := +5
    fmt.Printf("%+d %+d", i, j)
}
A. -5 +5
B. +5 +5
C. 0  0
//A

//3.下面这段代码输出什么？

type People struct{}

func (p *People) ShowA() {
    fmt.Println("showA")
    p.ShowB()
}
func (p *People) ShowB() {
    fmt.Println("showB")
}

type Teacher struct {
    People
}

func (t *Teacher) ShowB() {
    fmt.Println("teacher showB")
}

func main() {
    t := Teacher{}
    t.ShowB()
}
//teacher showB
//当外部类型声明自己的方法时会覆盖内部类型方法