package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"
)

var c int

func counter() int {
	c++
	return c
}

func count() {
	x := 0
	for i := 0; i < math.MaxUint32; i++ {
		x += i
	}
	println(x)
}

func test1(n int) {
	for i := 0; i < n; i++ {
		count()
	}
}

func test2(n int) {
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			count()
			wg.Done()
		}()
	}
	wg.Wait()
}

func main() {
	{ //只需要在函数调用前添加go关键字即可创建并发任务
		go println("hello, world!")
		go func(s string) {
			println(s)
		}("hello, world!")
	}
	{
		a := 100
		go func(x, y int) {
			time.Sleep(time.Second)
			println("go:", x, y)
		}(a, counter()) //立即计算并复制参数
		a += 100
		println("main:", a, counter())
		time.Sleep(time.Second * 3)
	}
	{
		exit := make(chan struct{})
		go func() {
			time.Sleep(time.Second)
			println("goroutine done.")
			close(exit)
		}()
		println("main...")
		<-exit
		println("main exit.")
	}
	{
		var wg sync.WaitGroup
		for i := 0; i < 10; i++ {
			wg.Add(1) //累计加数
			go func(id int) {
				defer wg.Done() //递减计数
				time.Sleep(time.Second)
				println("goroutine", id, "done.")
			}(i)
		}
		println("main...")
		wg.Wait() //阻塞，直到计数归零
		println("main.exit.")
	}
	{
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			wg.Wait() //等待归零，解除阻塞
			println("wait exit.")
		}()
		go func() {
			time.Sleep(time.Second)
			println("done.")
			wg.Done() //递归计数
		}()
		wg.Wait()
		println("main.eixt.") //等待归零，解除阻塞
	}
	{
		n := runtime.GOMAXPROCS(0) //用于修改线程数量，如果为0则默认与处理器数量相等
		test1(n)
		test2(n)
	}
	{
		var wg sync.WaitGroup
		var gs [5]struct { //用于实现局部存贮(TLS)功能
			id     int //编号
			result int //返回值
		}
		for i := 0; i < len(gs); i++ {
			wg.Add(1)
			go func(id int) { //使用参数避免闭包延迟求值
				defer wg.Done()
				gs[id].id = id
				gs[id].result = (id + 1) * 100
			}(i)
		}
		wg.Wait()
		fmt.Printf("%+v\n", gs)
	}
	{
		runtime.GOMAXPROCS(1)
		exit := make(chan struct{})
		go func() { //任务a
			defer close(exit)
			go func() { //任务b放在此处，是为了确保a优先执行
				println("b")
			}()
			for i := 0; i < 4; i++ {
				println("a: ", i)
				if i == 1 {
					runtime.Gosched() //释放线程执行其他任务
				}
			}
		}()
		<-exit
	}
	{
		exit := make(chan struct{})
		go func() {
			defer close(exit)      //执行
			defer fmt.Println("a") //执行
			func() {
				defer func() {
					println("b", recover() == nil) //执行，recover为nil
				}()
				func() {
					println("c")
					runtime.Goexit()   //立即终止整个调用堆栈
					println("c done.") //不会执行
				}()
				println("b done.") //不会执行
			}()
		}()
		<-exit
		println("main exit.")
	}
}
