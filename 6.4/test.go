package main

import "fmt"

type N int

func (n N) test() {
	fmt.Printf("test.n:%p, %d\n", &n, n)
}

func (n *N) test1() {
	fmt.Printf("test.n:%p, %d\n", n, *n)
}

func call(m func()) {
	m()
}

func (N) value()    {}
func (*N) pointer() {}

func main() {
	{
		var n N = 25
		fmt.Printf("main.n:%p, %d\n", &n, n)
		f1 := N.test //func(n N)
		f1(n)
		f2 := (*N).test //func(n *N)
		f2(&n)
	}
	{ //表达式方法调用
		var n N = 25
		N.test(n)
		(*N).test(&n)
	}
	{ //当方法被赋值给变量或作为参数传递时，会立即计算并复制该方法执行所需的receiver对象
		var n N = 25
		p := &n
		n++
		f1 := n.test //此处复制reveiver，复制n，等于101
		n++
		f2 := p.test //复制*p，等于102
		n++
		fmt.Printf("main.n: %p, %v\n", p, n)
		f1()
		f2()
	}
	{
		var n N = 100
		p := &n
		fmt.Printf("main.n: %p, %v\n", p, n)
		n++
		call(n.test)
		n++
		call(p.test)
	}
	{ //如果方法的receiver是指针类型，那么被复制的仅是指针
		var n N = 100
		p := &n
		n++
		f1 := n.test1
		n++
		f2 := n.test1
		n++
		fmt.Printf("main.n: %p, %v\n", p, n) //n = 103
		f1()
		f2()
	}
	{ //只要receiver参数类型正确，使用nil同样也可以执行
		var p *N
		p.pointer()
		(*N)(nil).pointer()
		(*N).pointer(nil)
	}
}
