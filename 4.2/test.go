package main

import "fmt"

func test(x *int) {
	fmt.Printf("pointer: %p, target: %v\n", &x, x) //输出形参 x 的地址
}

func test1(p *int) {
	go func() { //延长了p的生命周期，导致其所指对象被分配到栈上，会增加垃圾回收所需要的时间
		println(p) //逃逸
	}()
}

func test2(s string, a ...int) { //变参本质上是一个切片
	fmt.Printf("%T, %v\n", a, a) //显示类型和值
}

func test3(a ...int) { //变参复制的是切片而非底层数组，所以可以直接修改底层数组的数据
	for i := range a {
		a[i] += 100
	}
}

func main() {
	{
		a := 0x100
		p := &a
		fmt.Printf("pointer: %p, target: %v\n", &p, p) //输出实参p的地址
		test(p)                                        //此处可见函数为值拷贝传递
	}
	{
		test2("Abc", 1, 2, 3, 4)
		a := [3]int{10, 20, 30}
		test2("123", a[:]...) //转换为slice后展开
	}
	{
		a := []int{1, 2, 3}
		test3(a...)
		fmt.Println(a)
	}
}
