package main

func main() {
	{
		x, y := 1, 2
		defer func(a int) {
			println("defer x, y = ", a, y) //y为闭包引用
		}(x) //注册时复制调用参数

		x += 100 //对x的修改不会影响延迟调用，因为函数注册时已经将x复制了
		y += 200
		println(x, y)
	}
	{
		defer println("a")
		defer println("b") //按FILO顺序执行
	}
}
