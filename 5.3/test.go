package main

import (
	"errors"
	"fmt"
	"reflect"
	"unsafe"
)

type sliceStruct struct { //slice内部通过指针引用底层数组
	array unsafe.Pointer
	len   int //len表示用于限定可读的写元素数量
	cap   int //cap表示切片所引用数组片段的真实长度
}

func main() {
	{
		x := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		s := x[2:5]
		for i := 0; i < len(s); i++ {
			println(s[i])
		}
	}
	{
		s1 := make([]int, 3, 5)    //指定len、cap，底层数组初始化为零值
		s2 := make([]int, 3)       //省略cap，和len相等
		s3 := []int{10, 20, 5: 30} //按初始化元素分配底层数组，并设置len、cap
		fmt.Println(s1, len(s1), cap(s1))
		fmt.Println(s2, len(s2), cap(s2))
		fmt.Println(s3, len(s3), cap(s3))
	}
	{
		var a []int                 //未初始化的切片。注意[]为切片，数组需要填值
		b := []int{}                //完成了初始化的切片
		println(a == nil, b == nil) //切片只能和nil比较
		fmt.Printf("a: %#v\n", (*reflect.SliceHeader)(unsafe.Pointer(&a)))
		fmt.Printf("a: %#v\n", (*reflect.SliceHeader)(unsafe.Pointer(&b)))
		fmt.Printf("a size: %d\n", unsafe.Sizeof(a))
	}
	{
		s := []int{0, 1, 2, 3, 4}
		p := &s     //取header地址
		p0 := &s[0] //取array[0]地址
		p1 := &s[1]
		println(p, p0, p1)
		(*p)[0] += 100 //*[]int不支持索引操作，需要先返回[]int对象
		*p1 += 100     //直接用元素指针操作
		fmt.Println(s)
	}
	{
		x := [][]int{
			{1, 2},
			{10, 20, 30},
			{100},
		}
		fmt.Println(x[1])
		x[2] = append(x[2], 200, 300)
		fmt.Println(x[2])
	}
	{
		d := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		s1 := d[3:7]
		s2 := s1[1:3]       //将切片视作新切片的数据源，不能超出cap，但不受len限制
		for i := range s2 { //新建切片对象依旧指向原底层数组
			s2[i] += 100
		}
		fmt.Println(d)
		fmt.Println(s1)
		fmt.Println(s2)
	}
	{
		stack := make([]int, 0, 5)
		push := func(x int) error {
			n := len(stack)
			if n == cap(stack) {
				return errors.New("stack is full")
			}
			stack = stack[:n+1]
			stack[n] = x
			return nil
		}
		pop := func() (int, error) {
			n := len(stack)
			if n == 0 {
				return 0, errors.New("stack is empty")
			}
			x := stack[n-1]
			stack = stack[:n-1]
			return x, nil
		}
		for i := 0; i < 7; i++ {
			fmt.Printf("push %d: %v, %v\n", i, push(i), stack)
		}
		for i := 0; i < 7; i++ {
			x, err := pop()
			fmt.Printf("pop: %d, %v, %v\n", x, err, stack)
		}
	}
	{
		s := make([]int, 0, 5)
		s1 := append(s, 10)
		s2 := append(s1, 20, 30)
		fmt.Println(s, len(s), cap(s))
		fmt.Println(s1, len(s1), cap(s1))
		fmt.Println(s2, len(s2), cap(s2))
	}
	{
		s := make([]int, 0, 100)
		s1 := s[:2:4]
		s2 := append(s1, 1, 2, 3, 4, 5, 6) //超出s1cap限制，分配新底层数组，注意，是超出cap而非原数组长度
		fmt.Printf("s1: %p: %v\n", &s1[0], s1)
		fmt.Printf("s1: %p: %v\n", &s2[0], s2)
		fmt.Printf("s data:%v\n", s[:10])                       //append并未向原数组写回数据（因为新分配了数组）
		fmt.Printf("s1 cap:%d, s2 cap: %d\n", cap(s1), cap(s2)) //新数组的长度为cap*2
	}
	{
		var s []int
		s = append(s, 1, 2, 3) //向nil切片追加数据时，会为其分配底层数组内存
		fmt.Println(s)
	}
	{
		s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

		s1 := s[5:8]
		n := copy(s[4:], s1) //将s1copy到了s[4]及往后位置，n为copy的长度
		fmt.Println(n, s)
		s2 := make([]int, 6)
		n = copy(s2, s)
		fmt.Println(n, s2)
	}
	{
		b := make([]byte, 3)
		n := copy(b, "abcde")
		fmt.Println(n, b)
	}
}
