package main

import "fmt"

func main () {
rome := map [rune] int {
	'I' : 1,
	'V' : 5,
	'X' : 10,
	'L' : 50,
	'C' : 100,
	'D' : 500,
	'M' : 1000,
}
// want to convert s to arabian numer
s := "MMDCCLXXXIX"
arab := make([]int, 0)
sum := 0

// making slice with arab numbers
		for _, c := range s {
			x := rome [c]
			arab = append(arab, x)
		}
// summing arab numbers depending on their position against max
		max := arab [0]	
		for _, num := range arab {
			if num > max {
			sum = sum + (num - 2*max)
			max = num
			} else {
				sum = sum + num
				max = num
			}
		}
	fmt.Println (sum)
}
