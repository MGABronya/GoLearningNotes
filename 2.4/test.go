package main

import (
	"fmt"
	"math"
	"strconv"
)

func test(x byte) {
	println(x)
}

func add(x, y int) int {
	return x + y
}

func main() {
	{
		a, b, c := 100, 0144, 0x64

		fmt.Println(a, b, c)

		fmt.Printf("0b%b, %#o, %#c\n", a, a, a)

		fmt.Println(math.MinInt8, math.MaxInt8)
	}

	{
		a, _ := strconv.ParseInt("1100100", 2, 32)
		b, _ := strconv.ParseInt("0144", 8, 32)
		c, _ := strconv.ParseInt("64", 16, 32)
		println(a, b, c)

		println("0b" + strconv.FormatInt(a, 2))
		println("0" + strconv.FormatInt(a, 8))
		println("0x" + strconv.FormatInt(a, 16))
	}

	{ //默认浮点类型为float64
		var a float32 = 1.123456789
		var b float32 = 1.12345678
		var c float32 = 1.12345681

		println(a, b, c)
		println(a == b, a == c)
		fmt.Printf("%v %v, %v\n", a, b, c)
	}

	{ //别名： byte  -uint8        rune -int32
		var a byte = 0x11
		var b uint8 = a
		var c uint8 = a + b
		test(c)
		/*   并不是拥有相同底层结构的就属于别名，就算在64位平台上int64和int底层结构完全一致，但也需要显示转换
		var x int = 100
		var y int64 = x
		add(x, y)
		*/
	}

}
