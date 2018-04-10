package main

import (
	"golang.org/x/tour/pic"
	"fmt"
)

func Pic(dx, dy int) [][]uint8 {

	slicing := make([][]uint8, dy)

	for i, _ := range slicing {
		slicing[i] = make([]uint8, dx)
	}

	fmt.Printf("len %d and cap %d \n", len(slicing), cap(slicing))
	fmt.Printf("len %d and cap %d of each slice \n", len(slicing[0]), cap(slicing[0]))

	for i, val := range slicing{
		for j, _ := range val{
			val[j] = uint8(j*i / 2 +3)
		}
	}

	return slicing
}

func main() {
	pic.Show(Pic)
}

