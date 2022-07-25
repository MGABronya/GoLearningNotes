package main

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func add(x, y int) int {
	return x + y
}

//测试函数以Test为名称前缀，测试命令为go test
func TestAdd1(t *testing.T) {
	if add(1, 2) != 3 {
		t.FailNow()
	}
}

func TestAdd2(t *testing.T) {
	var tests = []struct {
		x      int
		y      int
		expect int
	}{
		{1, 1, 2},
		{2, 2, 4},
		{3, 2, 5},
	}
	for _, tt := range tests {
		actual := add(tt.x, tt.y)
		if actual != tt.expect {
			t.Errorf("add(%d, %d): expect %d, actual %d", tt.x, tt.y, tt.expect, actual)
		}
	}
}

//go test会该为执行TestMain函数，而不再是具体的测试用例, 通过MainStart自行构建M对象
func TestMain(m *testing.M) {
	match := func(pat, str string) (bool, error) { //pat，命令行参数-run提供的过滤条件
		return true, nil //str， InternalTest.Name
	}
	tests := []testing.InternalTest{ //用例列表，可排序
		{"b", TestB},
		{"a", TestA},
	}
	benchmarks := []testing.InternalBenchmark{}
	examples := []testing.InternalExample{}
	m = testing.MainStart(match, tests, benchmarks, examples)
	os.Exit(m.Run())
}

//使用Parallel可有效利用多核并行优势，缩短测试时间
func TestA(t *testing.T) {
	t.Parallel()
	time.Sleep(time.Second * 2)
}

func TestB(t *testing.T) {
	if os.Args[len(os.Args)-1] == "b" {
		t.Parallel()
	}
	time.Sleep(time.Second * 2)
}

//例代码，如果没有ouput注释则不会执行，它通过比对输出stdout和内部output注释是否一致来判断是否成功
func ExampleAdd() {
	fmt.Println(add(1, 2))
	fmt.Println(add(2, 2))
	//Output:
	//3
	//4
}

func main() {

}
