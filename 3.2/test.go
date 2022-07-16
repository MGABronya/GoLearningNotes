package main

import "fmt"

func main() {
	{
		const v = 20 //无显示类型声明的常量
		var a byte = 10
		b := v + a //v自动转换为 byte/uint8 类型
		fmt.Printf("%T, %v\n", b, b)

		const c float32 = 1.2
		d := c + v //v自动转换为 float32 类型
		fmt.Printf("%T, %v\n", d, d)
	}
	{
		a := 1.0 << 3
		fmt.Printf("%T, %v\n", a, a) //int， 8
		var s uint = 3
		//b := 1.0 << s    //无效操作，因为b没有提供类型，那么编译器通过1.0推断，显然无法对浮点数做位移操作

		var c int32 = 1.0 << s       //自动将1.0转换为int32类型
		fmt.Printf("%T, %v\n", c, c) //int32，8
	}
	{
		x := 10
		var p *int = &x //获取地址，保存到指针变量
		*p += 20        //用指针间接引用，并更新对象
		println(p, *p)  //输出指针所存储的地址，以及目标对象
	} //指针支持相等运算符，但不能做加减法运算和类型转换
	{ //指针没有专门指向成员的->运算符，统一使用.选择表达式
		a := struct {
			x int
		}{}
		a.x = 100
		p := &a
		p.x += 100
		println(p.x)
	}
	{
		var a, b struct{}
		println(&a, &b)
		println(&a == &b, &a == nil) //即使长度为0，可对象依然是合法存在的，拥有合法内存地址
	}

}
