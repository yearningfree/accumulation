//1.下面代码中，x 已声明，y 没有声明，判断每条语句的对错。

1. x, _ := f()
2. x, _ = f()
3. x, y := f()
4. x, y = f()
//错 对 对 错
//1.x已声明，不能使用:=
//3.当多赋值时，:=左边不论是否声明都可以使用

//2.下面代码输出什么？

func increaseA() int {
    var i int
    defer func() {
        i++
    }()
    return i
}

func increaseB() (r int) {
    defer func() {
        r++
    }()
    return r
}

func main() {
    fmt.Println(increaseA())
    fmt.Println(increaseB())
}
A. 1 1
B. 0 1
C. 1 0
D. 0 0
//B

//3.下面代码输出什么？

type A interface {
    ShowA() int
}

type B interface {
    ShowB() int
}

type Work struct {
    i int
}

func (w Work) ShowA() int {
    return w.i + 10
}

func (w Work) ShowB() int {
    return w.i + 20
}

func main() {
    var a A = Work{3}
    s := a.(Work)
    fmt.Println(s.ShowA())
    fmt.Println(s.ShowB())
}
A. 13 23
B. compilation error
//A