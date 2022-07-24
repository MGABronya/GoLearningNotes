package main

import (
	"sync"
	"time"
)

type data1 struct {
	sync.Mutex
}

func (d *data1) test1(s string) { //必须为pointer-receiver类型，不然会锁失效
	d.Lock()
	defer d.Unlock()
	for i := 0; i < 5; i++ {
		println(s, i)
		time.Sleep(time.Second)
	}
}

func main() {
	{
		var wg sync.WaitGroup
		wg.Add(2)
		var d data1
		go func() {
			defer wg.Done()
			d.test1("read")
		}()
		go func() {
			defer wg.Done()
			d.test1("write")
		}()
		wg.Wait()
	}
}
