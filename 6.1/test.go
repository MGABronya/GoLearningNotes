package main

import "fmt"

type N int

func (n N) toString() string {
	return fmt.Sprintf("%#x", n)
}

func (N) test() {
	println("hi!")
}

func (n N) value() {
	n++
	fmt.Printf("v: %p, %v\n", &n, n)
}

func (n *N) pointer() {
	(*n)++
	fmt.Printf("v: %p, %v\n", n, *n)
}

func main() {
	{
		var a N = 25
		println(a.toString())
	}
	{
		var a N = 25
		a.value() //此处对象实例被复制
		a.pointer()
		fmt.Printf("a: %p, %v\n", &a, a)
	}
	{
		var a N = 25
		p := &a //可以使用对象实例的指针调用方法,但不能用多级指针调用方法
		a.value()
		a.pointer()
		p.value()
		p.pointer()
	}
	{

	}
}
