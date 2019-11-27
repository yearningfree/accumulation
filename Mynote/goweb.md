[《Go Web 编程
》](https://astaxie.gitbooks.io/build-web-application-with-golang/zh/)

###Go语言基础
####Go基础
#####iota枚举
	const(
		x = iota//0
		y = iota//1
		z = iota//2
		h       //3，常量声明省略值时，默认和前一个值相同。此处为隐式的iota
	)
	
	const v = iota//0，每遇到一个const，iota就会重置
	
	const(
		x, y, z = iota, iota, iota//x=0,y=0,z=0，iota在同一行值相同
	)
	
	const(
		a = iota//0
		b = "B"
		c = iota//2
		d, e, f = iota, iota, iota//d=3,e=3,f=3
		g = iota//4
	)

#####array、slice、map
array 数组，定义：

	var arr [n]type//长度为n
	arr := [3]type{1,2,3}
	arr := [...]type//省略长度，go自动计算
	//二维数组
	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}
	//简写忽略内部类型
	easyArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}

slice 切片，“动态数组”

	var slice []type//[]中不用制定长度
	len()//获取slice长度
	cap()//获取slice容量
	append(slice, elems)//向一个slice追加一个或多个元素，会影响原数组
	copy(dst, src)//从源slice的src中复制元素到目标dst，并且返回复制的元素的个数

map 字典

	mapArray := make(map[string]int)//定义
	mapArray["one"] = 1//赋值
	i, ok := mapArray["two"]//判断two是否存在map中
	delete(mapArray, "one")//删除元素
	

* map是无序的，每次打印出来的map都会不一样，它不能通过index获取，而必须通过key获取
* map的长度是不固定的，也就是和slice一样，也是一种引用类型
* 内置的len函数同样适用于map，返回map拥有的key的数量
* map的值可以很方便的修改，通过numbers["one"]=11可以很容易的把key为one的字典值改为11
* map和其他基本型别不同，它不是thread-safe，在多个go-routine存取时，必须使用mutex lock机制



#####make、new操作
make用于内建类型（map、slice 和channel）的内存分配。new用于各种类型的内存分配，返回指针。

####流程和函数
#####goto
用goto跳转到必须在当前函数内定义的标签。

	func myFunc() {
	    i := 0
	Here:   //这行的第一个词，以冒号结束作为标签
	    println(i)
	    i++
	    goto Here   //跳转到Here去
	}
	标签名是大小写敏感的。

#####init和main
每个package都有init()函数，可选是否使用，main()函数仅在package main中可用；函数引入包初始化流程：
->package main->import pkg->const...->var...->init()->main()
#####import

1.相对路径

	import “./model”//当前文件同一目录的model目录

2.绝对路径

	impoet "pkg/model"//gopath/src/pkg/model

---
1.点操作

	import(
		. "fmt"
	)
	可将包名省略
	fmt.Println("hello world")可以省略的写成Println("hello world")

2.别名操作

	import(
			f "fmt"
	)
	
	fmt.Println("hello world")->f.Println("hello world")
	
3._操作

	import(
			_ “pkg/model”
	)
	_操作其实是引入该包，而不直接使用包里面的函数，而是调用了该包里面的init函数。
	
####面向对象
######method
method是附属在一个给定的类型上的，他的语法和函数的声明语法几乎一样，只是在func后面增加了一个receiver(也就是method所依从的主体)。method的语法如下：

	func (r ReceiverType) funcName(parameters) (results)
---
	package main
	
	import (
	    "fmt"
	    "math"
	)
	
	type Rectangle struct {
	    width, height float64
	}
	
	type Circle struct {
	    radius float64
	}
	
	func (r Rectangle) area() float64 {
	    return r.width*r.height
	}
	
	func (c Circle) area() float64 {
	    return c.radius * c.radius * math.Pi
	}
在使用method的时候重要注意几点

* 虽然method的名字一模一样，但是如果接收者不一样，那么method就不一样
* method里面可以访问接收者的字段
* 调用method通过.访问，就像struct里面访问字段一样

#####指针作为receiver
当需要改变类型的属性值时，需要将类型的指针作为receiver，否则传的参数只是类型的copy。

	func (r *ReceiverType) funcName(parameters) (results)

#####method继承

如果匿名字段实现了一个method，那么包含这个匿名字段的struct也能调用该method。

	func main() {
		xiaoming := student{human{"xiaoming", 18, "男"}, 7, "class 1"}
		xiaoming.sayHi()
	}
	
	type human struct {
		name string
		age  int
		sex  string
	}
	
	type student struct {
		human
		id    int
		class string
	}
	
	func (h human) sayHi() {
		fmt.Printf("hi i am %s sex:%s age:%d\n", h.name, h.sex, h.age)
	}

#####method重写

如果匿名字段实现了一个method，那么该类型可以根据需要自己写一个以该类型作为receiver的method。

	type human struct {
		name string
		age  int
		sex  string
	}
	
	type student struct {
		human
		id    int
		class string
	}
	
	func (s student) sayHi() {
		fmt.Printf("hi i am %s sex:%s age:%d id:%d class:%s\n", s.name, s.sex, s.age, s.id, s.class)
	}
	
	func (h human) sayHi() {
		fmt.Printf("hi i am %s sex:%s age:%d\n", h.name, h.sex, h.age)
	}

