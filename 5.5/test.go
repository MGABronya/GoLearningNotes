package main

import (
	"fmt"
	"os"
	"reflect"
	"unsafe"
)

type attr struct{
	perm int
}

type file struct{
	name string
	attr
}

type data struct{
	os.File
}

type data1 struct{
	*int
	string
}

type file2 struct{
	name string
}

type data2 struct{
	file2
	name string
}

type user struct {
	name string `昵称`
	sex byte `性别`
}

func main() {
	{
		type node struct {         //字段名必须唯一，可用_补位，支持使用自身指针类型成员
			_    int
			id   int
			next *node
		}
		n1 := node{
			id: 1,
		}
		n2 := node{
			id:   2,
			next: &n1,
		}
		fmt.Println(n1, n2)
	}
	{
		type user struct {
			name string
			age byte
		}
		u1 := user{"Tom", 12}      //可按顺序初始化全部字段
		fmt.Println(u1)
	}
	{
		u := struct {
			name string
			age byte
		}{
			name: "Tom",
			age: 12,
		}
		type file struct{
			name string
			attr struct {
				owner int
				perm int
			}
		}
		f := file{
			name: "test.dat",
			//attr:{                                //错误，缺少类型表示，作为字段类型时无法直接初始化
			//	owner:1,
			//	perm:0775,
			//},
		}
		f.attr.owner = 1
		f.attr.perm = 0755
		fmt.Println(u, f)
	}
	{                                               //struct 只有在所有的字段均支持==操作时，才可以使用==操作
		type data struct{
			x int
		}
		d1 := data{
			x : 100,
		}
		d2 := data{
			x : 100,
		}
		println(d1 == d2)
	}
	{
		type user struct {
			name string
			age int
		}
		p := &user{                                //可使用指针直接操作结构字段，但不能是多级指针
			name: "Tom",
			age: 20,
		}
		p.name = "Mary"
		p.age++
	}
	{
		var a struct{}
		var b [100]struct{}
		println(unsafe.Sizeof(a), unsafe.Sizeof(b))           //空结构体，长度为0
	}
	{                             //尽管没有分配内存，但依然可以操作元素，对于切片len、cap属性也正常
		var d [100]struct{}
		s := d[:]
		d[1] = struct{}{}
		s[2] = struct{}{}
		fmt.Println(s[3], len(s), cap(s))
	}
	{
		exit := make(chan struct{})
		go func() {                 
			println("hello, world!")
			exit <- struct{}{}
		}()
		<-exit
		println("end.")
	}
	{
		f := file{
			name: "test.dat",
			attr: attr{                          //显示初始化匿名字段
				perm: 0755,
			},
		}
		f.perm = 0644                            //直接设置匿名字段成员
		println(f.perm)                          //直接读取匿名字段成员
	}
	{
		d := data{
			File: os.File{},                     //嵌入其他包中的类型，则隐式字段名字不包含包名
		}
		fmt.Printf("%v\n", d)
	}
	{                                            //除接口指针和多级指针以外的任何命名类型都可作为匿名字段
		x := 100
		d := data1{
			int: &x,                             //使用基础类型作为字段名
			string: "abc",
		}
		fmt.Printf("%#v\n", d)
	}
	{
		d := data2{
			name: "data",
			file2: file2{"file"},
		}
		d.name = "data2"                         //访问data.name
		d.file2.name = "file2"                   //使用显示字段名访问data.file.name
	}
	{
		u := user{"Tom", 1}
		v := reflect.ValueOf(u)
		t := v.Type()
		for i, n := 0, t.NumField(); i < n; i++ {
			fmt.Printf("%s: %v\n", t.Field(i).Tag, v.Field(i))
		}             //可以用反射获取标签信息
	}
}