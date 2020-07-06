package _main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	aa := make([][]uint8, dx, dy)
	for y := 0; y < dy; y++ {
		a := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			a[x] = uint8(x * y)
		}
		aa[y] = a
	}
	return aa
}

func main() {
	pic.Show(Pic)
}