####interface
#####什么是interface
简单的说，interface是一组method签名的组合，我们通过interface来定义对象的一组行为。
#####interface类型
interface类型定义了一组方法，如果某个对象实现了某个接口的所有方法，则此对象就实现了此接口。

	type Human struct {
	    name string
	    age int
	    phone string
	}
	
	type Student struct {
	    Human //匿名字段Human
	    school string
	    loan float32
	}
	
	type Employee struct {
	    Human //匿名字段Human
	    company string
	    money float32
	}
	
	//Human对象实现Sayhi方法
	func (h *Human) SayHi() {
	    fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
	}
	
	// Human对象实现Sing方法
	func (h *Human) Sing(lyrics string) {
	    fmt.Println("La la, la la la, la la la la la...", lyrics)
	}
	
	//Human对象实现Guzzle方法
	func (h *Human) Guzzle(beerStein string) {
	    fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
	}
	
	// Employee重载Human的Sayhi方法
	func (e *Employee) SayHi() {
	    fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
	        e.company, e.phone) //此句可以分成多行
	}
	
	//Student实现BorrowMoney方法
	func (s *Student) BorrowMoney(amount float32) {
	    s.loan += amount // (again and again and...)
	}
	
	//Employee实现SpendSalary方法
	func (e *Employee) SpendSalary(amount float32) {
	    e.money -= amount // More vodka please!!! Get me through the day!
	}
	
	// 定义interface
	type Men interface {
	    SayHi()
	    Sing(lyrics string)
	    Guzzle(beerStein string)
	}
	
	type YoungChap interface {
	    SayHi()
	    Sing(song string)
	    BorrowMoney(amount float32)
	}
	
	type ElderlyGent interface {
	    SayHi()
	    Sing(song string)
	    SpendSalary(amount float32)
	}
	
Human实现了Men interface，所以Student和Employee都实现了Men interface；Student还实现了YoungChap interface。

#####反射

。。。

####web
网页优化方面有一项措施是减少HTTP请求次数，就是把尽量多的css和js资源合并在一起，目的是尽量减少网页请求静态资源的次数，提高网页加载速度，同时减缓服务器的压力。

#####表单
######处理表单的输入
显式的调用r.ParseForm()后才能对表单数据进行操作（r.Form["xxx"]）。r.Form里面包含了所有请求的参数，比如URL中query-string、POST的数据、PUT的数据，所以当你在URL中的query-string字段和POST冲突时，会保存成一个slice，里面存储了多个值.

	func sayhelloName(w http.ResponseWriter, r *http.Request) {
	    r.ParseForm()       //解析url传递的参数，对于POST则解析响应包的主体（request body）
	    //注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	    fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	    fmt.Println("path", r.URL.Path)
	    fmt.Println("scheme", r.URL.Scheme)
	    fmt.Println(r.Form["url_long"])
	    for k, v := range r.Form {
	        fmt.Println("key:", k)
	        fmt.Println("val:", strings.Join(v, ""))
	    }
	    fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
	}
	
request.Form是一个url.Values类型，里面存储的是对应的类似key=value的信息，下面展示了可以对form数据进行的一些操作:

	v := url.Values{}
	v.Set("name", "Ava")
	v.Add("friend", "Jess")
	v.Add("friend", "Sarah")
	v.Add("friend", "Zoe")
	// v.Encode() == "name=Ava&friend=Jess&friend=Sarah&friend=Zoe"
	fmt.Println(v.Get("name"))
	fmt.Println(v.Get("friend"))
	fmt.Println(v["friend"])
Tips: Request本身也提供了FormValue()函数来获取用户提交的参数。如r.Form["username"]也可写成r.FormValue("username")。调用r.FormValue时会自动调用r.ParseForm，所以不必提前调用。r.FormValue只会返回同名参数中的第一个，若参数不存在则返回空字符串。

######验证表单的输入
如果是map值用r.Form[""][]获取；单个值可以用r.Form.Get("")获取，如果字段不存在取到空值。

######预防跨站脚本
Go的html/template里面带有下面几个函数可以帮你转义，默认帮你过滤了html标签。
	
	func HTMLEscape(w io.Writer, b []byte) //把b进行转义之后写到w
	func HTMLEscapeString(s string) string //转义s之后返回结果字符串
	func HTMLEscaper(args ...interface{}) string //支持多个参数一起转义，返回结果字符串

假设输入username是输入的username是`<script>alert()</script>`，

	template.HTMLEscape(w, []byte(r.Form.Get("username")))
输出：

	&lt;script&gt;alert()&lt;/script&gt;
text/template正常输出html标签：
	
	import "text/template"
	...
	t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	err = t.ExecuteTemplate(out, "T", "<script>alert('you have been pwned')</script>")
输出

	Hello, <script>alert('you have been pwned')</script>!

或者使用template.HTML类型

	import "html/template"
	...
	t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	err = t.ExecuteTemplate(out, "T", template.HTML("<script>alert('you have been pwned')</script>"))
输出

	Hello, <script>alert('you have been pwned')</script>!
转换成template.HTML后，变量的内容也不会被转义
转义的例子：

	import "html/template"
	...
	t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	err = t.ExecuteTemplate(out, "T", "<script>alert('you have been pwned')</script>")
转义之后的输出：

	Hello, &lt;script&gt;alert(&#39;you have been pwned&#39;)&lt;/script&gt;!