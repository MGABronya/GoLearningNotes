package main

type tester1 interface {
	test()
	string() string
}

type data1 struct {}

func(*data1) test() {}
func(data1) string() string { return "" }

type stringer1 interface {
	string() string
}

type tester2 interface {
	stringer1
	test()
}

type data2 struct {}

func (*data2) test() {}

func (data2) string() string {
	return ""
}

func pp(a stringer1){
	println(a.string())
}

type node struct {
	data1 interface {                     //匿名接口类型
		string() string
	}
}

func main() {
	{
		var d data1
		var t tester1 = &d
		t.test()
		println(t.string())
	}
	{
		var t1, t2 interface{}
		println(t1 == nil, t1 == t2)
		t1, t2 = 100, 100
		println(t1 == t2)
		t1, t2 = map[string]int{}, map[string]int{}
		println(t1 == t2)                              //此处报错，实现接口的类型不支持相等运算
	}
	{
		var d data2
		var t tester2 = &d
		pp(t)                              //隐式转换为子集接口
		var s stringer1 = t                //转换为子集
		println(s.string)
	}
	{
		var t interface{                          //定义匿名接口变量
			string() string
		} = data1{}
		n := node{
			data1: t,
		}
		println(n.data1.string())
	}
}