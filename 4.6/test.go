package main

import (
	"fmt"
	"log"
)

type DivError struct { //自定义错误类型
	x, y int
}

func (DivError) Error() string { //实现 error 接口方法
	return "division by zero"
}

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, DivError{x, y} //返回自定义错误类型
	}
	return x / y, nil
}

func test1() {
	defer println("test.1")
	defer println("test.2")
	panic("i am dead")
}

func test2(x, y int) {
	z := 0
	func() {
		defer func() {
			if recover() != nil { //利用匿名函数保护x / y
				z = 0
			}
		}()
		z = x / y
	}()
	println("x / y = ", z)
}

func main() {
	{
		z, err := div(5, 0)
		if err != nil {
			switch e := err.(type) { //根据类型匹配
			case DivError:
				fmt.Println(e, e.x, e.y)
			default:
				fmt.Println(e)
			}
			log.Fatalln(err)
		}
		println(z)
	}
	{
		defer func() {
			if err := recover(); err != nil { //捕获错误
				log.Fatalln(err)
			}
		}()
		panic("i am dead") //引发错误
	}
	{
		defer func() {
			log.Println(recover()) //被外层捕获
		}()
		test1()
	}
	{ //recover必须在延迟调用函数中才能正常使用
		defer func() {
			for {
				if err := recover(); err != nil {
					log.Println(err)
				} else {
					log.Fatalln("gatal")
				}
			}
		}()
		defer func() {
			panic("you are dead") //类似重新抛出
		}() //可先recover捕获，包装后重新抛出
		panic("i am dead")
	}
	{

	}
}
