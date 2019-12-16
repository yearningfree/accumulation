//1.关于 bool 变量 b 的赋值，下面错误的用法是？

A. b = true
B. b = 1
C. b = bool(1)
D. b = (1 == 2)
//BC

//2.关于变量的自增和自减操作，下面语句正确的是？

A.
i := 1
i++

B.
i := 1
j = i++

C.
i := 1
++i

D.
i := 1
i--
//AD

//3.关于GetPodAction定义，下面赋值正确的是

type Fragment interface {
	Exec(transInfo *TransInfo) error
}
type GetPodAction struct {
}
func (g GetPodAction) Exec(transInfo *TransInfo) error {
	...
	return nil
}
A. var fragment Fragment = new(GetPodAction)
B. var fragment Fragment = GetPodAction
C. var fragment Fragment = &GetPodAction{}
D. var fragment Fragment = GetPodAction{}
//ACD
