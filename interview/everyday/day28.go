//1.下面的代码有什么问题？

func main() {
    fmt.Println([...]int{1} == [2]int{1})
    fmt.Println([]int{1} == []int{1})
}
//数组长度不同，不能比较
//切片不能比较

//2.一道很有代表性的题目，很多老司机都因此翻车！

//下面这段代码输出什么？如果编译错误的话，为什么？

var p *int

func foo() (*int, error) {
    var i int = 5
    return &i, nil
}

func bar() {
    //use p
    fmt.Println(*p)
}

func main() {
    p, err := foo()
    if err != nil {
        fmt.Println(err)
        return
    }
    bar()
    fmt.Println(*p)
}
A. 5 5
B. runtime error
//B
//变量作用域，main函数中的:=会定义一个新的p变量，会遮住全局变量p，导致调用bar()时p仍是nil