package main

import "sync"

type data struct {
	sync.Mutex
	buf [1024]byte
}

type user struct{}

type maneger struct {
	user
}

func (user) toString() string {
	return "user"
}

func (m maneger) toString() string {
	return m.user.toString() + ";manager"
}

func main() {
	{ //可以向访问匿名字段成员那样调用其方法，由编译器负责查找
		d := data{}
		d.Lock()
		defer d.Unlock()
	}
	{
		var m maneger
		println(m.toString())
		println(m.user.toString())
	}
	{

	}
}
