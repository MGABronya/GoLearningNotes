package main

import "strconv"

/*
	命名建议
	以字母或下划线开始，由多个字母、数字和下划线组合而成
	区分大小写
	使用驼峰拼写格式
	局部变量优先使用短名
	不要使用保留关键字
	不建议使用与预定义常量、类型、内置函数相同的名字
	专有名词通常会全部大写，例如HTML
*/
func main() {
	var c int                 //用c代替count
	for i := 0; i < 10; i++ { //用i代替index
		c++
	}
	println(c)

	x, _ := strconv.Atoi("12") //忽略 Atoi的err返回值

	println(x)
}

//首字母大小写决定了作用域，大写的为导出成员，可以被包外引用
