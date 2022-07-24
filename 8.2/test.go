package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
	"unsafe"
)

var exits = &struct {
	sync.RWMutex
	funcs   []func()
	signals chan os.Signal
}{}

func atexit(f func()) {
	exits.Lock()
	defer exits.Unlock()
	exits.funcs = append(exits.funcs, f)
}

func waitExit() {
	if exits.signals == nil {
		exits.signals = make(chan os.Signal)
		signal.Notify(exits.signals, syscall.SIGINT, syscall.SIGTERM)
	}
	exits.RLock()
	for _, f := range exits.funcs {
		defer f() //即便某些函数panic，延迟调用依旧可以确保后续函数执行
	} //延迟调用按FILO顺序执行
	<-exits.signals
}

type receiver struct {
	sync.WaitGroup
	data chan int
}

func newReceiver() *receiver {
	r := &receiver{
		data: make(chan int),
	}
	r.Add(1)
	go func() {
		defer r.Done()
		for x := range r.data { //接收消息，直到通道被关闭
			println("recv: ", x)
		}
	}()
	return r
}

type pool chan []byte

func newPool(cap int) pool {
	return make(chan []byte, cap)
}

func (p pool) get() []byte {
	var v []byte
	select {
	case v = <-p: //返回
	default:
		v = make([]byte, 10) //返回失败，新建
	}
	return v
}

func (p pool) put(b []byte) {
	select {
	case p <- b: //放回
	default: //放回失败，放弃
	}
}

