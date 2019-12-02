package main

//1.切片 a、b、c 的长度和容量分别是多少？

func main() {

	s := [3]int{1, 2, 3}
	a := s[:0]
	b := s[:2]
	c := s[1:2:cap(s)]
}

//len(a)=0 cap(a)=3
//len(b)=2 cap(b)=3
//len(c)=1 cap(c)=2

//2.下面代码中 A B 两处应该怎么修改才能顺利编译？

func main() {
    var m map[string]int        //A
    m["a"] = 1
    if v := m["b"]; v != nil {  //B
        fmt.Println(v)
    }
}
//A  m:=make(map[string]int)
//B  if v, ok := m["b"]; ok {

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
    c := Work{3}
    var a A = c
    var b B = c
    fmt.Println(a.ShowB())
    fmt.Println(b.ShowA())
}
A. 23 13
B. compilation error
//B
//a的静态类型为A，并未实现ShowB()，同理b也未实现ShowA()