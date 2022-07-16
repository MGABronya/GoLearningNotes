package main

import "fmt"

type flags byte

const (
	read flags = 1 << iota
	write
	exec
)

func main() {
	{
		f := read | exec
		fmt.Printf("%b\n", f)
	}
	{
		type ( //组
			user struct { //结构体
				name string
				age  uint8
			}
			event func(string) bool
		)
		u := user{"Tom", 20}
		fmt.Println(u)

		var f event = func(s string) bool {
			println(s)
			return s != ""
		}
		f("abc")
	}
	{ //自定义类型不能视作别名，不能隐式转换，不能直接用于比较表达式
		type data int
		var d data = 10
		println(d)
	}
	{ //未命名类型
		type data [2]int
		var d data = [2]int{1, 2} //基础类型相同，右值为未命名类型
		fmt.Println(d)
		a := make(chan int, 2)
		var b chan<- int = a //双向通道转换为单向通道，其中b为未命名类型
		b <- 2
	}
}
