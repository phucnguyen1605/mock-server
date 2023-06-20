package main

import (
	"fmt"
)

func main() {
	a := []int{1, 5, 6, 20}
	b := []int{2, 3, 5, 7, 100, 101}
	maxA := len(a)
	maxB := len(b)
	c := []int{}

	i, j := 0, 0
	for i < maxA && j < maxB {
		if a[i] < b[j] {
			c = append(c, a[i])
			i++
		} else if a[i] > b[j] {
			c = append(c, b[j])
			j++
		} else {
			c = append(c, a[i], b[j])
			i++
			j++
		}
	}

	for i < maxA {
		c = append(c, a[i])
		i++
	}
	for j < maxB {
		c = append(c, b[j])
		j++
	}

	fmt.Println(c)
}
