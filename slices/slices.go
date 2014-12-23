package main

import "fmt"
import "math"
import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
	rows := make([][]uint8, dy)

	for y := range rows {
		rows[y] = make([]uint8, dx)
		for x := range rows[y] {
			rows[y][x] = uint8(math.Pow(float64(x), float64(y)))
		}
	}

	fmt.Println(rows)

	return rows
}

func main() {
	pic.Show(Pic)
}
