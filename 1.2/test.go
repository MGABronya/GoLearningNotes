package main

import (
	"errors"
	"fmt"
	"time"
)

func div(a, b int) (int, error) {
	defer println("dispose...") //可以定义多个defer，按照FILO顺序执行
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func test(x int) func() { //返回函数类型
	return func() { //匿名函数
		println(x) //闭包
	}
}

type user struct { //结构体类型
	name string
	age  byte
}

type maneger struct { //匿名嵌入其它类型
	user
	title string
}

type X int

func (x *X) inc() { //名称前的参数称作 receiver， 作用类似于python self
	*x++
}

func (u user) Tostring() string {
	return fmt.Sprintf("%+v", u)
}

func (u user) Print() {
	fmt.Printf("%+v", u)
}

type Printer interface {
	Print()
}

func task(id int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d: %d\n", id, i)
		time.Sleep(time.Second)
	}
}

//消费者
func consumer(data chan int, done chan bool) {
	for x := range data { //接收数据，直到通道关闭
		println("recv: ", x)
	}
	done <- true //通知main，消费结束
}

//生产者
func producer(data chan int) {
	for i := 0; i < 4; i++ {
		data <- i //发送数据
	}
	close(data) //生产结束，关闭通道
}

func main() {
	//源文件
	println("hello, world!")
	fmt.Print("hello, world!")

	//变量
	var x int32
	var s = "hello, world!"
	println(x, s)

	//表达式if
	x = 100

	if x > 0 {
		println("x")
	} else if x < 0 {
		println("-x")
	} else {
		println("0")
	}

	//表达式switch
	switch {
	case x > 0:
		println("x")
	case x < 0:
		println("-x")
	default:
		println("0")
	}

	//表达式for
	for i := 0; i > 5; i++ {
		println(i)
	}

	for i := 4; i >= 0; i-- {
		println(i)
	}

	x = 0

	for x < 5 { //相当于 while (x < 5) ...
		println(x)
		x++
	}

	for { //相当于 while(true)
		println(x)
		x--
		if x < 0 {
			break
		}
	}

	y := []int{100, 101, 102}

	for i, n := range y { //i为索引，n为值
		println(i, ":", n)
	}

	//函数
	fmt.Println(div(10, 2))
	f := test(int(x))
	f()

	//数据 切片
	a := make([]int, 0, 5) //创建容量为5的切片

	for i := 0; i < 8; i++ { //追加数据，当超出容量限制时，自动分配更大的存储空间
		a = append(a, i)
	}
	fmt.Println(a)

	//数据 字典
	m := make(map[string]int) //创建字典类型对象

	m["a"] = 1      //添加或设置
	t, ok := m["b"] //使用 ok-idiom获取值，可知道key/value是否存在
	fmt.Println(t, ok)
	delete(m, "a") //删除

	//结构体
	var ma maneger
	ma.name = "Tom" //直接访问匿名字段的成员
	ma.age = 29
	ma.title = "CTO"
	fmt.Println(ma)

	//方法
	var xx X
	xx.inc()
	println(xx)

	println(ma.Tostring()) //调用user.Tostring

	//接口
	var p Printer = ma //只要包含接口所需的全部方法，即实现了该接口
	p.Print()

	//并发
	go task(1) //创建 goroutine
	go task(2)

	time.Sleep(time.Second * 6)

	done := make(chan bool) //用于接收消费结束信号
	data := make(chan int)  //数据管道
	go consumer(data, done) //启动消费者
	go producer(data)       //启动生产者

	<-done //阻塞，直到消费者发回结束信号
}
