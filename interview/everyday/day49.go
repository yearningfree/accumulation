//1.下面代码输出什么？
func main(){
	var ch chan int
	select{
	case v,ok := <-ch:
		println(v,ok)
	default:
		println("default")
	}
}
//default ch为nil,读写都会阻塞

//2.下面这段代码输出什么？
type People struct{
	name string `json:"name"`
}

func main()  {
	js := `{
		"name":"seekload"
	}`
	var p People
	err := json.Unmarshal([]byte(js),&p)
	if err != nil{
		fmt.Println("err: ", err)
	}
	fmt.Println(p)
}
//{},因为name首字母小写，其他包无法访问