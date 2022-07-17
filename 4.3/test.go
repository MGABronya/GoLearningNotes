package main

import "errors"

func div(x, y int) (z int, err error) { //命名返回值和参数一样，可以当作函数局部变量使用，最后由return隐式返回
	if y == 0 {
		err = errors.New("division by zero")
		return
	}
	z = x / y
	return
}

func main() {

}
