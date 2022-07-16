package main

import "unsafe"

const (
	parSize = unsafe.Sizeof(uintptr(0)) //常量值可以是某些编译器能计算出结果的表达式
	strSize = len("hello, world!")
)

type color byte //自定义类型

const (
	black color = iota
	red
	blue
)

func main() {
	const x = 123 //常量未使用不会报错
	const y = 1.23
	{
		const x = "abc"
	}
	const (
		a uint16 = 120
		b        //与上一行的类型、右值相同
		c = "abc"
		d //与上一行的类型、右值相同
	)
	const (
		e = iota //0
		f        //1
		g        //2
	)
	const (
		_  = iota             //0
		KB = 1 << (10 * iota) //1 << 10
		MB = 1 << (10 * iota) //1 << 20
		GB = 1 << (10 * iota) //1 << 30
	)
	const (
		_, _ = iota, iota * 10 //每个iota单独计数
		h, i
		j, k
	)
	const ( //iota以行计数
		l = iota //0
		m        //1
		n = 100  //100
		o        //100
		p = iota //4
		q        //5
	)
	/*
		不同于变量在运行期分配存储内存（非优化状态），常量通常会被编译器在预处理阶段直接展开，作为指令数据使用
	*/

}
