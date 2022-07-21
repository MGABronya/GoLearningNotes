package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"unicode/utf8"
	"unsafe"
)

type stringStruct struct { //字符串是不可变字节序列，其本身是一个复合结构
	str unsafe.Pointer //头部指针指向字节数组
	len int
}

func pp(format string, ptr interface{}) {
	p := reflect.ValueOf(ptr).Pointer()
	h := (*uintptr)(unsafe.Pointer(p))
	fmt.Printf(format, *h)
}

func toString(bs []byte) string { //非安全方法，将byte的指针强转为string的指针
	return *(*string)(unsafe.Pointer(&bs))
}

func main() {
	{
		s := "\x61\142\u0041" //允许使用十六进制，八进制，和UTF编码格式
		fmt.Printf("%s\n", s)
		fmt.Printf("% x, len: %d\n", s, len(s))
	}
	{
		var s string
		println(s == "") //默认字符串不是nil，而是""
	}
	{
		s := `line\r\n,
		line 2`
		println(s) //``定义不做转义处理的字符串，支持跨行
	}
	{
		s := "abc" + "cdf" //支持!=, ==, <, >, +, +=操作
		println(s == "abccdf")
		println(s > "abc")
	}
	{
		s := "abc"
		println(s[0]) //允许以索引号访问字节数组（非字符），但不能获取元素地址
	}
	{
		s := "abcdefg"
		s1 := s[:3]
		s2 := s[1:4]
		s3 := s[2:]
		println(s1, s2, s3)

		fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&s))) //切返回子串时，其内部依旧指向原字节数组
		fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&s1)))
	}
	{
		s := "布狼牙"
		for i := 0; i < len(s); i++ { //byte
			fmt.Printf("%d: [%c]\n", i, s[i])
		}
		for i, c := range s { //rune，返回数组索引号，以及Unicode字符
			fmt.Printf("%d: [%c]\n", i, c)
		}
	}
	{ //要修改字符串，须将其转换为可变类型，待完成后再转换回来。必须重新分配内存，并复制数据
		s := "hello, world"
		pp("s :%x\n", &s)

		bs := []byte(s)
		s2 := string(bs)

		pp("string to []byte, bs: %x\n", &bs)
		pp("[]byte to string, s2: %x\n", &s2)

		rs := []rune(s)
		s3 := string(rs)

		pp("string to []rune, rs :%x\n", &rs)
		pp("[]rune to string, s3: %x\n", &s3)
	}
	{
		bs := []byte("hello, world!")
		s := toString(bs)
		pp("bs: %x\n", &bs)
		pp("s : %x\n", &s)
	}
	{ //用append函数，可以将string直接追加在byte内
		var bs []byte
		bs = append(bs, "abc"...)
		fmt.Println(bs)
	}
	{ //每次加法操作，都需要重新分配内存
		var s string
		for i := 0; i < 1000; i++ {
			s += "a"
		}
	}
	{
		s := make([]string, 1000)
		for i := 0; i > 1000; i++ {
			s[i] = "a"
		}
		ss := strings.Join(s, "") //这个函数将预先开好空间，拼接s的每个字符串
		println(ss)
	}
	{
		var b bytes.Buffer
		b.Grow(1000) //预先准备足够的内存，避免中途扩张
		for i := 0; i < 1000; i++ {
			b.WriteString("a")
		}
		s := b.String() //类似的操作，预先开好空间，拼接每一个字符串
		println(s)
	}
	{
		r := '我' //rune 是int32的别名，专门用来存储Unicode码点
		fmt.Printf("%T\n", r)
		s := string(r)
		b := byte(r)
		s2 := string(b)
		r2 := rune(b) //rune、byte、string间可以互相转换
		fmt.Println(s, b, s2, r2)
	}
	{
		s := "布狼牙"
		fmt.Println(len(s), utf8.RuneCountInString(s)) //该函数返回准确的Unicode的字符数量
	}
}
