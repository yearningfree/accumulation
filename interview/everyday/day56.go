//1.下面的代码有什么问题？
package main
 
 import "fmt"
 
 func main() {
     s := make([]int, 3, 9)
     fmt.Println(len(s)) 
     s2 := s[4:8]
     fmt.Println(len(s2)) 
 }
 //输出3 4
 //从一个基础切片派生出的子切片的长度可能大于基础切片的长度
 //当基础切片的容量大于长度时，子切片的索引值可以大于基础切片的长度

 //2.下面代码输出什么？
type N int

func (n N) test(){
    fmt.Println(n)
}

func main()  {
    var n N = 10
    p := &n

    n++
    f1 := n.test

    n++
    f2 := p.test

    n++
    fmt.Println(n)

    f1()
    f2()
}
//13 11 12
//方法值。当指针值赋值给变量或者作为函数参数传递时，会立即计算并复制该方法执行所需的接收者对象，与其绑定，以便在稍后执行时，能隐式地传入接收者参数。