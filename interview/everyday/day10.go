//1.下面这段代码输出什么？

func main() {  
    a := 5
    b := 8.1
    fmt.Println(a + b)
}
A.13.1  
B.13
C.compilation error 
//C

//2.下面这段代码输出什么？

package main

import (  
    "fmt"
)

func main() {  
    a := [5]int{1, 2, 3, 4, 5}
    t := a[3:4:4]
    fmt.Println(t[0])
}
A.3
B.4
C.compilation error 
//B

//3.下面这段代码输出什么？

func main() {
    a := [2]int{5, 6}
    b := [3]int{5, 6}
    if a == b {
        fmt.Println("equal")
    } else {
        fmt.Println("not equal")
    }
}
A. compilation error  
B. equal  
C. not equal
//C