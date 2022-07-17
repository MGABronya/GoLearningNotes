package main

func hello() {
	println("hello, world!")
}

func exec(f func()) {
	f()
}

//定义函数类型
type FormatFunc func(string, ...interface{}) (string, error)

//如不使用命名类型，这个函数签名会长到没法看
func format(f FormatFunc, s string, a ...interface{}) (string, error) {
	return f(s, a...)
}

func a() {}
func b() {}

func test() *int { //从函数返回局部变量指针是安全的，编译器会通过逃逸分析来决定是否在堆上分配内存
	a := 0x100
	return &a
}

func main() {
	{
		f := hello //函数属于第一类对象，具备相同签名（参数及返回值列表）的视作同一类型
		exec(f)
	}
	{
		println(a == nil) //函数只能判断其是否为nil，不支持其它比较操作
	}
	{
		var a *int = test()
		println(a, *a)
	}
}
