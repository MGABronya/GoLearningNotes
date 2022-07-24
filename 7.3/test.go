package main

import "fmt"

type data int

func (d data) String() string {
	return fmt.Sprintf("data: %d", d)
}

func main() {
	{
		var d data = 15
		var x interface{} = d
		if n, ok := x.(fmt.Stringer); ok{      //转为更具体的接口类型
			fmt.Println(n)
		}
		if d2, ok := x.(data); ok {            //转回原始类型
			fmt.Println(d2)
		}
		e := x.(error)                         //此处转换会报错
		fmt.Println(e)
	}
	{          //使用ok-idiom模式，即便转换失败也不会引发panic，还可用switch语句在多种类型间做出推断匹配
		var x interface{} = func(x int) string {
			return fmt.Sprintf("d:%d", x)
		}
		switch v := x.(type){
		case nil:
			println("nil")
		case *int:
			println(*v)
		case func(int) string:
			println(v(100))
		case fmt.Stringer:
			fmt.Println(v)
		default:
			println("unknown")
		}
	}        //type switch不支持 fallthrought
}