//1.下面这段代码输出结果正确正确吗？

type Foo struct {
    bar string
}
func main() {
    s1 := []Foo{
        {"A"},
        {"B"},
        {"C"},
    }
    s2 := make([]*Foo, len(s1))
    for i, value := range s1 {
        s2[i] = &value
    }
    fmt.Println(s1[0], s1[1], s1[2])
    fmt.Println(s2[0], s2[1], s2[2])
}
输出：
{A} {B} {C}
&{A} &{B} &{C}
//不正确，第二行输出&{C} &{C} &{C}


//2.下面代码里的 counter 的输出值？

func main() {

    var m = map[string]int{
        "A": 21,
        "B": 22,
        "C": 23,
    }
    counter := 0
    for k, v := range m {
        if counter == 0 {
            delete(m, "A")
        }
        counter++
        fmt.Println(k, v)
    }
    fmt.Println("counter is ", counter)
}
A. 2
B. 3
C. 2 或 3
//C，因为map是无序的，如果第1次循环到A则输出3，否则输出2