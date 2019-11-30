//1.关于字符串连接，下面语法正确的是？

A. str := 'abc' + '123'
B. str := "abc" + "123"
C. str := '123' + "abc"
D. fmt.Sprintf("abc%d", 123)
//BD

//2.下面这段代码能否编译通过？如果可以，输出什么？

const(
	x = iota
	_
	y
	z = "zz"
	k
	p = iota
)

func main(){
	fmt.Println(x,y,z,k,p)
}
//0 2 zz zz 5

//3.下面赋值正确的是()

A. var x = nil
B. var x interface{} = nil
C. var x string = nil
D. var x error = nil
//BD
//A为不确定类型，C的string零值为“”