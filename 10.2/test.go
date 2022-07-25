package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	{
		a := 100
		va, vp := reflect.ValueOf(a), reflect.ValueOf(&a).Elem()
		fmt.Println(va.CanAddr(), va.CanSet()) //false false
		fmt.Println(vp.CanAddr(), vp.CanSet()) //true true
	}
	{
		type User struct {
			Name string
			code int //不能对非导出字段进行设置操作，无论当前包还是外包
		}
		p := new(User)
		v := reflect.ValueOf(p).Elem()
		name := v.FieldByName("Name")
		code := v.FieldByName("code")
		fmt.Printf("name: canaddr = %v, canset = %v\n", name.CanAddr(), name.CanSet()) //true true
		fmt.Printf("code: canaddr = %v, canset = %v\n", code.CanAddr(), code.CanSet()) //true false
		if name.CanSet() {
			name.SetString("Tom")
		}
		if code.CanAddr() {
			*(*int)(unsafe.Pointer(code.UnsafeAddr())) = 100
		}
		fmt.Printf("%+v\n", *p)
	}
	{ //使用Interface方法进行类型的识别和转换
		type user struct {
			Name string
			Age  int
		}
		u := user{
			"q.yuhen",
			60,
		}
		v := reflect.ValueOf(&u)
		if !v.CanInterface() {
			println("CanInterface: fail.")
			return
		}
		p, ok := v.Interface().(*user)
		if !ok {
			println("Interface: fail.")
			return
		}
		p.Age++
		fmt.Printf("%+v\n", u)
	}
	{ //复合类型对象设置设立
		c := make(chan int, 4)
		v := reflect.ValueOf(c)
		if v.TrySend(reflect.ValueOf(100)) {
			fmt.Println(v.TryRecv())
		}
	}
	{
		var a interface{} = nil
		var b interface{} = (*int)(nil)
		fmt.Println(a == nil)                             //true
		fmt.Println(b == nil, reflect.ValueOf(b).IsNil()) //false true
	}
	{
		var b interface{} = (*int)(nil)
		iface := (*[2]uintptr)(unsafe.Pointer(&b))
		fmt.Println(iface, iface[1] == 0)
	}
	{
		v := reflect.ValueOf(struct{ name string }{})
		println(v.FieldByName("name").IsValid()) //true
		println(v.FieldByName("xxx").IsValid())  //false
	}
}
