package main

import "fmt"

// need to find two numbers from given array that sum = target
func main() {
	var target int = 10
	nums := make([]int, 0, 10)
	nums = append(nums, 1, 3, 4, 5, 7, 8)

	for i, x := range nums {
		for i2, y := range nums {
			if x+y == target && i2 > i {
				fmt.Println(i, i2)
			}

		}

	}
}
