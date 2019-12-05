//1.下面这段代码输出什么？为什么？

func main() {

    s1 := []int{1, 2, 3}
    s2 := s1[1:]
    s2[1] = 4
    fmt.Println(s1)
    s2 = append(s2, 5, 6, 7)
    fmt.Println(s1)
}
//[1 2 4] [1 2 4]

//2.下面选项正确的是？

func main() {
    if a := 1; false {
    } else if b := 2; false {
    } else {
        println(a, b)
    }
}
A. 1 2
B. compilation error
//A