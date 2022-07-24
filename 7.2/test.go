package main

type Ner interface{
	a()
	b(int)
	c(string)string
}

type N int

func (N) a() {}
func (*N) b(int) {}
func (*N) c(string) string {return ""}

type data struct {
	x int
}

func main() {
	{
		var n N
		var t Ner = &n
		t.a()
	}
	{
		d := data{100}
		var t interface{} = d                  //将对象赋予接口变量时，会复制该变量
		println(t.(data).x)
	}
	{
		d := data{100}
		var t interface{} = &d                 //将对象指针赋值给接口，接口内存储的就是指针的复制品
		t.(*data).x = 200
		println(t.(*data).x)
	}
	{
		var a interface{} = nil
		var b interface{} = (*int)(nil)
		println(a == nil, b == nil)             //只有当接口内部的两个指针(itab, data)都为nil时，接口才等于nil
	}
}