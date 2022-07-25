package main

import (
	"reflect"
	"strings"
)

func add(args []reflect.Value) (results []reflect.Value) {
	if len(args) == 0 {
		return nil
	}
	var ret reflect.Value
	switch args[0].Kind() {
	case reflect.Int:
		n := 0
		for _, a := range args {
			n += int(a.Int())
		}
		ret = reflect.ValueOf(n)
	case reflect.String:
		ss := make([]string, 0, len(args))
		for _, s := range args {
			ss = append(ss, s.String())
		}
		ret = reflect.ValueOf(strings.Join(ss, ""))
	}
	results = append(results, ret)
	return
}

//将函数指针参数指向通用算法函数
func makeAdd(fptr interface{}) {
	fn := reflect.ValueOf(fptr).Elem()    //这是关键
	v := reflect.MakeFunc(fn.Type(), add) //指向通用算法函数
	fn.Set(v)
}

func main() {
	{
		var intAdd func(x, y int) int
		var strAdd func(a, b string) string

		makeAdd(&intAdd)
		makeAdd(&strAdd)

		println(intAdd(100, 200))
		println(strAdd("hello, ", "world!"))
	}
}
