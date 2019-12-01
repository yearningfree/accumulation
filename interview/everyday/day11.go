//1.关于 cap() 函数的适用类型，下面说法正确的是()

A. array
B. slice
C. map
D. channel
//ABD

//2.下面这段代码输出什么？

func main() {  
    var i interface{}
    if i == nil {
        fmt.Println("nil")
        return
    }
    fmt.Println("not nil")
}
A. nil
B. not nil
C. compilation error  
//A

//3.下面这段代码输出什么？

func main() {  
    s := make(map[string]int)
    delete(s, "h")
    fmt.Println(s["h"])
}
A. runtime panic
B. 0
C. compilation error 
//B