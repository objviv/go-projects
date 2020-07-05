package fibonacci

import "fmt"

func fibonacci(i int) uint64 {
	if i < 0 {
		panic("Got a number less than 0")
	}

	var ret uint64

	if 0 == i {
		ret = 0
	} else if 1 == i {
		ret = 1
	} else {
		ret = fibonacci(i-1) + fibonacci(i-2)
	}
	return ret
}

func main() {
	var r = fibonacci(10)
	fmt.Printf("10th fibonacci term is %v", r)
}
