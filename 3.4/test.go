package main

import "fmt"

func main() {
	{
		x := 3
		if x > 5 {
			println("a")
		} else if x < 5 && x > 0 {
			println("b")
		} else {
			println("z")
		}
	}
	{
		if x := 10; x == 0 {
			println("a")
		}
	}
	{
		a, b, c, x := 1, 2, 3, 2
		switch x {
		case a, b:
			println("a | b")
		case c:
			println("c")
		case 4:
			println("d")
		default:
			println("z")
		}
	}
	{
		switch x := 5; x { //编译器确保不会先执行default块
		default:
			x += 100
			println(x)
		case 5:
			x += 50
			println(x)
		}
	}
	{
		switch x := 5; x {
		default:
			println(x)
		case 5:
			x += 10
			println(x)  //注意，fallthrough必须放在case块结尾，可以使用break语句阻止
			fallthrough //继续执行下一个case，但不在匹配条件表达式
		case 6:
			x += 20
			println(x)
		}
	}
	{
		switch x := 5; { //相当于 "switch x := 5; true { ... }"
		case x > 5:
			println("a")
		case x > 0 && x <= 5: //不能写成 x > 0, x <= 5， 因为多条件是or关系
			println("b")
		default:
			println("z")
		}
	}
	{
		data := [3]string{"a", "b", "c"}
		for i := range data { //只返回1st value
			println(i, data[i])
		}

		for _, s := range data { //忽略1st value
			println(s)
		}
		for range data { //仅迭代，不返回，可用来执行清空channel等操作

		}
	}
	{
		data := [3]int{10, 20, 30}
		for i, x := range data { //第一次进循环时，数组data已经被复制了
			if i == 0 {
				data[0] += 100
				data[1] += 200
				data[2] += 300
			}
			fmt.Printf("x:%d, data:%d\n", x, data[i])
		}
		for i, x := range data[:] { //每次进循环时复制一次切片
			if i == 0 {
				data[0] += 100
				data[1] += 200
				data[2] += 300
			}
			fmt.Printf("x:%d, data:%d\n", x, data[i])
		}
	}
	{
		for i := 0; i < 3; i++ {
			println(i)
			if i > 1 {
				goto exit //跳转
			}
		}
	exit:
		println("exit.")
	}
	{
		for i := 0; i < 10; i++ {
			if i%2 == 0 {
				continue
			}
			if i > 5 {
				break
			}
			println(i)
		}
	}
	{
	outer:
		for x := 0; x < 5; x++ {
			for y := 0; y < 10; y++ {
				if y > 2 {
					println()
					continue outer
				}
				if x > 2 {
					break outer
				}
				print(x, ":", y, " ")
			}
		}
	}
}