func main() {
	{
		done := make(chan struct{}) //结束事件
		c := make(chan string)      //数据传输通道
		go func() {
			s := <-c //接收消息
			println(s)
			close(done) //关闭通道，作为结束通知
		}()
		c <- "hi" //发送消息
		<-done    //阻塞，直到有数据或管道关闭
	}
	{
		c := make(chan int, 3) //创建带有三个缓冲槽的异步通道
		c <- 1                 //缓冲区未满，不会阻塞
		c <- 2
		println(<-c) //缓冲区尚有数据，不会阻塞
		println(<-c)
	}
	{
		var a, b chan int = make(chan int, 3), make(chan int)
		var c chan bool
		println(a == b) //缓冲区大小仅是内部属性，不输于类型组成部分
		println(c == nil)
		fmt.Printf("%p, %d\n", a, unsafe.Sizeof(a))
	}
	{
		a, b := make(chan int), make(chan int, 3)
		b <- 1
		b <- 2
		println("a:", len(a), cap(a)) //0 0
		println("b:", len(b), cap(b)) //2 3
	}
	{
		done := make(chan struct{})
		c := make(chan int)
		go func() {
			defer close(done)
			for {
				x, ok := <-c //据此判断通道是否被关闭
				if !ok {
					return
				}
				println(x)
			}
		}()
		c <- 1
		c <- 2
		c <- 3
		close(c)
		<-done
	}
	{
		done := make(chan struct{})
		c := make(chan int)
		go func() {
			defer close(done)
			for x := range c { //循环获取消息，直到通道被关闭
				println(x)
			}
		}()
		c <- 1
		c <- 2
		c <- 3
		close(c)
		<-done
	}
	{
		var wg sync.WaitGroup
		ready := make(chan struct{})
		for i := 0; i < 3; i++ {
			wg.Add(1)
			go func(id int) {
				defer wg.Done()
				println(id, ":ready.") //运动员准备就绪
				<-ready                //等待发令
				println(id, ":running...")
			}(i)

			time.Sleep(time.Second)
			println("Ready? Go!")
			wg.Wait()
		}
	}
	{
		c := make(chan int, 3)
		c <- 10
		c <- 20
		close(c)
		for i := 0; i < cap(c)+1; i++ {
			x, ok := <-c //从已关闭接收数据，返回已缓冲数据或者零值
			println(i, ":", ok, x)
		}
	}
	{
		var wg sync.WaitGroup
		wg.Add(2)
		c := make(chan int)
		var send chan<- int = c //单向传入通道
		var recv <-chan int = c //单向传出通道
		go func() {
			defer wg.Done()
			for x := range recv {
				println(x)
			}
		}()
		go func() {
			defer wg.Done()
			defer close(c)
			for i := 0; i < 3; i++ {
				send <- i
			}
		}()
		wg.Wait()
	}
	{
		var wg sync.WaitGroup
		wg.Add(2)
		a, b := make(chan int), make(chan int)
		go func() { //接收端
			defer wg.Done()
			for {
				var (
					name string
					x    int
					ok   bool
				)
				select { //随机选择可用channel接收数据
				case x, ok = <-a:
					name = "a"
				case x, ok = <-b:
					name = "b"
				}
				if !ok { //如果任意通道关闭，则终止接收
					return
				}
				println(name, x) //输出接收的数据信息
			}
		}()
		go func() { //发送端
			defer wg.Done()
			defer close(a)
			defer close(b)
			for i := 0; i < 10; i++ {
				select { //随机选择发送channel
				case a <- i:
				case b <- i * 10:
				}
			}
		}()
		wg.Wait()
	}
	{
		var wg sync.WaitGroup
		wg.Add(3)
		a, b := make(chan int), make(chan int)
		go func() { //接收端
			defer wg.Done()
			for {
				select {
				case x, ok := <-a:
					if !ok { //如果通道关闭，则设置为nil，阻塞，不会被select选中
						a = nil
						break
					}
					println("a", x)
				case x, ok := <-b:
					if !ok {
						b = nil
						break
					}
					println("b", x)
				}
				if a == nil && b == nil { //全部结束，退出循环
					return
				}
			}
		}()
		go func() { //发送端 a
			defer wg.Done()
			defer close(a)
			for i := 0; i < 3; i++ {
				a <- i
			}
		}()
		go func() { //发送端b
			defer wg.Done()
			defer close(b)
			for i := 0; i < 5; i++ {
				b <- i * 10
			}
		}()
		wg.Wait()
	}
	{
		var wg sync.WaitGroup
		wg.Add(2)
		c := make(chan int)
		go func() { //接收端
			defer wg.Done()
			for {
				var v int
				var ok bool
				select {
				case v, ok = <-c: //随机选择case
					println("a1:", v)
				case v, ok = <-c:
					println("a2:", v)
				}
				if !ok {
					return
				}
			}
		}()
		go func() {
			defer wg.Done()
			defer close(c)
			for i := 0; i < 10; i++ {
				select {
				case c <- i:
				case c <- i * 10:
				}
			}
		}()
		wg.Wait()
	}
	{
		done := make(chan struct{})
		c := make(chan int)
		go func() {
			defer close(done)
			for {
				select {
				case x, ok := <-c:
					if !ok {
						return
					}
					fmt.Println("data: ", x)
				default: //避免select阻塞
				}
				fmt.Println(time.Now())
				time.Sleep(time.Second)
			}
		}()
		time.Sleep(time.Second * 5)
		c <- 100
		close(c)
		<-done
	}
	{
		done := make(chan struct{})
		data := []chan int{ //数据缓冲区
			make(chan int, 3),
		}
		go func() {
			defer close(done)
			for i := 0; i < 10; i++ {
				select { //生产数据
				case data[len(data)-1] <- i:
				default:
					data = append(data, make(chan int, 3)) //如果通道已满，生成新的缓存通道
				}
			}
		}()
		<-done
		for i := 0; i < len(data); i++ { //显示所有数据
			c := data[i]
			close(c)
			for x := range c {
				println(x)
			}
		}
	}
	{
		r := newReceiver()
		r.data <- 1
		r.data <- 2
		close(r.data) //关闭通道，发出结束通知
		r.Wait()      //等待接收者处理结束
	}
	{
		runtime.GOMAXPROCS(4)
		var wg sync.WaitGroup
		sem := make(chan struct{}, 2) //最多允许两个并发同时执行
		for i := 0; i < 5; i++ {
			wg.Add(1)
			go func(id int) {
				defer wg.Done()
				sem <- struct{}{}        //获取信号
				defer func() { <-sem }() //释放信号
				time.Sleep(time.Second * 2)
				fmt.Println(id, time.Now())
			}(i)
		}
		wg.Wait()
	}
	{
		go func() {
			for {
				select {
				case <-time.After(time.Second * 5):
					fmt.Println("timeout...")
					os.Exit(0)
				}
			}
		}()
		go func() {
			tick := time.Tick(time.Second)
			for {
				select {
				case <-tick:
					fmt.Println(time.Now())
				}
			}
			<-(chan struct{})(nil)
		}()
	}
	{
		atexit(func() { println("exit1...") })
		atexit(func() { println("exit2...") })
		waitExit()
	}
}
