package main

var x = 100

func main() {
	/*
		var x int               //自动初始化为0
		var y = false           //自动推断为bool类型
		var X, Y int            //相同类型的多个变量
		//var a, s = 100, "abc"   //不同类型初始化值
	*/
	//简短模式，定义变量，同时显式初始化，不能提供数据类型，只能用在函数内部
	//a, s := 100, "abc"

	println(&x, x) //全局变量

	x := "abc"
	println(&x, x) //局部变量

	x, y := "abc", 100 //x退化为赋值操作，仅有y是变量定义
	println(&x, x)
	println(y)
	//退化赋值的前提条件：至少有一个新变量被定义，且必须是同一作用域

	//多变量赋值
	a, b := 1, 2
	a, b = b+3, a+2 //先计算出右值，再对变量赋值
	println(a, b)
}
