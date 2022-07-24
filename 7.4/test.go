package main

import "fmt"

type FuncString func() string

func (f FuncString) String() string {
	return f()
}

func main() {
	var t fmt.Stringer = FuncString(func() string {       //将一个匿名函数类型转化为FuncString，赋给接口
		return "hello, world!"
	})
	fmt.Println(t)
}