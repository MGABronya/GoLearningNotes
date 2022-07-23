package main

import (
	"fmt"
	"reflect"
)

type S struct{}

type T struct {
	S //匿名嵌入字段
}

func (S) sVal()  {}
func (*S) sPtr() {}
func (T) tVal()  {}
func (*T) tPtr() {}

func methodSet(a interface{}) { //显示方法集里所有方法的名字
	t := reflect.TypeOf(a)
	for i, n := 0, t.NumMethod(); i < n; i++ {
		m := t.Method(i)
		fmt.Println(m.Name, m.Type)
	}
}

func main() {
	{
		var t T
		methodSet(t)
		println("-------")
		methodSet(&t)
	}
}
