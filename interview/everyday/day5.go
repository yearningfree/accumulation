package main

import "fmt"

//1.下面这段代码能否通过编译？不能的话，原因是什么？如果通过，输出什么？

func main() {
	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	if sn1 == sn2 {
		fmt.Println("sn1==sn2")
	}

	sm1 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	sm2 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	if sm1 == sm2 {
		fmt.Println("sm1==sm2")
	}
}

//sm1==sm2编译不通过，map、slice为不可比较类型，只能与nil比较
//在比较struct的时候，只能比较是否相等，不能比较大小，
//如果struct中的属性类型相同但顺序不同也是不相等的
