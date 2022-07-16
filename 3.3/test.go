package main

func main() {
	type data struct {
		x int
		s string
	}
	b := data{
		1,
		"abc",
	}
	var a data = data{1, "abc"}
	print(a, b)
}
