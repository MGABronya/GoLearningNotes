package main

func test1(f func()) {
	f()
}

func test2() func(int, int) int {
	return func(x, y int) int {
		return x + y
	}
}

func testStruct() {
	type calc struct { //定义结构体类型
		mul func(x, y int) int //函数类型字段
	}

	x := calc{
		mul: func(x, y int) int {
			return x * y
		},
	}

	println(x.mul(2, 3))
}

func testChannel() {
	c := make(chan func(int, int) int, 2)

	c <- func(x, y int) int {
		return x + y
	}
	println((<-c)(1, 2))
}

func test3(x int) func() {
	println(&x)
	return func() {
		println(x, &x) //闭包，函数+引用环境的组合体，该匿名函数引用了上下文中的x
	}
}

func test4() []func() {
	var s []func()
	for i := 0; i < 2; i++ {
		s = append(s, func() { //将多个匿名函数添加至列表
			println(&i, i)
		})
	}
	return s
}

func main() {
	{ //匿名函数直接执行
		func(s string) {
			println(s)
		}("hello world")
	}
	{ //赋值给变量
		add := func(x, y int) int {
			return x + y
		}
		println(add(1, 2))
	}
	{ //作为参数
		test1(func() {
			println("hello world")
		})
	}
	{ //作为返回值
		add := test2()
		println(add(1, 2))
	}
	{
		for _, f := range test4() {
			f() //此处打印的结果将次次相同，因为匿名函数中的i始终为同一个i
		}
	}
}
