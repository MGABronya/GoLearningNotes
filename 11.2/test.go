package MainStart

import (
	"testing"
	"time"
)

func add(x, y int) int {
	return x + y
}

//性能测试函数以Benchmark为名称前缀，测试工具默认不会执行性能测试，需使用bench参数
func BenchmarkAdd(b *testing.B) {
	b.ReportAllocs() //可将测试函数设置为总时输出内存分配星系，无论使用benchmem参数与否
	time.Sleep(time.Second)
	b.ResetTimer() //重置事件
	for i := 0; i < b.N; i++ {
		add(1, 2)
		if i == 1 {
			b.StopTimer() //当测试函数中需要做一些额外的操作，用这个函数阻止计时器工作
			time.Sleep(time.Second)
			b.StartTimer() //恢复计时器
		}
	}
}

func main() {

}
