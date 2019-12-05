//1.下面这段代码输出什么？为什么？

func (i int) PrintInt ()  {
    fmt.Println(i)
}

func main() {
    var i int = 1
    i.PrintInt()
}
A. 1
B. compilation error
//B
//基于类型定义的方法必须定义在同一个包内

//2.下面这段代码输出什么？为什么？

type People interface {
    Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
    if think == "speak" {
        talk = "speak"
    } else {
        talk = "hi"
    }
    return
}

func main() {
    var peo People = Student{}
    think := "speak"
    fmt.Println(peo.Speak(think))
}
A. speak
B. compilation error
//B
//值类型Student没有实现Speak方法，而是指针类型*Student实现该方法
