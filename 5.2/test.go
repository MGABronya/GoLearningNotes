package main

import "fmt"

func test1(x [2]int) {
	fmt.Printf("x :%p, %v\n", &x, x)
}

func test2(x *[2]int) {
	fmt.Printf("x: %p, %v\n", x, *x)
	x[1] += 100
}

func main() {
	{
		var a [4]int
		b := [4]int{2, 5}
		c := [4]int{5, 3: 10} //可指定索引初始化
		d := [...]int{1, 2, 3}
		e := [...]int{10, 3: 100}
		fmt.Println(a, b, c, d, e)
	}
	{ //对于结构等复合类型，可省略元素初始化类型标签
		type user struct {
			name string
			age  byte
		}
		d := [...]user{
			{"Tom", 20},
			{"Mary", 18},
		}
		fmt.Printf("%#v\n", d)
	}
	{ //定义多维数组时，仅第一维允许使用...
		a := [2][2]int{
			{1, 2},
			{3, 4},
		}
		b := [...][2]int{
			{10, 20},
			{30, 40},
		}
		c := [...][2][2]int{
			{
				{1, 2},
				{3, 4},
			},
			{
				{10, 20},
				{30, 40},
			},
		}
		fmt.Println(a)
		fmt.Println(b)
		fmt.Println(c)
	}
	{ //内置函数len和cap都返回第一维度长度
		a := [2]int{}
		b := [...][2]int{
			{10, 20},
			{30, 40},
			{50, 60},
		}
		println(len(a), cap(a))
		println(len(b), cap(b))
		println(len(b[0]), cap(b[0]))
	}
	{
		var a, b [2]int
		println(a == b)
		c := [2]int{1, 2}
		d := [2]int{0, 1}
		println(c != d) //支持==，!=操作，比较的是类型，即存储的元素类型和长度是否一致
	}
	{
		x, y := 10, 20
		a := [...]*int{&x, &y}
		p := &a
		fmt.Printf("%T, %v\n", a, a)
		fmt.Printf("%T, %v\n", p, p)
	}
	{
		a := [...]int{1, 2}
		println(&a, &a[0], &a[1])
	}
	{
		a := [...]int{1, 2}
		p := &a
		p[1] += 10
		println(p[1])
	}
	{ //与C数组变量隐式作为指针使用不同，Go数组是值类型，赋值和传参操作都会复制整个数组数据
		a := [2]int{10, 20}
		var b [2]int
		b = a
		fmt.Printf("a :%p, %v\n", &a, a)
		fmt.Printf("b :%p, %v\n", &b, b)
		test1(a)
	}
	{ //可以改用数组或切片，避免数据复制
		a := [2]int{10, 20}
		test2(&a)
		fmt.Printf("a: %p, %v\n", &a, a)
	}
}
