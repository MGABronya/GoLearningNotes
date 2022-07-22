package main

import (
	"fmt"
	"sync"
	"time"
	"unsafe"
)

func test1(x map[string]int) {
	fmt.Printf("x: %p\n", x)
}

func main() {
	{
		m := make(map[string]int)
		m["a"] = 1
		m["b"] = 2
		m2 := map[int]struct {
			x int
		}{
			1: {x: 100},
			2: {x: 200},
		}
		fmt.Println(m, m2)
	}
	{
		m := map[string]int{
			"a": 1,
			"b": 2,
		}
		m["a"] = 10              //修改
		m["c"] = 30              //新增
		if v, ok := m["d"]; ok { //用ok-idiom判断key是否存在，返回值
			println(v)
		}
		delete(m, "d") //删除键值对，不存在时不会报错
	}
	{
		m := make(map[string]int)
		for i := 0; i < 8; i++ {
			m[string('a'+i)] = i
		}
		for i := 0; i < 4; i++ {
			for k, v := range m { //对字典进行迭代，每次返回的键值次序都不相同
				print(k, ":", v, " ")
			}
			println()
		}
	}
	{
		type user struct {
			name string
			age  byte
		}
		m := map[int]user{
			1: {"Tom", 19},
		} //字典被设计为not addressable，不能直接修改value 的成员
		u := m[1] //只能返回整个value后修改再赋值回去
		u.age += 1
		m[1] = u
		m2 := map[int]*user{ //或者使用指针类型
			1: &user{"Jack", 20},
		}
		m2[1].age++ //m2[1]返回的是指针，可透过指针修改目标对象
	}
	{
		var m map[string]int
		println(m["a"]) //返回零值
		//m["a"] = 1                              //不能对nil字典进行写操作，但却能读
	}
	{
		var m map[string]int
		m2 := map[string]int{}       //为空，但是已经完成了初始化
		println(m == nil, m2 == nil) //为空的字典和nil是不同的
	}
	{
		m := make(map[int]int)
		for i := 0; i < 10; i++ {
			m[i] = i + 10
		}
		for k := range m {
			if k == 5 {
				m[100] = 1000
			}
			delete(m, k)      //不能保证迭代操作会删除新增的键值
			fmt.Println(k, m) //迭代期间删除或新增键值是安全的
		}
	}
	{ //如果字典正在写入，对字典执行其它的并发操作时，程序会崩溃
		m := make(map[string]int)
		go func() {
			for {
				m["a"] += 1 //写操作
				time.Sleep(time.Microsecond)
			}
		}()
		go func() {
			for {
				_ = m["b"] //读操作
				time.Sleep(time.Microsecond)
			}
		}()
		//select{}                                         //阻止进程退出
	}
	{
		var lock sync.RWMutex //使用读写锁，以获得最佳性能
		m := make(map[string]int)
		go func() {
			for {
				lock.Lock()
				m["a"] += 1
				lock.Unlock()
				time.Sleep(time.Microsecond)
			}
		}()
		go func() {
			for {
				lock.RLock()
				_ = m["b"]
				lock.RUnlock()
				time.Sleep(time.Microsecond)
			}
		}()
		//select{}
	}
	{
		m := make(map[string]int)
		test1(m)
		fmt.Printf("m : %p, %d\n", m, unsafe.Sizeof(m)) //字典对象本身是指针包装，传参时无需再次取地址
		m2 := map[string]int{}
		test1(m2)
		fmt.Printf("m : %p, %d\n", m2, unsafe.Sizeof(m2))
	}
	{
		m := make(map[int]int, 1000) //预先准备足够的空间
		for i := 0; i < 1000; i++ {
			m[i] = i
		}
	}
}
