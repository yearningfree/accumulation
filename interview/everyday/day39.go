//1.关于无缓冲和有冲突的channel，下面说法正确的是？

A. 无缓冲的channel是默认的缓冲为1的channel；
B. 无缓冲的channel和有缓冲的channel都是同步的；
C. 无缓冲的channel和有缓冲的channel都是非同步的；
D. 无缓冲的channel是同步的，而有缓冲的channel是非同步的；
//D

//2.下面代码是否能编译通过？如果通过，输出什么？

 func Foo(x interface{}) {
    if x == nil {
        fmt.Println("empty interface")
        return
    }
    fmt.Println("non-empty interface")
}
func main() {
    var x *int = nil
    Foo(x)
}
//non-empty interface

//3.下面代码输出什么？
func main(){
	ch := make(chan int, 100)
	//A
	go func(){
		for i := 0; i < 10; i++ {
			ch<-i
		}
	}()
	//B
	go fun(){
		for{
			a, ok := <- ch
			if !ok{
				fmt.println("close")
				return
			}
			fmt.Println("a:", a)
		}
	}()
	close(ch)
	fmt.Println("ok")
	time.Sleep(time.Second * 10)
}
//panic，在A协程使用的ch被主程关闭，当A往ch里发送数据时会panic