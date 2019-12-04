//1.下面这段代码正确的输出是什么？

func f() {
    defer fmt.Println("D")
    fmt.Println("F")
}

func main() {
    f()
    fmt.Println("M")
}
A. F M D
B. D F M
C. F D M
//C

//2.下面代码输出什么？

type Person struct {
    age int
}

func main() {
    person := &Person{28}

    // 1.
    defer fmt.Println(person.age)

    // 2.
    defer func(p *Person) {
        fmt.Println(p.age)
    }(person)

    // 3.
    defer func() {
        fmt.Println(person.age)
    }()

    person = &Person{29}
}
//29 28 28
//1.存入缓存的person.age=28，存在栈中等待最后取出
//2.此处的参数是地址，传入的也是p=28的地址
//3.闭包引用