package main

import (
	"fmt"
	"net/http"
	"reflect"
)

type A int
type B struct {
	A
}

func (A) av() {}

func (B) bv()  {}
func (*B) bp() {}

type X int

func (X) String() string {
	return ""
}

func main() {
	{
		type X int
		var a X = 100
		t := reflect.TypeOf(a)
		fmt.Println(t.Name(), t.Kind()) //前者表示真实类型，后者表示底层类型
		//              X        int
	}
	{
		type X int
		type Y int
		var a, b X = 100, 200
		var c Y = 300
		ta, tb, tc := reflect.TypeOf(a), reflect.TypeOf(b), reflect.TypeOf(c)
		fmt.Println(ta == tb, ta == tc)     //true false
		fmt.Println(ta.Kind() == tc.Kind()) //true
	}
	{ //直接构造一些基础复合类型
		a := reflect.ArrayOf(10, reflect.TypeOf(byte(0)))
		m := reflect.MapOf(reflect.TypeOf(""), reflect.TypeOf(0))
		fmt.Println(a, m) //[10]uint8   map[string]int
	}
	{ //指针类型和基类型不是同一类型
		x := 100
		tx, tp := reflect.TypeOf(x), reflect.TypeOf(&x)
		fmt.Println(tx, tp, tx == tp)     //int int* false
		fmt.Println(tx.Kind(), tp.Kind()) //int ptr
		fmt.Println(tx == tp.Elem())      //Elem返回指针、数组、切片、字典（值）或通道的基类型
	}
	{
		fmt.Println(reflect.TypeOf(map[string]int{}).Elem()) //int
		fmt.Println(reflect.TypeOf([]int32{}).Elem())        //int32
	}
	{
		type user struct {
			name string
			age  int
		}
		type manager struct {
			user
			title string
		}
		var m manager
		t := reflect.TypeOf(&m)
		if t.Kind() == reflect.Ptr {
			t = t.Elem() //获取指针的基类型
		}
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			fmt.Println(f.Name, f.Type, f.Offset)
			if f.Anonymous { //输出匿名字段结构
				for x := 0; x < f.Type.NumField(); x++ {
					af := f.Type.Field(x)
					fmt.Println(" ", af.Name, af.Type)
				}
			}
		}
	}
	{
		type user struct {
			name string
			age  int
		}
		type manager struct {
			user
			title string
		}
		var m manager
		t := reflect.TypeOf(m)
		name, _ := t.FieldByName("name") //按名称查找
		fmt.Println(name.Name, name.Type)
		age := t.FieldByIndex([]int{0, 1}) //按多级索引查找
		fmt.Println(age.Name, age.Type)
	}
	{
		var b B
		t := reflect.TypeOf(b)
		s := []reflect.Type{t, t.Elem()}
		for _, t := range s {
			fmt.Println(t, ":")
			for i := 0; i < t.NumMethod(); i++ {
				fmt.Println(" ", t.Method(i))
			}
		}
	}
	{
		var s http.Server
		t := reflect.TypeOf(s)
		for i := 0; i < t.NumField(); i++ {
			fmt.Println(t.Field(i).Name)
		}
	}
	{
		type user struct {
			name string `field:"name" type:"varchar(50)"`
			age  int    `field:"age" type:"int"`
		}
		var u user
		t := reflect.TypeOf(u)
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			fmt.Printf("%s: %s %s\n", f.Name, f.Tag.Get("field"), f.Tag.Get("type"))
		}
	}
	{
		var a X
		t := reflect.TypeOf(a)
		st := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
		fmt.Println(t.Implements(st)) //true
		it := reflect.TypeOf(0)
		fmt.Println(t.ConvertibleTo(it))                    //true
		fmt.Println(t.AssignableTo(st), t.AssignableTo(it)) //true false
	}
}
